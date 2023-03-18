package controllers

import (
	"math"
	"math/rand"
	"sync"
)

// AssetRentCalc - Calculates up to NOI. calculates units concurrently, then sums at entity level.
func (e *EntityModel) AssetRentCalc(mc bool, compute string) {
	e.COA = map[int]FloatCOA{}
	wg := sync.WaitGroup{}
	for _, u := range e.ChildUnitModels {
		wg.Add(1)
		go func(uu *UnitModel, ee *EntityModel, mcmc bool) {
			defer wg.Done()
			// TODO - implement unit overrides
			uu.InitialRentScheduleCalc()
			uu.COA = map[int]FloatCOA{}
			if uu.PercentSoldRent > 0.0 {
				uu.BondIndex = Indexation{RentSchedule: &uu.RentSchedule}
				uu.BondIndex.IndexationCalc(e, uu.RentSchedule.StartDate, false)
				leaselength := dateintdiff(uu.LeaseExpiryDate.Dateint, ee.StartDate.Dateint)
				monthstosell := math.Min(float64(leaselength), float64(ee.HoldPeriod))
				uu.BondProceeds = uu.PassingRent / 12 * float64(monthstosell) * uu.PercentSoldRent
			}
			bondincome := uu.BondIncome
			for date := ee.StartDate; date.Dateint <= ee.EndDate.Dateint; date.Add(1) {
				passingrent, indexation := IndexCalc(uu, date, ee, mcmc, compute)
				// CAPEX
				capex := 0.0
				vacancy, void := VacancyCalc(mcmc, uu, date, ee, &capex)
				// BPUPLIFT
				var bpuplift float64
				var bondexpense float64
				var interestexpense float64
				bpuplift, bondexpense, interestexpense, bondincome = BPUplift(passingrent, indexation, uu, void, ee, date, bondincome)
				firstPass := FloatCOA{
					MarketValue:             0,
					TotalERV:                uu.ERVAmount * uu.ERVArea * ee.Growth["ERV"][date.Dateint] / 12,
					OccupiedERV:             uu.ERVAmount * uu.ERVArea * ee.Growth["ERV"][date.Dateint] * void / 12,
					VacantERV:               uu.ERVAmount * uu.ERVArea * ee.Growth["ERV"][date.Dateint] * (1 - void) / 12,
					TopSlice:                passingrent + indexation + bpuplift + vacancy - (uu.ERVAmount * uu.ERVArea * ee.Growth["ERV"][date.Dateint] / 12),
					TotalArea:               uu.ERVArea,
					OccupiedArea:            uu.ERVArea * void,
					VacantArea:              uu.ERVArea * (1 - void),
					PassingRent:             passingrent,
					Indexation:              indexation,
					TheoreticalRentalIncome: passingrent + indexation,
					BPUplift:                bpuplift,
					Vacancy:                 vacancy,
					ContractRent:            passingrent + indexation + bpuplift + vacancy,
					TurnoverRent:            0,
					MallRent:                0,
					ParkingIncome:           0,
					OtherIncome:             0,
					OperatingIncome:         passingrent + indexation + bpuplift + vacancy,
					Capex:                   capex,
					InterestExpense:         interestexpense,
					BondIncome:              bondincome,
					BondExpense:             bondexpense,
				}
				// SECOND PASS - necessary because in the first pass, Area and ERV arent assigned yet. Unit level Capex can refer to ERV or Area
				secondPass := uu.CostInput.CostMapCalc(date, uu.Parent.Growth, firstPass)
				firstPass.Add(secondPass)
				// THIRD PASS - necessary to sum up to NCF
				thirdPass := FloatCOA{
					NetOperatingIncome: firstPass.RentFree + firstPass.OperatingExpenses,
				}
				firstPass.Add(thirdPass)
				uu.COA[date.Dateint] = firstPass
			}
		}(u, e, mc)
	}
	wg.Wait()
	tempcpi := e.GrowthInput["CPI"]
	tempcpi.EndingIndex = e.Growth["CPI"][e.EndDate.Dateint]
	temperv := e.GrowthInput["ERV"]
	temperv.EndingIndex = e.Growth["ERV"][e.EndDate.Dateint]
	e.GrowthInput["CPI"] = tempcpi
	e.GrowthInput["ERV"] = temperv
	for date := e.StartDate; date.Dateint <= e.EndDate.Dateint; date.Add(1) {
		for _, u := range e.ChildUnitModels {
			// u.Mutex.Lock()
			e.COA[date.Dateint] = AddCOA(e.COA[date.Dateint], u.COA[date.Dateint])
			// u.Mutex.Unlock()
		}
	}
}

