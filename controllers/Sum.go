package controllers

// SumCOALines - returns the sum of COA lines from specified start and end dates
func SumCOALines(setup FloatCOA, COA IntFloatCOAMap, start Datetype, end Datetype) (sum float64) {
	for date := start; date.Dateint <= end.Dateint; date.Add(1) {
		tempcoa := MultiplyCOA(setup, COA[date.Dateint])
		sum = sum + SumCOADown(tempcoa)
	}
	return sum
}

// SumCOA - Yearly sum of COA based on their monthly values. Sums an entities monthly values.
func (e *Entity) SumCOA() {
	x := FloatCOA{}
	for date := Dateadd(e.StartDate, -1); date.Dateint <= e.SalesDate.Dateint; date = Dateadd(date, 1) {
		if date.Month == 1 {
			x = FloatCOA{}
		}
		x = SumCOAAcross(x, e.COA[date.Dateint])
		if date.Month == 12 || date.Dateint == e.SalesDate.Dateint {
			e.COA[date.Year] = x
		}
	}
}

// SumNCF -
func (e *Entity) SumNCF() {
	cashbalance := e.COA[Dateadd(e.StartDate, -1).Dateint].NetCashFlow
	for date := e.StartDate; date.Dateint <= e.SalesDate.Dateint; date.Add(1) {
		ncf := e.COA[date.Dateint].NetOperatingIncome + e.COA[date.Dateint].Capex + e.COA[date.Dateint].AcqDispProperty + e.COA[date.Dateint].AcqDispCosts + e.COA[date.Dateint].InterestExpense + e.COA[date.Dateint].Debt + e.COA[date.Dateint].Tax + e.COA[date.Dateint].Fees
		if date.Dateint == e.SalesDate.Dateint && e.Strategy == "Balloon" {
			ncf = ncf + (e.COA[date.Dateint].BPUplift-e.COA[Dateadd(e.StartDate, -1).Dateint].BPUplift)*e.BalloonPercent
		}
		temp := FloatCOA{
			NetCashFlow: ncf,
			CashBalance: ncf + e.COA[Dateadd(date, -1).Dateint].CashBalance + cashbalance,
		}
		temp.Add(e.COA[date.Dateint])
		e.COA[date.Dateint] = temp
		cashbalance = 0
	}
}

// // SumChildren -
// func (e *Entity) SumChildren() {
// 	e.COA = IntFloatCOAMap{}
// 	for _, v := range e.ChildUnits.Array {
// 		for date := e.StartDate; date.Dateint <= e.EndDate.Dateint; date = Dateadd(date, 1) {
// 			e.COA[date.Dateint] = AddCOA(e.COA[date.Dateint], v.COA[date.Dateint])
// 		}
// 	}
// 	for _, v := range e.ChildEntities.Array {
// 		for date := e.StartDate; date.Dateint <= e.EndDate.Dateint; date = Dateadd(date, 1) {
// 			e.COA[date.Dateint] = AddCOA(e.COA[date.Dateint], v.COA[date.Dateint])
// 		}
// 	}
// }

