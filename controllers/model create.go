package controllers

import (
	"sync"
)

// PopulateModels - moves data from datastores to Models. Calls CreateEntity for each model.
func PopulateModels() {
	for _, v := range EntityDataStore {
		temp := CreateEntity(*v)
		EntityMap[v.MasterID] = EntityMutex{
			Mutex:       &sync.Mutex{},
			EntityModel: &temp,
		}
		if v.Parent != 0 {
			ModelsList[temp.Name] = temp.MasterID
			EntitiesList[temp.Name] = temp.MasterID
		} else {
			FundsList[temp.Name] = temp.MasterID
			EntitiesList[temp.Name] = temp.MasterID
		}
	}
}

// CreateEntity -
func CreateEntity(v EntityData) (e EntityModel) {
	startdate := Dateadd(Datetype{Month: v.StartMonth, Year: v.StartYear}, 0)
	salesdate := Dateadd(Datetype{Month: v.SalesMonth, Year: v.SalesYear}, 0)
	growth := map[string]HModel{}
	growth["CPI"] = v.CPIGrowth
	growth["ERV"] = v.ERVGrowth
	e = EntityModel{
		Mutex:         &sync.Mutex{},
		MasterID:      v.MasterID,
		Name:          v.Name,
		ChildEntities: map[int]*EntityModel{},
		ChildUnits:    map[int]*Unit{},
		Metrics:       Metrics{},
		ParentID:      v.Parent,
		Parent:        &EntityModel{},
		StartDate:     startdate,
		HoldPeriod:    dateintdiff(salesdate.Dateint, startdate.Dateint) / 12,
		SalesDate:     salesdate,
		EndDate:       Dateadd(Datetype{Month: v.EndMonth, Year: v.EndYear}, 0),
		GrowthInput:   growth,
		Growth:        map[string]map[int]float64{},
		DebtInput:     DebtInput{LTV: v.LTV, InterestRate: v.LoanRate},
		OpEx:          CostInput{Amount: 0, AmountPerTotalArea: 0, AmountPerOccupiedArea: 0, AmountPerVacantArea: 0, PercentOfERV: v.OpExpercent, PercentOfTRI: v.OpExpercent, PercentOfContractRent: 0},
		Fees:          CostInput{Amount: 0, AmountPerTotalArea: 0, AmountPerOccupiedArea: 0, AmountPerVacantArea: 0, PercentOfERV: 0, PercentOfTRI: 0, PercentOfContractRent: 0, PercentOfNAV: 0, PercentOfGAV: v.Fees},
		Capex:         map[int]CostInput{},
		GLA: Unit{
			Probability:     v.GLA.Probability,
			PercentSoldRent: v.percentIncometosell,
			DiscountRate:    v.GLA.DiscountRate,
			RentRevisionERV: v.GLA.RentRevisionERV,
			EXTDuration:     v.GLA.EXTDuration,
			Default: Default{
				Hazard: v.GLA.Default.Hazard,
			},
			RentIncentives: CostInput{
				PercentOfContractRent: v.GLA.RentIncentives.PercentOfContractRent,
				IsCapitalized:         false,
				Duration:              v.GLA.RentIncentives.Duration,
			},
			FitOutCosts: CostInput{
				AmountPerTotalArea: v.GLA.FitOutCosts.AmountPerTotalArea,
			},
			Void: v.GLA.Void,
		},
		MCSetup: MCSetup{
			Sims:        v.Sims,
			ERV:         v.CPISigma,
			CPI:         v.CPISigma,
			YieldShift:  v.YieldShiftSigma,
			Void:        v.VoidSigma,
			Probability: v.ProbabilitySigma,
			OpEx:        v.OpExSigma,
			Hazard:      v.Hazard,
		},
		MCSlice: []*EntityModel{},
		Tax: Tax{
			MinValue:        v.WOZpercent,
			LandValue:       v.Landvalue,
			UsablePeriod:    v.DeprPeriod,
			RETT:            v.RETT,
			CIT:             map[float64]float64{},
			VAT:             v.VAT,
			CarryBackYrs:    v.CarryBackYrs,
			CarryForwardYrs: v.CarryForwardYrs,
			DTA:             map[int]float64{},
		},
		COA:            map[int]FloatCOA{},
		Valuation:      Valuation{EntryYield: v.EntryYield, YieldShift: v.YieldShift, ExitYield: v.ExitYield, AcqFees: map[string]Fee{}, DispFees: map[string]Fee{}, IncomeCapSetup: FloatCOA{TotalERV: 1}, IncomeDeduction: FloatCOA{}},
		TableHeader:    HeaderType{},
		Table:          []TableJSON{},
		Strategy:       v.Strategy,
		BalloonPercent: v.BalloonPercent,
	}
	e.CalculateModel(false, "Internal")
	return e
}

