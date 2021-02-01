package controllers

// PopulateModels - moves data from datastores to Models
func PopulateModels() {
	for _, v := range EntityStore {
		temp := CreateEntity(*v)
		Models[v.MasterID] = &temp
		ModelsList[temp.Name] = temp.MasterID
	}
}

// CreateEntity -
func CreateEntity(v EntityData) (e Entity) {
	startdate := Dateadd(Datetype{Month: v.StartMonth, Year: v.StartYear}, 0)
	salesdate := Dateadd(Datetype{Month: v.SalesMonth, Year: v.SalesYear}, 0)
	growth := map[string]float64{}
	growth["CPI"] = v.CPIGrowth
	growth["ERV"] = v.ERVGrowth
	e = Entity{
		MasterID:       v.MasterID,
		Name:           v.Name,
		ChildEntities:  map[int]Entity{},
		ChildUnits:     map[int]*Unit{},
		Metrics:        Metrics{},
		Parent:         &Entity{},
		StartDate:      startdate,
		HoldPeriod:     dateintdiff(salesdate.Dateint, startdate.Dateint) / 12,
		SalesDate:      salesdate,
		EndDate:        Dateadd(Datetype{Month: v.EndMonth, Year: v.EndYear}, 0),
		GrowthInput:    growth,
		Growth:         map[string]map[int]float64{},
		DebtInput:      DebtInput{LTV: v.LTV, InterestRate: v.LoanRate},
		OpEx:           CostInput{Amount: 0, AmountPerTotalArea: 0, AmountPerOccupiedArea: 0, AmountPerVacantArea: 0, PercentOfERV: v.OpExpercent, PercentOfTRI: v.OpExpercent, PercentOfContractRent: 0},
		Fees:           CostInput{Amount: 0, AmountPerTotalArea: 0, AmountPerOccupiedArea: 0, AmountPerVacantArea: 0, PercentOfERV: 0, PercentOfTRI: 0, PercentOfContractRent: 0, PercentOfNAV: 0, PercentOfGAV: v.Fees},
		Capex:          map[int]CostInput{},
		GLA:            Unit{},
		MCSlice:        []*Entity{},
		Tax:            Tax{MinValue: v.WOZpercent, LandValue: v.Landvalue, UsablePeriod: v.HoldPeriod, RETT: .07, CIT: map[float64]float64{}, VAT: 0, CarryBackYrs: v.CarryBackYrs, CarryForwardYrs: v.CarryForwardYrs},
		COA:            map[int]FloatCOA{},
		Valuation:      Valuation{EntryYield: v.EntryYield, YieldShift: v.YieldShift, ExitYield: v.ExitYield, AcqFees: map[string]Fee{}, DispFees: map[string]Fee{}, IncomeCapSetup: FloatCOA{TotalERV: 1}, IncomeDeduction: FloatCOA{}},
		TableHeader:    HeaderType{},
		Table:          []TableJSON{},
		Strategy:       v.Strategy,
		BalloonPercent: v.BalloonPercent,
	}
	e.CalculateModel(false)
	return e
}

// UpdateEntity -
func (e *Entity) UpdateEntity(mc bool, v *EntityData) {
	startdate := Dateadd(Datetype{Month: v.StartMonth, Year: v.StartYear}, 0)
	salesdate := Dateadd(startdate, v.HoldPeriod*12-1)
	enddate := Dateadd(salesdate, 60)
	growthinput := map[string]float64{}
	growthinput["CPI"] = v.CPIGrowth
	growthinput["ERV"] = v.ERVGrowth
	temp := Entity{
		MasterID:      v.MasterID,
		Name:          v.Name,
		ChildEntities: map[int]Entity{},
		ChildUnits:    map[int]*Unit{},
		Metrics:       Metrics{},
		Parent:        &Entity{},
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
			MasterID:              0,
			Name:                  "",
			LeaseStartDate:        startdate,
			LeaseExpiryDate:       Datetype{},
			UnitStatus:            "",
			Tenant:                "",
			PassingRent:           0,
			RentSchedule:          RentSchedule{},
			RSStore:               []RentSchedule{},
			Parent:                &Entity{},
			Probability:           v.GLA.Probability,
			PercentSoldRent:       v.GLA.PercentSoldRent,
			BondIncome:            0,
			BondExpense:           0,
			Default:               Default{Hazard: v.GLA.Default.Hazard, DefaultEnd: Datetype{}},
			RentRevisionERV:       v.GLA.RentRevisionERV,
			EXTDuration:           v.GLA.EXTDuration,
			IndexDetails:          IndexDetails{Frequency: v.GLA.IndexDetails.Frequency, Type: v.GLA.IndexDetails.Type, StartMonth: 0, Anniversary: ""},
			RentIncentivesMonths:  v.GLA.RentIncentivesMonths,
			RentIncentivesPercent: v.GLA.RentIncentivesPercent,
			Void:                  v.GLA.Void,
			FitOutCosts:           CostInput{Amount: v.GLA.FitOutCosts.Amount, AmountPerTotalArea: v.GLA.FitOutCosts.AmountPerTotalArea, AmountPerOccupiedArea: v.GLA.FitOutCosts.AmountPerOccupiedArea, AmountPerVacantArea: v.GLA.FitOutCosts.AmountPerVacantArea, PercentOfERV: v.GLA.FitOutCosts.PercentOfERV, PercentOfTRI: v.GLA.FitOutCosts.PercentOfTRI, PercentOfContractRent: v.GLA.FitOutCosts.PercentOfContractRent},
			DiscountRate:          v.GLA.DiscountRate,
			ERVArea:               0,
			ERVAmount:             0,
		},
		Tax:            Tax{MinValue: v.WOZpercent, LandValue: v.Landvalue, UsablePeriod: v.HoldPeriod, RETT: v.RETT, CIT: map[float64]float64{}, VAT: v.VAT, CarryBackYrs: v.CarryBackYrs, CarryForwardYrs: v.CarryForwardYrs},
		COA:            map[int]FloatCOA{},
		Valuation:      Valuation{EntryYield: v.EntryYield, YieldShift: v.YieldShift, ExitYield: v.ExitYield, AcqFees: map[string]Fee{}, DispFees: map[string]Fee{}, IncomeCapSetup: FloatCOA{TotalERV: 1}, IncomeDeduction: FloatCOA{}},
		TableHeader:    HeaderType{},
		Table:          []TableJSON{},
		Strategy:       v.Strategy,
		BalloonPercent: v.BalloonPercent,
	}
	temp.CalculateModel(mc)
	(*e) = temp
}

