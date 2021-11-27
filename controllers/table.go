package controllers

// MakeTable - Date, COA, Amount
func (e *Entity) MakeTable(coas BoolCOA, quarterly bool, yearly bool) {
	e.CreateTableHeader(Monthly, quarterly, yearly)
	e.Table = make([]TableJSON, 0)
	marketvalue := make(map[int]string)
	totalerv := make(map[int]string)
	occupiederv := make(map[int]string)
	vacanterv := make(map[int]string)
	topslice := make(map[int]string)
	totalarea := make(map[int]string)
	occupiedarea := make(map[int]string)
	vacantarea := make(map[int]string)
	blank := make(map[int]string)
	netoperatingincome := make(map[int]string)
	capex := make(map[int]string)
	acqdispproperty := make(map[int]string)
	debt := make(map[int]string)
	acqdispcosts := make(map[int]string)
	tax := make(map[int]string)
	taxableincome := make(map[int]string)
	taxableincomecarryback := make(map[int]string)
	dta := make(map[int]string)
	depreciation := make(map[int]string)
	fees := make(map[int]string)
	operatingincome := make(map[int]string)
	operatingexpenses := make(map[int]string)
	contractrent := make(map[int]string)
	rentfree := make(map[int]string)
	turnoverrent := make(map[int]string)
	mallrent := make(map[int]string)
	parkingincome := make(map[int]string)
	otherincome := make(map[int]string)
	theoreticalrentalincome := make(map[int]string)
	vacancy := make(map[int]string)
	passingrent := make(map[int]string)
	indexation := make(map[int]string)
	bpuplift := make(map[int]string)
	loanproceeds := make(map[int]string)
	interestexpense := make(map[int]string)
	loanbalance := make(map[int]string)
	netcashflow := make(map[int]string)
	cashbalance := make(map[int]string)
	bondincome := make(map[int]string)
	bondexpense := make(map[int]string)
	for date := Dateadd(e.StartDate, -1); date.Dateint <= e.SalesDate.Dateint; date = Dateadd(date, 1) {
		if date.Dateint == e.StartDate.Dateint || date.Month == 1 || date == Dateadd(e.StartDate, -1) {
			marketvalue[date.Year] = RenderFloat("#,###.", e.COA[date.Year].MarketValue)
			totalerv[date.Year] = RenderFloat("#,###.", e.COA[date.Year].TotalERV)
			occupiederv[date.Year] = RenderFloat("#,###.", e.COA[date.Year].OccupiedERV)
			vacanterv[date.Year] = RenderFloat("#,###.", e.COA[date.Year].VacantERV)
			topslice[date.Year] = RenderFloat("#,###.", e.COA[date.Year].TopSlice)
			totalarea[date.Year] = RenderFloat("#,###.", e.COA[date.Year].TotalArea)
			occupiedarea[date.Year] = RenderFloat("#,###.", e.COA[date.Year].OccupiedArea)
			vacantarea[date.Year] = RenderFloat("#,###.", e.COA[date.Year].VacantArea)
			blank[date.Year] = ""
			netoperatingincome[date.Year] = RenderFloat("#,###.", e.COA[date.Year].NetOperatingIncome)
			capex[date.Year] = RenderFloat("#,###.", e.COA[date.Year].Capex)
			acqdispproperty[date.Year] = RenderFloat("#,###.", e.COA[date.Year].AcqDispProperty)
			debt[date.Year] = RenderFloat("#,###.", e.COA[date.Year].Debt)
			acqdispcosts[date.Year] = RenderFloat("#,###.", e.COA[date.Year].AcqDispCosts)
			tax[date.Year] = RenderFloat("#,###.", e.COA[date.Year].Tax)
			taxableincome[date.Year] = RenderFloat("#,###.", e.COA[date.Year].TaxableIncome)
			taxableincomecarryback[date.Year] = RenderFloat("#,###.", e.COA[date.Year].TaxableIncomeCarryBack)
			dta[date.Year] = RenderFloat("#,###.", e.COA[date.Year].DTA)
			depreciation[date.Year] = RenderFloat("#,###.", e.COA[date.Year].Depreciation)
			fees[date.Year] = RenderFloat("#,###.", e.COA[date.Year].Fees)
			operatingincome[date.Year] = RenderFloat("#,###.", e.COA[date.Year].OperatingIncome)
			operatingexpenses[date.Year] = RenderFloat("#,###.", e.COA[date.Year].OperatingExpenses)
			contractrent[date.Year] = RenderFloat("#,###.", e.COA[date.Year].ContractRent)
			rentfree[date.Year] = RenderFloat("#,###.", e.COA[date.Year].RentFree)
			turnoverrent[date.Year] = RenderFloat("#,###.", e.COA[date.Year].TurnoverRent)
			mallrent[date.Year] = RenderFloat("#,###.", e.COA[date.Year].MallRent)
			parkingincome[date.Year] = RenderFloat("#,###.", e.COA[date.Year].ParkingIncome)
			otherincome[date.Year] = RenderFloat("#,###.", e.COA[date.Year].OtherIncome)
			theoreticalrentalincome[date.Year] = RenderFloat("#,###.", e.COA[date.Year].TheoreticalRentalIncome)
			vacancy[date.Year] = RenderFloat("#,###.", e.COA[date.Year].Vacancy)
			passingrent[date.Year] = RenderFloat("#,###.", e.COA[date.Year].PassingRent)
			indexation[date.Year] = RenderFloat("#,###.", e.COA[date.Year].Indexation)
			bpuplift[date.Year] = RenderFloat("#,###.", e.COA[date.Year].BPUplift)
			loanproceeds[date.Year] = RenderFloat("#,###.", e.COA[date.Year].LoanProceeds)
			interestexpense[date.Year] = RenderFloat("#,###.", e.COA[date.Year].InterestExpense)
			loanbalance[date.Year] = RenderFloat("#,###.", e.COA[date.Year].LoanBalance)
			netcashflow[date.Year] = RenderFloat("#,###.", e.COA[date.Year].NetCashFlow)
			blank[date.Year] = ""
			cashbalance[date.Year] = RenderFloat("#,###.", e.COA[date.Year].CashBalance)
			bondincome[date.Year] = RenderFloat("#,###.", e.COA[date.Year].BondIncome)
			bondexpense[date.Year] = RenderFloat("#,###.", e.COA[date.Year].BondExpense)
		}
		marketvalue[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].MarketValue)
		totalerv[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].TotalERV)
		occupiederv[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].OccupiedERV)
		vacanterv[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].VacantERV)
		topslice[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].TopSlice)
		totalarea[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].TotalArea)
		occupiedarea[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].OccupiedArea)
		vacantarea[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].VacantArea)
		blank[date.Dateint] = ""
		netoperatingincome[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].NetOperatingIncome)
		capex[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].Capex)
		acqdispproperty[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].AcqDispProperty)
		debt[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].Debt)
		acqdispcosts[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].AcqDispCosts)
		tax[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].Tax)
		taxableincome[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].TaxableIncome)
		taxableincomecarryback[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].TaxableIncomeCarryBack)
		dta[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].DTA)
		depreciation[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].Depreciation)
		fees[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].Fees)
		operatingincome[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].OperatingIncome)
		operatingexpenses[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].OperatingExpenses)
		contractrent[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].ContractRent)
		rentfree[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].RentFree)
		turnoverrent[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].TurnoverRent)
		mallrent[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].MallRent)
		parkingincome[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].ParkingIncome)
		otherincome[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].OtherIncome)
		theoreticalrentalincome[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].TheoreticalRentalIncome)
		vacancy[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].Vacancy)
		passingrent[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].PassingRent)
		indexation[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].Indexation)
		bpuplift[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].BPUplift)
		loanproceeds[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].LoanProceeds)
		interestexpense[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].InterestExpense)
		loanbalance[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].LoanBalance)
		netcashflow[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].NetCashFlow)
		blank[date.Dateint] = ""
		cashbalance[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].CashBalance)
		bondincome[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].BondIncome)
		bondexpense[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].BondExpense)
	}
	// build e.Table based on COA selection
	if coas.MarketValue == true {
		x := TableJSON{
			COA:   "Market Value",
			Value: marketvalue,
		}
		e.Table = append(e.Table, x)
	}
	if coas.TotalERV == true {
		x := TableJSON{
			COA:   "Total ERV",
			Value: totalerv,
		}
		e.Table = append(e.Table, x)
	}
	if coas.OccupiedERV == true {
		x := TableJSON{
			COA:   "Occupied ERV",
			Value: occupiederv,
		}
		e.Table = append(e.Table, x)
	}
	if coas.VacantERV == true {
		x := TableJSON{
			COA:   "Vacant ERV",
			Value: vacanterv,
		}
		e.Table = append(e.Table, x)
	}
	if coas.TopSlice == true {
		x := TableJSON{
			COA:   "Top Slice",
			Value: topslice,
		}
		e.Table = append(e.Table, x)
	}
	if coas.TotalArea == true {
		x := TableJSON{
			COA:   "Total Area",
			Value: totalarea,
		}
		e.Table = append(e.Table, x)
	}
	if coas.OccupiedArea == true {
		x := TableJSON{
			COA:   "Occupied Area",
			Value: occupiedarea,
		}
		e.Table = append(e.Table, x)
	}
	if coas.VacantArea == true {
		x := TableJSON{
			COA:   "Vacant Area",
			Value: vacantarea,
		}
		e.Table = append(e.Table, x)
	}
	e.Table = append(e.Table, TableJSON{COA: "", Value: blank})
	if coas.PassingRent == true {
		x := TableJSON{
			COA:   "Passing Rent",
			Value: passingrent,
		}
		e.Table = append(e.Table, x)
	}
	if coas.Indexation == true {
		x := TableJSON{
			COA:   "Indexation",
			Value: indexation,
		}
		e.Table = append(e.Table, x)
	}
	if coas.TheoreticalRentalIncome == true {
		x := TableJSON{
			COA:   "Theoretical Rental Income",
			Value: theoreticalrentalincome,
		}
		e.Table = append(e.Table, x)
	}
	if coas.BPUplift == true {
		x := TableJSON{
			COA:   "BP Uplift",
			Value: bpuplift,
		}
		e.Table = append(e.Table, x)
	}
	if coas.Vacancy == true {
		x := TableJSON{
			COA:   "Vacancy",
			Value: vacancy,
		}
		e.Table = append(e.Table, x)
	}
	if coas.ContractRent == true {
		x := TableJSON{
			COA:   "Contract Rent",
			Value: contractrent,
		}
		e.Table = append(e.Table, x)
	}
	if coas.RentFree == true {
		x := TableJSON{
			COA:   "Rent Free",
			Value: rentfree,
		}
		e.Table = append(e.Table, x)
	}
	if coas.TurnoverRent == true {
		x := TableJSON{
			COA:   "Turnover Rent",
			Value: turnoverrent,
		}
		e.Table = append(e.Table, x)
	}
	if coas.MallRent == true {
		x := TableJSON{
			COA:   "Mall Rent",
			Value: mallrent,
		}
		e.Table = append(e.Table, x)
	}
	if coas.ParkingIncome == true {
		x := TableJSON{
			COA:   "Parking Income",
			Value: parkingincome,
		}
		e.Table = append(e.Table, x)
	}
	if coas.OtherIncome == true {
		x := TableJSON{
			COA:   "Other Income",
			Value: otherincome,
		}
		e.Table = append(e.Table, x)
	}
	if coas.OperatingIncome == true {
		x := TableJSON{
			COA:   "Operating Income",
			Value: operatingincome,
		}
		e.Table = append(e.Table, x)
	}
	if coas.OperatingExpenses == true {
		x := TableJSON{
			COA:   "Operating Expenses",
			Value: operatingexpenses,
		}
		e.Table = append(e.Table, x)
	}
	if coas.NetOperatingIncome == true {
		x := TableJSON{
			COA:   "Net Operating Income",
			Value: netoperatingincome,
		}
		e.Table = append(e.Table, x)
	}
	if coas.Capex == true {
		x := TableJSON{
			COA:   "Capex",
			Value: capex,
		}
		e.Table = append(e.Table, x)
	}
	if coas.AcqDispProperty == true {
		x := TableJSON{
			COA:   "Acq/Disp Property",
			Value: acqdispproperty,
		}
		e.Table = append(e.Table, x)
	}
	if coas.AcqDispCosts == true {
		x := TableJSON{
			COA:   "Acq/Disp Costs",
			Value: acqdispcosts,
		}
		e.Table = append(e.Table, x)
	}
	if coas.LoanProceeds == true {
		x := TableJSON{
			COA:   "Loan Proceeds",
			Value: loanproceeds,
		}
		e.Table = append(e.Table, x)
	}
	if coas.InterestExpense == true {
		x := TableJSON{
			COA:   "Interest Expense",
			Value: interestexpense,
		}
		e.Table = append(e.Table, x)
	}
	if coas.LoanBalance == true {
		x := TableJSON{
			COA:   "Loan Balance",
			Value: loanbalance,
		}
		e.Table = append(e.Table, x)
	}
	if coas.Debt == true {
		x := TableJSON{
			COA:   "Debt",
			Value: debt,
		}
		e.Table = append(e.Table, x)
	}
	if coas.Tax == true {
		x := TableJSON{
			COA:   "Tax",
			Value: tax,
		}
		e.Table = append(e.Table, x)
	}
	if coas.TaxableIncome == true {
		x := TableJSON{
			COA:   "Taxable Income",
			Value: taxableincome,
		}
		e.Table = append(e.Table, x)
	}
	if coas.TaxableIncomeCarryBack == true {
		x := TableJSON{
			COA:   "Taxable Income Carry Back",
			Value: taxableincomecarryback,
		}
		e.Table = append(e.Table, x)
	}
	if coas.DTA == true {
		x := TableJSON{
			COA:   "DTA",
			Value: dta,
		}
		e.Table = append(e.Table, x)
	}
	if coas.Depreciation == true {
		x := TableJSON{
			COA:   "Depreciation",
			Value: depreciation,
		}
		e.Table = append(e.Table, x)
	}
	if coas.Fees == true {
		x := TableJSON{
			COA:   "Fees",
			Value: fees,
		}
		e.Table = append(e.Table, x)
	}
	if coas.NetCashFlow == true {
		x := TableJSON{
			COA:   "Net Cash Flow",
			Value: netcashflow,
		}
		e.Table = append(e.Table, x)
	}
	e.Table = append(e.Table, TableJSON{COA: "", Value: blank})
	if coas.CashBalance == true {
		x := TableJSON{
			COA:   "Cash Balance",
			Value: cashbalance,
		}
		e.Table = append(e.Table, x)
	}
	if coas.BondIncome == true {
		x := TableJSON{
			COA:   "Bond Income",
			Value: bondincome,
		}
		e.Table = append(e.Table, x)
	}
	if coas.BondExpense == true {
		x := TableJSON{
			COA:   "Bond Expense",
			Value: bondexpense,
		}
		e.Table = append(e.Table, x)
	}

}

