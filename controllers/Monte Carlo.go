package controllers

import (
	"html/template"
	"math"
	"time"

	//"math/rand"
	"sort"
	"sync"

	"golang.org/x/exp/rand"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot/plotter"
)

// MCSetup -
type MCSetup struct {
	Sims        int     `json:"sims,omitempty"`
	ERV         HModel  `json:"erv,omitempty"`
	CPI         HModel  `json:"cpi,omitempty"`
	YieldShift  float64 `json:"yieldshift,omitempty"`
	Void        float64 `json:"void,omitempty"`
	Probability float64 `json:"probability,omitempty"`
	OpEx        float64 `json:"opex,omitempty"`
	Hazard      float64 `json:"hazard,omitempty"`
}

// MCResultSlice - Fields with 'Fund' at the end are used by the Asset to store data that is consumed by the Fund MC. It is the same data as the other field, but with the arrays reversed
type MCResultSlice struct {
	SimNumber        int         `json:"SimNumber,omitempty"` // used by the Azure function
	EndCash          []float64   `json:"EndCash,omitempty"`
	CashBalance      [][]float64 `json:"CashBalance,omitempty"`     // Month, Sim
	CashBalanceFund  [][]float64 `json:"CashBalanceFund,omitempty"` // Sim, Month
	EndNCF           []float64   `json:"MastEndNCFerID,omitempty"`
	NCF              [][]float64 `json:"NCF,omitempty"`     // Month, Sim
	NCFFund          [][]float64 `json:"NCFFund,omitempty"` // Sim, Month
	EndMarketValue   []float64   `json:"EndMarketValue,omitempty"`
	MarketValue      [][]float64 `json:"MarketValue,omitempty"`     // Month, Sim
	MarketValueFund  [][]float64 `json:"MarketValueFund,omitempty"` // Sim, Month
	IRR              []float64   `json:"IRR,omitempty"`
	EM               []float64   `json:"EM,omitempty"`
	YTM              []float64   `json:"YTM,omitempty"`
	BondExpense      [][]float64 `json:"BondExpense,omitempty"` // Sim, Month
	Duration         []float64   `json:"Duration,omitempty"`
	YTMDUR           []float64   `json:"YTMDUR,omitempty"`
	Void             []float64   `json:"Void,omitempty"`
	Probability      []float64   `json:"Probability,omitempty"`
	NumberOfDefaults []float64   `json:"NumberOfDefaults,omitempty"`
	OpEx             []float64   `json:"OpEx,omitempty"`
	CPI              []float64   `json:"CPI,omitempty"`
	ERV              []float64   `json:"ERV,omitempty"`
	Hazard           []float64   `json:"Hazard,omitempty"`
	YieldShift       []float64   `json:"YieldShift,omitempty"`
}

// MCResults - used for final display. NOI, MarketValue and NCF have stats per year, IRR is for the hold period
type MCResults struct {
	EndCash        MCStats       `json:"EndCash,omitempty"`
	CashBalance    []Ribbon      `json:"CashBalance,omitempty"`
	CashBalanceVaR VaRPercentile `json:"CashBalanceVaR,omitempty"`
	EndNCF         MCStats       `json:"EndNCF,omitempty"`
	NCF            []Ribbon      `json:"NCF,omitempty"`
	NCFVaR         VaRPercentile `json:"NCFVaR,omitempty"`
	EndMarketValue MCStats       `json:"EndMarketValue,omitempty"`
	MarketValue    []Ribbon      `json:"MarketValue,omitempty"`
	MarketValueVaR VaRPercentile `json:"MarketValueVaR,omitempty"`
	IRR            MCStats       `json:"IRR,omitempty"`
	EM             MCStats       `json:"EM,omitempty"`
	YTM            MCStats       `json:"YTM,omitempty"`
	Duration       MCStats       `json:"Duration,omitempty"`
	YTMDUR         MCStats       `json:"YTMDUR,omitempty"`
}

// MCStats -
type MCStats struct {
	Mean     float64 `json:"Mean,omitempty"`
	Variance float64 `json:"Variance,omitempty"`
	StDev    float64 `json:"StDev,omitempty"`
	Skew     float64 `json:"Skew,omitempty"`
	Kurtosis float64 `json:"Kurtosis,omitempty"`
	P1       float64 `json:"P1,omitempty"`
	P5       float64 `json:"P5,omitempty"`
	P10      float64 `json:"P10,omitempty"`
	P25      float64 `json:"P25,omitempty"`
	P50      float64 `json:"P50,omitempty"`
	P75      float64 `json:"P75,omitempty"`
	P90      float64 `json:"P90,omitempty"`
	P95      float64 `json:"P95,omitempty"`
	P99      float64 `json:"P99,omitempty"`
	LRalpha  float64 `json:"LRalpha,omitempty"`
	LRbeta   float64 `json:"LRbeta,omitempty"`
	Hist     Hist    `json:"Hist,omitempty"`
}

