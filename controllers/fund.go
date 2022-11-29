package controllers

// CalculateFund - need to fix cash balance: if one asset is sold after another, the cash balance of the first will end with its sale
// TODO - include sensitivity analysis (to asset exposure) by implementing a map that stores the memory address of each assets COA array, and a function that loops over them to do calculations, and then storing it all in appropriate place
func (e *EntityModel) CalculateFund() {
	e.UpdateEntity(false, EntityDataStore[e.MasterID], "Internal")
	e.COA = IntFloatCOAMap{}
	for _, v := range e.ChildEntities {
		// set the sales date to that of the latest asset
		if e.SalesDate.Dateint < v.SalesDate.Dateint {
			e.SalesDate = v.SalesDate
			e.EndDate = Dateadd(v.SalesDate, 120)
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