// AddCOA - Area is excluded
func AddCOA(x, y FloatCOA) FloatCOA {
	return FloatCOA{
		MarketValue:             x.MarketValue + y.MarketValue,
		TotalERV:                x.TotalERV + y.TotalERV,
		OccupiedERV:             x.OccupiedERV + y.OccupiedERV,
		VacantERV:               x.VacantERV + y.VacantERV,
		TopSlice:                x.TopSlice + y.TopSlice,
		TotalArea:               x.TotalArea + y.TotalArea,
		OccupiedArea:            x.OccupiedArea + y.OccupiedArea,
		VacantArea:              x.VacantArea + y.VacantArea,
		PassingRent:             x.PassingRent + y.PassingRent,
		Indexation:              x.Indexation + y.Indexation,
		BPUplift:                x.BPUplift + y.BPUplift,
		TheoreticalRentalIncome: x.TheoreticalRentalIncome + y.TheoreticalRentalIncome,
		Vacancy:                 x.Vacancy + y.Vacancy,
		ContractRent:            x.ContractRent + y.ContractRent,
		RentFree:                x.RentFree + y.RentFree,
		TurnoverRent:            x.TurnoverRent + y.TurnoverRent,
		MallRent:                x.MallRent + y.MallRent,
		ParkingIncome:           x.ParkingIncome + y.ParkingIncome,
		OtherIncome:             x.OtherIncome + y.OtherIncome,
		OperatingIncome:         x.OperatingIncome + y.OperatingIncome,
		OperatingExpenses:       x.OperatingExpenses + y.OperatingExpenses,
		NetOperatingIncome:      x.NetOperatingIncome + y.NetOperatingIncome,
		Capex:                   x.Capex + y.Capex,
		AcqDispProperty:         x.AcqDispProperty + y.AcqDispProperty,
		AcqDispCosts:            x.AcqDispCosts + y.AcqDispCosts,
		LoanProceeds:            x.LoanProceeds + y.LoanProceeds,
		InterestExpense:         x.InterestExpense + y.InterestExpense,
		LoanBalance:             x.LoanBalance + y.LoanBalance,
		Debt:                    x.Debt + y.Debt,
		Tax:                     x.Tax + y.Tax,
		Fees:                    x.Fees + y.Fees,
		NetCashFlow:             x.NetCashFlow + y.NetCashFlow,
		CashBalance:             x.CashBalance + y.CashBalance,
		BondIncome:              x.BondIncome + y.BondIncome,
		BondExpense:             x.BondExpense + y.BondExpense,
	}
}

// MultiplyCOA - Area is excluded
func MultiplyCOA(x, y FloatCOA) FloatCOA {
	return FloatCOA{
		MarketValue: x.MarketValue * y.MarketValue,
		TotalERV:    x.TotalERV * y.TotalERV,
		OccupiedERV: x.OccupiedERV * y.OccupiedERV,
		VacantERV:   x.VacantERV * y.VacantERV,
		TopSlice:    x.TopSlice * y.TopSlice,
		// TotalArea:               y.TotalArea,
		// OccupiedArea:            y.OccupiedArea,
		// VacantArea:              y.VacantArea,
		PassingRent:             x.PassingRent * y.PassingRent,
		Indexation:              x.Indexation * y.Indexation,
		BPUplift:                x.BPUplift * y.BPUplift,
		TheoreticalRentalIncome: x.TheoreticalRentalIncome * y.TheoreticalRentalIncome,
		Vacancy:                 x.Vacancy * y.Vacancy,
		ContractRent:            x.ContractRent * y.ContractRent,
		RentFree:                x.RentFree * y.RentFree,
		TurnoverRent:            x.TurnoverRent * y.TurnoverRent,
		MallRent:                x.MallRent * y.MallRent,
		ParkingIncome:           x.ParkingIncome * y.ParkingIncome,
		OtherIncome:             x.OtherIncome * y.OtherIncome,
		OperatingIncome:         x.OperatingIncome * y.OperatingIncome,
		OperatingExpenses:       x.OperatingExpenses * y.OperatingExpenses,
		NetOperatingIncome:      x.NetOperatingIncome * y.NetOperatingIncome,
		Capex:                   x.Capex * y.Capex,
		AcqDispProperty:         x.AcqDispProperty * y.AcqDispProperty,
		AcqDispCosts:            x.AcqDispCosts * y.AcqDispCosts,
		LoanProceeds:            x.LoanProceeds * y.LoanProceeds,
		InterestExpense:         x.InterestExpense * y.InterestExpense,
		LoanBalance:             x.LoanBalance * y.LoanBalance,
		Debt:                    x.Debt * y.Debt,
		Tax:                     x.Tax * y.Tax,
		Fees:                    x.Fees * y.Fees,
		NetCashFlow:             x.NetCashFlow * y.NetCashFlow,
		CashBalance:             x.CashBalance * y.CashBalance,
		BondIncome:              x.BondIncome * y.BondIncome,
		BondExpense:             x.BondExpense * y.BondExpense,
	}
}

