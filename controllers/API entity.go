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

// ChangeEntityController -
type ChangeEntityController struct {
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

// GetString -
func GetString(c *ViewEntityController, field string) string {
	c.Data[field] = c.GetString(field)
	return c.Data[field].(string)
}

// GetFloat -
func GetFloat(c *ViewEntityController, field string) (result float64) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.ParseFloat(temp, 64)
	return result
}

// GetInt -
func GetInt(c *ViewEntityController, field string) (result int) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.Atoi(temp)
	return result
}

// Post -
func (c *ViewEntityController) Post() {
	Key = ModelsList[GetString(c, "name")]
	EntityStore[Key].HoldPeriod = GetInt(c, "holdperiod")
	EntityStore[Key].EntryYield = GetFloat(c, "entryyield") / 100
	EntityStore[Key].LTV = GetFloat(c, "ltv") / 100
	EntityStore[Key].LoanRate = GetFloat(c, "rate") / 100
	EntityStore[Key].GLA.DiscountRate = GetFloat(c, "discount") / 100
	EntityStore[Key].GLA.PercentSoldRent = GetFloat(c, "soldrent") / 100
	EntityStore[Key].Strategy = GetString(c, "strategy")
	EntityStore[Key].BalloonPercent = GetFloat(c, "balpercent") / 100
	EntityStore[Key].ERVGrowth = GetFloat(c, "erv") / 100
	EntityStore[Key].CPIGrowth = GetFloat(c, "cpi") / 100
	EntityStore[Key].YieldShift = GetFloat(c, "yieldshift")
	EntityStore[Key].GLA.Void = GetInt(c, "void")
	EntityStore[Key].GLA.EXTDuration = GetInt(c, "duration")
	EntityStore[Key].GLA.RentRevisionERV = GetFloat(c, "rentrevision") / 100
	EntityStore[Key].GLA.Probability = GetFloat(c, "probability") / 100
	EntityStore[Key].OpExpercent = GetFloat(c, "opex") / 100
	EntityStore[Key].Fees = GetFloat(c, "fees")
	EntityStore[Key].GLA.Default.Hazard = GetFloat(c, "hazard") / 100
	mcsetup := MCSetup{
		Sims:        GetInt(c, "sims"),
		ERV:         GetFloat(c, "ervsigma") / 100,
		CPI:         GetFloat(c, "cpisigma") / 100,
		YieldShift:  GetFloat(c, "yieldshiftsigma"),
		Void:        GetInt(c, "voidsigma"),
		Probability: GetFloat(c, "probabilitysigma") / 100,
		OpEx:        GetFloat(c, "opexsigma") / 100,
		Hazard:      GetFloat(c, "hazardsigma") / 100,
	}
	temp := make(map[interface{}]interface{})
	Models[Key].UpdateEntity(false, EntityStore[Models[Key].MasterID])
	Models[Key].MCSetup = mcsetup
	Models[Key].MonteCarlo()
	Models[Key].MultiplyInputs()
	temp["entity"] = Models[Key]
	temp["modelslist"] = ModelsList
	c.TplName = "EntityView.tpl"
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
	e.GLA.RentRevisionERV = math.Round(e.GLA.RentRevisionERV*100000) / 1000
	e.GLA.Probability = math.Round(e.GLA.Probability*100000) / 1000
	e.OpEx.PercentOfTRI = math.Round(e.OpEx.PercentOfTRI*100000) / 1000
	e.DebtInput.LTV = math.Round(e.DebtInput.LTV*100000) / 1000
	e.DebtInput.InterestRate = math.Round(e.DebtInput.InterestRate*100000) / 1000
	e.GLA.Default.Hazard = math.Round(e.GLA.Default.Hazard*100000) / 1000
	e.MCSetup.ERV = math.Round(e.MCSetup.ERV*100000) / 1000
	e.MCSetup.CPI = math.Round(e.MCSetup.CPI*100000) / 1000
	e.MCSetup.Probability = math.Round(e.MCSetup.Probability*100000) / 1000
	e.MCSetup.OpEx = math.Round(e.MCSetup.OpEx*100000) / 1000
	e.MCSetup.Hazard = math.Round(e.MCSetup.Hazard*100000) / 1000
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
