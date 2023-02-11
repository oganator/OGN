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
	key := AssetModelsList[GetStringRentSchedule(c, "name")]
	indexstring := GetStringRentSchedule(c, "index")
	index, _ := strconv.Atoi(indexstring)
	unit := GetIntRentSchedule(c, "unit")
	temp["data"] = EntityModelsMap[key].EntityModel.ChildUnitModels[unit].RSStore
	if indexstring != "" {
		temp["data"] = EntityModelsMap[key].EntityModel.MCSlice[index].ChildUnitModels[unit].RSStore
	}
	// temp["baseURL"] = BaseURL
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
	modelkey := AssetModelsList[GetStringUnitCF(c, "name")]
	unit := GetIntUnitCF(c, "unit")
	tempentity := EntityModel{
		StartDate: EntityModelsMap[modelkey].EntityModel.StartDate,
		SalesDate: EntityModelsMap[modelkey].EntityModel.SalesDate,
		EndDate:   EntityModelsMap[modelkey].EntityModel.EndDate,
		COA:       EntityModelsMap[modelkey].EntityModel.ChildUnitModels[unit].COA,
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
	}, false, true)
	temp["entity"] = &tempentity
	// temp["baseURL"] = BaseURL
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
	key := AssetModelsList[GetStringUnitTable(c, "name")]
	index := GetIntUnitTable(c, "index")
	if index == -1 {
		temp["entity"] = EntityModelsMap[key].EntityModel
	} else {
		temp["entity"] = EntityModelsMap[key].EntityModel.MCSlice[index]
	}
	// temp["baseURL"] = BaseURL
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
	unit := UnitModelData{}
	parentmasterid := AssetModelsList[GetStringAddChildUnit(c, "parent")]
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
	EntityModelsMap[parentmasterid].Mutex.Lock()
	EntityModelsMap[parentmasterid].EntityModel.CalculateModel(false, "Internal")
	EntityModelsMap[parentmasterid].Mutex.Unlock()
	//
	c.TplName = "EntityView.tpl"
	temp["entity"] = EntityModelsMap[parentmasterid].EntityModel
	temp["modelslist"] = AssetModelsList
	temp["fundslist"] = FundModelsList
	// temp["baseURL"] = BaseURL
	c.Data = temp
}

//////////////////////////////////////////////////////////////////////

// UpdateUnitController -
type UpdateUnitController struct {
	beego.Controller
}

// GetStringAddChildUnit -
func GetStringUpdateUnit(c *UpdateUnitController, field string) string {
	c.Data[field] = c.GetString(field)
	return c.Data[field].(string)
}

// GetIntAddChildUnit -
func GetIntUpdateUnit(c *UpdateUnitController, field string) (result int) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.Atoi(temp)
	return result
}

// GetFloat -
func GetFloatUpdateUnit(c *UpdateUnitController, field string) (result float64) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.ParseFloat(temp, 64)
	return result
}

// Post - updates a single value for a single unit
func (c *UpdateUnitController) Post() {
	unit := GetIntUpdateUnit(c, "unit")
	field := GetStringUpdateUnit(c, "field")
	value := GetStringUpdateUnit(c, "value")
	WriteDBUnitModelSingleValue(unit, field, value)
	EntityModelsMap[Units[unit].Parent.MasterID].EntityModel.UpdateEntityModel()
	c.TplName = "test.tpl"
}
