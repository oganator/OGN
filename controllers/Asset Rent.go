package controllers

import (
	"math"
)

// AssetRentCalc - Calculates up to NOI
func (e *Entity) AssetRentCalc() {
	e.COA = map[int]FloatCOA{}
	erv := 0.0
	area := 0.0
	for _, u := range e.ChildUnits {
		erv = erv + u.ERVAmount*u.ERVArea/12
		area = area + u.ERVArea
		u.InitialRentScheduleCalc()
		for date := e.StartDate; date.Dateint <= e.EndDate.Dateint; date.Add(1) {
			if date.Dateint == u.RentSchedule.EndDate.Dateint {
				u.RentSchedule.EndContractRent = u.RentSchedule.PassingRent + u.RentSchedule.RenewIndex.Amount + u.RentSchedule.RotateIndex.Amount
				u.RentScheduleCalc()
			}
			if date.Dateint == u.RentSchedule.RenewIndex.EndDate.Dateint {
				u.RentSchedule.RenewIndex.IndexationCalc(*e, date)
			}
			if date.Dateint == u.RentSchedule.RotateIndex.EndDate.Dateint {
				u.RentSchedule.RotateIndex.IndexationCalc(*e, date)
			}
			// determine vacancy via if statement
			// u.RandomDefault
			passingrent := e.COA[date.Dateint].PassingRent + u.RentSchedule.PassingRent
			indexation := e.COA[date.Dateint].Indexation + u.RentSchedule.RenewIndex.Final*u.RentSchedule.PassingRent
			bpuplift := (passingrent + indexation) * -u.PercentSoldRent
			rentfree := 0.0
			vacancy := 0.0
			bondincome := 0.0
			bondexpense := 0.0
			interestexpense := 0.0
			switch e.Strategy {
			case "Pure Discount":
				discount := math.Pow(math.Pow(1+u.DiscountRate, .0833333333333333), float64(dateintdiff(date.Dateint, e.StartDate.Dateint)))
				bondincome = (passingrent * u.PercentSoldRent) / discount
				bondexpense = -bpuplift
			case "Amortized Coupon":
				bondincome = (passingrent * u.PercentSoldRent)
				interestexpense = bondincome * float64(e.HoldPeriod) * -u.DiscountRate //* (1 + u.RentSchedule.RenewIndex.Final)
				bondexpense = -bpuplift - interestexpense
			case "Balloon":
				bondincome = (passingrent * u.PercentSoldRent)
				interestexpense = bondincome * float64(e.HoldPeriod) * -u.DiscountRate * (1 + u.RentSchedule.RenewIndex.Final)
				bondexpense = -interestexpense - bpuplift*(1-e.BalloonPercent)
				bpuplift = bpuplift * (1 - e.BalloonPercent) // formerly 0.0
			}
			e.COA[date.Dateint] = FloatCOA{
				MarketValue:             0,
				TotalERV:                erv * e.Growth["ERV"][date.Dateint],
				OccupiedERV:             erv * e.Growth["ERV"][date.Dateint],
				VacantERV:               0,
				TopSlice:                0,
				TotalArea:               area,
				OccupiedArea:            area,
				VacantArea:              0,
				PassingRent:             passingrent,
				Indexation:              indexation,
				TheoreticalRentalIncome: passingrent + indexation,
				BPUplift:                bpuplift,
				Vacancy:                 0,
				ContractRent:            passingrent + indexation + bpuplift + vacancy,
				RentFree:                0,
				TurnoverRent:            0,
				MallRent:                0,
				ParkingIncome:           0,
				OtherIncome:             0,
				OperatingIncome:         passingrent + indexation + bpuplift + rentfree,
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
	indexyear := u.Parent.StartDate.Year
	if u.Parent.StartDate.Month > u.LeaseStartDate.Month {
		indexyear = u.Parent.StartDate.Year + 1
	}
	indexdate := Datetype{
		Month: u.LeaseStartDate.Month,
		Year:  indexyear,
	}
	indexdate.Add(0)
	u.RentSchedule = RentSchedule{
		StartDate:         u.Parent.StartDate,
		VacancyEnd:        Dateadd(u.LeaseStartDate, u.Void),
		RentIncentivesEnd: Dateadd(u.LeaseStartDate, u.RentIncentivesMonths),
		EndDate:           Dateadd(u.LeaseExpiryDate, 0),
		PassingRent:       u.PassingRent / 12,
		ParentUnit:        u,
		RenewIndex: Indexation{
			IndexNumber: 0,
			StartDate:   indexdate,
			EndDate:     Dateadd(indexdate, 12),
			Amount:      1,
		},
	}
	u.RentSchedule.RenewIndex.StartDate.Add(0)
}

// RentScheduleCalc -
func (u *Unit) RentScheduleCalc() {
	u.RentSchedule = RentSchedule{
		EXTNumber:         u.RentSchedule.EXTNumber + 1,
		StartDate:         Dateadd(u.RentSchedule.StartDate, 12),
		VacancyEnd:        Datetype{},
		RentIncentivesEnd: Datetype{},
		DefaultDate:       Datetype{},
		EndDate:           Datetype{},
		OriginalEndDate:   Datetype{},
		RenewRent:         0,
		RotateRent:        0,
		PassingRent:       0,
		EndContractRent:   0,
		RentRevisionERV:   0,
		Probability:       0,
		RenewIndex:        Indexation{},
		RotateIndex:       Indexation{},
		ParentUnit:        u,
	}
}

// IndexationCalc - Calculates the next index
func (i *Indexation) IndexationCalc(e Entity, date Datetype) {
	i.IndexNumber++
	i.StartDate = date
	i.EndDate = Dateadd(i.StartDate, 12)
	i.Amount = (e.Growth["CPI"][i.StartDate.Dateint]) / (e.Growth["CPI"][e.StartDate.Dateint])
	i.Final = i.Amount - 1
}
