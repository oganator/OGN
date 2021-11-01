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
	key := FundsList[name]
	Entities[key].CalculateFund()
	temp := make(map[interface{}]interface{})
	temp["entity"] = Entities[key]
	temp["modelslist"] = ModelsList
	temp["fundslist"] = FundsList
	c.TplName = "EntityView.tpl"
	c.Data = temp
}

// Post -
func (c *FundController) Post() {
	name := GetStringFund(c, "name")
	key := FundsList[name]
	Entities[key].CalculateFund()
	Entities[key].MCSetup.Sims = GetIntFund(c, "sims")
	Entities[key].FundMonteCarlo()
	temp := make(map[interface{}]interface{})
	temp["entity"] = Entities[key]
	temp["modelslist"] = ModelsList
	temp["fundslist"] = FundsList
	c.TplName = "EntityView.tpl"
	fmt.Println(Entities[key].Strategy)
	c.Data = temp
}