// UpdateEntity -
func (e *EntityModel) UpdateEntity(mc bool, v *EntityData, compute string) {
	startdate := Dateadd(Datetype{Month: v.StartMonth, Year: v.StartYear}, 0)
	salesdate := Dateadd(startdate, v.HoldPeriod*12-1)
	enddate := Dateadd(salesdate, 132)
	growthinput := map[string]HModel{}
	growthinput["CPI"] = v.CPIGrowth
	growthinput["ERV"] = v.ERVGrowth
	parentFinal := EntityModel{}
	parentID := v.Parent
	childunits := make(map[int]*Unit)
	childunitsMC := make(map[int]Unit)
	if compute == "Internal" {
		parentFinal = *EntityMap[v.Parent].EntityModel
	} else if compute == "Azure" {
		parentID = -1
		childunits = e.ChildUnits
		childunitsMC = e.ChildUnitsMC
	}
	// StructPrint("UpdateEntity - ", e.ChildUnits)
	childentitiesmap := map[int]*EntityModel{}
	*e = EntityModel{
		Mutex:         &sync.Mutex{},
		MasterID:      v.MasterID,
		Name:          v.Name,
		ChildEntities: childentitiesmap,
		ChildUnits:    childunits,
		ChildUnitsMC:  childunitsMC,
		Metrics:       Metrics{},
		ParentID:      parentID,
		Parent:        &parentFinal,
		StartDate:     startdate,
		HoldPeriod:    v.HoldPeriod,
		SalesDate:     salesdate,
		EndDate:       enddate,
		GrowthInput:   growthinput,
		Growth:        map[string]map[int]float64{},
		DebtInput:     DebtInput{LTV: v.LTV, InterestRate: v.LoanRate},
		OpEx:          CostInput{Amount: 0, AmountPerTotalArea: 0, AmountPerOccupiedArea: 0, AmountPerVacantArea: 0, PercentOfERV: v.OpExpercent, PercentOfTRI: v.OpExpercent, PercentOfContractRent: 0},
		Fees:          CostInput{Amount: 0, AmountPerTotalArea: 0, AmountPerOccupiedArea: 0, AmountPerVacantArea: 0, PercentOfERV: 0, PercentOfTRI: 0, PercentOfContractRent: 0, PercentOfNAV: 0, PercentOfGAV: v.Fees},
		Capex:         map[int]CostInput{},
		GLA: Unit{
			MasterID:        0,
			Name:            "",
			LeaseStartDate:  startdate,
			LeaseExpiryDate: Datetype{},
			UnitStatus:      "",
			Tenant:          "",
			PassingRent:     0,
			RentSchedule:    RentSchedule{},
			RSStore:         []RentSchedule{},
			Parent:          &EntityModel{},
			Probability:     v.GLA.Probability,
			PercentSoldRent: v.GLA.PercentSoldRent,
			BondIncome:      0,
			BondExpense:     0,
			Default:         Default{Hazard: v.GLA.Default.Hazard, DefaultEnd: Datetype{}},
			RentRevisionERV: v.GLA.RentRevisionERV,
			EXTDuration:     v.GLA.EXTDuration,
			IndexDetails:    IndexDetails{Frequency: v.GLA.IndexDetails.Frequency, Type: v.GLA.IndexDetails.Type, StartMonth: 0, Anniversary: ""},
			RentIncentives: CostInput{
				PercentOfContractRent: v.GLA.RentIncentives.PercentOfContractRent,
				IsCapitalized:         false,
				Duration:              v.GLA.RentIncentives.Duration,
			},
			Void:         v.GLA.Void,
			FitOutCosts:  CostInput{Amount: v.GLA.FitOutCosts.Amount, AmountPerTotalArea: v.GLA.FitOutCosts.AmountPerTotalArea, AmountPerOccupiedArea: v.GLA.FitOutCosts.AmountPerOccupiedArea, AmountPerVacantArea: v.GLA.FitOutCosts.AmountPerVacantArea, PercentOfERV: v.GLA.FitOutCosts.PercentOfERV, PercentOfTRI: v.GLA.FitOutCosts.PercentOfTRI, PercentOfContractRent: v.GLA.FitOutCosts.PercentOfContractRent},
			DiscountRate: v.GLA.DiscountRate,
			ERVArea:      0,
			ERVAmount:    0,
		},
		MCSetup: MCSetup{
			Sims: v.Sims},
		Tax: Tax{
			MinValue:        v.WOZpercent,
			LandValue:       v.Landvalue,
			UsablePeriod:    v.DeprPeriod,
			RETT:            v.RETT,
			CIT:             map[float64]float64{},
			VAT:             v.VAT,
			CarryBackYrs:    v.CarryBackYrs,
			CarryForwardYrs: v.CarryForwardYrs,
			DTA:             map[int]float64{},
		},
		COA:            map[int]FloatCOA{},
		Valuation:      Valuation{EntryYield: v.EntryYield, YieldShift: v.YieldShift, ExitYield: v.ExitYield, AcqFees: map[string]Fee{}, DispFees: map[string]Fee{}, IncomeCapSetup: FloatCOA{TotalERV: 1}, IncomeDeduction: FloatCOA{}},
		TableHeader:    HeaderType{},
		Table:          []TableJSON{},
		Strategy:       v.Strategy,
		BalloonPercent: v.BalloonPercent,
	}
	e.CalculateModel(mc, compute)
	e.PopulateChildEntities()
	// fmt.Println("UpdateEntity IRR: ", e.Metrics.IRR.NetLeveredAfterTax)
}