// Hist -
type Hist struct {
	Keys []float64 `json:"Keys,omitempty"`
	Vals []float64 `json:"Vals,omitempty"`
}

// MonteCarlo - compute can be either "Internal" or "Azure"
func (e *EntityModel) MonteCarlo(compute string) {
	if compute == "Azure" {
		e.Mutex = &sync.Mutex{}
		e.EntityData.Mutex = &sync.Mutex{}
		for index, key := range e.ChildUnitsMC {
			key.Parent = e
			e.ChildUnitModels[index] = &key
			// e.ChildUnits[index].Mutex = &sync.Mutex{}
		}
	}
	duration := e.MCDataObjectsCreate(1)
	wg := sync.WaitGroup{}
	for i := 1; i <= e.MCSetup.Sims; i++ {
		if e.MCSetup.Sims < 100 && compute == "Internal" {
			continue
		}
		wg.Add(1)
		go func(ee *EntityModel, index int) {
			defer wg.Done()
			temp := CreateShellEntity(ee, compute)
			tempentitydata := &EntityModelData{}
			if compute == "Azure" {
				tempentitydata = &ee.EntityData
			} else if compute == "Internal" {
				tempentitydata = EntityDataStore[e.MasterID]
			}
			tempentitydata.SampleForEntity(&temp)
			temp.MC = true
			temp.UpdateEntity(true, tempentitydata, compute)
			// StructPrint("MonteCarlo - ", temp)
			temp.MCResults.EndCash.Mean = temp.COA[temp.SalesDate.Dateint].CashBalance
			temp.MCResults.EndNCF.Mean = temp.COA[temp.SalesDate.Dateint].NetCashFlow
			temp.MCResults.EndMarketValue.Mean = temp.COA[temp.SalesDate.Dateint].MarketValue
			date := Dateadd(temp.StartDate, -1)
			e.Mutex.Lock()
			e.MCSlice[index-1] = &temp //
			// assign results for every month. ii is the month, index is the simulation
			for ii := 0; ii < duration; ii++ {
				e.MCResultSlice.CashBalance[ii][index-1] = temp.COA[date.Dateint].CashBalance
				e.MCResultSlice.NCF[ii][index-1] = temp.COA[date.Dateint].NetCashFlow
				e.MCResultSlice.MarketValue[ii][index-1] = temp.COA[date.Dateint].MarketValue
				// fmt.Println(ii, " : ", e.MCResultSlice.BondExpense[ii][index-1])
				date.Add(1)
			}
			// e.MCResultSlice.CashBalanceFund[index-1] = ReturnCOAArray(temp.COA, FloatCOA{CashBalance: 1.0}, temp.StartDate, temp.SalesDate)
			// e.MCResultSlice.NCFFund[index-1] = ReturnCOAArray(temp.COA, FloatCOA{NetCashFlow: 1.0}, temp.StartDate, temp.SalesDate)
			// e.MCResultSlice.MarketValueFund[index-1] = ReturnCOAArray(temp.COA, FloatCOA{MarketValue: 1.0}, temp.StartDate, temp.SalesDate)
			if e.Strategy != "Standard" {
				e.MCResultSlice.BondExpense[index-1] = ReturnCOAArray(temp.COA, FloatCOA{BondExpense: 1.0}, Dateadd(temp.StartDate, -1), temp.SalesDate, false)
				e.MCResultSlice.YTM[index-1] = e.MCSlice[index-1].Metrics.BondHolder.YTM
				e.MCResultSlice.Duration[index-1] = e.MCSlice[index-1].Metrics.BondHolder.Duration
				e.MCResultSlice.YTMDUR[index-1] = e.MCSlice[index-1].Metrics.BondHolder.YTMDUR
			}
			//
			e.MCResultSlice.EndCash[index-1] = e.MCSlice[index-1].COA[temp.SalesDate.Dateint].CashBalance
			e.MCResultSlice.EndNCF[index-1] = e.MCSlice[index-1].COA[temp.SalesDate.Dateint].NetCashFlow
			e.MCResultSlice.EndMarketValue[index-1] = e.MCSlice[index-1].COA[temp.SalesDate.Dateint].MarketValue
			irr := e.MCSlice[index-1].Metrics.IRR.NetLeveredAfterTax
			e.MCResultSlice.IRR[index-1] = irr
			em := e.MCSlice[index-1].Metrics.EM.NetLeveredAfterTax
			e.MCResultSlice.EM[index-1] = em
			//
			e.MCResultSlice.Void[index-1] = float64(e.MCSlice[index-1].GLA.Void)
			e.MCResultSlice.Probability[index-1] = e.MCSlice[index-1].GLA.Probability
			e.MCResultSlice.NumberOfDefaults[index-1] = float64(e.MCSlice[index-1].GLA.Default.NumberOfDefaults)
			// e.MCResultSlice.OpEx[index-1] = e.MCSlice[index-1].OpEx.PercentOfTRI
			e.MCResultSlice.CPI[index-1] = e.MCSlice[index-1].Growth["CPI"][e.SalesDate.Dateint]
			e.MCResultSlice.ERV[index-1] = e.MCSlice[index-1].Growth["ERV"][e.SalesDate.Dateint]
			e.MCResultSlice.Hazard[index-1] = e.MCSlice[index-1].GLA.Default.Hazard
			e.MCResultSlice.YieldShift[index-1] = e.MCSlice[index-1].Valuation.YieldShift
			//
			e.Mutex.Unlock()
			// clear unnecessary cash flow and unit data - use for debugging monte carlo issues
			temp.COA = IntFloatCOAMap{}
			e.MCSlice[index-1].COA = IntFloatCOAMap{}
			temp.ChildUnitModels = make(map[int]*UnitModel)
			e.MCSlice[index-1].ChildUnitModels = make(map[int]*UnitModel)
		}(e, i)
	}
	wg.Wait()
	if compute == "Internal" {
		e.MCCalc(duration)
	}
}

