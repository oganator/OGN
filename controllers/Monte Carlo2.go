package controllers

import (
	"sync"
)

// returns a new entity based on the input e. Removes ChildEntities, Metrics, Growth, MCResults/slice, table
func CreateShellEntity2(e *EntityModel, compute string, index int) (temp *EntityModel) {
	temp = &EntityModel{
		Mutex:             &sync.Mutex{},
		MasterID:          index,
		Entity:            e.Entity,
		EntityID:          e.EntityID,
		Name:              e.Name,
		Version:           "",
		ChildEntityModels: map[int]*EntityModel{},
		ChildUnitModels:   map[int]*UnitModel{},
		ChildUnitsMC:      map[int]UnitModel{},
		Metrics:           Metrics[float64]{},
		ParentID:          e.ParentID,
		Parent:            e.Parent,
		StartDate:         Dateadd(e.StartDate, 0),
		HoldPeriod:        e.HoldPeriod,
		SalesDate:         Dateadd(e.SalesDate, 0),
		EndDate:           Dateadd(e.EndDate, 0),
		GrowthInput:       e.GrowthInput,
		Growth:            map[string]map[int]float64{},
		DebtInput:         e.DebtInput,
		CostInput:         e.CostInput,
		GLA:               e.GLA,
		MC:                true,
		MCSetup:           e.MCSetup,
		MCSlice:           []*EntityModel{},
		MCResultSlice:     MCResultSlice{},
		MCResults:         MCResults{},
		FactorAnalysis:    []FactorIndependant{},
		Tax:               e.Tax,
		COA:               map[int]FloatCOA{},
		Valuation:         e.Valuation,
		TableHeader:       HeaderType{},
		Table:             []TableJSON{},
		RetoolTable:       []interface{}{},
		Strategy:          e.Strategy,
		UOM:               "",
		BalloonPercent:    e.BalloonPercent,
	}
	childunitsMC := make(map[int]UnitModel, len(e.ChildUnitModels))
	childunits := make(map[int]*UnitModel, len(e.ChildUnitModels))
	// Copy the units from the base EntityModel into the simulation, first into the MC field, then into the regular field
	for i, v := range e.ChildUnitModels {
		tempUnit := UnitModel{
			MasterID:        v.MasterID,
			Name:            v.Name,
			LeaseStartDate:  v.LeaseStartDate,
			LeaseExpiryDate: v.LeaseExpiryDate,
			UnitStatus:      v.UnitStatus,
			Tenant:          v.Tenant,
			PassingRent:     v.PassingRent,
			RentSchedule:    RentSchedule{},
			RSStore:         []RentSchedule{},
			Parent:          temp,
			Probability:     v.Probability,
			PercentSoldRent: v.PercentSoldRent,
			DiscountRate:    v.DiscountRate,
			BondProceeds:    v.BondProceeds,
			BondIncome:      v.BondIncome,
			BondIndex:       Indexation{},
			BondExpense:     v.BondExpense,
			Default:         Default{Hazard: v.Default.Hazard},
			RentRevisionERV: v.RentRevisionERV,
			EXTDuration:     v.EXTDuration,
			IndexDetails:    IndexDetails{},
			Void:            v.Void,
			ERVArea:         v.ERVArea,
			ERVAmount:       v.ERVAmount,
			COA:             map[int]FloatCOA{},
			CostInput:       v.CostInput,
			MCSetup:         v.MCSetup,
		}
		childunits[i] = &tempUnit
		childunitsMC[i] = tempUnit
	}
	temp.ChildUnitModels = childunits
	temp.ChildUnitsMC = childunitsMC
	return temp
}

