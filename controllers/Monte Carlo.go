package controllers

import (
	"math"
	"sort"
	"sync"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
	"gonum.org/v1/plot/plotter"
)

// MCSetup -
type MCSetup struct {
	Sims        int
	ERV         float64
	CPI         float64
	YieldShift  float64
	Void        int
	Probability float64
	OpEx        float64
	Hazard      float64
}

// MCResultSlice -
type MCResultSlice struct {
	EndCash     []float64
	CashBalance [][]float64
	EndNCF      []float64
	NCF         [][]float64
	IRR         []float64
	EM          []float64
	YTM         []float64
	Duration    []float64
}

// MCResults - used for final display. NOI, MarketValue and NCF have stats per year, IRR is for the hold period
type MCResults struct {
	EndCash     MCStats
	CashBalance []Ribbon
	EndNCF      MCStats
	NCF         []Ribbon
	IRR         MCStats
	EM          MCStats
	YTM         MCStats
	Duration    MCStats
}

// MCStats -
type MCStats struct {
	Mean     float64
	Variance float64
	StDev    float64
	Skew     float64
	Kurtosis float64
	LRalpha  float64
	LRbeta   float64
	Hist     Hist
}

// Hist -
type Hist struct {
	Keys []float64
	Vals []float64
}

// MonteCarlo -
func (e *Entity) MonteCarlo() {
	e.MCSlice = make([]*Entity, e.MCSetup.Sims)
	duration := dateintdiff(e.SalesDate.Dateint, e.StartDate.Dateint)
	e.MCResultSlice = MCResultSlice{EndCash: make([]float64, e.MCSetup.Sims), CashBalance: make([][]float64, duration-2), EndNCF: make([]float64, e.MCSetup.Sims), NCF: make([][]float64, duration-2), IRR: make([]float64, e.MCSetup.Sims), EM: make([]float64, e.MCSetup.Sims), YTM: make([]float64, e.MCSetup.Sims), Duration: make([]float64, e.MCSetup.Sims)}
	for i := 0; i < duration-2; i++ {
		e.MCResultSlice.CashBalance[i] = make([]float64, e.MCSetup.Sims)
		e.MCResultSlice.NCF[i] = make([]float64, e.MCSetup.Sims)
	}
	wg := sync.WaitGroup{}
	for i := 1; i <= e.MCSetup.Sims; i++ {
		wg.Add(1)
		go func(e *Entity, index int) {
			defer wg.Done()
			temp := *e
			tempentitydata := EntityStore[e.MasterID]
			tempentitydata.ERVGrowth = NormalSample(e.GrowthInput["ERV"], e.MCSetup.ERV)
			tempentitydata.CPIGrowth = NormalSample(e.GrowthInput["CPI"], e.MCSetup.CPI)
			tempentitydata.OpExpercent = NormalSample(e.OpEx.PercentOfTRI, e.MCSetup.OpEx)
			tempentitydata.YieldShift = NormalSample(e.Valuation.YieldShift, e.MCSetup.YieldShift)
			tempentitydata.GLA.Void = int(NormalSample(float64(e.GLA.Void), float64(e.MCSetup.Void)))
			tempentitydata.GLA.Probability = NormalSample(e.GLA.Probability, e.MCSetup.Probability)
			tempentitydata.GLA.Default.Hazard = LogNormalSample(float64(e.GLA.Default.Hazard), float64(e.MCSetup.Hazard))
			//
			temp.UpdateEntity(true, tempentitydata)
			temp.MCResults.EndCash.Mean = temp.COA[temp.SalesDate.Dateint].CashBalance
			temp.MCResults.EndNCF.Mean = temp.COA[temp.SalesDate.Dateint].NetCashFlow
			date := Dateadd(temp.StartDate, 1)
			// mutex lock
			e.Mutex.Lock()
			e.MCSlice[index-1] = &temp
			e.MCResultSlice.EndCash[index-1] = temp.COA[temp.SalesDate.Dateint].CashBalance
			// assign results for every month
			for ii := 0; ii < duration-2; ii++ {
				e.MCResultSlice.CashBalance[ii][index-1] = temp.COA[date.Dateint].CashBalance
				e.MCResultSlice.NCF[ii][index-1] = temp.COA[date.Dateint].NetCashFlow
				date.Add(1)
			}
			e.MCResultSlice.EndNCF[index-1] = temp.COA[temp.SalesDate.Dateint].NetCashFlow
			e.MCResultSlice.IRR[index-1] = temp.Metrics.IRR.NetLeveredAfterTax
			e.MCResultSlice.EM[index-1] = temp.Metrics.EM.NetLeveredAfterTax
			e.MCResultSlice.YTM[index-1] = temp.Metrics.BondHolder.YTM
			e.MCResultSlice.Duration[index-1] = temp.Metrics.BondHolder.Duration
			e.Mutex.Unlock()
		}(e, i)
	}
	wg.Wait()
	e.MCResults.EndCash = MCStatsCalc(e.MCResultSlice.EndCash, e.MCSetup.Sims)
	e.MCResults.CashBalance = RibbonPlot(e.MCResultSlice.CashBalance, duration-2, 50, e.MCSetup.Sims)
	e.MCResults.EndNCF = MCStatsCalc(e.MCResultSlice.EndNCF, e.MCSetup.Sims)
	e.MCResults.NCF = RibbonPlot(e.MCResultSlice.NCF, duration-2, 50, e.MCSetup.Sims)
	switch e.Strategy {
	case "Standard":
		e.MCResults.IRR = MCStatsCalc(e.MCResultSlice.IRR, e.MCSetup.Sims)
		e.MCResults.EM = MCStatsCalc(e.MCResultSlice.EM, e.MCSetup.Sims)
	case "Balloon", "Pure Discount":
		e.MCResults.YTM = MCStatsCalc(e.MCResultSlice.YTM, e.MCSetup.Sims)
		e.MCResults.Duration = MCStatsCalc(e.MCResultSlice.Duration, e.MCSetup.Sims)
	}
}