// SumCOAAcross - Area is excluded
func SumCOAAcross(x, y FloatCOA) FloatCOA {
	return FloatCOA{
		MarketValue:             y.MarketValue,
		TotalERV:                x.TotalERV + y.TotalERV,
		OccupiedERV:             x.OccupiedERV + y.OccupiedERV,
		VacantERV:               x.VacantERV + y.VacantERV,
		TopSlice:                x.TopSlice + y.TopSlice,
		TotalArea:               y.TotalArea,
		OccupiedArea:            y.OccupiedArea,
		VacantArea:              y.VacantArea,
		PassingRent:             x.PassingRent + y.PassingRent,
		Indexation:              x.Indexation + y.Indexation,
		BPUplift:                x.BPUplift + y.BPUplift,
		TheoreticalRentalIncome: x.TheoreticalRentalIncome + y.TheoreticalRentalIncome,
		Vacancy:                 x.Vacancy + y.Vacancy,
		ContractRent:            x.ContractRent + y.ContractRent,
		RentFree:                x.RentFree + y.RentFree,
		TurnoverRent:            x.TurnoverRent + y.TurnoverRent,
		MallRent:                x.MallRent + y.MallRent,
		ParkingIncome:           x.ParkingIncome + y.ParkingIncome,
		OtherIncome:             x.OtherIncome + y.OtherIncome,
		OperatingIncome:         x.OperatingIncome + y.OperatingIncome,
		OperatingExpenses:       x.OperatingExpenses + y.OperatingExpenses,
		NetOperatingIncome:      x.NetOperatingIncome + y.NetOperatingIncome,
		Capex:                   x.Capex + y.Capex,
		AcqDispProperty:         x.AcqDispProperty + y.AcqDispProperty,
		AcqDispCosts:            x.AcqDispCosts + y.AcqDispCosts,
		LoanProceeds:            x.LoanProceeds + y.LoanProceeds,
		InterestExpense:         x.InterestExpense + y.InterestExpense,
		LoanBalance:             x.LoanBalance + y.LoanBalance,
		Debt:                    x.Debt + y.Debt,
		Tax:                     x.Tax + y.Tax,
		Fees:                    x.Fees + y.Fees,
		NetCashFlow:             x.NetCashFlow + y.NetCashFlow,
		CashBalance:             y.CashBalance,
		BondIncome:              x.BondIncome + y.BondIncome,
		BondExpense:             x.BondExpense + y.BondExpense,
	}
}

// SumCOADown - will add all the COA items together
func SumCOADown(x FloatCOA) float64 {
	return x.MarketValue +
		x.TotalERV +
		x.OccupiedERV +
		x.VacantERV +
		x.TopSlice +
		x.TotalArea +
		x.OccupiedArea +
		x.VacantArea +
		x.PassingRent +
		x.Indexation +
		x.BPUplift +
		x.TheoreticalRentalIncome +
		x.Vacancy +
		x.ContractRent +
		x.RentFree +
		x.TurnoverRent +
		x.MallRent +
		x.ParkingIncome +
		x.OtherIncome +
		x.OperatingIncome +
		x.OperatingExpenses +
		x.NetOperatingIncome +
		x.Capex +
		x.AcqDispProperty +
		x.AcqDispCosts +
		x.LoanProceeds +
		x.InterestExpense +
		x.LoanBalance +
		x.Debt +
		x.Tax +
		x.Fees +
		x.NetCashFlow +
		x.CashBalance +
		x.BondIncome +
		x.BondExpense
}

// Add -
func (coa *FloatCOA) Add(x FloatCOA) {
	*coa = AddCOA(*coa, x)
}

// Multiply -
func (coa *FloatCOA) Multiply(x FloatCOA) {
	*coa = MultiplyCOA(*coa, x)
}

// ReturnCOAArray -
func ReturnCOAArray(coa IntFloatCOAMap, setup FloatCOA, start, end Datetype) (array []float64) {
	duration := dateintdiff(end.Dateint, start.Dateint)
	array = make([]float64, duration)
	for i := 0; i < duration; i++ {
		array[i] = SumCOADown(MultiplyCOA(coa[start.Dateint], setup))
		start.Add(1)
	}
	return array
}

// CreateTupleArray - creates an array of tuples from a source array. both values of the tuple are the same. this is used for the ribbonplot
func CreateTupleArray(slice []float64, xaxis bool) (tuplearray [][]float64) {
	tuplearray = make([][]float64, len(slice))
	for i, v := range slice {
		v2 := v
		if xaxis {
			v2 = v - 1
		}
		tuplearray[i] = []float64{v2, v}
	}
	return tuplearray
}