// EntitySample - Randomizes the values of the EntityModel. Uses the mean values that are already in the fields, and samples based on a sigma found in MCSetup
func (e *EntityModel) EntitySample(source *EntityModel) {
	cpi := HModel{}
	erv := HModel{}
	erv.ShortTermRate = NormalSample(source.GrowthInput["ERV"].ShortTermRate, source.MCSetup.ERV.ShortTermRate, 0.0, 10.0)
	erv.ShortTermPeriod = int(NormalSample(float64(source.GrowthInput["ERV"].ShortTermPeriod), float64(source.MCSetup.ERV.ShortTermPeriod), 0.0, 10.0))
	erv.TransitionPeriod = int(NormalSample(float64(source.GrowthInput["ERV"].TransitionPeriod), float64(source.MCSetup.ERV.TransitionPeriod), 0.0, 10.0))
	erv.LongTermRate = NormalSample(source.GrowthInput["ERV"].LongTermRate, source.MCSetup.ERV.LongTermRate, 0.0, 10.0)
	cpi.ShortTermRate = NormalSample(source.GrowthInput["CPI"].ShortTermRate, source.MCSetup.CPI.ShortTermRate, 0.0, 10.0)
	cpi.ShortTermPeriod = int(NormalSample(float64(source.GrowthInput["CPI"].ShortTermPeriod), float64(source.MCSetup.CPI.ShortTermPeriod), 0.0, 10.0))
	cpi.TransitionPeriod = int(NormalSample(float64(source.GrowthInput["CPI"].TransitionPeriod), float64(source.MCSetup.CPI.TransitionPeriod), 0.0, 10.0))
	cpi.LongTermRate = NormalSample(source.GrowthInput["CPI"].LongTermRate, source.MCSetup.CPI.LongTermRate, 0.0, 10.0)
	// source.OpEx.PercentOfTRI = NormalSample(source.OpEx.PercentOfTRI, source.MCSetup.OpEx, 0.0, 100.0)
	e.Valuation.YieldShift = NormalSample(source.Valuation.YieldShift, source.MCSetup.YieldShift, -99.9, 100.0)
	e.GLA.Void = int(NormalSample(float64(source.GLA.Void), source.MCSetup.Void, 0.0, 100.0))
	e.GLA.Probability = NormalSample(source.GLA.Probability, source.MCSetup.Probability, 0.0, 1.0)
	e.GLA.Default.Hazard = NormalSample(float64(source.GLA.Default.Hazard), float64(source.MCSetup.Hazard), 0.0, 10.0)
	e.GrowthInput["CPI"] = cpi
	e.GrowthInput["ERV"] = erv
}

// Monte Carlo simulation based on new data structures. Reduced memory consumption by reusing EntityModels
// func (e *EntityModel) MonteCarlo2(compute string) {
// 	duration := e.MCDataObjectsCreate(1)
// 	wg := sync.WaitGroup{}
// 	ch := make(chan *EntityModel)
// 	go e.MCListener(ch, duration)
// 	// outerLoop - sims^(1/2)
// 	for i := 1; i <= int(math.Sqrt(float64(e.MCSetup.Sims))); i++ {
// 		tempEM := CreateShellEntity2(e, compute)
// 		wg.Add(1)
// 		go func(ee *EntityModel) {
// 			defer wg.Done()
// 			// innerLoop - sims^(1/2)
// 			for ii := 1; ii <= int(math.Sqrt(float64(e.MCSetup.Sims))); ii++ {
// 				ee.EntitySample()
// 				ee.EntityModelCalc(true, compute)
// 				ch <- ee
// 			}
// 		}(&tempEM)
// 		wg.Wait()
// 	}
// 	close(ch)
// 	e.MCCalc(duration)
// }

// Monte Carlo simulation based on new data structures.
func (e *EntityModel) MonteCarlo3(compute string) {
	duration := e.MCDataObjectsCreate(1)
	wg := sync.WaitGroup{}
	ch := make(chan EntityModel)
	go e.MCListener(ch, duration)
	for i := 1; i <= e.MCSetup.Sims; i++ {
		wg.Add(1)
		go func(ee *EntityModel, ii int) {
			defer wg.Done()
			tempEM := CreateShellEntity2(ee, compute, ii)
			// tempEM.EntitySample()
			tempEM.EntityModelCalc(true, compute)
			ch <- *tempEM
		}(e, i)
		wg.Wait()
	}
	close(ch)
	e.MCCalc(duration)
}

// MCListener - receives pointers to completed EntityModel simulations, and puts the data into the MCResults
func (e *EntityModel) MCListener(ch chan EntityModel, duration int) {
	index := 1
	for temp := range ch {
		temp.MCResults.EndCash.Mean = temp.COA[temp.SalesDate.Dateint].CashBalance
		temp.MCResults.EndNCF.Mean = temp.COA[temp.SalesDate.Dateint].NetCashFlow
		temp.MCResults.EndMarketValue.Mean = temp.COA[temp.SalesDate.Dateint].MarketValue
		date := Dateadd(temp.StartDate, -1)
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
		// e.MCResultSlice.CPI[index-1] = e.MCSlice[index-1].Growth["CPI"][e.SalesDate.Dateint]
		// e.MCResultSlice.ERV[index-1] = e.MCSlice[index-1].Growth["ERV"][e.SalesDate.Dateint]
		e.MCResultSlice.Hazard[index-1] = e.MCSlice[index-1].GLA.Default.Hazard
		e.MCResultSlice.YieldShift[index-1] = e.MCSlice[index-1].Valuation.YieldShift
		//
		// clear unnecessary cash flow and unit data - use for debugging monte carlo issues
		temp.COA = IntFloatCOAMap{}
		e.MCSlice[index-1].COA = IntFloatCOAMap{}
		temp.ChildUnitModels = make(map[int]*UnitModel)
		e.Mutex.Lock()
		e.MCSlice[index-1].ChildUnitModels = make(map[int]*UnitModel)
		e.Mutex.Unlock()
		index++
	}
}