// TableJSON -  used for output. Value is map[dateint]formattedvalue
type TableJSON struct {
	COA   string
	Value IntStringMap
}

// IntStringMap - used for value field of TableJSON
type IntStringMap map[int]string

// BoolCOA -
type BoolCOA struct {
	MarketValue             bool
	TotalERV                bool
	OccupiedERV             bool
	VacantERV               bool
	TopSlice                bool
	TotalArea               bool
	OccupiedArea            bool
	VacantArea              bool
	PassingRent             bool
	Indexation              bool
	TheoreticalRentalIncome bool
	BPUplift                bool
	Vacancy                 bool
	ContractRent            bool
	RentFree                bool
	TurnoverRent            bool
	MallRent                bool
	ParkingIncome           bool
	OtherIncome             bool
	OperatingIncome         bool
	OperatingExpenses       bool
	NetOperatingIncome      bool
	Capex                   bool
	AcqDispProperty         bool
	AcqDispCosts            bool
	LoanProceeds            bool
	InterestExpense         bool
	LoanBalance             bool
	Debt                    bool
	Tax                     bool
	TaxableIncome           bool
	TaxableIncomeCarryBack  bool
	DTA                     bool
	Depreciation            bool
	Fees                    bool
	NetCashFlow             bool
	CashBalance             bool
	BondIncome              bool
	BondExpense             bool
}