// MCStatsCalc -
func MCStatsCalc(slice []float64, sims int) (stats MCStats) {
	variance := stat.Variance(slice, nil)
	stdev := math.Pow(variance, .5)
	mean := stat.Mean(slice, nil)
	// Sample from normal dist to compare to observed IRR array
	norm := make([]float64, len(slice))
	for i := range norm {
		norm[i] = NormalSample(mean, stdev)
	}
	sort.Float64s(norm)
	sort.Float64s(slice)
	lra, lrb := stat.LinearRegression(norm, slice, nil, false)
	stats = MCStats{
		Mean:     mean,
		Variance: variance,
		StDev:    stdev,
		Skew:     stat.Skew(slice, nil),
		Kurtosis: stat.ExKurtosis(slice, nil),
		LRalpha:  lra,
		LRbeta:   lrb,
	}
	histslice := make(plotter.Values, sims)
	histslice = slice
	raw, _ := plotter.NewHist(histslice, 100)
	stats.Hist.Vals = make([]float64, 100)
	stats.Hist.Keys = make([]float64, 100)
	for i := 0; i <= 99; i++ {
		if i >= len(raw.Bins) {
			break
		}
		stats.Hist.Vals[i] = raw.Bins[i].Weight
		stats.Hist.Keys[i] = raw.Bins[i].Max
	}
	stats.Hist.Keys = roundslice(2, 0, stats.Hist.Keys)
	return stats
}

// Ribbon - represents one histogram
type Ribbon struct {
	X          [][]float64 `json:"x"`
	Y          [][]float64 `json:"y"`
	Z          [][]float64 `json:"z"`
	Name       string      `json:"name"`
	Colorscale [][]string  `json:"colorscale"`
	Plottype   string      `json:"type"`
	Showscale  string      `json:"showscale"`
}

// RibbonData -
type RibbonData struct {
}

// RibbonPlot -
func RibbonPlot(matrix [][]float64, duration int, bucketnum int, sims int) (ribbon []Ribbon) {
	ribbon = make([]Ribbon, duration)
	for i, v := range matrix {
		sort.Float64s(v)
		histslice := make(plotter.Values, duration)
		histslice = v
		raw, _ := plotter.NewHist(histslice, bucketnum)
		vals := make([]float64, bucketnum)
		keys := make([]float64, bucketnum)
		tempval := 0.0
		for iii := 0; iii < bucketnum; iii++ {
			if iii >= len(raw.Bins) {
				break
			}
			vals[iii] = (raw.Bins[iii].Weight + tempval) / float64(sims)
			keys[iii] = raw.Bins[iii].Max
			tempval = raw.Bins[iii].Weight + tempval
		}
		keys = roundslice(2, 0, keys)
		// create slice for x axis - used as the index in the ribbon
		y := make([]float64, bucketnum)
		for index := range y {
			y[index] = float64((1 + i))
		}
		ribbon[i] = Ribbon{
			X: CreateTupleArray(vals, false),
			Y: CreateTupleArray(y, true),
			Z: CreateTupleArray(keys, false),
			// Colorscale: [][]string{},
			Plottype:  "surface",
			Showscale: "false",
		}
	}
	return ribbon
}

// BetaSample - give mu and stdev, alpha and beta are calculated, and then used to randomly sample the beta distribution
func BetaSample(mu, stdev float64) float64 {
	if stdev == 0 || mu == 0 {
		return mu
	}
	v := math.Pow(stdev, 2)
	alpha := ((1-mu)/v - 1/mu) * math.Pow(mu, 2)
	beta := ((1-mu)/v - 1/mu) * mu * (1 - mu)
	x := distuv.Beta{
		Alpha: alpha,
		Beta:  beta,
	}
	return x.Rand()
}

func AlphaCheck(mu, sigma float64) (alphaZero float64) {
	var alpha = ((1-mu)/math.Pow(sigma, 2) - 1/mu) * math.Pow(mu, 2)
	alphaZero = sigma
	if alpha <= 0 {
		mtwo := math.Pow(mu, 2)
		mthree := math.Pow(mu, 3)
		x := (mtwo - mthree)
		alphaZero = math.Pow(x/mu, .5) - .00001
		return alphaZero
	}
	return alphaZero
}

// BinomialSample - Skew must be between 0 and 1. a value of .5 indicates no Skew, <.5 is positive Skew and >.5 is negative Skew.
func BinomialSample(expval float64, Skew float64) float64 {
	if Skew == 0 {
		return expval
	}
	n := expval / Skew
	x := distuv.Binomial{
		N: n,
		P: Skew,
	}
	return math.Round(x.Rand())
}

// NormalSample -
func NormalSample(mean float64, stdev float64) float64 {
	if stdev == 0 {
		return mean
	}
	x := distuv.Normal{
		Mu:    mean,
		Sigma: stdev,
	}
	return x.Rand()
}

// LogNormalSample -
func LogNormalSample(mean float64, stdev float64) float64 {
	if stdev == 0 {
		return mean
	}
	x := distuv.LogNormal{
		Mu:    mean,
		Sigma: stdev,
	}
	return x.Rand()
}
