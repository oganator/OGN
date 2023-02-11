package controllers

import (
	"math"
	"sync"
)

// DirectCapCalc -
func (e *EntityModel) DirectCapCalc() {
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

// MonthlyDCFCalc -
func MonthlyDCFCalc(e *EntityModel, start Datetype, ch chan DateFloat) {
	shift := e.Valuation.YieldShift / 120000
	yield := e.Valuation.EntryYield
	discount := e.Valuation.DiscountRate + 1
	// discount := 1.08
	end := Dateadd(start, 120) // 10 year DCF
	i := 0.0
	tempvalue := 0.0
	for date := start; date.Dateint <= end.Dateint; date.Add(1) {
		yield = yield + shift
		e.Mutex.Lock()
		monthincome := SumCOADown(MultiplyCOA(e.Valuation.IncomeCapSetup, e.COA[date.Dateint]))
		e.Mutex.Unlock()
		if date.Dateint == end.Dateint {
			e.Mutex.Lock()
			exitval := SumCOALines(e.Valuation.IncomeCapSetup, e.COA, Dateadd(date, 1), Dateadd(date, 12)) / yield
			e.Mutex.Unlock()
			monthincome = monthincome + exitval
		}
		monthincome = monthincome / math.Pow(discount, i/12) // apply discount
		i++
		tempvalue = tempvalue + monthincome
	}
	ch <- DateFloat{Date: start, Float: tempvalue}
}

type DateFloat struct {
	Date  Datetype
	Float float64
}

// SumDCF -
func (e *EntityModel) SumDCF(ch chan DateFloat) {
	for v := range ch {
		e.Mutex.Lock()
		temp := FloatCOA{MarketValue: v.Float}
		temp.Add(e.COA[v.Date.Dateint])
		e.COA[v.Date.Dateint] = temp
		e.Mutex.Unlock()
	}
}

// DCFCalc - Look in to oddities with Balloon strategy, and include exit yield for sales period
func (e *EntityModel) DCFCalc() {
	wg := sync.WaitGroup{}
	ch := make(chan DateFloat)
	go e.SumDCF(ch)
	for date := Dateadd(e.StartDate, -1); date.Dateint <= e.SalesDate.Dateint; date.Add(1) {
		wg.Add(1)
		go func(date2 Datetype, ee *EntityModel) {
			defer wg.Done()
			MonthlyDCFCalc(ee, date2, ch)
		}(date, e)
	}
	wg.Wait()
	close(ch)
}

// func (e *Entity) DCFCalc() {
// 	wg := sync.WaitGroup{}
// 	for date := Dateadd(e.StartDate, -1); date.Dateint <= e.SalesDate.Dateint; date.Add(1) {
// 		wg.Add(1)
// 		go func(date2 Datetype, ee *Entity) {
// 			defer wg.Done()
// 			ee.Mutex.Lock()
// 			temp := FloatCOA{MarketValue: MonthlyDCFCalc(ee, date2)}
// 			// temp := FloatCOA{MarketValue: 25.0}
// 			temp.Add(ee.COA[date2.Dateint])
// 			ee.COA[date2.Dateint] = temp
// 			ee.Mutex.Unlock()
// 		}(date, e)
// 	}
// 	wg.Wait()
// }