// FloatCOA -
type FloatCOA struct {
	MarketValue             float64
	TotalERV                float64
	OccupiedERV             float64
	VacantERV               float64
	TopSlice                float64
	TotalArea               float64
	OccupiedArea            float64
	VacantArea              float64
	PassingRent             float64
	Indexation              float64
	TheoreticalRentalIncome float64
	BPUplift                float64 // accounts for the effect of sold rent
	Vacancy                 float64
	ContractRent            float64
	RentFree                float64
	TurnoverRent            float64
	MallRent                float64
	ParkingIncome           float64
	OtherIncome             float64
	OperatingIncome         float64
	OperatingExpenses       float64
	NetOperatingIncome      float64
	Capex                   float64
	AcqDispProperty         float64
	AcqDispCosts            float64
	LoanProceeds            float64
	InterestExpense         float64
	LoanBalance             float64
	Debt                    float64
	Tax                     float64
	TaxableIncome           float64
	TaxableIncomeCarryBack  float64
	DTA                     float64
	Depreciation            float64
	Fees                    float64
	NetCashFlow             float64
	CashBalance             float64
	BondIncome              float64 // income for me
	BondExpense             float64 // expense for me, or income for investor
}

// CreateTableHeader -
func (e *Entity) CreateTableHeader(month, quarter, year bool) {
	// Month
	if month {
		ForecastMonthly := make([]Datetype, 0)
		for mdate := Dateadd(e.StartDate, -1); mdate.Dateint <= e.EndDate.Dateint; mdate = Dateadd(mdate, 1) {
			if mdate.Dateint <= e.SalesDate.Dateint {
				mdate.Bool = true
			}
			if mdate.Dateint >= Dateadd(e.StartDate, -1).Dateint {
				ForecastMonthly = append(ForecastMonthly, mdate)
			}
		}
		e.TableHeader.Monthly = ForecastMonthly
	}
	// Quarter
	// Year
	if year {
		ForecastYearly := make([]Datetype, 0)
		for ydate := Dateadd(e.StartDate, -1); ydate.Year <= e.EndDate.Year; ydate = Dateadd(ydate, 12) {
			if ydate.Year <= e.SalesDate.Year {
				ydate.Bool = true
			}
			if ydate.Year >= Dateadd(e.StartDate, -1).Year {
				ForecastYearly = append(ForecastYearly, ydate)
			}
		}
		e.TableHeader.Yearly = ForecastYearly
	}
}