// func RentIncentivesCalc(date Datetype, uu *UnitModel) (rentfree, capex float64) {
// 	// RENEW
// 	if date.Dateint <= uu.RentSchedule.RentIncentivesEndRenew.Dateint {

// 		if uu.CostInput["Rent Incentives Renewal"].Name == "" {
// 			return rentfree, capex
// 		}
// 		if uu.CostInput["Rent Incentives Renewal"].IsCapitalized {
// 			capex = CostCalc(uu.CostInput["Rent Incentives Renewal"], uu.COA[date.Dateint], uu.Parent.Growth, date)
// 		} else {
// 			rentfree = CostCalc(uu.CostInput["Rent Incentives Renewal"], uu.COA[date.Dateint], uu.Parent.Growth, date)
// 		}
// 	}
// 	// ROTATE
// 	if date.Dateint <= uu.RentSchedule.RentIncentivesEndRotate.Dateint {
// 		if uu.CostInput["Rent Incentives Rotation"].Name == "" {
// 			return rentfree, capex
// 		}
// 		if uu.CostInput["Rent Incentives Rotation"].IsCapitalized {
// 			capex = CostCalc(uu.CostInput["Rent Incentives Rotation"], uu.COA[date.Dateint], uu.Parent.Growth, date)
// 		} else {
// 			rentfree = CostCalc(uu.CostInput["Rent Incentives Rotation"], uu.COA[date.Dateint], uu.Parent.Growth, date)
// 		}
// 	}
// 	return rentfree, capex
// }

func BPUplift(passingrent float64, indexation float64, uu *UnitModel, void float64, ee *EntityModel, date Datetype, bondincome float64) (float64, float64, float64, float64) {
	bpuplift := (passingrent + indexation) * -uu.PercentSoldRent * void
	bondexpense := 0.0
	interestexpense := 0.0
	switch ee.Strategy {
	case "Standard":
	case "Pure Discount":
		discount := math.Pow(math.Pow(1+uu.DiscountRate, .0833333333333333), float64(dateintdiff(date.Dateint, ee.StartDate.Dateint)))
		bondincome = uu.BondIncome / discount
		bondexpense = -(bondincome * uu.BondIndex.Final)
	case "Amortized Coupon":
		interestexpense = bondincome * float64(ee.HoldPeriod/12) * -uu.DiscountRate * void
		bondexpense = -bondincome - interestexpense
	case "Balloon":
		interestexpense = (uu.BondProceeds * -uu.DiscountRate / 12) * uu.BondIndex.Amount * void
		bondexpense = ((bondincome*uu.BondIndex.Amount)*(1-ee.BalloonPercent) - interestexpense) * void
		bpuplift = bpuplift * (1 - ee.BalloonPercent)
	}
	return bpuplift, bondexpense, interestexpense, bondincome
}

func IndexCalc(uu *UnitModel, date Datetype, ee *EntityModel, mcmc bool, compute string) (float64, float64) {
	renewrent := uu.RentSchedule.RenewRent
	rotaterent := uu.RentSchedule.RotateRent
	isdefault := false
	if date.Dateint <= ee.SalesDate.Dateint && date.Dateint > uu.RentSchedule.VacancyEnd.Dateint && mcmc {
		isdefault = uu.RandomDefault(date, ee.COA[date.Dateint].PassingRent+uu.RentSchedule.PassingRent)
	}
	if date.Dateint == Dateadd(uu.RentSchedule.EndDate, 1).Dateint && !isdefault {
		uu.RentSchedule.EndContractRent = renewrent*uu.RentSchedule.RenewIndex.Amount*uu.Probability + rotaterent*uu.RentSchedule.RotateIndex.Amount*(1-uu.Probability)
		uu.RentScheduleCalc(date, mcmc, compute)
	}
	if date.Dateint == uu.RentSchedule.RenewIndex.EndDate.Dateint {
		uu.RentSchedule.RenewIndex.IndexationCalc(ee, date, false)
	}
	if date.Dateint == uu.RentSchedule.RotateIndex.EndDate.Dateint {
		uu.RentSchedule.RotateIndex.IndexationCalc(ee, date, false)
	}
	if date.Dateint == uu.BondIndex.EndDate.Dateint {
		uu.BondIndex.IndexationCalc(ee, date, false)
	}
	passingrent := uu.RentSchedule.PassingRent
	indexation := uu.RentSchedule.RenewIndex.Final*renewrent + uu.RentSchedule.RotateIndex.Final*rotaterent
	return passingrent, indexation
}

