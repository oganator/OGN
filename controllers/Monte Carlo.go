package controllers

import (
	"math"

	//"math/rand"
	"sort"
	"sync"

	"golang.org/x/exp/rand"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot/plotter"
)

// MCSetup -
type MCSetup struct {
	Sims        int
	ERV         HModel
	CPI         HModel
	YieldShift  float64
	Void        float64
	Probability float64
	OpEx        float64
	Hazard      float64
}

// MCResultSlice -
type MCResultSlice struct {
	EndCash        []float64
	CashBalance    [][]float64
	EndNCF         []float64
	NCF            [][]float64
	EndMarketValue []float64
	MarketValue    [][]float64
	IRR            []float64
	EM             []float64
	YTM            []float64
	Duration       []float64
}

// MCResults - used for final display. NOI, MarketValue and NCF have stats per year, IRR is for the hold period
type MCResults struct {
	EndCash        MCStats
	CashBalance    []Ribbon
	CashBalanceVaR VaRPercentile
	EndNCF         MCStats
	NCF            []Ribbon
	NCFVaR         VaRPercentile
	EndMarketValue MCStats
	MarketValue    []Ribbon
	MarketValueVaR VaRPercentile
	IRR            MCStats
	EM             MCStats
	YTM            MCStats
	Duration       MCStats
}

// MCStats -
type MCStats struct {
	Mean     float64
	Variance float64
	StDev    float64
	Skew     float64
	Kurtosis float64
	P1       float64
	P5       float64
	P10      float64
	P25      float64
	P50      float64
	P75      float64
	P90      float64
	P95      float64
	P99      float64
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
	duration := e.MCDataObjectsCreate(-2)
	wg := sync.WaitGroup{}
	for i := 1; i <= e.MCSetup.Sims; i++ {
		wg.Add(1)
		go func(e *Entity, index int) {
			defer wg.Done()
			mu := sync.Mutex{}
			mu.Lock()
			rand.Seed(uint64(int64(index)))
			temp := *e
			tempentitydata := EntityStore[e.MasterID]
			tempentitydata.ERVGrowth.ShortTermRate = NormalSample(e.GrowthInput["ERV"].ShortTermRate, e.MCSetup.ERV.ShortTermRate, 0.0)
			tempentitydata.ERVGrowth.ShortTermPeriod = int(NormalSample(float64(e.GrowthInput["ERV"].ShortTermPeriod), float64(e.MCSetup.ERV.ShortTermPeriod), 0.0))
			tempentitydata.ERVGrowth.TransitionPeriod = int(NormalSample(float64(e.GrowthInput["ERV"].TransitionPeriod), float64(e.MCSetup.ERV.TransitionPeriod), 0.0))
			tempentitydata.ERVGrowth.LongTermRate = NormalSample(e.GrowthInput["ERV"].LongTermRate, e.MCSetup.ERV.LongTermRate, 0.0)
			tempentitydata.CPIGrowth.ShortTermRate = NormalSample(e.GrowthInput["CPI"].ShortTermRate, e.MCSetup.CPI.ShortTermRate, 0.0)
			tempentitydata.CPIGrowth.ShortTermPeriod = int(NormalSample(float64(e.GrowthInput["CPI"].ShortTermPeriod), float64(e.MCSetup.CPI.ShortTermPeriod), 0.0))
			tempentitydata.CPIGrowth.TransitionPeriod = int(NormalSample(float64(e.GrowthInput["CPI"].TransitionPeriod), float64(e.MCSetup.CPI.TransitionPeriod), 0.0))
			tempentitydata.CPIGrowth.LongTermRate = NormalSample(e.GrowthInput["CPI"].LongTermRate, e.MCSetup.CPI.LongTermRate, 0.0)
			tempentitydata.OpExpercent = NormalSample(e.OpEx.PercentOfTRI, e.MCSetup.OpEx, 0.0)
			tempentitydata.YieldShift = NormalSample(e.Valuation.YieldShift, e.MCSetup.YieldShift, -99.9)
			tempentitydata.GLA.Void = int(NormalSample(float64(e.GLA.Void), e.MCSetup.Void, 0.0))
			tempentitydata.GLA.Probability = NormalSample(e.GLA.Probability, e.MCSetup.Probability, 0.0)
			tempentitydata.GLA.Default.Hazard = NormalSample(float64(e.GLA.Default.Hazard), float64(e.MCSetup.Hazard), 0.0)
			mu.Unlock()
			//
			temp.MC = true
			temp.UpdateEntity(true, tempentitydata)
			temp.MCResults.EndCash.Mean = temp.COA[temp.SalesDate.Dateint].CashBalance
			temp.MCResults.EndNCF.Mean = temp.COA[temp.SalesDate.Dateint].NetCashFlow
			temp.MCResults.EndMarketValue.Mean = temp.COA[temp.SalesDate.Dateint].MarketValue
			date := Dateadd(temp.StartDate, 1)
			// mutex lock
			e.Mutex.Lock()
			e.MCSlice[index-1] = &temp
			// assign results for every month. ii is the month, index is the simulation
			for ii := 0; ii < duration; ii++ {
				e.MCResultSlice.CashBalance[ii][index-1] = temp.COA[date.Dateint].CashBalance
				e.MCResultSlice.NCF[ii][index-1] = temp.COA[date.Dateint].NetCashFlow
				e.MCResultSlice.MarketValue[ii][index-1] = temp.COA[date.Dateint].MarketValue
				date.Add(1)
			}
			e.MCResultSlice.EndCash[index-1] = temp.COA[temp.SalesDate.Dateint].CashBalance
			e.MCResultSlice.EndNCF[index-1] = temp.COA[temp.SalesDate.Dateint].NetCashFlow
			e.MCResultSlice.EndMarketValue[index-1] = temp.COA[temp.SalesDate.Dateint].MarketValue
			e.MCResultSlice.IRR[index-1] = temp.Metrics.IRR.NetLeveredAfterTax
			e.MCResultSlice.EM[index-1] = temp.Metrics.EM.NetLeveredAfterTax
			e.MCResultSlice.YTM[index-1] = temp.Metrics.BondHolder.YTM
			e.MCResultSlice.Duration[index-1] = temp.Metrics.BondHolder.Duration
			e.Mutex.Unlock()
		}(e, i)
	}
	wg.Wait()
	e.MCResults.EndCash = MCStatsCalc(e.MCResultSlice.EndCash, e.MCSetup.Sims)
	e.MCResults.CashBalance, e.MCResults.CashBalanceVaR = RibbonPlot(e.MCResultSlice.CashBalance, duration, 100, e.MCSetup.Sims)

	e.MCResults.EndNCF = MCStatsCalc(e.MCResultSlice.EndNCF, e.MCSetup.Sims)
	e.MCResults.NCF, e.MCResults.NCFVaR = RibbonPlot(e.MCResultSlice.NCF, duration, 100, e.MCSetup.Sims)

	e.MCResults.EndMarketValue = MCStatsCalc(e.MCResultSlice.EndMarketValue, e.MCSetup.Sims)
	e.MCResults.MarketValue, e.MCResults.MarketValueVaR = RibbonPlot(e.MCResultSlice.MarketValue, duration, 100, e.MCSetup.Sims)

	switch e.Strategy {
	case "Standard":
		e.MCResults.IRR = MCStatsCalc(e.MCResultSlice.IRR, e.MCSetup.Sims)
		e.MCResults.EM = MCStatsCalc(e.MCResultSlice.EM, e.MCSetup.Sims)
	case "Balloon", "Pure Discount":
		e.MCResults.YTM = MCStatsCalc(e.MCResultSlice.YTM, e.MCSetup.Sims)
		e.MCResults.Duration = MCStatsCalc(e.MCResultSlice.Duration, e.MCSetup.Sims)
	}
}

