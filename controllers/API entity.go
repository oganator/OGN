package controllers

import (
	"math"
	"strconv"

	beego "github.com/astaxie/beego"
)

// ViewEntityController - default route - "/". Executed from submit button
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
	EntityStore[Key].ERVGrowth.ShortTermRate = GetFloat(c, "ervshorttermrate") / 100
	EntityStore[Key].ERVGrowth.ShortTermPeriod = GetInt(c, "ervshorttermperiod")
	EntityStore[Key].ERVGrowth.TransitionPeriod = GetInt(c, "ervtransitionperiod")
	EntityStore[Key].ERVGrowth.LongTermRate = GetFloat(c, "ervlongtermrate") / 100
	EntityStore[Key].CPIGrowth.ShortTermRate = GetFloat(c, "cpishorttermrate") / 100
	EntityStore[Key].CPIGrowth.ShortTermPeriod = GetInt(c, "cpishorttermperiod")
	EntityStore[Key].CPIGrowth.TransitionPeriod = GetInt(c, "cpitransitionperiod")
	EntityStore[Key].CPIGrowth.LongTermRate = GetFloat(c, "cpilongtermrate") / 100
	EntityStore[Key].YieldShift = GetFloat(c, "yieldshift")
	EntityStore[Key].GLA.Void = GetInt(c, "void")
	EntityStore[Key].GLA.EXTDuration = GetInt(c, "duration")
	EntityStore[Key].GLA.RentRevisionERV = GetFloat(c, "rentrevision") / 100
	EntityStore[Key].GLA.Probability = GetFloat(c, "probability") / 100
	EntityStore[Key].OpExpercent = GetFloat(c, "opex") / 100
	EntityStore[Key].Fees = GetFloat(c, "fees")
	EntityStore[Key].GLA.Default.Hazard = GetFloat(c, "hazard") / 100
	mcsetup := MCSetup{
		Sims: GetInt(c, "sims"),
		ERV: HModel{
			ShortTermRate:    GetFloat(c, "ervshorttermratesigma") / 100,
			ShortTermPeriod:  GetInt(c, "ervshorttermperiodsigma"),
			TransitionPeriod: GetInt(c, "ervtransitionperiodsigma"),
			LongTermRate:     GetFloat(c, "ervlongtermratesigma") / 100,
		},
		CPI: HModel{
			ShortTermRate:    GetFloat(c, "cpishorttermratesigma") / 100,
			ShortTermPeriod:  GetInt(c, "cpishorttermperiodsigma"),
			TransitionPeriod: GetInt(c, "cpitransitionperiodsigma"),
			LongTermRate:     GetFloat(c, "cpilongtermratesigma") / 100,
		},
		YieldShift:  GetFloat(c, "yieldshiftsigma"),
		Void:        GetFloat(c, "voidsigma"),
		Probability: GetFloat(c, "probabilitysigma") / 100,
		OpEx:        GetFloat(c, "opexsigma") / 100,
		Hazard:      GetFloat(c, "hazardsigma") / 100,
	}
	// TAX
	EntityStore[Key].RETT = GetFloat(c, "rett") / 100
	EntityStore[Key].Landvalue = GetFloat(c, "landvalue") / 100
	EntityStore[Key].WOZpercent = GetFloat(c, "minvalue") / 100
	EntityStore[Key].DeprPeriod = GetInt(c, "usableperiod")
	EntityStore[Key].VAT = GetFloat(c, "vat") / 100
	EntityStore[Key].CarryBackYrs = GetInt(c, "carrybackyrs")
	EntityStore[Key].CarryForwardYrs = GetInt(c, "carryforwardyrs")
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
	temperv := HModel{}
	temperv.ShortTermRate = math.Round(e.GrowthInput["ERV"].ShortTermRate*100000) / 1000
	temperv.ShortTermPeriod = int(math.Round(float64(e.GrowthInput["ERV"].ShortTermPeriod)))
	temperv.TransitionPeriod = int(math.Round(float64(e.GrowthInput["ERV"].TransitionPeriod)))
	temperv.LongTermRate = math.Round(e.GrowthInput["ERV"].LongTermRate*100000) / 1000
	e.GrowthInput["ERV"] = temperv
	tempcpi := HModel{}
	tempcpi.ShortTermRate = math.Round(e.GrowthInput["CPI"].ShortTermRate*100000) / 1000
	tempcpi.ShortTermPeriod = int(math.Round(float64(e.GrowthInput["CPI"].ShortTermPeriod)))
	tempcpi.TransitionPeriod = int(math.Round(float64(e.GrowthInput["CPI"].TransitionPeriod)))
	tempcpi.LongTermRate = math.Round(e.GrowthInput["CPI"].LongTermRate*100000) / 1000
	e.GrowthInput["CPI"] = tempcpi
	e.GLA.RentRevisionERV = math.Round(e.GLA.RentRevisionERV*100000) / 1000
	e.GLA.Probability = math.Round(e.GLA.Probability*100000) / 1000
	e.OpEx.PercentOfTRI = math.Round(e.OpEx.PercentOfTRI*100000) / 1000
	e.DebtInput.LTV = math.Round(e.DebtInput.LTV*100000) / 1000
	e.DebtInput.InterestRate = math.Round(e.DebtInput.InterestRate*100000) / 1000
	e.GLA.Default.Hazard = math.Round(e.GLA.Default.Hazard*100000) / 1000

	tempervmc := HModel{}
	tempervmc.ShortTermRate = math.Round(e.MCSetup.ERV.ShortTermRate*100000) / 1000
	tempervmc.ShortTermPeriod = int(math.Round(float64(e.MCSetup.ERV.ShortTermPeriod)))
	tempervmc.TransitionPeriod = int(math.Round(float64(e.MCSetup.ERV.TransitionPeriod)))
	tempervmc.LongTermRate = math.Round(e.MCSetup.ERV.LongTermRate*100000) / 1000
	e.MCSetup.ERV = tempervmc

	tempcpimc := HModel{}
	tempcpimc.ShortTermRate = math.Round(e.MCSetup.CPI.ShortTermRate*100000) / 1000
	tempcpimc.ShortTermPeriod = int(math.Round(float64(e.MCSetup.CPI.ShortTermPeriod)))
	tempcpimc.TransitionPeriod = int(math.Round(float64(e.MCSetup.CPI.TransitionPeriod)))
	tempcpimc.LongTermRate = math.Round(e.MCSetup.CPI.LongTermRate*100000) / 1000
	e.MCSetup.CPI = tempcpimc

	e.MCSetup.Probability = math.Round(e.MCSetup.Probability*100000) / 1000
	e.MCSetup.OpEx = math.Round(e.MCSetup.OpEx*100000) / 1000
	e.MCSetup.Hazard = math.Round(e.MCSetup.Hazard*100000) / 1000

	// TAX
	e.Tax.RETT = math.Round(e.Tax.RETT*100000) / 1000
	e.Tax.LandValue = math.Round(e.Tax.LandValue*100000) / 1000
	e.Tax.MinValue = math.Round(e.Tax.MinValue*100000) / 1000
	e.Tax.UsablePeriod = int(math.Round(float64(e.Tax.UsablePeriod)))
	e.Tax.VAT = math.Round(e.Tax.VAT*100000) / 1000
	e.Tax.CarryBackYrs = int(math.Round(float64(e.Tax.CarryBackYrs)))
	e.Tax.CarryForwardYrs = int(math.Round(float64(e.Tax.CarryForwardYrs)))
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
