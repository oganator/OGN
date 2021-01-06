package controllers

import (
	"math"
)

// GrowthCalc -
func (e *Entity) GrowthCalc() {
	duration := dateintdiff(e.EndDate.Dateint, e.StartDate.Dateint)
	for i, v := range e.GrowthInput {
		e.Growth[i] = make(map[int]float64, duration)
		prevdate := e.StartDate
		e.Growth[i][prevdate.Dateint] = math.Pow(1+v, .083333333333333333333333333333333333)
		for date := Dateadd(prevdate, 1); date.Dateint <= e.EndDate.Dateint; {
			e.Growth[i][date.Dateint] = math.Pow(1+v, .083333333333333333333333333333333333) * e.Growth[i][prevdate.Dateint]
			prevdate = date
			date = Dateadd(date, 1)
		}
	}
}