// CalculateModel - mc == MonteCarlo; if true then table is not made
func (e *Entity) CalculateModel(mc bool) {
	e.StartDate.Add(0)
	e.SalesDate.Add(0)
	e.EndDate.Add(0)
	e.Growth = make(map[string]map[int]float64)
	e.GrowthCalc()
	e.PopulateUnits()
	// e.CalculateUnits()
	e.AssetRentCalc(mc)
	e.ValuationCalc()
	e.Acquisition()
	e.PropertyCFCalc()
	e.Disposal()
	e.SumCOA()
	// e.CIT()
	e.SumNCF()
	e.SumCOA()
	if mc == false {
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
			Capex:                   false,
			AcqDispProperty:         true,
			AcqDispCosts:            false,
			LoanProceeds:            false,
			InterestExpense:         true,
			LoanBalance:             false,
			Debt:                    true,
			Tax:                     true,
			Fees:                    true,
			NetCashFlow:             true,
			CashBalance:             true,
			BondIncome:              true,
			BondExpense:             true,
		}
		e.MakeTable(coas)
	}
	e.MetricsCalc()
}

// PopulateUnits -
func (e *Entity) PopulateUnits() {
	for _, v := range UnitStore {
		if v.ParentMasterID == e.MasterID {
			// from UnitStore
			temp := Unit{
				MasterID:              v.MasterID,
				Name:                  v.Name,
				LeaseStartDate:        Datetype{Month: v.LeaseStartMonth, Year: v.LeaseStartYear},
				LeaseExpiryDate:       Datetype{Month: v.LeaseEndMonth, Year: v.LeaseEndYear},
				UnitStatus:            v.UnitStatus,
				Tenant:                v.Tenant,
				PassingRent:           v.PassingRent,
				RentSchedule:          RentSchedule{},
				RSStore:               []RentSchedule{},
				Parent:                e,
				Probability:           e.GLA.Probability,
				PercentSoldRent:       e.GLA.PercentSoldRent,
				BondIncome:            0,
				BondExpense:           0,
				Default:               e.GLA.Default,
				RentRevisionERV:       e.GLA.RentRevisionERV,
				EXTDuration:           e.GLA.EXTDuration,
				IndexDetails:          e.GLA.IndexDetails,
				RentIncentivesMonths:  e.GLA.RentIncentivesMonths,
				RentIncentivesPercent: e.GLA.RentIncentivesPercent,
				Void:                  e.GLA.Void,
				FitOutCosts:           e.GLA.FitOutCosts,
				DiscountRate:          e.GLA.DiscountRate,
				ERVArea:               v.ERVArea,
				ERVAmount:             v.ERVAmount,
			}
			temp.LeaseStartDate.Add(0)
			temp.LeaseExpiryDate.Add(0)
			e.ChildUnits[v.MasterID] = &temp
		}
	}
}

// CalculateUnits -
func (e *Entity) CalculateUnits() {
	for _, u := range e.ChildUnits {
		u.LeaseStartDate.Add(0)
		u.LeaseExpiryDate.Add(0)
		// u.RentScheduleCalc()
		// fmt.Printf("%+v\n", u)
	}
}
