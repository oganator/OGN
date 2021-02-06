package controllers

import (
	"strconv"

	beego "github.com/astaxie/beego"
)

// ViewRentScheduleController -
type ViewRentScheduleController struct {
	beego.Controller
}

// GetStringRentSchedule -
func GetStringRentSchedule(c *ViewRentScheduleController, field string) string {
	c.Data[field] = c.GetString(field)
	return c.Data[field].(string)
}

// GetIntRentSchedule -
func GetIntRentSchedule(c *ViewRentScheduleController, field string) (result int) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.Atoi(temp)
	return result
}

// Post -
func (c *ViewRentScheduleController) Post() {
	temp := make(map[interface{}]interface{})
	key := ModelsList[GetStringRentSchedule(c, "name")]
	unit := GetIntRentSchedule(c, "unit")
	temp["data"] = Models[key].ChildUnits[unit].RSStore
	c.TplName = "RentSchedule.tpl"
	c.Data = temp
}

////////////////////////////////////////////////////////////////////////

// ViewUnitCFController -
type ViewUnitCFController struct {
	beego.Controller
}

// GetStringUnitCF -
func GetStringUnitCF(c *ViewUnitCFController, field string) string {
	c.Data[field] = c.GetString(field)
	return c.Data[field].(string)
}

// GetIntUnitCF -
func GetIntUnitCF(c *ViewUnitCFController, field string) (result int) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.Atoi(temp)
	return result
}

// Post -
func (c *ViewUnitCFController) Post() {
	temp := make(map[interface{}]interface{})
	modelkey := ModelsList[GetStringUnitCF(c, "name")]
	unit := GetIntUnitCF(c, "unit")
	tempentity := Entity{
		StartDate: Models[modelkey].StartDate,
		SalesDate: Models[modelkey].SalesDate,
		EndDate:   Models[modelkey].EndDate,
		COA:       Models[modelkey].ChildUnits[unit].COA,
	}
	tempentity.SumCOA()
	tempentity.MakeTable(BoolCOA{
		TotalERV:                true,
		TotalArea:               true,
		PassingRent:             true,
		Indexation:              true,
		TheoreticalRentalIncome: true,
		Vacancy:                 true,
		ContractRent:            true,
		RentFree:                true,
		TurnoverRent:            true,
		MallRent:                true,
		ParkingIncome:           true,
		OtherIncome:             true,
		OperatingIncome:         true,
		Capex:                   true,
	})
	temp["entity"] = &tempentity
	c.TplName = "CFTable.tpl"
	c.Data = temp
}
