package controllers

import "math"

// ValuationCalc -
func (e *Entity) ValuationCalc() {
	// e.Valuation.AcqPrice = e.COA[Dateadd(e.StartDate, -1).Dateint].AcqDispProperty
	shift := e.Valuation.YieldShift / 120000
	yield := e.Valuation.EntryYield
	for date := Dateadd(e.StartDate, -1); date.Dateint <= e.EndDate.Dateint; date.Add(1) {
		mv := SumCOALines(e.Valuation.IncomeCapSetup, e.COA, Dateadd(date, 1), Dateadd(date, 12)) / yield
		yield = yield + shift
		temp := FloatCOA{MarketValue: mv}
		temp.Add(e.COA[date.Dateint])
		e.COA[date.Dateint] = temp
		if date.Dateint == e.SalesDate.Dateint {
			e.Valuation.ExitYield = math.Round(yield*10000) / 10000
		}
	}
}