// func FitOutCostsCalc(uu *UnitModel, date Datetype) float64 {
// 	endingindex := 1.0
// 	// uu.FitOutCosts.AmountPerTotalArea = 15.0
// 	if uu.CostInput["Fit Out Costs"].GrowthItem != "None" {
// 		endingindex = uu.Parent.Growth["CPI"][date.Dateint]
// 	}
// 	voidslice := []float64{float64(uu.Void), 1.0}
// 	void := floats.Max(voidslice)
// 	return uu.CostInput["Fit Out Costs"].AmountPerTotalArea * uu.ERVArea * -endingindex / void
// }

func VacancyCalc(mcmc bool, uu *UnitModel, date Datetype, ee *EntityModel, capex *float64) (vacancy float64, void float64) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	void = 1.0
	switch {
	case mcmc && uu.RSStore[len(uu.RSStore)-1].DefaultDate.Year > 1: // case - tenant has just defaulted
		if date.Dateint <= uu.RentSchedule.VacancyEnd.Dateint {
			vacancy = -uu.ERVAmount * uu.ERVArea * ee.Growth["ERV"][date.Dateint] / 12
			// duration := dateintdiff(uu.RSStore[len(uu.RSStore)-1].DefaultDate.Dateint, uu.RSStore[len(uu.RSStore)-1].StartDate.Dateint)
			// *capex = *capex + FitOutCostsCalc(uu, date)*float64((duration/uu.EXTDuration))
			// *capex = *capex + uu.CostInput.CostMapCalc(date, "Fit Out Costs", uu.COA[date.Dateint], uu.Parent.Growth)*float64((duration/uu.EXTDuration))
			void = 0.0
		}
	case mcmc:
		if date.Dateint <= uu.RentSchedule.VacancyEnd.Dateint {
			vacancy = -uu.ERVAmount * uu.ERVArea * uu.RentSchedule.ProbabilitySim / 12 * ee.Growth["ERV"][date.Dateint]
			void = 0.0
			// *capex = *capex + FitOutCostsCalc(uu, date)
			// *capex = *capex + uu.CostInput.CostMapCalc(date, "Fit Out Costs", uu.COA[date.Dateint], uu.Parent.Growth)
		}
	case !mcmc:
		if date.Dateint <= uu.RentSchedule.VacancyEnd.Dateint {
			vacancy = -uu.ERVAmount * uu.ERVArea * (1 - uu.RentSchedule.Probability) / 12 * ee.Growth["ERV"][date.Dateint]
			void = 0.0
			// *capex = *capex + FitOutCostsCalc(uu, date)
			// *capex = *capex + uu.CostInput.CostMapCalc(date, "Fit Out Costs", uu.COA[date.Dateint], uu.Parent.Growth)
		}
	}
	return vacancy, void
}