func (e *EntityModel) MCCalc(duration int) {
	e.MCResults.EndCash = MCStatsCalc(&e.MCResultSlice.EndCash, e.MCSetup.Sims)
	e.MCResults.CashBalance, e.MCResults.CashBalanceVaR = RibbonPlot(&e.MCResultSlice.CashBalance, duration-2, 100, e.MCSetup.Sims)
	e.MCResults.EndNCF = MCStatsCalc(&e.MCResultSlice.EndNCF, e.MCSetup.Sims)
	e.MCResults.NCF, e.MCResults.NCFVaR = RibbonPlot(&e.MCResultSlice.NCF, duration-2, 100, e.MCSetup.Sims)
	e.MCResults.EndMarketValue = MCStatsCalc(&e.MCResultSlice.EndMarketValue, e.MCSetup.Sims)
	e.MCResults.MarketValue, e.MCResults.MarketValueVaR = RibbonPlot(&e.MCResultSlice.MarketValue, duration-2, 100, e.MCSetup.Sims)

	switch e.Strategy {
	case "Standard":
		e.MCResults.IRR = MCStatsCalc(&e.MCResultSlice.IRR, e.MCSetup.Sims)
		e.MCResults.EM = MCStatsCalc(&e.MCResultSlice.EM, e.MCSetup.Sims)
	case "Balloon", "Pure Discount":
		e.MCResults.IRR = MCStatsCalc(&e.MCResultSlice.IRR, e.MCSetup.Sims)
		e.MCResults.EM = MCStatsCalc(&e.MCResultSlice.EM, e.MCSetup.Sims)
		e.MCResults.YTM = MCStatsCalc(&e.MCResultSlice.YTM, e.MCSetup.Sims)
		e.MCResults.Duration = MCStatsCalc(&e.MCResultSlice.Duration, e.MCSetup.Sims)
		e.MCResults.YTMDUR = MCStatsCalc(&e.MCResultSlice.YTMDUR, e.MCSetup.Sims)
	}
	e.FactorAnalysisCalc(&e.MCResultSlice)
}

