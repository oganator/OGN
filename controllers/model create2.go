package controllers

// UpdateEntityModel - Updates a model after reading the DB, used with API requests
func (e *EntityModel) EntityModelCalc(mc bool, compute string) {
	e.HoldPeriod = dateintdiff(e.SalesDate.Dateint, e.StartDate.Dateint)
	e.EndDate = Dateadd(e.SalesDate, 144)
	e.GrowthCalc(mc)
	e.CostMapSetup()
	e.Merge()
	e.DebtSetup()
	e.AssetRentCalc(mc, compute)
	e.Valuation.IncomeCapSetup = FloatCOA{TotalERV: 1}
	switch e.Valuation.Method {
	case "DirectCap":
		e.DirectCapCalc()
	case "DCF":
		e.DCFCalc()
	}
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
			OccupiedERV:             true,
			VacantERV:               true,
			TopSlice:                true,
			TotalArea:               true,
			OccupiedArea:            true,
			VacantArea:              true,
			PassingRent:             true,
			Indexation:              true,
			TheoreticalRentalIncome: true,
			BPUplift:                true,
			Vacancy:                 true,
			ContractRent:            true,
			RentFree:                true,
			TurnoverRent:            false,
			MallRent:                false,
			ParkingIncome:           false,
			OtherIncome:             false,
			OperatingIncome:         false,
			OperatingExpenses:       true,
			NetOperatingIncome:      true,
			Capex:                   true,
			AcqDispProperty:         true,
			AcqDispCosts:            true,
			LoanProceeds:            true,
			InterestExpense:         true,
			PrincipalRepayment:      true,
			LoanBalance:             true,
			Debt:                    true,
			Tax:                     true,
			TaxableIncome:           true,
			TaxableIncomeCarryBack:  true,
			DTA:                     true,
			Depreciation:            true,
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
	if mc {
		for _, v := range e.ChildUnitModels {
			*v = UnitModel{}
		}
	}
}

// UpdateFundModel - currently this is hardcoded to use internal compute
func (e *EntityModel) UpdateFundModel() {
	e.StartDate.Dateint = 300011
	e.COA = IntFloatCOAMap{}
	for _, v := range e.ChildEntityModels {
		// set the sales date to that of the latest asset
		if e.SalesDate.Dateint < v.SalesDate.Dateint {
			e.SalesDate = v.SalesDate
			e.EndDate = v.SalesDate
		}
		// set the start date to that of the earliest asset
		if e.StartDate.Dateint > v.StartDate.Dateint {
			e.StartDate = v.StartDate
		}
		for date := Dateadd(v.StartDate, -1); date.Dateint <= v.SalesDate.Dateint; date.Add(1) {
			e.COA[date.Dateint] = AddCOA(e.COA[date.Dateint], v.COA[date.Dateint])
		}
		for date := v.StartDate.Year - 1; date <= v.SalesDate.Year; date++ {
			e.COA[date] = AddCOA(e.COA[date], v.COA[date])
		}
	}
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
	switch e.Strategy {
	case "Standard":
		coas.BPUplift = false
		coas.RentFree = true
		coas.Debt = true
		coas.BondIncome = false
		coas.BondExpense = false
	}
	e.MakeTable(coas, false, true)
	e.MetricsCalc()
}