// InitialRentScheduleCalc -
func (u *UnitModel) InitialRentScheduleCalc() {
	u.RSStore = make([]RentSchedule, 0)
	indexyear := u.Parent.StartDate.Year - 1
	if u.Parent.StartDate.Month > u.LeaseStartDate.Month {
		indexyear = u.Parent.StartDate.Year
	}
	indexdate := Datetype{
		Month: u.LeaseStartDate.Month,
		Year:  indexyear,
	}
	indexdate.Add(0)
	renewrent := u.PassingRent / 12
	rotaterent := u.PassingRent / 12
	prob := u.Probability
	vacancyvoid := 0
	if u.UnitStatus == "Vacant" {
		renewrent = u.PassingRent / 12
		rotaterent = u.PassingRent / 12
		indexdate.Add(u.Void + 1)
		prob = 0.0
		vacancyvoid = u.Void - 1
	}
	temp := RentSchedule{
		EXTNumber:  0,
		StartDate:  u.Parent.StartDate,
		VacancyEnd: Dateadd(u.LeaseStartDate, u.Void),
		// RentIncentivesEndRenew:  Dateadd(u.LeaseStartDate, u.CostInput["Rent Incentives Renewal"].Duration),
		// RentIncentivesEndRotate: Dateadd(u.LeaseStartDate, u.CostInput["Rent Incentives Rotation"].Duration+vacancyvoid+1),
		DefaultDate:     Datetype{},
		EndDate:         Dateadd(u.LeaseExpiryDate, vacancyvoid),
		OriginalEndDate: Dateadd(u.LeaseExpiryDate, vacancyvoid),
		RenewRent:       renewrent,
		RotateRent:      rotaterent,
		PassingRent:     u.PassingRent / 12,
		EndContractRent: 0,
		RentRevisionERV: u.RentRevisionERV,
		Probability:     prob,
		RenewIndex:      Indexation{IndexNumber: 0, StartDate: indexdate, EndDate: Dateadd(indexdate, 12), Amount: 1, RentSchedule: &u.RentSchedule},
		RotateIndex:     Indexation{IndexNumber: 0, StartDate: indexdate, EndDate: Dateadd(indexdate, 12), Amount: 1, RentSchedule: &u.RentSchedule},
		ParentUnit:      u,
	}
	if u.UnitStatus == "Vacant" {
		rotation := u.CostInput["Rent Incentives Rotation"]
		rotation.Start = Dateadd(temp.StartDate, u.Void)
		rotation.End = Dateadd(temp.StartDate, rotation.Duration+u.Void-1)
		u.CostInput["Rent Incentives Rotation"] = rotation
		fitoutcosts := u.CostInput["Fit Out Costs"]
		fitoutcosts.Start = temp.StartDate
		fitoutcosts.End = Dateadd(temp.StartDate, 0)
		u.CostInput["Fit Out Costs"] = fitoutcosts
	}
	//
	u.RentSchedule = temp
	u.BondIncome = u.PassingRent / 12 * u.PercentSoldRent
	u.RentSchedule.RenewIndex.StartDate.Add(0)
	u.RSStore = append(u.RSStore, temp)
}

// RentScheduleCalc -
func (u *UnitModel) RentScheduleCalc(date Datetype, mc bool, compute string) {
	u.RSStore[len(u.RSStore)-1].EndContractRent = u.RentSchedule.EndContractRent
	rotate := u.Parent.Growth["ERV"][date.Dateint] * u.ERVAmount * u.ERVArea
	renew := rotate*u.RentSchedule.RentRevisionERV + u.RentSchedule.EndContractRent*(1-u.RentSchedule.RentRevisionERV)*12
	indexyear := date.Year
	indexdate := Datetype{
		Month: u.LeaseStartDate.Month,
		Year:  indexyear,
	}
	indexdate.Add(0)
	duration := u.EXTDuration + int(u.Probability*float64(u.Void))
	prob := u.Probability
	sample := rand.Float64()
	if mc {
		prob = 1.0
		if sample < u.Probability {
			prob = 0.0
		}
	}
	temp := RentSchedule{
		EXTNumber:  u.RentSchedule.EXTNumber + 1,
		StartDate:  Dateadd(u.RentSchedule.EndDate, 1),
		VacancyEnd: Dateadd(u.RentSchedule.EndDate, u.Void),
		// RentIncentivesEndRenew:  Dateadd(u.RentSchedule.EndDate, u.CostInput["Rent Incentives"].Duration),
		// RentIncentivesEndRotate: Dateadd(u.RentSchedule.EndDate, u.CostInput["Rent Incentives"].Duration+u.Void),
		DefaultDate:     Datetype{},
		EndDate:         Dateadd(u.RentSchedule.EndDate, duration),
		OriginalEndDate: Dateadd(u.RentSchedule.EndDate, duration),
		RenewRent:       renew / 12,
		RotateRent:      rotate / 12,
		PassingRent:     renew*prob/12 + rotate*(1-prob)/12,
		EndContractRent: 0,
		RentRevisionERV: u.RentSchedule.RentRevisionERV,
		Probability:     u.Probability,
		ProbabilitySim:  prob,
		RenewIndex:      Indexation{IndexNumber: 0, StartDate: indexdate, EndDate: Dateadd(indexdate, 12), Amount: 1, RentSchedule: &u.RentSchedule},
		RotateIndex:     Indexation{IndexNumber: 0, StartDate: Dateadd(indexdate, u.Void), EndDate: Dateadd(indexdate, 12+u.Void), Amount: 1, RentSchedule: &u.RentSchedule},
		ParentUnit:      u,
	}
	u.RentSchedule = temp
	u.RSStore = append(u.RSStore, temp)
	renewal := u.CostInput["Rent Incentives Renewal"]
	renewal.Start = temp.StartDate
	renewal.End = Dateadd(temp.StartDate, renewal.Duration)
	u.CostInput["Rent Incentives Renewal"] = renewal
	rotation := u.CostInput["Rent Incentives Rotation"]
	rotation.Start = Dateadd(temp.StartDate, u.Void)
	rotation.End = Dateadd(temp.StartDate, rotation.Duration+u.Void)
	u.CostInput["Rent Incentives Rotation"] = rotation
	fitoutcosts := u.CostInput["Fit Out Costs"]
	fitoutcosts.Start = temp.StartDate
	fitoutcosts.End = Dateadd(temp.StartDate, fitoutcosts.Duration)
	u.CostInput["Fit Out Costs"] = fitoutcosts
}