// CalculateModel - mc == MonteCarlo; if true then table is not made
func (e *EntityModel) CalculateModel(mc bool, compute string) {
	if e.ParentID != 0 {
		e.StartDate.Add(0)
		e.SalesDate.Add(0)
		e.EndDate.Add(0)
		e.Growth = make(map[string]map[int]float64)
		e.GrowthCalc(mc)
		if compute == "Internal" {
			e.PopulateUnits()
		}
		// e.CalculateUnits()
		// fmt.Println("Calculate Model: ", e.Growth)
		e.AssetRentCalc(mc, compute)
		// StructPrint("CalculateModel: ", e.COA[202101])
		e.DirectCapCalc()
		// e.DCFCalc()
		e.Acquisition()
		e.PropertyCFCalc()
		e.Disposal()
		e.SumCOA()
		e.CIT()
		e.SumNCF()
		e.SumCOA()
		if !mc {
			coas := BoolCOA{
				MarketValue:             true,
				TotalERV:                true,
				OccupiedERV:             false,
				VacantERV:               false,
				TopSlice:                false,
				TotalArea:               false,
				OccupiedArea:            false,
				VacantArea:              false,
				PassingRent:             true,
				Indexation:              true,
				TheoreticalRentalIncome: true,
				BPUplift:                true,
				Vacancy:                 true,
				ContractRent:            true,
				RentFree:                false,
				TurnoverRent:            false,
				MallRent:                false,
				ParkingIncome:           false,
				OtherIncome:             false,
				OperatingIncome:         false,
				OperatingExpenses:       true,
				NetOperatingIncome:      true,
				AcqDispProperty:         true,
				AcqDispCosts:            false,
				LoanProceeds:            false,
				InterestExpense:         true,
				LoanBalance:             false,
				Debt:                    false,
				Tax:                     true,
				TaxableIncome:           false,
				TaxableIncomeCarryBack:  false,
				DTA:                     false,
				Depreciation:            false,
				Capex:                   true,
				Fees:                    true,
				NetCashFlow:             true,
				CashBalance:             true,
				BondIncome:              true,
				BondExpense:             true,
			}
			switch e.Strategy {
			case "Standard":
				coas.BPUplift = false
				coas.RentFree = true
				coas.Debt = true
				coas.BondIncome = false
				coas.BondExpense = false
			}
			e.MakeTable(coas, false, true)
		}
		e.MetricsCalc()
		// for _, v := range e.ChildUnits {
		// 	fmt.Println("\n", "Unit: ", &v)
		// 	for _, vv := range v.RSStore {
		// 		fmt.Println("________________________________________________________________________")
		// 		fmt.Println("EXTNumber", vv.EXTNumber)
		// 		fmt.Println("StartDate", vv.StartDate)
		// 		fmt.Println("VacancyEnd", vv.VacancyEnd)
		// 		fmt.Println("VacancyAmount", vv.VacancyAmount)
		// 		fmt.Println("RentIncentivesEndRenew", vv.RentIncentivesEndRenew)
		// 		fmt.Println("RentIncentivesEndRotate", vv.RentIncentivesEndRotate)
		// 		fmt.Println("DefaultDate", vv.DefaultDate)
		// 		fmt.Println("EndDate", vv.EndDate)
		// 		fmt.Println("OriginalEndDate", vv.OriginalEndDate)
		// 		fmt.Println("RenewRent", vv.RenewRent)
		// 		fmt.Println("RotateRent", vv.RotateRent)
		// 		fmt.Println("PassingRent", vv.PassingRent)
		// 		fmt.Println("EndContractRent", vv.EndContractRent)
		// 		fmt.Println("RentRevisionERV", vv.RentRevisionERV)
		// 		fmt.Println("Probability", vv.Probability)
		// 		fmt.Println("ProbabilitySim", vv.ProbabilitySim)
		// 		fmt.Println("RenewIndex", vv.RenewIndex)
		// 		fmt.Println("RotateIndex", vv.RotateIndex)
		// 		fmt.Println("ParentUnit", &vv.ParentUnit)
		// 	}
		// }
		if mc {
			for _, v := range e.ChildUnits {
				*v = Unit{}
			}
		}
	}
	// for date := e.StartDate.Year; date <= e.SalesDate.Year; date++ {
	// 	fmt.Println(date, ": ", e.COA[date].PassingRent)
	// }
}

