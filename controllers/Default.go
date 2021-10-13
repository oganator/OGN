package controllers

import (
	"math/rand"
)

// RandomDefault -
func (u *Unit) RandomDefault(date Datetype, endrent float64) (isdefault bool) {
	randomdefault := rand.Float64()
	if randomdefault < u.Default.Hazard {
		isdefault = true
		u.RentSchedule.DefaultDate = Dateadd(date, 0)
		u.RentSchedule.EndDate = Dateadd(date, 0)
		u.RentSchedule.Probability = 0.0
		u.RentSchedule.ProbabilitySim = 0.0
		u.RentSchedule.EndContractRent = endrent
		index := len(u.RSStore) - 1
		u.RSStore[index] = u.RentSchedule
		if u.Parent.Strategy == "Balloon" && date.Dateint <= u.Parent.SalesDate.Dateint {
			u.Parent.SalesDate.Add(u.Void)
			u.Parent.EndDate.Add(u.Void)
		}
		u.RentScheduleDefaultCalc(date)
		u.Default.NumberOfDefaults++
		u.Parent.GLA.Default.NumberOfDefaults++
		return isdefault
	}
	return isdefault
}
