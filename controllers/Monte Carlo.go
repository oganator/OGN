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
	EndCash  []float64
	EndNCF   []float64
	IRR      []float64
	EM       []float64
	YTM      []float64
	Duration []float64
}

// MCResults - used for final display. NOI, MarketValue and NCF have stats per year, IRR is for the hold period
type MCResults struct {
	EndCash  MCStats
	EndNCF   MCStats
	IRR      MCStats
	EM       MCStats
	YTM      MCStats
	Duration MCStats
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
	e.MCSlice = []*Entity{}
	e.MCResultSlice = MCResultSlice{EndCash: []float64{}, EndNCF: []float64{}, IRR: []float64{}, EM: []float64{}, YTM: []float64{}, Duration: []float64{}}
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
			tempentitydata.GLA.Default.Hazard = LogNormalSample(float64(e.GLA.Default.Hazard), float64(e.MCSetup.Hazard))
			//
			temp.UpdateEntity(true, tempentitydata)
			temp.MCResults.EndCash.Mean = temp.COA[temp.SalesDate.Dateint].CashBalance
			temp.MCResults.EndNCF.Mean = temp.COA[temp.SalesDate.Dateint].NetCashFlow
			// fmt.Println(temp.Metrics.IRR.NetLeveredAfterTax)
			//
			e.Mutex.Lock()
			e.MCSlice = append(e.MCSlice, &temp)
			e.MCResultSlice.EndCash = append(e.MCResultSlice.EndCash, temp.COA[temp.SalesDate.Dateint].CashBalance)
			e.MCResultSlice.EndNCF = append(e.MCResultSlice.EndNCF, temp.COA[temp.SalesDate.Dateint].NetCashFlow)
			e.MCResultSlice.IRR = append(e.MCResultSlice.IRR, temp.Metrics.IRR.NetLeveredAfterTax)
			e.MCResultSlice.EM = append(e.MCResultSlice.EM, temp.Metrics.EM.NetLeveredAfterTax)
			e.MCResultSlice.YTM = append(e.MCResultSlice.YTM, temp.Metrics.BondHolder.YTM)
			e.MCResultSlice.Duration = append(e.MCResultSlice.Duration, temp.Metrics.BondHolder.Duration)
			e.Mutex.Unlock()
		}(e, i)
	}
	wg.Wait()
	e.MCResults.EndCash = MCStatsCalc(e.MCResultSlice.EndCash, e.MCSetup.Sims)
	e.MCResults.EndNCF = MCStatsCalc(e.MCResultSlice.EndNCF, e.MCSetup.Sims)
	e.MCResults.IRR = MCStatsCalc(e.MCResultSlice.IRR, e.MCSetup.Sims)
	e.MCResults.EM = MCStatsCalc(e.MCResultSlice.EM, e.MCSetup.Sims)
	e.MCResults.YTM = MCStatsCalc(e.MCResultSlice.YTM, e.MCSetup.Sims)
	e.MCResults.Duration = MCStatsCalc(e.MCResultSlice.Duration, e.MCSetup.Sims)
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
	// sort.Sort(EntitySlice(e.MCSlice))
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
	raw, _ := plotter.NewHist(histslice, 99)
	stats.Hist.Vals = make([]float64, 99)
	stats.Hist.Keys = make([]float64, 99)
	for i := 0; i < 99; i++ {
		if i >= len(raw.Bins)-1 {
			break
		}
		stats.Hist.Vals[i] = raw.Bins[i].Weight
		stats.Hist.Keys[i] = raw.Bins[i].Max
	}
	stats.Hist.Keys = roundslice(2, 0, stats.Hist.Keys)
	return stats
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
