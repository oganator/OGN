package controllers

import (
	"fmt"
)

// MakeTable - for timeframe there are two options:
//   - 0 indicates yearly values for all years
//   - a 4 digit integer indicates monthly values for a specific year
//
// table is map[coa]map[date]value
func (e *EntityModel) MakeTable2(coas BoolCOA, timeframe int) (table []map[string]string) {
	table = make([]map[string]string, 0)
	marketvalue := make(map[string]string)
	marketvalue["0"] = "Market Value"
	totalerv := make(map[string]string)
	totalerv["0"] = "Total ERV"
	occupiederv := make(map[string]string)
	occupiederv["0"] = "Occupied ERV"
	vacanterv := make(map[string]string)
	vacanterv["0"] = "Vacant ERV"
	topslice := make(map[string]string)
	topslice["0"] = "Top Slice"
	totalarea := make(map[string]string)
	totalarea["0"] = "Total Area"
	occupiedarea := make(map[string]string)
	occupiedarea["0"] = "Occupied Area"
	vacantarea := make(map[string]string)
	vacantarea["0"] = "Vacant Area"
	blank := make(map[string]string)
	blank["0"] = ""
	netoperatingincome := make(map[string]string)
	netoperatingincome["0"] = "Net Operating Income"
	capex := make(map[string]string)
	capex["0"] = "Capex"
	acqdispproperty := make(map[string]string)
	acqdispproperty["0"] = "Acquisitions"
	debt := make(map[string]string)
	debt["0"] = "Debt"
	acqdispcosts := make(map[string]string)
	acqdispcosts["0"] = "Acquisition Costs"
	tax := make(map[string]string)
	tax["0"] = "Tax"
	taxableincome := make(map[string]string)
	taxableincome["0"] = "Taxable Income"
	taxableincomecarryback := make(map[string]string)
	taxableincomecarryback["0"] = "Taxable Income Carry Back"
	dta := make(map[string]string)
	dta["0"] = "Deferred Tax Asset"
	depreciation := make(map[string]string)
	depreciation["0"] = "Depreciation"
	fees := make(map[string]string)
	fees["0"] = "Fees"
	operatingincome := make(map[string]string)
	operatingincome["0"] = "Operating Income"
	operatingexpenses := make(map[string]string)
	operatingexpenses["0"] = "Operating Expenses"
	contractrent := make(map[string]string)
	contractrent["0"] = "Contract Rent"
	rentfree := make(map[string]string)
	rentfree["0"] = "Rent Free"
	turnoverrent := make(map[string]string)
	turnoverrent["0"] = "Turnover Rent"
	mallrent := make(map[string]string)
	mallrent["0"] = "Mall Rent"
	parkingincome := make(map[string]string)
	parkingincome["0"] = "Parking Income"
	otherincome := make(map[string]string)
	otherincome["0"] = "Other Income"
	theoreticalrentalincome := make(map[string]string)
	theoreticalrentalincome["0"] = "Theoretical Rental Income"
	vacancy := make(map[string]string)
	vacancy["0"] = "Vacancy"
	passingrent := make(map[string]string)
	passingrent["0"] = "Passing Rent"
	indexation := make(map[string]string)
	indexation["0"] = "Indexation"
	bpuplift := make(map[string]string)
	bpuplift["0"] = "BP Uplift"
	loanproceeds := make(map[string]string)
	loanproceeds["0"] = "Loan Proceeds"
	interestexpense := make(map[string]string)
	interestexpense["0"] = "Interest Expense"
	loanbalance := make(map[string]string)
	loanbalance["0"] = "Loan Balance"
	netcashflow := make(map[string]string)
	netcashflow["0"] = "Net Cash Flow"
	cashbalance := make(map[string]string)
	cashbalance["0"] = "Cash Balance"
	bondincome := make(map[string]string)
	bondincome["0"] = "Bond Income"
	bondexpense := make(map[string]string)
	bondexpense["0"] = "Bond Expense"
	startdate := Datetype{Month: 1, Year: timeframe}
	salesdate := Datetype{Month: 12, Year: timeframe}
	if timeframe == 0 {
		startdate.Year = e.StartDate.Year
		salesdate.Year = e.SalesDate.Year
	} else {
		startdate.Add(1)
	}
	startdate.Add(0)
	salesdate.Add(0)
	// loops through the months, if it is the start of the year then the yearly logic executes, then january thereafter
	for date := Dateadd(startdate, -1); date.Dateint <= salesdate.Dateint; date = Dateadd(date, 1) { //Monthly Loop
		if date.Dateint == startdate.Dateint || date.Month == 1 || date == Dateadd(startdate, -1) {
			if timeframe == 0 {

				// Yearly Loop
				// YEARLY
				if coas.MarketValue {
					marketvalue[fmt.Sprint(date.Year)] = fmt.Sprint(e.COA[date.Year].MarketValue)
				}
				if coas.TotalERV {
					totalerv[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].TotalERV)
				}
				if coas.OccupiedERV {
					occupiederv[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].OccupiedERV)
				}
				if coas.VacantERV {
					vacanterv[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].VacantERV)
				}
				if coas.TopSlice {
					topslice[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].TopSlice)
				}
				if coas.TotalArea {
					totalarea[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].TotalArea)
				}
				if coas.OccupiedArea {
					occupiedarea[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].OccupiedArea)
				}
				if coas.VacantArea {
					vacantarea[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].VacantArea)
				}
				blank[fmt.Sprint(date.Year)] = ""
				if coas.NetOperatingIncome {
					netoperatingincome[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].NetOperatingIncome)
				}
				if coas.Capex {
					capex[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].Capex)
				}
				if coas.AcqDispProperty {
					acqdispproperty[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].AcqDispProperty)
				}
				if coas.Debt {
					debt[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].Debt)
				}
				if coas.AcqDispCosts {
					acqdispcosts[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].AcqDispCosts)
				}
				if coas.Tax {
					tax[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].Tax)
				}
				if coas.TaxableIncome {
					taxableincome[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].TaxableIncome)
				}
				if coas.TaxableIncomeCarryBack {
					taxableincomecarryback[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].TaxableIncomeCarryBack)
				}
				if coas.DTA {
					dta[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].DTA)
				}
				if coas.Depreciation {
					depreciation[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].Depreciation)
				}
				if coas.Fees {
					fees[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].Fees)
				}
				if coas.OperatingIncome {
					operatingincome[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].OperatingIncome)
				}
				if coas.OperatingExpenses {
					operatingexpenses[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].OperatingExpenses)
				}
				if coas.ContractRent {
					contractrent[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].ContractRent)
				}
				if coas.RentFree {
					rentfree[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].RentFree)
				}
				if coas.TurnoverRent {
					turnoverrent[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].TurnoverRent)
				}
				if coas.MallRent {
					mallrent[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].MallRent)
				}
				if coas.ParkingIncome {
					parkingincome[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].ParkingIncome)
				}
				if coas.OtherIncome {
					otherincome[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].OtherIncome)
				}
				if coas.TheoreticalRentalIncome {
					theoreticalrentalincome[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].TheoreticalRentalIncome)
				}
				if coas.Vacancy {
					vacancy[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].Vacancy)
				}
				if coas.PassingRent {
					passingrent[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].PassingRent)
				}
				if coas.Indexation {
					indexation[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].Indexation)
				}
				if coas.BPUplift {
					bpuplift[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].BPUplift)
				}
				if coas.LoanProceeds {
					loanproceeds[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].LoanProceeds)
				}
				if coas.InterestExpense {
					interestexpense[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].InterestExpense)
				}
				if coas.LoanBalance {
					loanbalance[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].LoanBalance)
				}
				if coas.NetCashFlow {
					netcashflow[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].NetCashFlow)
				}
				blank[fmt.Sprint(date.Year)] = ""
				if coas.CashBalance {
					cashbalance[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].CashBalance)
				}
				if coas.BondIncome {
					bondincome[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].BondIncome)
				}
				if coas.BondExpense {
					bondexpense[fmt.Sprint(date.Year)] = RenderFloat("#,###.", e.COA[date.Year].BondExpense)
				}
			}
		}
		// MONTHLY
		if timeframe != 0 {
			if coas.MarketValue {
				marketvalue[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].MarketValue)
			}
			if coas.TotalERV {
				totalerv[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].TotalERV)
			}
			if coas.OccupiedERV {
				occupiederv[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].OccupiedERV)
			}
			if coas.VacantERV {
				vacanterv[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].VacantERV)
			}
			if coas.TopSlice {
				topslice[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].TopSlice)
			}
			if coas.TotalArea {
				totalarea[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].TotalArea)
			}
			if coas.OccupiedERV {
				occupiedarea[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].OccupiedArea)
			}
			if coas.VacantERV {
				vacantarea[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].VacantArea)
			}
			blank[fmt.Sprint(date.Dateint)] = ""
			if coas.NetOperatingIncome {
				netoperatingincome[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].NetOperatingIncome)
			}
			if coas.Capex {
				capex[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].Capex)
			}
			if coas.AcqDispProperty {
				acqdispproperty[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].AcqDispProperty)
			}
			if coas.Debt {
				debt[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].Debt)
			}
			if coas.AcqDispCosts {
				acqdispcosts[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].AcqDispCosts)
			}
			if coas.Tax {
				tax[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].Tax)
			}
			if coas.TaxableIncome {
				taxableincome[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].TaxableIncome)
			}
			if coas.TaxableIncomeCarryBack {
				taxableincomecarryback[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].TaxableIncomeCarryBack)
			}
			if coas.DTA {
				dta[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].DTA)
			}
			if coas.Depreciation {
				depreciation[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].Depreciation)
			}
			if coas.Fees {
				fees[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].Fees)
			}
			if coas.OperatingIncome {
				operatingincome[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].OperatingIncome)
			}
			if coas.OperatingExpenses {
				operatingexpenses[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].OperatingExpenses)
			}
			if coas.ContractRent {
				contractrent[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].ContractRent)
			}
			if coas.RentFree {
				rentfree[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].RentFree)
			}
			if coas.TurnoverRent {
				turnoverrent[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].TurnoverRent)
			}
			if coas.MallRent {
				mallrent[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].MallRent)
			}
			if coas.ParkingIncome {
				parkingincome[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].ParkingIncome)
			}
			if coas.OtherIncome {
				otherincome[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].OtherIncome)
			}
			if coas.TheoreticalRentalIncome {
				theoreticalrentalincome[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].TheoreticalRentalIncome)
			}
			if coas.Vacancy {
				vacancy[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].Vacancy)
			}
			if coas.PassingRent {
				passingrent[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].PassingRent)
			}
			if coas.Indexation {
				indexation[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].Indexation)
			}
			if coas.BPUplift {
				bpuplift[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].BPUplift)
			}
			if coas.LoanProceeds {
				loanproceeds[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].LoanProceeds)
			}
			if coas.InterestExpense {
				interestexpense[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].InterestExpense)
			}
			if coas.LoanBalance {
				loanbalance[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].LoanBalance)
			}
			if coas.NetCashFlow {
				netcashflow[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].NetCashFlow)
			}
			blank[fmt.Sprint(date.Dateint)] = ""
			if coas.CashBalance {
				cashbalance[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].CashBalance)
			}
			if coas.BondIncome {
				bondincome[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].BondIncome)
			}
			if coas.BondExpense {
				bondexpense[fmt.Sprint(date.Dateint)] = RenderFloat("#,###.", e.COA[date.Dateint].BondExpense)
			}
		}
	}
	// build e.Table based on COA selection
	if coas.MarketValue == true {
		table = append(table, marketvalue)
	}
	if coas.TotalERV == true {
		table = append(table, totalerv)
	}
	if coas.OccupiedERV == true {
		table = append(table, occupiederv)
	}
	if coas.VacantERV == true {
		table = append(table, vacanterv)
	}
	if coas.TopSlice == true {
		table = append(table, topslice)
	}
	if coas.TotalArea == true {
		table = append(table, totalarea)
	}
	if coas.OccupiedArea == true {
		table = append(table, occupiedarea)
	}
	if coas.VacantArea == true {
		table = append(table, vacantarea)
	}
	table = append(table, blank)
	if coas.PassingRent == true {
		table = append(table, passingrent)
	}
	if coas.Indexation == true {
		table = append(table, indexation)
	}
	if coas.TheoreticalRentalIncome == true {
		table = append(table, theoreticalrentalincome)
	}
	if coas.BPUplift == true {
		table = append(table, bpuplift)
	}
	if coas.Vacancy == true {
		table = append(table, vacancy)
	}
	if coas.ContractRent == true {
		table = append(table, contractrent)
	}
	if coas.RentFree == true {
		table = append(table, rentfree)
	}
	if coas.TurnoverRent == true {
		table = append(table, turnoverrent)
	}
	if coas.MallRent == true {
		table = append(table, mallrent)
	}
	if coas.ParkingIncome == true {
		table = append(table, parkingincome)
	}
	if coas.OtherIncome == true {
		table = append(table, otherincome)
	}
	if coas.OperatingIncome == true {
		table = append(table, operatingincome)
	}
	if coas.OperatingExpenses == true {
		table = append(table, operatingexpenses)
	}
	if coas.NetOperatingIncome == true {
		table = append(table, netoperatingincome)
	}
	if coas.Capex == true {
		table = append(table, capex)
	}
	if coas.AcqDispProperty == true {
		table = append(table, acqdispproperty)
	}
	if coas.AcqDispCosts == true {
		table = append(table, acqdispcosts)
	}
	if coas.LoanProceeds == true {
		table = append(table, loanproceeds)
	}
	if coas.InterestExpense == true {
		table = append(table, interestexpense)
	}
	if coas.LoanBalance == true {
		table = append(table, loanbalance)
	}
	if coas.Debt == true {
		table = append(table, debt)
	}
	if coas.Tax == true {
		table = append(table, tax)
	}
	if coas.TaxableIncome == true {
		table = append(table, taxableincome)
	}
	if coas.TaxableIncomeCarryBack == true {
		table = append(table, taxableincomecarryback)
	}
	if coas.DTA == true {
		table = append(table, dta)
	}
	if coas.Depreciation == true {
		table = append(table, depreciation)
	}
	if coas.Fees == true {
		table = append(table, fees)
	}
	if coas.NetCashFlow == true {
		table = append(table, netcashflow)
	}
	table = append(table, blank)
	if coas.CashBalance == true {
		table = append(table, cashbalance)
	}
	if coas.BondIncome == true {
		table = append(table, bondincome)
	}
	if coas.BondExpense == true {
		table = append(table, bondexpense)
	}
	return table
}