// Used in Monte Carlo methods - creates MCResultSlice
func (e *Entity) MCDataObjectsCreate(addperiods int) int {
	e.MCSlice = make([]*Entity, e.MCSetup.Sims)
	duration := dateintdiff(e.SalesDate.Dateint, e.StartDate.Dateint) + addperiods
	e.MCResultSlice = MCResultSlice{
		EndCash:        make([]float64, e.MCSetup.Sims),
		CashBalance:    make([][]float64, duration),
		EndNCF:         make([]float64, e.MCSetup.Sims),
		NCF:            make([][]float64, duration),
		EndMarketValue: make([]float64, e.MCSetup.Sims),
		MarketValue:    make([][]float64, duration),
		IRR:            make([]float64, e.MCSetup.Sims),
		EM:             make([]float64, e.MCSetup.Sims),
		YTM:            make([]float64, e.MCSetup.Sims),
		Duration:       make([]float64, e.MCSetup.Sims),
	}
	for i := 0; i < duration; i++ {
		e.MCResultSlice.CashBalance[i] = make([]float64, e.MCSetup.Sims)
		e.MCResultSlice.NCF[i] = make([]float64, e.MCSetup.Sims)
		e.MCResultSlice.MarketValue[i] = make([]float64, e.MCSetup.Sims)
	}
	return duration
}

