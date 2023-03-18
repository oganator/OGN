package controllers

// MakeTable -
func (e *EntityModel) MakeTable(coas BoolCOA, quarterly bool, yearly bool) {
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
	principalrepayments := make(map[int]string)
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
			principalrepayments[date.Year] = RenderFloat("#,###.", e.COA[date.Year].PrincipalRepayment)
			loanbalance[date.Year] = RenderFloat("#,###.", e.COA[date.Year].LoanBalance)
			netcashflow[date.Year] = RenderFloat("#,###.", e.COA[date.Year].NetCashFlow)
			blank[date.Year] = ""
			cashbalance[date.Year] = RenderFloat("#,###.", e.COA[date.Year].CashBalance)
			bondincome[date.Year] = RenderFloat("#,###.", e.COA[date.Year].BondIncome)
			bondexpense[date.Year] = RenderFloat("#,###.", e.COA[date.Year].BondExpense)
		}
		if Monthly == true {
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
			principalrepayments[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].PrincipalRepayment)
			loanbalance[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].LoanBalance)
			netcashflow[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].NetCashFlow)
			blank[date.Dateint] = ""
			cashbalance[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].CashBalance)
			bondincome[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].BondIncome)
			bondexpense[date.Dateint] = RenderFloat("#,###.", e.COA[date.Dateint].BondExpense)
		}
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
	if coas.PrincipalRepayment == true {
		x := TableJSON{
			COA:   "Principal Repayment",
			Value: principalrepayments,
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
	COA      string       `json:"COA,omitempty"`      //
	Value    IntStringMap `json:"Value,omitempty"`    //
	Selected bool         `json:"Selected,omitempty"` //
}

// IntStringMap - used for value field of TableJSON
type IntStringMap map[int]string

// BoolCOA -
type BoolCOA struct {
	MarketValue             bool `json:"MarketValue,omitempty"`             //
	TotalERV                bool `json:"TotalERV,omitempty"`                //
	OccupiedERV             bool `json:"OccupiedERV,omitempty"`             //
	VacantERV               bool `json:"VacantERV,omitempty"`               //
	TopSlice                bool `json:"TopSlice,omitempty"`                //
	TotalArea               bool `json:"TotalArea,omitempty"`               //
	OccupiedArea            bool `json:"OccupiedArea,omitempty"`            //
	VacantArea              bool `json:"VacantArea,omitempty"`              //
	PassingRent             bool `json:"PassingRent,omitempty"`             //
	Indexation              bool `json:"Indexation,omitempty"`              //
	TheoreticalRentalIncome bool `json:"TheoreticalRentalIncome,omitempty"` //
	BPUplift                bool `json:"BPUplift,omitempty"`                //
	Vacancy                 bool `json:"Vacancy,omitempty"`                 //
	ContractRent            bool `json:"ContractRent,omitempty"`            //
	RentFree                bool `json:"RentFree,omitempty"`                //
	TurnoverRent            bool `json:"TurnoverRent,omitempty"`            //
	MallRent                bool `json:"MallRent,omitempty"`                //
	ParkingIncome           bool `json:"ParkingIncome,omitempty"`           //
	OtherIncome             bool `json:"OtherIncome,omitempty"`             //
	OperatingIncome         bool `json:"OperatingIncome,omitempty"`         //
	OperatingExpenses       bool `json:"OperatingExpenses,omitempty"`       //
	NetOperatingIncome      bool `json:"NetOperatingIncome,omitempty"`      //
	Capex                   bool `json:"Capex,omitempty"`                   //
	AcqDispProperty         bool `json:"AcqDispProperty,omitempty"`         //
	AcqDispCosts            bool `json:"AcqDispCosts,omitempty"`            //
	LoanProceeds            bool `json:"LoanProceeds,omitempty"`            //
	InterestExpense         bool `json:"InterestExpense,omitempty"`         //
	PrincipalRepayment      bool `json:"PrincipalRepayment,omitempty"`      //
	LoanBalance             bool `json:"LoanBalance,omitempty"`             //
	Debt                    bool `json:"Debt,omitempty"`                    //
	Tax                     bool `json:"Tax,omitempty"`                     //
	TaxableIncome           bool `json:"TaxableIncome,omitempty"`           //
	TaxableIncomeCarryBack  bool `json:"TaxableIncomeCarryBack,omitempty"`  //
	DTA                     bool `json:"DTA,omitempty"`                     //
	Depreciation            bool `json:"Depreciation,omitempty"`            //
	Fees                    bool `json:"Fees,omitempty"`                    //
	NetCashFlow             bool `json:"NetCashFlow,omitempty"`             //
	CashBalance             bool `json:"CashBalance,omitempty"`             //
	BondIncome              bool `json:"BondIncome,omitempty"`              //
	BondExpense             bool `json:"BondExpense,omitempty"`             //
}

// FloatCOA -
type FloatCOA struct {
	MarketValue             float64 `json:"MarketValue,omitempty"`             //
	TotalERV                float64 `json:"TotalERV,omitempty"`                //
	OccupiedERV             float64 `json:"OccupiedERV,omitempty"`             //
	VacantERV               float64 `json:"VacantERV,omitempty"`               //
	TopSlice                float64 `json:"TopSlice,omitempty"`                //
	TotalArea               float64 `json:"TotalArea,omitempty"`               //
	OccupiedArea            float64 `json:"OccupiedArea,omitempty"`            //
	VacantArea              float64 `json:"VacantArea,omitempty"`              //
	PassingRent             float64 `json:"PassingRent,omitempty"`             //
	Indexation              float64 `json:"Indexation,omitempty"`              //
	TheoreticalRentalIncome float64 `json:"TheoreticalRentalIncome,omitempty"` //
	BPUplift                float64 `json:"BPUplift,omitempty"`                // // accounts for the effect of sold rent
	Vacancy                 float64 `json:"Vacancy,omitempty"`                 //
	ContractRent            float64 `json:"ContractRent,omitempty"`            //
	RentFree                float64 `json:"RentFree,omitempty"`                //
	TurnoverRent            float64 `json:"TurnoverRent,omitempty"`            //
	MallRent                float64 `json:"MallRent,omitempty"`                //
	ParkingIncome           float64 `json:"ParkingIncome,omitempty"`           //
	OtherIncome             float64 `json:"OtherIncome,omitempty"`             //
	OperatingIncome         float64 `json:"OperatingIncome,omitempty"`         //
	OperatingExpenses       float64 `json:"OperatingExpenses,omitempty"`       //
	NetOperatingIncome      float64 `json:"NetOperatingIncome,omitempty"`      //
	Capex                   float64 `json:"Capex,omitempty"`                   //
	AcqDispProperty         float64 `json:"AcqDispProperty,omitempty"`         //
	AcqDispCosts            float64 `json:"AcqDispCosts,omitempty"`            //
	LoanProceeds            float64 `json:"LoanProceeds,omitempty"`            //
	InterestExpense         float64 `json:"InterestExpense,omitempty"`         //
	PrincipalRepayment      float64 `json:"PrincipalRepayment,omitempty"`      //
	LoanBalance             float64 `json:"LoanBalance,omitempty"`             // Ending balance
	Debt                    float64 `json:"Debt,omitempty"`                    //
	Tax                     float64 `json:"Tax,omitempty"`                     //
	TaxableIncome           float64 `json:"TaxableIncome,omitempty"`           //
	TaxableIncomeCarryBack  float64 `json:"TaxableIncomeCarryBack,omitempty"`  //
	DTA                     float64 `json:"DTA,omitempty"`                     //
	Depreciation            float64 `json:"Depreciation,omitempty"`            //
	Fees                    float64 `json:"Fees,omitempty"`                    //
	NetCashFlow             float64 `json:"NetCashFlow,omitempty"`             //
	CashBalance             float64 `json:"CashBalance,omitempty"`             //
	BondIncome              float64 `json:"BondIncome,omitempty"`              //// income for me
	BondExpense             float64 `json:"BondExpense,omitempty"`             //// expense for me, or income for investor
}

// CreateTableHeader -
func (e *EntityModel) CreateTableHeader(month, quarter, year bool) {
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

// func (e *EntityModel) MakeRetoolTable() {
// 	e.RetoolTable = make([]interface{}, len(e.Table))
// 	for i, v := range e.Table {
// 		temp := make(map[string]interface{})
// 		temp[v.COA] = v.Value
// 		e.RetoolTable[i] = temp
// 	}
// 	// fmt.Println(e.RetoolTable)
// }
