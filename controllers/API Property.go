package controllers

import (
	"strconv"

	beego "github.com/astaxie/beego"
)

// ViewEntityController - default route - "/". Executed from submit button
type AddCostController struct {
	beego.Controller
}

// GetString -
func GetCostString(c *AddCostController, field string) string {
	c.Data[field] = c.GetString(field)
	return c.Data[field].(string)
}

// GetFloat -
func GetCostFloat(c *AddCostController, field string) (result float64) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.ParseFloat(temp, 64)
	return result
}

// GetInt -
func GetCostInt(c *AddCostController, field string) (result int) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.Atoi(temp)
	return result
}

// Post -
func (c *AddCostController) Post() {
	entityModel := EntityModelsList[GetCostString(c, "entityModel")]
	unitModel := GetCostInt(c, "unitModel")
	tempCost := AddCostInput(entityModel, unitModel)
	EntityModelsMap[entityModel].EntityModel.CostInput[tempCost.MasterID] = tempCost
	c.TplName = "test.tpl"
}
