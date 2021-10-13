package controllers

import (
	"math"
)

// PropertyCFCalc -
func (e *Entity) PropertyCFCalc() {
	for date := e.StartDate; date.Dateint <= e.EndDate.Dateint; date.Add(1) {
		opex := e.COA[date.Dateint].TheoreticalRentalIncome * -e.OpEx.PercentOfTRI
		interestexpense := e.DebtInput.InterestRate * e.DebtInput.LTV * e.COA[Dateadd(e.StartDate, -1).Dateint].AcqDispProperty / 12
		fees := e.COA[date.Dateint].MarketValue * -e.Fees.PercentOfGAV / 1200
		temp := FloatCOA{
			OperatingExpenses:  opex,
			NetOperatingIncome: opex + e.COA[date.Dateint].OperatingIncome,
			Capex:              0,
			InterestExpense:    interestexpense,
			LoanBalance:        0,
			Debt:               0,
			Fees:               fees,
		}
		temp.Add(e.COA[date.Dateint])
		e.COA[date.Dateint] = temp
	}
}

// Acquisition -
func (e *Entity) Acquisition() {
	soldrent := 0.0
	e.Mutex.Lock()
	acq := e.COA[Dateadd(e.StartDate, -1).Dateint].MarketValue
	e.Mutex.Unlock()
	e.Valuation.AcqPrice = acq
	rett := acq * -e.Tax.RETT
	switch e.Strategy {
	case "Pure Discount":
		soldrent = SumCOALines(FloatCOA{BondIncome: 1.0}, e.COA, e.StartDate, e.SalesDate)
	case "Amortized Coupon", "Balloon":
		for _, u := range e.ChildUnits {
			leaselength := dateintdiff(u.LeaseExpiryDate.Dateint, e.StartDate.Dateint)
			monthstosell := math.Min(float64(leaselength), float64(e.HoldPeriod*12))
			soldrent = soldrent + u.PassingRent/12*float64(monthstosell)*u.PercentSoldRent
		}
	}
	fees := soldrent * 0.0 //-.025
	cf0 := Dateadd(e.StartDate, -1)
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
		AcqDispProperty:         -acq,
		AcqDispCosts:            -acq * .01,
		LoanProceeds:            acq * e.DebtInput.LTV,
		InterestExpense:         0,
		LoanBalance:             0,
		Debt:                    acq * e.DebtInput.LTV,
		Tax:                     rett,
		Fees:                    fees,
		NetCashFlow:             soldrent - acq*1.01 + rett + acq*e.DebtInput.LTV + fees,
		CashBalance:             0,
		BondIncome:              soldrent,
		BondExpense:             -soldrent,
	}
}

// Disposal -
func (e *Entity) Disposal() {
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
		Debt:            -e.COA[Dateadd(e.StartDate, -1).Dateint].Debt,
		NetCashFlow:     ncf,
		BondExpense:     bondexp,
	}
	temp.Add(e.COA[e.SalesDate.Dateint])
	e.COA[e.SalesDate.Dateint] = temp
}

// SumNCF -
func (e *Entity) SumNCF() {
	cashbalance := e.COA[Dateadd(e.StartDate, -1).Dateint].NetCashFlow
	for date := e.StartDate; date.Dateint <= e.SalesDate.Dateint; date.Add(1) {
		ncf := e.COA[date.Dateint].NetOperatingIncome + e.COA[date.Dateint].Capex + e.COA[date.Dateint].AcqDispProperty + e.COA[date.Dateint].AcqDispCosts + e.COA[date.Dateint].InterestExpense + e.COA[date.Dateint].Debt + e.COA[date.Dateint].Tax + e.COA[date.Dateint].Fees
		if date.Dateint == e.SalesDate.Dateint && e.Strategy == "Balloon" {
			ncf = ncf + (e.COA[date.Dateint].BPUplift-e.COA[Dateadd(e.StartDate, -1).Dateint].BPUplift)*e.BalloonPercent
		}
		temp := FloatCOA{
			NetCashFlow: ncf,
			CashBalance: ncf + e.COA[Dateadd(date, -1).Dateint].CashBalance + cashbalance,
		}
		temp.Add(e.COA[date.Dateint])
		e.COA[date.Dateint] = temp
		cashbalance = 0
	}
}