// FundMonteCarlo - Random sample from the ChildEntities
func (e *EntityModel) FundMonteCarlo() {
	// make slices
	duration := e.MCDataObjectsCreate(1)
	for sim := 0; sim < e.MCSetup.Sims; sim++ {
		// get CPI
		for _, v := range e.ChildEntityModels {
			if v.MCSetup.Sims == 0 {
				continue
			}
			// select one of the simulations
			samplefloat := rand.Float64()
			simsfloat := float64(e.MCSetup.Sims)
			simratio := float64(v.MCSetup.Sims) / float64(e.MCSetup.Sims)
			sampleint := int(samplefloat * simsfloat * simratio)
			// childduration := dateintdiff(v.SalesDate.Dateint, v.StartDate.Dateint) + 1
			for i := 0; i < duration; i++ {
				e.MCResultSlice.CashBalance[i][sim] = e.MCResultSlice.CashBalance[i][sim] + v.MCResultSlice.CashBalance[i][sampleint]
				e.MCResultSlice.NCF[i][sim] = e.MCResultSlice.NCF[i][sim] + v.MCResultSlice.NCF[i][sampleint]
				e.MCResultSlice.MarketValue[i][sim] = e.MCResultSlice.MarketValue[i][sim] + v.MCResultSlice.MarketValue[i][sampleint]
			}
			e.MCResultSlice.EndCash[sim] = e.MCResultSlice.EndCash[sim] + v.MCResultSlice.EndCash[sampleint]
			e.MCResultSlice.EndNCF[sim] = e.MCResultSlice.EndNCF[sim] + v.MCResultSlice.EndNCF[sampleint]
			e.MCResultSlice.EndMarketValue[sim] = e.MCResultSlice.EndMarketValue[sim] + v.MCResultSlice.EndMarketValue[sampleint]
			// e.MCResultSlice.YTM[sim] = (math.Pow(IRRCalc(v.MCResultSlice.BondExpense[sampleint])+1, 12) - 1) * 100
			// e.MCResultSlice.Duration[sim] = v.Metrics.BondHolder.Duration
			// e.MCResultSlice.YTMDUR[sim] = e.MCResultSlice.YTM[sim] / e.MCResultSlice.Duration[sim]
			if e.Strategy == "Balloon" {
				for i, v := range v.MCResultSlice.BondExpense[sampleint] {
					e.MCResultSlice.BondExpense[sim][i] = e.MCResultSlice.BondExpense[sim][i] + v
				}
			}
		} // end range children
		// create slice for IRR and EM
		irrslice := make([]float64, duration)
		emslice := map[int]FloatCOA{}
		for i, date := 0, Dateadd(e.StartDate, -1); i < duration; i, date = i+1, Dateadd(date, 1) {
			irrslice[i] = e.MCResultSlice.NCF[i][sim]
			emslice[date.Dateint] = FloatCOA{NetCashFlow: e.MCResultSlice.NCF[i][sim]}
		}
		e.MCResultSlice.IRR[sim] = (math.Pow(IRRCalc(irrslice)+1, 12) - 1) * 100
		e.MCResultSlice.EM[sim] = EquityMultipleCalc(e.StartDate, e.SalesDate, emslice)
		if e.Strategy != "Standard" {
			e.MCResultSlice.YTM[sim] = (math.Pow(IRRCalc(e.MCResultSlice.BondExpense[sim])+1, 12) - 1) * 100
			e.MCResultSlice.Duration[sim] = Duration(e.MCResultSlice.BondExpense[sim][1:])
			e.MCResultSlice.YTMDUR[sim] = e.MCResultSlice.YTM[sim] / e.MCResultSlice.Duration[sim]
		}
	} // end range sims
	//
	e.MCResults.EndCash = MCStatsCalc(&e.MCResultSlice.EndCash, e.MCSetup.Sims)
	e.MCResults.CashBalance, e.MCResults.CashBalanceVaR = RibbonPlot(&e.MCResultSlice.CashBalance, duration-2, 100, e.MCSetup.Sims)
	e.MCResults.EndNCF = MCStatsCalc(&e.MCResultSlice.EndNCF, e.MCSetup.Sims)
	e.MCResults.NCF, e.MCResults.NCFVaR = RibbonPlot(&e.MCResultSlice.NCF, duration-2, 100, e.MCSetup.Sims)
	e.MCResults.EndMarketValue = MCStatsCalc(&e.MCResultSlice.EndMarketValue, e.MCSetup.Sims)
	e.MCResults.MarketValue, e.MCResults.MarketValueVaR = RibbonPlot(&e.MCResultSlice.MarketValue, duration-2, 100, e.MCSetup.Sims)
	switch e.Strategy {
	case "Standard":
		e.MCResults.IRR = MCStatsCalc(&e.MCResultSlice.IRR, e.MCSetup.Sims)
		e.MCResults.EM = MCStatsCalc(&e.MCResultSlice.EM, e.MCSetup.Sims)
	case "Balloon", "Pure Discount":
		e.MCResults.IRR = MCStatsCalc(&e.MCResultSlice.IRR, e.MCSetup.Sims)
		e.MCResults.EM = MCStatsCalc(&e.MCResultSlice.EM, e.MCSetup.Sims)
		e.MCResults.YTM = MCStatsCalc(&e.MCResultSlice.YTM, e.MCSetup.Sims)
		e.MCResults.Duration = MCStatsCalc(&e.MCResultSlice.Duration, e.MCSetup.Sims)
		e.MCResults.YTMDUR = MCStatsCalc(&e.MCResultSlice.YTMDUR, e.MCSetup.Sims)
	}
}