// FundMonteCarlo - Random sample from the ChildEntities
func (e *Entity) FundMonteCarlo() {
	// make slices
	duration := e.MCDataObjectsCreate(1)
	for sim := 0; sim < e.MCSetup.Sims; sim++ {
		// get CPI
		for _, v := range e.ChildEntities {
			// select one of the simulations
			samplefloat := rand.Float64()
			simsfloat := float64(e.MCSetup.Sims)
			simratio := float64(v.MCSetup.Sims) / float64(e.MCSetup.Sims)
			sampleint := int(samplefloat * simsfloat * simratio)
			e.MCResultSlice.CashBalance[0][sim] = e.MCResultSlice.CashBalance[0][sim] + v.MCSlice[sampleint].COA[Dateadd(e.StartDate, -1).Dateint].CashBalance
			e.MCResultSlice.NCF[0][sim] = e.MCResultSlice.NCF[0][sim] + v.MCSlice[sampleint].COA[Dateadd(e.StartDate, -1).Dateint].NetCashFlow
			e.MCResultSlice.MarketValue[0][sim] = e.MCResultSlice.MarketValue[0][sim] + v.MCSlice[sampleint].COA[Dateadd(e.StartDate, -1).Dateint].MarketValue
			for i, date := 1, e.StartDate; i < duration; i, date = i+1, Dateadd(date, 1) {
				e.MCResultSlice.CashBalance[i-1][sim] = e.MCResultSlice.CashBalance[i][sim] + v.MCSlice[sampleint].COA[date.Dateint].NetCashFlow
				e.MCResultSlice.NCF[i][sim] = e.MCResultSlice.NCF[i][sim] + v.MCSlice[sampleint].COA[date.Dateint].NetCashFlow
				e.MCResultSlice.MarketValue[i][sim] = e.MCResultSlice.MarketValue[i][sim] + v.MCSlice[sampleint].COA[date.Dateint].MarketValue
			}
			e.MCResultSlice.EndCash[sim] = e.MCResultSlice.EndCash[sim] + v.MCSlice[sampleint].COA[v.SalesDate.Dateint].CashBalance
			e.MCResultSlice.EndNCF[sim] = e.MCResultSlice.EndNCF[sim] + v.MCSlice[sampleint].COA[v.SalesDate.Dateint].NetCashFlow
			e.MCResultSlice.EndMarketValue[sim] = e.MCResultSlice.EndMarketValue[sim] + v.MCSlice[sampleint].COA[v.SalesDate.Dateint].MarketValue
		}
		// create slice for IRR
		irrslice := make([]float64, duration)
		emslice := map[int]FloatCOA{}
		for i, date := 0, Dateadd(e.StartDate, -1); i < duration; i, date = i+1, Dateadd(date, 1) {
			irrslice[i] = e.MCResultSlice.NCF[i][sim]
			emslice[date.Dateint] = FloatCOA{NetCashFlow: e.MCResultSlice.NCF[i][sim]}
		}
		e.MCResultSlice.IRR[sim] = (math.Pow(IRRCalc(irrslice)+1, 12) - 1) * 100
		e.MCResultSlice.EM[sim] = EquityMultipleCalc(e.StartDate, e.SalesDate, emslice)

		// fmt.Println(sim, e.MCResultSlice.NCF[sim])
		// fmt.Println(e.MCResultSlice.IRR[sim])
		// fmt.Println("irr: ", math.Pow(IRRCalc(e.MCResultSlice.NCF[sim])+1, 12)-1) //v.MCSlice[sampleint].Metrics.IRR.NetLeveredAfterTax

		// e.MCResultSlice.EM[sim] = v.MCSlice[sampleint].Metrics.EM.NetLeveredAfterTax
		// e.MCResultSlice.YTM[sim] = v.MCSlice[sampleint].Metrics.BondHolder.YTM
		// e.MCResultSlice.Duration[sim] = v.MCSlice[sampleint].Metrics.BondHolder.Duration
	}
	// fmt.Println(e.MCResultSlice.NCF[0])
	e.MCResults.EndCash = MCStatsCalc(e.MCResultSlice.EndCash, e.MCSetup.Sims)
	e.MCResults.CashBalance, e.MCResults.CashBalanceVaR = RibbonPlot(e.MCResultSlice.CashBalance, duration-0, 100, e.MCSetup.Sims)

	e.MCResults.EndNCF = MCStatsCalc(e.MCResultSlice.EndNCF, e.MCSetup.Sims)
	e.MCResults.NCF, e.MCResults.NCFVaR = RibbonPlot(e.MCResultSlice.NCF, duration-0, 100, e.MCSetup.Sims)

	e.MCResults.EndMarketValue = MCStatsCalc(e.MCResultSlice.EndMarketValue, e.MCSetup.Sims)
	e.MCResults.MarketValue, e.MCResults.MarketValueVaR = RibbonPlot(e.MCResultSlice.MarketValue, duration-0, 100, e.MCSetup.Sims)

	e.MCResults.IRR = MCStatsCalc(e.MCResultSlice.IRR, e.MCSetup.Sims)
	e.MCResults.EM = MCStatsCalc(e.MCResultSlice.EM, e.MCSetup.Sims)
}

