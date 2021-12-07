package controllers

import (
	"math"
)

// MetricsCalc -
func (e *Entity) MetricsCalc() {
	e.Metrics.IRR.NetLeveredAfterTax = math.Round(IRR(&e.COA, Dateadd(e.StartDate, -1), e.SalesDate, FloatCOA{NetCashFlow: 1})*10000) / 10000
	e.Metrics.BondHolder.YTM = math.Round(IRR(&e.COA, Dateadd(e.StartDate, -1), e.SalesDate, FloatCOA{BondExpense: 1})*10000) / 10000
	e.Metrics.EM.NetLeveredAfterTax = EquityMultipleCalc(e.StartDate, e.SalesDate, e.COA)
	bondexpense := ExtractCOALine(e.StartDate, e.SalesDate, FloatCOA{BondExpense: 1}, &e.COA)
	e.Metrics.BondHolder.Duration = Duration(bondexpense)
	e.Metrics.BondHolder.YTMDUR = e.Metrics.BondHolder.YTM / e.Metrics.BondHolder.Duration
}

// IRR -
func IRR(cf *IntFloatCOAMap, start Datetype, end Datetype, setup FloatCOA) float64 {
	values := ExtractCOALine(start, end, setup, cf)
	irr := math.Pow(IRRCalc(values)+1, 12) - 1
	return irr * 100
}

func ExtractCOALine(start Datetype, end Datetype, setup FloatCOA, cf *IntFloatCOAMap) []float64 {
	var values []float64
	for date := start; date.Dateint <= end.Dateint; date.Add(1) {
		tempcoa := setup
		tempcoa.Multiply((*cf)[date.Dateint])
		values = append(values, SumCOADown(tempcoa))
	}
	return values
}

// IRRCalc -
func IRRCalc(values []float64) float64 {
	const irrMaxInterations = 20
	const irrAccuracy = 1e-9
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
func Duration(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		// if i == 0 {
		// 	continue
		// }
		sum = sum + v
	}

	// sum := SumCOALines(FloatCOA{BondExpense: 1}, e.COA, e.StartDate, e.SalesDate)
	dur := 0.0
	for i, v := range values {
		weight := v / sum
		dur = weight*float64(i) + dur
	}
	duration := math.Round((dur/12)*10000) / 10000
	if sum == 0.0 {
		duration = 0.0
	}
	return duration
}

// EquityMultipleCalc -
func EquityMultipleCalc(start Datetype, end Datetype, coa IntFloatCOAMap) float64 {
	numerator := 0.0
	denominator := 0.0
	for date := Dateadd(start, -1); date.Dateint <= end.Dateint; date.Add(1) {
		if coa[date.Dateint].NetCashFlow < 0 {
			denominator = denominator + coa[date.Dateint].NetCashFlow
		} else {
			numerator = numerator + coa[date.Dateint].NetCashFlow
		}
	}
	return math.Round((numerator/-denominator)*100000) / 100000
}