func (tempentitydata *EntityModelData) SampleForEntity(e *EntityModel) {
	// rand.Seed(uint64(time.Nanosecond))
	tempentitydata.ERVGrowth.ShortTermRate = NormalSample(e.GrowthInput["ERV"].ShortTermRate, e.MCSetup.ERV.ShortTermRate, 0.0, 10.0)
	tempentitydata.ERVGrowth.ShortTermPeriod = int(NormalSample(float64(e.GrowthInput["ERV"].ShortTermPeriod), float64(e.MCSetup.ERV.ShortTermPeriod), 0.0, 10.0))
	tempentitydata.ERVGrowth.TransitionPeriod = int(NormalSample(float64(e.GrowthInput["ERV"].TransitionPeriod), float64(e.MCSetup.ERV.TransitionPeriod), 0.0, 10.0))
	tempentitydata.ERVGrowth.LongTermRate = NormalSample(e.GrowthInput["ERV"].LongTermRate, e.MCSetup.ERV.LongTermRate, 0.0, 10.0)
	tempentitydata.CPIGrowth.ShortTermRate = NormalSample(e.GrowthInput["CPI"].ShortTermRate, e.MCSetup.CPI.ShortTermRate, 0.0, 10.0)
	tempentitydata.CPIGrowth.ShortTermPeriod = int(NormalSample(float64(e.GrowthInput["CPI"].ShortTermPeriod), float64(e.MCSetup.CPI.ShortTermPeriod), 0.0, 10.0))
	tempentitydata.CPIGrowth.TransitionPeriod = int(NormalSample(float64(e.GrowthInput["CPI"].TransitionPeriod), float64(e.MCSetup.CPI.TransitionPeriod), 0.0, 10.0))
	tempentitydata.CPIGrowth.LongTermRate = NormalSample(e.GrowthInput["CPI"].LongTermRate, e.MCSetup.CPI.LongTermRate, 0.0, 10.0)
	// tempentitydata.OpExpercent = NormalSample(e.OpEx.PercentOfTRI, e.MCSetup.OpEx, 0.0, 100.0)
	tempentitydata.YieldShift = NormalSample(e.Valuation.YieldShift, e.MCSetup.YieldShift, -99.9, 100.0)
	tempentitydata.GLA.Void = int(NormalSample(float64(e.GLA.Void), e.MCSetup.Void, 0.0, 100.0))
	tempentitydata.GLA.Probability = NormalSample(e.GLA.Probability, e.MCSetup.Probability, 0.0, 1.0)
	tempentitydata.GLA.Default.Hazard = NormalSample(float64(e.GLA.Default.Hazard), float64(e.MCSetup.Hazard), 0.0, 10.0)
}

func CopyMatrix(source *[][]float64) [][]float64 {
	copy := make([][]float64, len(*source))
	for i, v := range *source {
		copy[i] = CopyArray(&v)
	}
	return copy
}

func CopyArray(source *[]float64) []float64 {
	copyslice := make([]float64, len(*source))
	copy(copyslice, *source)
	return copyslice
}

