package controllers

import (
	"math/rand"
)

// RandomDefault -
func (u *Unit) RandomDefault(date Datetype, endrent float64) {
	randomdefault := rand.Float64()
	if randomdefault <= u.Default.Hazard/12 {
		u.RentSchedule.DefaultDate = date
		u.RentSchedule.EndDate = Dateadd(date, -1)
		u.RentSchedule.Probability = 0.0
		u.RentSchedule.EndContractRent = endrent
		index := len(u.RSStore) - 1
		u.RSStore[index].DefaultDate = date
		u.RSStore[index].EndDate = date
		u.RSStore[index].Probability = 0.0
		if u.Parent.Strategy == "Balloon" && date.Dateint <= u.Parent.SalesDate.Dateint {
			u.Parent.SalesDate.Add(u.Void)
		}
	}
}
