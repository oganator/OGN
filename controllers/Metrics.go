package controllers

import (
	"math"
)

// MetricsCalc -
func (e *Entity) MetricsCalc() {

	nlatirr := IRR(&e.COA, Dateadd(e.StartDate, -1), e.SalesDate, FloatCOA{NetCashFlow: 1})
	e.Metrics.IRR.NetLeveredAfterTax = math.Round(nlatirr*100) / 100
	ytm := IRR(&e.COA, Dateadd(e.StartDate, -1), e.SalesDate, FloatCOA{BondExpense: 1})
	e.Metrics.BondHolder.YTM = math.Round(ytm*100) / 100
	e.Duration()
	e.EquityMultipleCalc()
}

// IRR -
func IRR(cf *IntFloatCOAMap, start Datetype, end Datetype, setup FloatCOA) float64 {
	var values []float64
	for date := start; date.Dateint <= end.Dateint; date.Add(1) {
		tempcoa := setup
		tempcoa.Multiply((*cf)[date.Dateint])
		values = append(values, SumCOADown(tempcoa))
	}
	irr := math.Pow(IRRCalc(values)+1, 12) - 1
	return irr * 100
}

// IRRCalc -
func IRRCalc(values []float64) float64 {
	const irrMaxInterations = 20
	const irrAccuracy = 1e-7
	const irrInitialGuess = 0
	x0 := float64(irrInitialGuess)
	var x1 float64
	for i := 0; i < irrMaxInterations; i++ {
		fValue := float64(0)
		fDerivative := float64(0)
		for k := 0; k < len(values); k++ {
			fk := float64(k)
			fValue += values[k] / math.Pow(1.0+x0, fk)
			fDerivative += -fk * values[k] * math.Pow(1.0+x0, -fk-1.0)
		}
		x1 = x0 - fValue/fDerivative
		if math.Abs(x1-x0) <= irrAccuracy {
			return x1
		}
		x0 = x1
	}
	return -1
}

// Duration -
func (e *Entity) Duration() {
	sum := SumCOALines(FloatCOA{BondExpense: 1}, e.COA, e.StartDate, e.SalesDate)
	dur := 0.0
	for date, i := e.StartDate, 1.0; date.Dateint <= e.SalesDate.Dateint; date, i = Dateadd(date, 1), i+1.0 {
		weight := e.COA[date.Dateint].BondExpense / sum
		dur = weight*i + dur
	}
	duration := math.Round((dur/12)*100) / 100
	if sum == 0.0 {
		duration = 0.0
	}
	e.Metrics.BondHolder.Duration = duration
}

// EquityMultipleCalc -
func (e *Entity) EquityMultipleCalc() {
	numerator := 0.0
	denominator := 0.0
	for date := Dateadd(e.StartDate, -1); date.Dateint <= e.SalesDate.Dateint; date.Add(1) {
		if e.COA[date.Dateint].NetCashFlow < 0 {
			denominator = denominator + e.COA[date.Dateint].NetCashFlow
		} else {
			numerator = numerator + e.COA[date.Dateint].NetCashFlow
		}
	}
	e.Metrics.EM.NetLeveredAfterTax = numerator / -denominator
}