// MCStatsCalc -
func MCStatsCalc(slice []float64, sims int) (stats MCStats) {
	if sims >= 100 {
		variance := stat.Variance(slice, nil)
		stdev := math.Pow(variance, .5)
		mean := stat.Mean(slice, nil)
		// Sample from normal dist to compare to observed IRR array
		norm := make([]float64, len(slice))
		for i := range norm {
			norm[i] = NormalSample(mean, stdev, -999.9)
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
			P1:       slice[int(float64(sims)*.01)],
			P5:       slice[int(float64(sims)*.05)],
			P10:      slice[int(float64(sims)*.10)],
			P25:      slice[int(float64(sims)*.25)],
			P50:      slice[int(float64(sims)*.50)],
			P75:      slice[int(float64(sims)*.75)],
			P90:      slice[int(float64(sims)*.90)],
			P95:      slice[int(float64(sims)*.95)],
			P99:      slice[int(float64(sims)*.99)],
			LRalpha:  lra,
			LRbeta:   lrb,
			Hist:     Hist{},
		}
		histslice := make(plotter.Values, sims)
		histslice = slice
		raw, _ := plotter.NewHist(histslice, 100)
		stats.Hist.Vals = make([]float64, 100)
		stats.Hist.Keys = make([]float64, 100)
		for i := 1; i <= 100; i++ {
			if i >= len(raw.Bins) {
				break
			}
			stats.Hist.Vals[i] = raw.Bins[i].Weight
			stats.Hist.Keys[i] = raw.Bins[i].Max
		}
		stats.Hist.Keys = roundslice(2, 0, stats.Hist.Keys)
		return stats
	}
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

// VaRPercentile -
type VaRPercentile struct {
	One        XYFloatSlice
	Five       XYFloatSlice
	Ten        XYFloatSlice
	TwentyFive XYFloatSlice
	Fifty      XYFloatSlice
}

// XYFloatSlice -
type XYFloatSlice struct {
	X []float64
	Y []float64
}

// RibbonPlot -
func RibbonPlot(matrix [][]float64, duration int, bucketnum int, sims int) (ribbonslice []Ribbon, varpslice VaRPercentile) {
	if sims >= 100 {
		ribbonslice = make([]Ribbon, duration)
		varpslice = VaRPercentile{
			One: XYFloatSlice{
				X: make([]float64, duration),
				Y: make([]float64, duration),
			},
			Five: XYFloatSlice{
				X: make([]float64, duration),
				Y: make([]float64, duration),
			},
			Ten: XYFloatSlice{
				X: make([]float64, duration),
				Y: make([]float64, duration),
			},
			TwentyFive: XYFloatSlice{
				X: make([]float64, duration),
				Y: make([]float64, duration),
			},
			Fifty: XYFloatSlice{
				X: make([]float64, duration),
				Y: make([]float64, duration),
			},
		}
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
			ribbonslice[i] = Ribbon{
				X:         CreateTupleArray(vals, false),
				Y:         CreateTupleArray(y, true),
				Z:         CreateTupleArray(keys, false),
				Plottype:  "surface",
				Showscale: "false",
			}
			varpslice.One.X[i] = y[0]
			one := int(float64(sims) * .01)
			varpslice.One.Y[i] = v[one-1]
			varpslice.Five.X[i] = y[4]
			five := int(float64(sims) * .05)
			varpslice.Five.Y[i] = v[five-1]
			varpslice.Ten.X[i] = y[9]
			ten := int(float64(sims) * .1)
			varpslice.Ten.Y[i] = v[ten-1]
			varpslice.TwentyFive.X[i] = y[24]
			twentyfive := int(float64(sims) * .25)
			varpslice.TwentyFive.Y[i] = v[twentyfive-1]
			varpslice.Fifty.X[i] = y[49]
			fifty := int(float64(sims) * .5)
			varpslice.Fifty.Y[i] = v[fifty-1]
		}
		return ribbonslice, varpslice
	}
	return ribbonslice, varpslice
}
