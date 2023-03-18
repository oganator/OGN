package controllers

import (
	"fmt"
	"strings"

	_ "modernc.org/sqlite"

	beego "github.com/astaxie/beego"
)

// ViewEntityController - default route - "/". Executed from submit button
type ChartController struct {
	beego.Controller
}

// GetString -
func GetChartString(c *ChartController, field string) string {
	c.Data[field] = c.GetString(field)
	return c.Data[field].(string)
}

func (c *ChartController) Post() {
	selections := GetChartString(c, "coa")
	entityModel := GetChartString(c, "entityModel")
	temp := make(map[interface{}]interface{})
	entityModelInt := EntityModelsList[entityModel]
	temp["coaArray"] = EntityModelsMap[entityModelInt].EntityModel.FormatChartData(selections)
	c.TplName = "Chart.tpl"
	c.Data = temp
}

func (e *EntityModel) FormatChartData(selections string) (coaArray []XYFloatSlice) {
	selectionArray := strings.Split(selections, ",")
	//
	x := make([]float64, 0)
	for year := e.StartDate.Year; year <= e.SalesDate.Year; year++ {
		x = append(x, float64(year))
	}
	//
	if stringInSlice("Market Value", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{MarketValue: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		fmt.Println(temp.Y)
		temp.X = x
		temp.Name = "Market Value"
		temp.Index = `marketValue`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Total ERV", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{TotalERV: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Total ERV"
		temp.Index = `totalErv`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Occupied ERV", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{OccupiedERV: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Occupied ERV"
		temp.Index = `occupiedERV`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Vacant ERV", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{VacantERV: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Vacant ERV"
		temp.Index = `vacantERV`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Top Slice", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{TopSlice: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Top Slice"
		temp.Index = `topSlice`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Total Area", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{TotalArea: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Total Area"
		temp.Index = `totalArea`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Occupied Area", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{OccupiedArea: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Occupied Area"
		temp.Index = `occupiedArea`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Vacant Area", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{VacantArea: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Vacant Area"
		temp.Index = `vacantArea`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Passing Rent", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{PassingRent: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Passing Rent"
		temp.Index = `passingRent`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Indexation", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{Indexation: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Indexation"
		temp.Index = `Indexation`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Theoretical Rental Income", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{TheoreticalRentalIncome: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Theoretical Rental Income"
		temp.Index = `theoreticalRentalIncome`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("BP Uplift", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{BPUplift: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "BP Uplift"
		temp.Index = `bpUplift`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Vacancy", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{Vacancy: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Vacancy"
		temp.Index = `vacancy`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Contract Rent", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{ContractRent: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Contract Rent"
		temp.Index = `contractRent`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Rent Free", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{RentFree: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Rent Free"
		temp.Index = `rentFree`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Turnover Rent", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{TurnoverRent: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Turnover Rent"
		temp.Index = `turnoverRent`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Mall Rent", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{MallRent: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Mall Rent"
		temp.Index = `mallRent`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Parking Income", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{ParkingIncome: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Parking Income"
		temp.Index = `parkingIncome`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Other Income", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{OtherIncome: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Other Income"
		temp.Index = `otherIncome`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Operating Income", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{OperatingIncome: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Operating Income"
		temp.Index = `operatingIncome`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Operating Expenses", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{OperatingExpenses: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Operating Expenses"
		temp.Index = `operatingExpenses`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Net Operating Income", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{NetOperatingIncome: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Net Operating Income"
		temp.Index = `netOperatingIncome`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Capex", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{Capex: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Capex"
		temp.Index = `capex`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Acq/Disp Property", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{AcqDispProperty: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Acq/Disp Property"
		temp.Index = `acqDispProperty`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Acq/Disp Costs", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{AcqDispCosts: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Acq/Disp Costs"
		temp.Index = `acqDispCosts`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Loan Proceeds", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{LoanProceeds: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Loan Proceeds"
		temp.Index = `loanProceeds`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Interest Expense", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{InterestExpense: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Interest Expense"
		temp.Index = `interestExpense`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Principal Repayment", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{PrincipalRepayment: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Principal Repayment"
		temp.Index = `principalRepayment`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Loan Balance", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{LoanBalance: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Loan Balance"
		temp.Index = `loanBalance`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Debt", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{Debt: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Debt"
		temp.Index = `debt`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Tax", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{Tax: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Tax"
		temp.Index = `tax`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Taxable Income", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{TaxableIncome: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Taxable Income"
		temp.Index = `taxableIncome`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Taxable Income Carryback", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{TaxableIncomeCarryBack: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Taxable Income Carryback"
		temp.Index = `taxableIncomeCarryback`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("DTA", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{DTA: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "DTA"
		temp.Index = `dta`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Depreciation", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{Depreciation: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Depreciation"
		temp.Index = `depreciation`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Fees", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{Fees: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Fees"
		temp.Index = `fees`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Net Cash Flow", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{NetCashFlow: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Net Cash Flow"
		temp.Index = `netCashFlow`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Cash Balance", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{CashBalance: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Cash Balance"
		temp.Index = `cashBalance`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Bond Income", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{BondIncome: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Bond Income"
		temp.Index = `bondIncome`
		coaArray = append(coaArray, temp)
	}
	if stringInSlice("Bond Expense", selectionArray) {
		temp := XYFloatSlice{}
		temp.Y = ReturnCOAArray(e.COA, FloatCOA{BondExpense: 1.0}, e.StartDate, Dateadd(e.SalesDate, 12), true)
		temp.X = x
		temp.Name = "Bond Expense"
		temp.Index = `bondExpense`
		coaArray = append(coaArray, temp)
	}

	return coaArray
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
