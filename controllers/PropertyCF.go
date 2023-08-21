package controllers

import (
	"math"
)

// PropertyCFCalc -
func (e *EntityModel) PropertyCFCalc() {
	for date := e.StartDate; date.Dateint <= e.EndDate.Dateint; date.Add(1) {
		e.Mutex.Lock()
		debt := FloatCOA{}
		if date.Dateint <= e.SalesDate.Dateint {
			debt = e.DebtCalc(date)
		}
		// fees := e.COA[date.Dateint].MarketValue * -e.Fees.PercentOfGAV / 1200
		firstPass := FloatCOA{
			NetOperatingIncome: e.COA[date.Dateint].OperatingIncome,
			LoanProceeds:       debt.LoanProceeds,
			InterestExpense:    -debt.InterestExpense,
			PrincipalRepayment: -debt.PrincipalRepayment,
			LoanBalance:        debt.LoanBalance,
			Debt:               debt.Debt,
		}
		secondPass := e.CostInput.CostMapCalc(date, e.Growth, e.COA[date.Dateint])
		firstPass.Add(secondPass)
		thirdPass := FloatCOA{
			NetOperatingIncome: firstPass.OperatingExpenses,
		}
		firstPass.Add(thirdPass)
		firstPass.Add(e.COA[date.Dateint])
		e.COA[date.Dateint] = firstPass
		e.Mutex.Unlock()
	}
}

// Acquisition -
func (e *EntityModel) Acquisition() {
	soldrent := 0.0
	// e.Mutex.Lock()
	// acq := e.COA[Dateadd(e.StartDate, -1).Dateint].MarketValue
	// e.Mutex.Unlock()
	// e.Valuation.AcqPrice = acq
	// rett := acq * -e.Tax.RETT
	rett := e.Valuation.AcqPrice * -e.Tax.RETT
	switch e.Strategy {
	case "Pure Discount":
		soldrent = SumCOALines(FloatCOA{BondIncome: 1.0}, e.COA, e.StartDate, e.SalesDate)
	case "Amortized Coupon", "Balloon":
		for _, u := range e.ChildUnitModels {
			leaselength := dateintdiff(u.LeaseExpiryDate.Dateint, e.StartDate.Dateint)
			monthstosell := math.Min(float64(leaselength), float64(e.HoldPeriod))
			soldrent = soldrent + u.PassingRent/12*float64(monthstosell)*u.PercentSoldRent
		}
	}
	// fees := soldrent * 0.0 //-.025
	cf0 := Dateadd(e.StartDate, -1)
	debt := e.DebtCalc(cf0)
	// fmt.Println(debt)
	e.Mutex.Lock()
	e.COA[cf0.Dateint] = FloatCOA{
		MarketValue:             e.COA[cf0.Dateint].MarketValue,
		TotalERV:                0,
		OccupiedERV:             0,
		VacantERV:               0,
		TopSlice:                0,
		TotalArea:               0,
		OccupiedArea:            0,
		VacantArea:              0,
		PassingRent:             0,
		Indexation:              0,
		TheoreticalRentalIncome: 0,
		BPUplift:                soldrent,
		Vacancy:                 0,
		ContractRent:            0,
		RentFree:                0,
		TurnoverRent:            0,
		MallRent:                0,
		ParkingIncome:           0,
		OtherIncome:             0,
		OperatingIncome:         0,
		OperatingExpenses:       0,
		NetOperatingIncome:      0,
		Capex:                   0,
		// AcqDispProperty:         -acq,
		AcqDispProperty: -e.Valuation.AcqPrice,
		// LoanProceeds:            acq * e.DebtInput.LTV,
		LoanProceeds:    debt.LoanProceeds,
		InterestExpense: 0,
		LoanBalance:     debt.LoanBalance,
		Debt:            debt.LoanProceeds,
		Tax:             rett,
		// Fees:        fees,
		// NetCashFlow: soldrent - acq*1.01 + rett + acq*e.DebtInput.LTV + fees,
		NetCashFlow: soldrent - e.Valuation.AcqPrice + rett + debt.LoanProceeds,
		CashBalance: 0,
		BondIncome:  soldrent,
		BondExpense: -soldrent,
	}
	e.Mutex.Unlock()
}

// Disposal -
func (e *EntityModel) Disposal() {
	bpuplift := 0.0
	intexp := 0.0
	ncf := 0.0
	bondexp := 0.0
	switch e.Strategy {
	case "Balloon":
		bpuplift = -e.COA[Dateadd(e.StartDate, -1).Dateint].BPUplift * (e.Growth["CPI"][e.SalesDate.Dateint] - 1) * e.BalloonPercent
		intexp = bpuplift
		bondexp = (e.COA[Dateadd(e.StartDate, -1).Dateint].BPUplift * e.Growth["CPI"][e.SalesDate.Dateint]) * e.BalloonPercent
	}
	temp := FloatCOA{
		BPUplift:        bpuplift,
		AcqDispProperty: e.COA[e.SalesDate.Dateint].MarketValue,
		AcqDispCosts:    e.COA[e.SalesDate.Dateint].MarketValue * -0.01,
		InterestExpense: intexp,
		// Debt:            -e.COA[Dateadd(e.StartDate, -1).Dateint].Debt, // NOT NEEDED AS THIS IS NOW DONE IN DEBT.GO
		NetCashFlow: ncf,
		BondExpense: bondexp,
	}
	temp.Add(e.COA[e.SalesDate.Dateint])
	e.COA[e.SalesDate.Dateint] = temp
}
