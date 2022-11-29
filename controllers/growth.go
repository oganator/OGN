package controllers

import (
	"math"
)

type HModel struct {
	ShortTermRate    float64 `json:"ShortTermRate,omitempty"`
	ShortTermPeriod  int     `json:"ShortTermPeriod,omitempty"`
	TransitionPeriod int     `json:"TransitionPeriod,omitempty"`
	LongTermRate     float64 `json:"LongTermRate,omitempty"`
	EndingIndex      float64 `json:"EndingIndex,omitempty"`
}

// GrowthCalc -
func (e *EntityModel) GrowthCalc(mc bool) {
	duration := dateintdiff(e.EndDate.Dateint, e.StartDate.Dateint)
	const tw = 1.0 / 12.0
	for i, v := range e.GrowthInput {
		e.Growth[i] = make(map[int]float64, duration)
		prevdate := e.StartDate
		e.Growth[i][prevdate.Dateint] = math.Pow(1.0+v.ShortTermRate, tw)
		shortTermEndDate := Dateadd(e.StartDate, v.ShortTermPeriod)
		longTermStartDate := Dateadd(e.StartDate, v.ShortTermPeriod+v.TransitionPeriod)
		marginaltransitionrate := (v.LongTermRate - v.ShortTermRate) / float64(v.TransitionPeriod)
		transitionrate := v.ShortTermRate
		for date := Dateadd(prevdate, 1); date.Dateint <= Dateadd(e.EndDate, 120).Dateint; date.Add(1) {
			switch {
			case date.Dateint <= shortTermEndDate.Dateint:
				e.Growth[i][date.Dateint] = math.Pow(1.0+v.ShortTermRate, tw) * e.Growth[i][prevdate.Dateint]
				prevdate = date
			case date.Dateint < longTermStartDate.Dateint:
				e.Growth[i][date.Dateint] = math.Pow(1.0+transitionrate, tw) * e.Growth[i][prevdate.Dateint]
				transitionrate = transitionrate + marginaltransitionrate
				prevdate = date
			case date.Dateint >= longTermStartDate.Dateint:
				e.Growth[i][date.Dateint] = math.Pow(1.0+v.LongTermRate, tw) * e.Growth[i][prevdate.Dateint]
				prevdate = date
			}
		}
	}
}
