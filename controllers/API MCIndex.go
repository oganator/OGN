package controllers

import (
	"strconv"

	beego "github.com/astaxie/beego"
)

// MCIndexController -
type MCIndexController struct {
	beego.Controller
}

// GetStringMCIndex -
func GetStringMCIndex(c *MCIndexController, field string) string {
	c.Data[field] = c.GetString(field)
	return c.Data[field].(string)
}

// GetIntMCIndex -
func GetIntMCIndex(c *MCIndexController, field string) (result int) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.Atoi(temp)
	return result
}

// Post -
func (c *MCIndexController) Post() {
	tempkey := ModelsList[GetStringMCIndex(c, "name")]
	index := GetIntMCIndex(c, "index")
	temp := make(map[interface{}]interface{})
	coas := BoolCOA{
		MarketValue:             true,
		TotalERV:                true,
		PassingRent:             true,
		Indexation:              true,
		TheoreticalRentalIncome: true,
		BPUplift:                true,
		Vacancy:                 true,
		ContractRent:            true,
		OperatingExpenses:       true,
		NetOperatingIncome:      true,
		AcqDispProperty:         true,
		InterestExpense:         true,
		Tax:                     true,
		Fees:                    true,
		Capex:                   true,
		NetCashFlow:             true,
		CashBalance:             true,
		BondIncome:              true,
		BondExpense:             true,
	}
	if Entities[tempkey].Strategy == "Standard" {
		coas.BPUplift = false
		coas.BondExpense = false
		coas.BondIncome = false
		coas.Debt = true
	}
	Entities[tempkey].MCSlice[index].MakeTable(coas, false, false, true)
	temp["entity"] = Entities[tempkey].MCSlice[index]
	temp["baseURL"] = BaseURL
	c.TplName = "CFTable.tpl"
	c.Data = temp
}
