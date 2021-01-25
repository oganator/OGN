package controllers

import (
	"math"
)

// AssetRentCalc - Calculates up to NOI
func (e *Entity) AssetRentCalc(mc bool) {
	e.COA = map[int]FloatCOA{}
	erv := 0.0
	area := 0.0
	for _, u := range e.ChildUnits {
		erv = erv + u.ERVAmount*u.ERVArea/12
		area = area + u.ERVArea
		u.InitialRentScheduleCalc()
		bondindex := Indexation{}
		bondindex.IndexationCalc(*u.Parent, u.RentSchedule.StartDate)
		//
		leaselength := dateintdiff(u.LeaseExpiryDate.Dateint, e.StartDate.Dateint)
		monthstosell := math.Min(float64(leaselength), float64(e.HoldPeriod*12))
		soldrent := u.PassingRent / 12 * float64(monthstosell) * u.PercentSoldRent
		for date := u.Parent.StartDate; date.Dateint <= u.Parent.EndDate.Dateint; date.Add(1) {
			if date.Dateint <= e.SalesDate.Dateint && date.Dateint > u.RentSchedule.VacancyEnd.Dateint && mc == true {
				u.RandomDefault(date, u.Parent.COA[date.Dateint].PassingRent+u.RentSchedule.PassingRent)
			}
			if date.Dateint == Dateadd(u.RentSchedule.EndDate, 1).Dateint {
				u.RentSchedule.EndContractRent = u.RentSchedule.RenewRent * u.RentSchedule.RenewIndex.Amount
				u.RentScheduleCalc(date)
			}
			if date.Dateint == u.RentSchedule.RenewIndex.EndDate.Dateint {
				u.RentSchedule.RenewIndex.IndexationCalc(*u.Parent, date)
			}
			if date.Dateint == u.RentSchedule.RotateIndex.EndDate.Dateint {
				u.RentSchedule.RotateIndex.IndexationCalc(*u.Parent, date)
			}
			if date.Dateint == bondindex.EndDate.Dateint {
				bondindex.IndexationCalc(*u.Parent, date)
			}
			passingrent := u.Parent.COA[date.Dateint].PassingRent + u.RentSchedule.PassingRent
			indexation := u.Parent.COA[date.Dateint].Indexation + u.RentSchedule.RenewIndex.Final*u.RentSchedule.PassingRent
			rentfree := 0.0
			vacancy := 0.0
			void := 1.0
			if date.Dateint <= u.RentSchedule.VacancyEnd.Dateint {
				vacancy = (-passingrent - indexation) * (1 - u.RentSchedule.Probability)
				void = 0.0
			}
			bpuplift := (passingrent + indexation) * -u.PercentSoldRent * void
			bondincome := u.BondIncome //monthly passing rent at start * percent rent sold
			bondexpense := 0.0
			interestexpense := 0.0
			switch u.Parent.Strategy {
			case "Pure Discount":
				discount := math.Pow(math.Pow(1+u.DiscountRate, .0833333333333333), float64(dateintdiff(date.Dateint, u.Parent.StartDate.Dateint)))
				bondincome = u.BondIncome / discount
				bondexpense = -(bondincome * bondindex.Final)
			case "Amortized Coupon":
				interestexpense = bondincome * float64(u.Parent.HoldPeriod) * -u.DiscountRate * void
				bondexpense = -bondincome - interestexpense
			case "Balloon":
				interestexpense = (soldrent * -u.DiscountRate / 12) * bondindex.Amount * void
				bondexpense = ((bondincome*bondindex.Amount)*(1-u.Parent.BalloonPercent) - interestexpense) * void
				bpuplift = bpuplift * (1 - u.Parent.BalloonPercent)
			}
			u.Parent.COA[date.Dateint] = FloatCOA{
				MarketValue:             0,
				TotalERV:                erv * u.Parent.Growth["ERV"][date.Dateint],
				OccupiedERV:             erv * u.Parent.Growth["ERV"][date.Dateint],
				VacantERV:               0,
				TopSlice:                0,
				TotalArea:               area,
				OccupiedArea:            area,
				VacantArea:              0,
				PassingRent:             passingrent,
				Indexation:              indexation,
				TheoreticalRentalIncome: passingrent + indexation,
				BPUplift:                bpuplift,
				Vacancy:                 vacancy,
				ContractRent:            passingrent + indexation + bpuplift + vacancy,
				RentFree:                0,
				TurnoverRent:            0,
				MallRent:                0,
				ParkingIncome:           0,
				OtherIncome:             0,
				OperatingIncome:         passingrent + indexation + bpuplift + rentfree + vacancy,
				Capex:                   0,
				InterestExpense:         interestexpense,
				BondIncome:              bondincome,
				BondExpense:             bondexpense,
			}
		}
	}
}