// PopulateUnits -
func (e *EntityModel) PopulateUnits() {
	for _, v := range UnitStore {
		if v.ParentMasterID == e.MasterID {
			// VACANCY
			expirydate := Datetype{
				Month: v.LeaseEndMonth,
				Year:  v.LeaseEndYear,
			}
			startdate := Datetype{
				Month: v.LeaseStartMonth,
				Year:  v.LeaseEndMonth,
			}
			passingrent := v.PassingRent
			if v.UnitStatus == "Vacant" {
				start := Dateadd(e.StartDate, -1)
				end := Dateadd(e.StartDate, e.GLA.EXTDuration)
				startdate.Month = start.Month
				startdate.Year = start.Year
				expirydate.Month = end.Month
				expirydate.Year = end.Year
				passingrent = v.ERVAmount * v.ERVArea
			}
			// from UnitStore
			temp := Unit{
				// Mutex:           &sync.Mutex{},
				MasterID:        v.MasterID,
				Name:            v.Name,
				LeaseStartDate:  startdate,
				LeaseExpiryDate: expirydate,
				UnitStatus:      v.UnitStatus,
				Tenant:          v.Tenant,
				PassingRent:     passingrent,
				RentSchedule:    RentSchedule{},
				RSStore:         []RentSchedule{},
				Parent:          e,
				Probability:     e.GLA.Probability,
				PercentSoldRent: e.GLA.PercentSoldRent,
				BondIncome:      0,
				BondExpense:     0,
				Default:         e.GLA.Default,
				RentRevisionERV: e.GLA.RentRevisionERV,
				EXTDuration:     e.GLA.EXTDuration,
				IndexDetails:    e.GLA.IndexDetails,
				RentIncentives: CostInput{
					PercentOfContractRent: e.GLA.RentIncentives.PercentOfContractRent,
					IsCapitalized:         false,
					Duration:              e.GLA.RentIncentives.Duration,
				},
				Void:         e.GLA.Void,
				FitOutCosts:  e.GLA.FitOutCosts,
				DiscountRate: e.GLA.DiscountRate,
				ERVArea:      v.ERVArea,
				ERVAmount:    v.ERVAmount,
			}
			temp.LeaseStartDate.Add(0)
			temp.LeaseExpiryDate.Add(0)
			e.ChildUnits[v.MasterID] = &temp
		}
		// fmt.Println("PopulateUnits - ", v.LeaseStartYear)
	}
}

func (e *EntityModel) PopulateChildEntities() {
	for _, v := range EntityDataStore {
		if v.Parent == e.MasterID {
			e.ChildEntities[v.MasterID] = EntityMap[v.MasterID].EntityModel
		}
	}
}
