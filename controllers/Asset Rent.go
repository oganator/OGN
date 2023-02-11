package controllers

import (
	"math"
	"math/rand"
	"sync"

	"gonum.org/v1/gonum/floats"
)

// AssetRentCalc - Calculates up to NOI. calculates units concurrently, then sums at entity level.
func (e *EntityModel) AssetRentCalc(mc bool, compute string) {
	e.COA = map[int]FloatCOA{}
	wg := sync.WaitGroup{}
	for _, u := range e.ChildUnitModels {
		wg.Add(1)
		go func(uu *UnitModel, ee *EntityModel, mcmc bool) {
			defer wg.Done()
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
				// CAPEX - create capex before a pointer to it is sent to vacancy and rent incentives, that way they can both modify it
				capex := 0.0
				// VACANCY
				vacancy, void := VacancyCalc(mcmc, uu, date, ee, &capex)
				rentfree := RentIncentivesCalc(date, uu, &capex)
				// BPUPLIFT
				var bpuplift float64
				var bondexpense float64
				var interestexpense float64
				bpuplift, bondexpense, interestexpense, bondincome = BPUplift(passingrent, indexation, uu, void, ee, date, bondincome)
				// uu.Mutex.Lock()
				uu.COA[date.Dateint] = FloatCOA{
					MarketValue:             0,
					TotalERV:                uu.ERVAmount * uu.ERVArea * ee.Growth["ERV"][date.Dateint] / 12,
					OccupiedERV:             uu.ERVAmount * uu.ERVArea * ee.Growth["ERV"][date.Dateint] * void / 12,
					VacantERV:               uu.ERVAmount * uu.ERVArea * ee.Growth["ERV"][date.Dateint] * (1 - void) / 12,
					TopSlice:                0,
					TotalArea:               uu.ERVArea,
					OccupiedArea:            uu.ERVArea,
					VacantArea:              0,
					PassingRent:             passingrent,
					Indexation:              indexation,
					TheoreticalRentalIncome: passingrent + indexation,
					BPUplift:                bpuplift,
					Vacancy:                 vacancy,
					ContractRent:            passingrent + indexation + bpuplift + vacancy,
					RentFree:                rentfree,
					TurnoverRent:            0,
					MallRent:                0,
					ParkingIncome:           0,
					OtherIncome:             0,
					OperatingIncome:         passingrent + indexation + bpuplift + rentfree + vacancy,
					Capex:                   capex,
					InterestExpense:         interestexpense,
					BondIncome:              bondincome,
					BondExpense:             bondexpense,
				}
				// uu.Mutex.Unlock()
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

func RentIncentivesCalc(date Datetype, uu *UnitModel, capex *float64) (rentfree float64) {
	// RENEW
	if date.Dateint <= uu.RentSchedule.RentIncentivesEndRenew.Dateint {
		if uu.RentIncentives.IsCapitalized {
			*capex = *capex + uu.RentSchedule.RenewRent*-uu.RentIncentives.PercentOfContractRent
		} else {
			rentfree = uu.RentSchedule.RenewRent * -uu.RentIncentives.PercentOfContractRent
		}
	}
	// ROTATE
	if date.Dateint <= uu.RentSchedule.RentIncentivesEndRotate.Dateint && date.Dateint > uu.RentSchedule.VacancyEnd.Dateint {
		if uu.RentIncentives.IsCapitalized {
			*capex = *capex + uu.RentSchedule.RotateRent*-uu.RentIncentives.PercentOfContractRent
		} else {
			rentfree = rentfree + uu.RentSchedule.RotateRent*-uu.RentIncentives.PercentOfContractRent
		}
	}
	return rentfree
}

func FitOutCostsCalc(uu *UnitModel, date Datetype) float64 {
	endingindex := 1.0
	// uu.FitOutCosts.AmountPerTotalArea = 15.0
	uu.FitOutCosts.IsIndexed = true
	if uu.FitOutCosts.IsIndexed {
		endingindex = uu.Parent.Growth["CPI"][date.Dateint]
	}
	voidslice := []float64{float64(uu.Void), 1.0}
	void := floats.Max(voidslice)
	return uu.FitOutCosts.AmountPerTotalArea * uu.ERVArea * -endingindex / void
}

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
			duration := dateintdiff(uu.RSStore[len(uu.RSStore)-1].DefaultDate.Dateint, uu.RSStore[len(uu.RSStore)-1].StartDate.Dateint)
			*capex = *capex + FitOutCostsCalc(uu, date)*float64((duration/uu.EXTDuration))
			void = 0.0
		}
	case mcmc:
		if date.Dateint <= uu.RentSchedule.VacancyEnd.Dateint {
			vacancy = -uu.ERVAmount * uu.ERVArea * uu.RentSchedule.ProbabilitySim / 12 * ee.Growth["ERV"][date.Dateint]
			void = 0.0
			*capex = *capex + FitOutCostsCalc(uu, date)
		}
	case !mcmc:
		if date.Dateint <= uu.RentSchedule.VacancyEnd.Dateint {
			vacancy = -uu.ERVAmount * uu.ERVArea * (1 - uu.RentSchedule.Probability) / 12 * ee.Growth["ERV"][date.Dateint]
			void = 0.0
			*capex = *capex + FitOutCostsCalc(uu, date)
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
		EXTNumber:               0,
		StartDate:               u.Parent.StartDate,
		VacancyEnd:              Dateadd(u.LeaseStartDate, u.Void),
		RentIncentivesEndRenew:  Dateadd(u.LeaseStartDate, u.RentIncentives.Duration),
		RentIncentivesEndRotate: Dateadd(u.LeaseStartDate, u.RentIncentives.Duration+vacancyvoid+1),
		DefaultDate:             Datetype{},
		EndDate:                 Dateadd(u.LeaseExpiryDate, vacancyvoid),
		OriginalEndDate:         Dateadd(u.LeaseExpiryDate, vacancyvoid),
		RenewRent:               renewrent,
		RotateRent:              rotaterent,
		PassingRent:             u.PassingRent / 12,
		EndContractRent:         0,
		RentRevisionERV:         u.RentRevisionERV,
		Probability:             prob,
		RenewIndex:              Indexation{IndexNumber: 0, StartDate: indexdate, EndDate: Dateadd(indexdate, 12), Amount: 1, RentSchedule: &u.RentSchedule},
		RotateIndex:             Indexation{IndexNumber: 0, StartDate: indexdate, EndDate: Dateadd(indexdate, 12), Amount: 1, RentSchedule: &u.RentSchedule},
		ParentUnit:              u,
	}
	u.RentSchedule = temp
	u.BondIncome = u.PassingRent / 12 * u.PercentSoldRent
	u.RentSchedule.RenewIndex.StartDate.Add(0)
	u.RSStore = append(u.RSStore, temp)
}

// RentScheduleCalc -
func (u *UnitModel) RentScheduleCalc(date Datetype, mc bool, compute string) {
	u.RSStore[len(u.RSStore)-1].EndContractRent = u.RentSchedule.EndContractRent
	// renew := (u.Parent.Growth["ERV"][date.Dateint]*u.ERVAmount*u.ERVArea-u.RentSchedule.EndContractRent*12)*u.RentSchedule.RentRevisionERV + u.RentSchedule.EndContractRent*12
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
		EXTNumber:               u.RentSchedule.EXTNumber + 1,
		StartDate:               Dateadd(u.RentSchedule.EndDate, 1),
		VacancyEnd:              Dateadd(u.RentSchedule.EndDate, u.Void),
		RentIncentivesEndRenew:  Dateadd(u.RentSchedule.EndDate, u.RentIncentives.Duration),
		RentIncentivesEndRotate: Dateadd(u.RentSchedule.EndDate, u.RentIncentives.Duration+u.Void),
		DefaultDate:             Datetype{},
		EndDate:                 Dateadd(u.RentSchedule.EndDate, duration),
		OriginalEndDate:         Dateadd(u.RentSchedule.EndDate, duration),
		// RenewRent:               renew * prob / 12,
		// RotateRent:              rotate * (1 - prob) / 12,
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

	// temp.RenewRent = temp.RenewRent / prob

	// temp.RotateRent = temp.RotateRent / (1 - prob)
	// fmt.Println("RentScheduleCalc: parent.growth: ", u.Parent.Growth["ERV"][date.Dateint], "ERVAmount", u.ERVAmount, "ERVArea", u.ERVArea, "prob", prob)
	u.RSStore = append(u.RSStore, temp)
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
		EXTNumber:               u.RentSchedule.EXTNumber + 1,
		StartDate:               date,
		VacancyEnd:              Dateadd(date, u.Void-1),
		RentIncentivesEndRotate: Dateadd(Dateadd(u.RentSchedule.EndDate, 1), u.RentIncentives.Duration+u.Void),
		DefaultDate:             Datetype{},
		EndDate:                 Dateadd(u.RentSchedule.EndDate, duration),
		OriginalEndDate:         Dateadd(u.RentSchedule.EndDate, duration),
		RotateRent:              rotate / 12,
		PassingRent:             rotate / 12,
		EndContractRent:         0,
		RentRevisionERV:         u.RentSchedule.RentRevisionERV,
		Probability:             u.RentSchedule.Probability,
		RotateIndex:             Indexation{IndexNumber: 0, StartDate: Dateadd(indexdate, u.Void), EndDate: Dateadd(indexdate, 12+u.Void), Amount: 1, RentSchedule: &u.RentSchedule},
		ParentUnit:              u,
	}
	u.RentSchedule = temp
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