// RentScheduleDefaultCalc -
func (u *UnitModel) RentScheduleDefaultCalc(date Datetype) {
	u.RSStore = append(u.RSStore, u.RentSchedule)
	rotate := u.Parent.Growth["ERV"][date.Dateint] * u.ERVAmount * u.ERVArea
	indexyear := date.Year
	indexdate := Datetype{
		Month: u.LeaseStartDate.Month,
		Year:  indexyear,
	}
	indexdate.Add(0)
	duration := u.EXTDuration
	temp := RentSchedule{
		EXTNumber:  u.RentSchedule.EXTNumber + 1,
		StartDate:  date,
		VacancyEnd: Dateadd(date, u.Void-1),
		// RentIncentivesEndRotate: Dateadd(Dateadd(u.RentSchedule.EndDate, 1), u.CostInput["Rent Incentives"].Duration+u.Void),
		DefaultDate:     Datetype{},
		EndDate:         Dateadd(u.RentSchedule.EndDate, duration),
		OriginalEndDate: Dateadd(u.RentSchedule.EndDate, duration),
		RotateRent:      rotate / 12,
		PassingRent:     rotate / 12,
		EndContractRent: 0,
		RentRevisionERV: u.RentSchedule.RentRevisionERV,
		Probability:     u.RentSchedule.Probability,
		RotateIndex:     Indexation{IndexNumber: 0, StartDate: Dateadd(indexdate, u.Void), EndDate: Dateadd(indexdate, 12+u.Void), Amount: 1, RentSchedule: &u.RentSchedule},
		ParentUnit:      u,
	}
	u.RentSchedule = temp
	rotation := u.CostInput["Rent Incentives Rotation"]
	rotation.Start = Dateadd(temp.StartDate, u.Void)
	rotation.End = Dateadd(temp.StartDate, rotation.Duration+u.Void)
	u.CostInput["Rent Incentives Rotation"] = rotation
	fitoutcosts := u.CostInput["Fit Out Costs"]
	fitoutcosts.Start = temp.StartDate
	fitoutcosts.End = Dateadd(temp.StartDate, fitoutcosts.Duration)
	u.CostInput["Fit Out Costs"] = fitoutcosts
}

// IndexationCalc - Calculates the next index. Date is startdate
func (i *Indexation) IndexationCalc(e *EntityModel, date Datetype, debug bool) {
	i.IndexNumber++
	i.StartDate = i.RentSchedule.StartDate
	i.EndDate = Dateadd(date, 12)
	// i.Amount = math.Pow(1+e.GrowthInput["CPI"], float64(i.IndexNumber))
	i.Amount = e.Growth["CPI"][i.EndDate.Dateint] / e.Growth["CPI"][i.StartDate.Dateint]
	// if !e.MC && debug {
	// 	fmt.Println(i.StartDate.Dateint, "-", i.EndDate.Dateint, " : ", e.Growth["CPI"][i.StartDate.Dateint], ": ", e.Growth["CPI"][i.EndDate.Dateint])
	// }
	i.Final = i.Amount - 1
}