// Monte Carlo simulation based on new data structures.
func (e *EntityModel) MonteCarlo2(compute string) {
	duration := e.MCDataObjectsCreate(0)
	wg := sync.WaitGroup{}
	for i := 1; i <= e.MCSetup.Sims; i++ {
		wg.Add(1)
		go func(ee *EntityModel, index int) {
			defer wg.Done()
			temp := CreateShellEntity2(ee, compute, index)
			temp.EntitySample(ee)
			temp.EntityModelCalc(true, compute)
			temp.MCResults.EndCash.Mean = temp.COA[temp.SalesDate.Dateint].CashBalance
			temp.MCResults.EndNCF.Mean = temp.COA[temp.SalesDate.Dateint].NetCashFlow
			temp.MCResults.EndMarketValue.Mean = temp.COA[temp.SalesDate.Dateint].MarketValue
			date := Dateadd(temp.StartDate, -1)
			ee.Mutex.Lock()
			ee.MCSlice[index-1] = temp //
			// assign results for every month. ii is the month, index is the simulation
			for ii := 0; ii < duration; ii++ {
				ee.MCResultSlice.CashBalance[ii][index-1] = temp.COA[date.Dateint].CashBalance
				ee.MCResultSlice.NCF[ii][index-1] = temp.COA[date.Dateint].NetCashFlow
				ee.MCResultSlice.MarketValue[ii][index-1] = temp.COA[date.Dateint].MarketValue
				// fmt.Println(ii, " : ", e.MCResultSlice.BondExpense[ii][index-1])
				date.Add(1)
			}
			// e.MCResultSlice.CashBalanceFund[index-1] = ReturnCOAArray(temp.COA, FloatCOA{CashBalance: 1.0}, temp.StartDate, temp.SalesDate)
			// e.MCResultSlice.NCFFund[index-1] = ReturnCOAArray(temp.COA, FloatCOA{NetCashFlow: 1.0}, temp.StartDate, temp.SalesDate)
			// e.MCResultSlice.MarketValueFund[index-1] = ReturnCOAArray(temp.COA, FloatCOA{MarketValue: 1.0}, temp.StartDate, temp.SalesDate)
			if ee.Strategy != "Standard" {
				ee.MCResultSlice.BondExpense[index-1] = ReturnCOAArray(temp.COA, FloatCOA{BondExpense: 1.0}, Dateadd(temp.StartDate, -1), temp.SalesDate, false)
				ee.MCResultSlice.YTM[index-1] = ee.MCSlice[index-1].Metrics.BondHolder.YTM
				ee.MCResultSlice.Duration[index-1] = ee.MCSlice[index-1].Metrics.BondHolder.Duration
				ee.MCResultSlice.YTMDUR[index-1] = ee.MCSlice[index-1].Metrics.BondHolder.YTMDUR
			}
			//
			ee.MCResultSlice.EndCash[index-1] = ee.MCSlice[index-1].COA[temp.SalesDate.Dateint].CashBalance
			ee.MCResultSlice.EndNCF[index-1] = ee.MCSlice[index-1].COA[temp.SalesDate.Dateint].NetCashFlow
			ee.MCResultSlice.EndMarketValue[index-1] = ee.MCSlice[index-1].COA[temp.SalesDate.Dateint].MarketValue
			irr := ee.MCSlice[index-1].Metrics.IRR.NetLeveredAfterTax
			ee.MCResultSlice.IRR[index-1] = irr
			em := ee.MCSlice[index-1].Metrics.EM.NetLeveredAfterTax
			ee.MCResultSlice.EM[index-1] = em
			//
			ee.MCResultSlice.Void[index-1] = float64(ee.MCSlice[index-1].GLA.Void)
			ee.MCResultSlice.Probability[index-1] = ee.MCSlice[index-1].GLA.Probability
			ee.MCResultSlice.NumberOfDefaults[index-1] = float64(ee.MCSlice[index-1].GLA.Default.NumberOfDefaults)
			// e.MCResultSlice.OpEx[index-1] = e.MCSlice[index-1].OpEx.PercentOfTRI
			// e.MCResultSlice.CPI[index-1] = e.MCSlice[index-1].Growth["CPI"][e.SalesDate.Dateint]
			// e.MCResultSlice.ERV[index-1] = e.MCSlice[index-1].Growth["ERV"][e.SalesDate.Dateint]
			ee.MCResultSlice.Hazard[index-1] = ee.MCSlice[index-1].GLA.Default.Hazard
			ee.MCResultSlice.YieldShift[index-1] = ee.MCSlice[index-1].Valuation.YieldShift
			//
			// clear unnecessary cash flow and unit data - use for debugging monte carlo issues
			// temp.COA = IntFloatCOAMap{}
			// e.MCSlice[index-1].COA = IntFloatCOAMap{}
			// temp.ChildUnitModels = make(map[int]*UnitModel)
			ee.MCSlice[index-1].ChildUnitModels = make(map[int]*UnitModel)
			ee.Mutex.Unlock()
		}(e, i)
		wg.Wait()
	}
	e.MCCalc(duration)
}
