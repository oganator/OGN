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
	indexstring := GetStringRentSchedule(c, "index")
	index, _ := strconv.Atoi(indexstring)
	unit := GetIntRentSchedule(c, "unit")
	temp["data"] = Entities[key].ChildUnits[unit].RSStore
	if indexstring != "" {
		temp["data"] = Entities[key].MCSlice[index].ChildUnits[unit].RSStore
	}
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
		StartDate: Entities[modelkey].StartDate,
		SalesDate: Entities[modelkey].SalesDate,
		EndDate:   Entities[modelkey].EndDate,
		COA:       Entities[modelkey].ChildUnits[unit].COA,
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
		OperatingIncome:         true,
		Capex:                   true,
		NetCashFlow:             true,
	})
	temp["entity"] = &tempentity
	c.TplName = "CFTable.tpl"
	c.Data = temp
}

////////////////////////////////////////////////////////////////////////

// ViewUnitTableController -
type ViewUnitTableController struct {
	beego.Controller
}

// GetStringUnitTable -
func GetStringUnitTable(c *ViewUnitTableController, field string) string {
	c.Data[field] = c.GetString(field)
	return c.Data[field].(string)
}

// GetIntUnitTable -
func GetIntUnitTable(c *ViewUnitTableController, field string) (result int) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.Atoi(temp)
	return result
}

// Post -
func (c *ViewUnitTableController) Post() {
	temp := make(map[interface{}]interface{})
	key := ModelsList[GetStringUnitTable(c, "name")]
	index := GetIntUnitTable(c, "index")
	if index == -1 {
		temp["entity"] = Entities[key]
	} else {
		temp["entity"] = Entities[key].MCSlice[index]
	}
	c.TplName = "UnitTable.tpl"
	c.Data = temp
}

//////////////////////////////////////////////////////////////////////

// AddChildUnitController -
type AddChildUnitController struct {
	beego.Controller
}

// GetStringAddChildUnit -
func GetStringAddChildUnit(c *AddChildUnitController, field string) string {
	c.Data[field] = c.GetString(field)
	return c.Data[field].(string)
}

// GetIntAddChildUnit -
func GetIntAddChildUnit(c *AddChildUnitController, field string) (result int) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.Atoi(temp)
	return result
}

// GetFloat -
func GetFloatAddChildUnit(c *AddChildUnitController, field string) (result float64) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.ParseFloat(temp, 64)
	return result
}

// Post -
func (c *AddChildUnitController) Post() {
	temp := make(map[interface{}]interface{})
	unit := UnitData{}
	parentmasterid := ModelsList[GetStringAddChildUnit(c, "parent")]
	unit.ParentMasterID = parentmasterid
	unit.Name = GetStringAddChildUnit(c, "unitname")
	unit.Tenant = GetStringAddChildUnit(c, "tenant")
	if unit.Tenant == "" {
		unit.UnitStatus = "Vacant"
	} else {
		unit.UnitStatus = "Occupied"
	}
	unit.PassingRent = GetFloatAddChildUnit(c, "rent")
	unit.LeaseStartMonth = GetIntAddChildUnit(c, "startmonth")
	unit.LeaseStartYear = GetIntAddChildUnit(c, "startyear")
	unit.LeaseEndMonth = GetIntAddChildUnit(c, "expirymonth")
	unit.LeaseEndYear = GetIntAddChildUnit(c, "expiryyear")
	unit.ERVAmount = GetFloatAddChildUnit(c, "amount")
	unit.ERVArea = GetFloatAddChildUnit(c, "area")
	UnitStore[unit.MasterID] = unit
	unit.WriteXLSXUnits()
	// Associations()
	Entities[parentmasterid].CalculateModel(false)
	//
	c.TplName = "EntityView.tpl"
	temp["entity"] = Entities[parentmasterid]
	temp["modelslist"] = ModelsList
	c.Data = temp
}
