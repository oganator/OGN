package controllers

import (
	"fmt"
	"strconv"

	beego "github.com/astaxie/beego"
)

type FundController struct {
	beego.Controller
}

// GetStringFund -
func GetStringFund(c *FundController, field string) string {
	c.Data[field] = c.GetString(field)
	return c.Data[field].(string)
}

// GetIntAddChildUnit -
func GetIntFund(c *FundController, field string) (result int) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.Atoi(temp)
	return result
}

// GetFloat -
func GetFloatFund(c *FundController, field string) (result float64) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.ParseFloat(temp, 64)
	return result
}

// Get -
func (c *FundController) Get() {
	name := GetStringFund(c, "name")
	key := FundModelsList[name]
	temp := make(map[interface{}]interface{})
	temp["entity"] = EntityModelsMap[key].EntityModel
	temp["modelslist"] = AssetModelsList
	temp["fundslist"] = FundModelsList
	c.TplName = "EntityView.tpl"
	fmt.Println("Fund Controller - ", EntityModelsMap[key].EntityModel.Name, " - ", EntityModelsMap[key].EntityModel.Metrics.IRR.NetLeveredAfterTax)
	c.Data = temp
}

// Post -
func (c *FundController) Post() {
	name := GetStringFund(c, "name")
	key := FundModelsList[name]
	EntityModelsMap[key].EntityModel.Strategy = GetStringFund(c, "strategy")
	EntityModelsMap[key].EntityModel.CalculateFund()
	EntityModelsMap[key].EntityModel.MCSetup.Sims = GetIntFund(c, "sims")
	EntityModelsMap[key].EntityModel.FundMonteCarlo()
	temp := make(map[interface{}]interface{})
	temp["entity"] = EntityModelsMap[key].EntityModel
	temp["modelslist"] = AssetModelsList
	temp["fundslist"] = FundModelsList
	c.TplName = "EntityView.tpl"
	c.Data = temp
	fmt.Println("Fund Controller - ", EntityModelsMap[key].EntityModel.Name, " - ", EntityModelsMap[key].EntityModel.Metrics.IRR.NetLeveredAfterTax)
	// go debug.FreeOSMemory()
}