// Used in Monte Carlo methods - creates MCResultSlice
func (e *EntityModel) MCDataObjectsCreate(addperiods int) int {
	e.MCSlice = make([]*EntityModel, e.MCSetup.Sims)
	duration := dateintdiff(e.SalesDate.Dateint, e.StartDate.Dateint) + addperiods
	e.MCResultSlice = MCResultSlice{
		EndCash:         make([]float64, e.MCSetup.Sims),
		CashBalance:     make([][]float64, duration),
		CashBalanceFund: make([][]float64, e.MCSetup.Sims),
		EndNCF:          make([]float64, e.MCSetup.Sims),
		NCF:             make([][]float64, duration),
		NCFFund:         make([][]float64, e.MCSetup.Sims),
		EndMarketValue:  make([]float64, e.MCSetup.Sims),
		MarketValue:     make([][]float64, duration),
		MarketValueFund: make([][]float64, e.MCSetup.Sims),
		IRR:             make([]float64, e.MCSetup.Sims),
		EM:              make([]float64, e.MCSetup.Sims),
		YTM:             make([]float64, e.MCSetup.Sims),
		BondExpense:     make([][]float64, e.MCSetup.Sims),
		Duration:        make([]float64, e.MCSetup.Sims),
		YTMDUR:          make([]float64, e.MCSetup.Sims),
		//
		Void:             make([]float64, e.MCSetup.Sims),
		Probability:      make([]float64, e.MCSetup.Sims),
		NumberOfDefaults: make([]float64, e.MCSetup.Sims),
		OpEx:             make([]float64, e.MCSetup.Sims),
		CPI:              make([]float64, e.MCSetup.Sims),
		ERV:              make([]float64, e.MCSetup.Sims),
		Hazard:           make([]float64, e.MCSetup.Sims),
		YieldShift:       make([]float64, e.MCSetup.Sims),
	}
	for i := 0; i < duration; i++ {
		e.MCResultSlice.CashBalance[i] = make([]float64, e.MCSetup.Sims)
		e.MCResultSlice.NCF[i] = make([]float64, e.MCSetup.Sims)
		e.MCResultSlice.MarketValue[i] = make([]float64, e.MCSetup.Sims)
	}
	for i := 0; i < e.MCSetup.Sims; i++ {
		e.MCResultSlice.CashBalanceFund[i] = make([]float64, duration)
		e.MCResultSlice.NCFFund[i] = make([]float64, duration)
		e.MCResultSlice.MarketValueFund[i] = make([]float64, duration)
		e.MCResultSlice.BondExpense[i] = make([]float64, int(float64(duration)*1.5))
	}
	return duration
}

// MCStatsCalc -
func MCStatsCalc(sourceslice *[]float64, sims int) (stats MCStats) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	if sims >= 100 {
		slice := CopyArray(sourceslice)
		variance := stat.Variance(slice, nil)
		stdev := math.Pow(variance, .5)
		mean := stat.Mean(slice, nil)
		// Sample from normal dist to compare to observed IRR array
		norm := make([]float64, len(slice))
		for i := range norm {
			norm[i] = NormalSample(mean, stdev, -9999.9, 1000.0)
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
		// raw := &plotter.Histogram{}
		raw, err := plotter.NewHist(histslice, 100)
		if err != nil {
			recover()
			return
		}
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
	X          [][]float64 `json:"x,omitempty"`
	Y          [][]float64 `json:"y,omitempty"`
	Z          [][]float64 `json:"z,omitempty"`
	Name       string      `json:"name,omitempty"`
	Colorscale [][]string  `json:"colorscale,omitempty"`
	Plottype   string      `json:"plottype,omitempty"`
	Showscale  string      `json:"showscale,omitempty"`
}

// VaRPercentile -
type VaRPercentile struct {
	One        XYFloatSlice `json:"one,omitempty"`
	Five       XYFloatSlice `json:"five,omitempty"`
	Ten        XYFloatSlice `json:"ten,omitempty"`
	TwentyFive XYFloatSlice `json:"twentyfive,omitempty"`
	Fifty      XYFloatSlice `json:"fifty,omitempty"`
}

// XYFloatSlice -
type XYFloatSlice struct {
	X     []float64   `json:"x,omitempty"`
	Y     []float64   `json:"y,omitempty"`
	Name  string      `json:"Name,omitempty"`
	Index template.JS `json:"Index,omitempty"`
}

// RibbonPlot -
func RibbonPlot(matrix *[][]float64, duration int, bucketnum int, sims int) (ribbonslice []Ribbon, varpslice VaRPercentile) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
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
		temp := *matrix
		temp = temp[1 : len(temp)-1]
		newmatrix := CopyMatrix(&temp)
		for i, v := range newmatrix {
			sort.Float64s(v)
			raw := &plotter.Histogram{}
			if _, err := plotter.NewHist(make(plotter.Values, duration), bucketnum); err != nil {
				time.Sleep(10000 * time.Millisecond)
				// fmt.Println("Ribbon Plot Error")
				raw, _ = plotter.NewHist(make(plotter.Values, duration), bucketnum)
			} else {
				raw, _ = plotter.NewHist(make(plotter.Values, duration), bucketnum)
			}
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