// InitialRentScheduleCalc -
func (u *Unit) InitialRentScheduleCalc() {
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
	temp := RentSchedule{
		EXTNumber:         0,
		StartDate:         u.Parent.StartDate,
		VacancyEnd:        Dateadd(u.LeaseStartDate, u.Void),
		RentIncentivesEnd: Dateadd(u.LeaseStartDate, u.RentIncentivesMonths),
		DefaultDate:       Datetype{},
		EndDate:           Dateadd(u.LeaseExpiryDate, 0),
		OriginalEndDate:   Datetype{},
		RenewRent:         u.PassingRent / 12,
		RotateRent:        0,
		PassingRent:       u.PassingRent / 12,
		EndContractRent:   0,
		RentRevisionERV:   u.RentRevisionERV,
		Probability:       u.Probability,
		RenewIndex:        Indexation{IndexNumber: 0, StartDate: indexdate, EndDate: Dateadd(indexdate, 12), Amount: 1},
		RotateIndex:       Indexation{},
		ParentUnit:        u,
	}
	u.RentSchedule = temp
	u.BondIncome = u.PassingRent / 12 * u.PercentSoldRent
	u.RentSchedule.RenewIndex.StartDate.Add(0)
	u.RSStore = append(u.RSStore, temp)
}

// RentScheduleCalc -
func (u *Unit) RentScheduleCalc(date Datetype) {
	renew := (u.Parent.Growth["ERV"][date.Dateint]*u.ERVAmount*u.ERVArea-u.RentSchedule.EndContractRent*12)*u.RentSchedule.RentRevisionERV + u.RentSchedule.EndContractRent*12
	// fmt.Println(date, u.Parent.Growth["ERV"][date.Dateint], u.ERVAmount, u.ERVArea, -u.RentSchedule.EndContractRent, u.RentSchedule.RentRevisionERV, u.RentSchedule.EndContractRent)
	rotate := u.Parent.Growth["ERV"][date.Dateint] * u.ERVAmount * u.ERVArea
	indexyear := date.Year
	indexdate := Datetype{
		Month: u.LeaseStartDate.Month,
		Year:  indexyear,
	}
	indexdate.Add(0)
	temp := RentSchedule{
		EXTNumber:         u.RentSchedule.EXTNumber + 1,
		StartDate:         Dateadd(u.RentSchedule.EndDate, 1),
		VacancyEnd:        Dateadd(u.RentSchedule.EndDate, u.Void),
		RentIncentivesEnd: Dateadd(Dateadd(u.RentSchedule.EndDate, 1), u.RentIncentivesMonths),
		DefaultDate:       Datetype{},
		EndDate:           Dateadd(u.RentSchedule.EndDate, u.EXTDuration),
		OriginalEndDate:   Dateadd(u.RentSchedule.EndDate, u.EXTDuration),
		RenewRent:         renew / 12,
		RotateRent:        rotate,
		PassingRent:       renew*u.Probability/12 + rotate*(1-u.Probability)/12,
		EndContractRent:   0,
		RentRevisionERV:   u.RentSchedule.RentRevisionERV,
		Probability:       u.RentSchedule.Probability,
		RenewIndex:        Indexation{IndexNumber: 0, StartDate: indexdate, EndDate: Dateadd(indexdate, 12), Amount: 1},
		RotateIndex:       Indexation{IndexNumber: 0, StartDate: Dateadd(indexdate, u.Void), EndDate: Dateadd(indexdate, 12+u.Void), Amount: 1},
		ParentUnit:        u,
	}
	u.RentSchedule = temp
	u.RSStore = append(u.RSStore, temp)
}

// IndexationCalcOld - Calculates the next index. Date is startdate
func (i *Indexation) IndexationCalcOld(e Entity, date Datetype) {
	i.IndexNumber++
	i.StartDate = date
	i.EndDate = Dateadd(i.StartDate, 12)
	i.Amount = (e.Growth["CPI"][i.StartDate.Dateint]) / (e.Growth["CPI"][e.StartDate.Dateint])
	i.Final = i.Amount - 1
}

// IndexationCalc - Calculates the next index. Date is startdate
func (i *Indexation) IndexationCalc(e Entity, date Datetype) {
	i.IndexNumber++
	i.StartDate = date
	i.EndDate = Dateadd(i.StartDate, 12)
	i.Amount = math.Pow(1+e.GrowthInput["CPI"], float64(i.IndexNumber))
	i.Final = i.Amount - 1
}
