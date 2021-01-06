package controllers

import (
	"math"
	"strconv"

	beego "github.com/astaxie/beego"
)

// ViewEntityController -
type ViewEntityController struct {
	beego.Controller
}

// ChanveEntityController -
type ChangeEntityController struct {
	beego.Controller
}

// ViewUnitController -
type ViewUnitController struct {
	beego.Controller
}

// Get -
func (c *ViewEntityController) Get() {
	temp := make(map[interface{}]interface{})
	temp["entity"] = Models[Key]
	temp["modelslist"] = ModelsList
	c.TplName = "EntityView.tpl"
	c.Data = temp
}

// Post -
func (c *ViewEntityController) Post() {
	// select model
	c.Data["name"] = c.GetString("name")
	name := c.Data["name"].(string)
	Key = ModelsList[name]
	// hold period
	c.Data["holdperiod"] = c.GetString("holdperiod")
	holdperiod := c.Data["holdperiod"].(string)
	EntityStore[Key].HoldPeriod, _ = strconv.Atoi(holdperiod)
	// Entry Yield
	c.Data["entryyield"] = c.GetString("entryyield")
	entryyield := c.Data["entryyield"].(string)
	EntityStore[Key].EntryYield, _ = strconv.ParseFloat(entryyield, 64)
	EntityStore[Key].EntryYield = EntityStore[Key].EntryYield / 100
	// Yield Shift
	c.Data["yieldshift"] = c.GetString("yieldshift")
	yieldshift := c.Data["yieldshift"].(string)
	EntityStore[Key].YieldShift, _ = strconv.ParseFloat(yieldshift, 64)
	// LTV
	c.Data["ltv"] = c.GetString("ltv")
	ltv := c.Data["ltv"].(string)
	EntityStore[Key].LTV, _ = strconv.ParseFloat(ltv, 64)
	// rate
	c.Data["rate"] = c.GetString("rate")
	rate := c.Data["rate"].(string)
	EntityStore[Key].LoanRate, _ = strconv.ParseFloat(rate, 64)
	// Discount rate
	c.Data["discount"] = c.GetString("discount")
	discount := c.Data["discount"].(string)
	EntityStore[Key].GLA.DiscountRate, _ = strconv.ParseFloat(discount, 64)
	EntityStore[Key].GLA.DiscountRate = EntityStore[Key].GLA.DiscountRate / 100
	// sold rent
	c.Data["soldrent"] = c.GetString("soldrent")
	soldrent := c.Data["soldrent"].(string)
	EntityStore[Key].GLA.PercentSoldRent, _ = strconv.ParseFloat(soldrent, 64)
	EntityStore[Key].GLA.PercentSoldRent = EntityStore[Key].GLA.PercentSoldRent / 100
	// strategy
	c.Data["strategy"] = c.GetString("strategy")
	strategy := c.Data["strategy"].(string)
	EntityStore[Key].Strategy = strategy
	// balloon percent
	c.Data["balpercent"] = c.GetString("balpercent")
	balpercent := c.Data["balpercent"].(string)
	EntityStore[Key].BalloonPercent, _ = strconv.ParseFloat(balpercent, 64)
	EntityStore[Key].BalloonPercent = EntityStore[Key].BalloonPercent / 100
	// erv
	c.Data["erv"] = c.GetString("erv")
	erv := c.Data["erv"].(string)
	EntityStore[Key].ERVGrowth, _ = strconv.ParseFloat(erv, 64)
	EntityStore[Key].ERVGrowth = EntityStore[Key].ERVGrowth / 100
	// cpi
	c.Data["cpi"] = c.GetString("cpi")
	cpi := c.Data["cpi"].(string)
	EntityStore[Key].CPIGrowth, _ = strconv.ParseFloat(cpi, 64)
	EntityStore[Key].CPIGrowth = EntityStore[Key].CPIGrowth / 100
	// opex
	c.Data["opex"] = c.GetString("opex")
	opex := c.Data["opex"].(string)
	EntityStore[Key].OpExpercent, _ = strconv.ParseFloat(opex, 64)
	EntityStore[Key].OpExpercent = EntityStore[Key].OpExpercent / 100
	// fees
	c.Data["fees"] = c.GetString("fees")
	fees := c.Data["fees"].(string)
	EntityStore[Key].Fees, _ = strconv.ParseFloat(fees, 64)
	//
	c.TplName = "EntityView.tpl"
	// c.TplName = "test.tpl"
	temp := make(map[interface{}]interface{})
	// wg := sync.WaitGroup{}
	// for i := 0; i < 1000; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		Models[Key].UpdateEntity()
	// 	}()
	// }
	// wg.Wait()
	Models[Key].UpdateEntity()
	Models[Key].MultiplyInputs()
	temp["entity"] = Models[Key]
	temp["modelslist"] = ModelsList
	c.Data = temp
}

// MultiplyInputs -
func (e *Entity) MultiplyInputs() {
	e.Valuation.EntryYield = math.Round(e.Valuation.EntryYield*100000) / 1000
	e.Valuation.ExitYield = math.Round(e.Valuation.ExitYield*100000) / 1000
	e.GLA.DiscountRate = math.Round(e.GLA.DiscountRate*100000) / 1000
	e.GLA.PercentSoldRent = math.Round(e.GLA.PercentSoldRent*100000) / 1000
	e.BalloonPercent = math.Round(e.BalloonPercent*100000) / 1000
	e.GrowthInput["ERV"] = math.Round(e.GrowthInput["ERV"]*100000) / 1000
	e.GrowthInput["CPI"] = math.Round(e.GrowthInput["CPI"]*100000) / 1000
	e.OpEx.PercentOfTRI = math.Round(e.OpEx.PercentOfTRI*100000) / 1000
}

// Post -
func (c *ChangeEntityController) Post() {
	// select model
	c.Data["modelname"] = c.GetString("modelname")
	name := c.Data["modelname"].(string)
	Key = ModelsList[name]
	//
	c.TplName = "EntityView.tpl"
	temp := make(map[interface{}]interface{})
	temp["entity"] = Models[Key]
	temp["modelslist"] = ModelsList
	c.Data = temp
}

// Post -
// func (c *ViewUnitController) Post() {
// 	// parent
// 	c.Data["parent"] = c.GetString("parent")
// 	parent := c.Data["parent"].(string)
// 	pkey := EntityStore.Keys[parent]
// 	// unit
// 	c.Data["name"] = c.GetString("name")
// 	name := c.Data["name"].(string)
// 	ukey := EntityStore.Array[pkey].ChildUnits.Keys[name]
// 	//
// 	temp := make(map[interface{}]interface{})
// 	temp["unit"] = EntityStore.Array[pkey].ChildUnits.Array[ukey]
// 	c.TplName = "UnitView.tpl"
// 	c.Data = temp
// }
