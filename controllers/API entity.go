package controllers

import (
	"runtime/debug"
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

// Get -
func (c *ViewEntityController) Get() {
	key := ModelsList[GetString(c, "name")]
	temp := make(map[interface{}]interface{})
	temp["entity"] = EntityMap[key].Entity
	temp["modelslist"] = ModelsList
	temp["baseURL"] = BaseURL
	c.TplName = "EntityView.tpl"
	c.Data = temp
}

// Post -
func (c *ViewEntityController) Post() {
	key := ModelsList[GetString(c, "name")]
	EntityDataStore[key].Mutex.Lock()
	EntityDataStore[key].HoldPeriod = GetInt(c, "holdperiod")
	EntityDataStore[key].EntryYield = GetFloat(c, "entryyield") / 100
	EntityDataStore[key].LTV = GetFloat(c, "ltv") / 100
	EntityDataStore[key].LoanRate = GetFloat(c, "rate") / 100
	EntityDataStore[key].GLA.DiscountRate = GetFloat(c, "discount") / 100
	EntityDataStore[key].GLA.PercentSoldRent = GetFloat(c, "soldrent") / 100
	EntityDataStore[key].Strategy = GetString(c, "strategy")
	EntityDataStore[key].BalloonPercent = GetFloat(c, "balpercent") / 100
	EntityDataStore[key].ERVGrowth.ShortTermRate = GetFloat(c, "ervshorttermrate") / 100
	EntityDataStore[key].ERVGrowth.ShortTermPeriod = GetInt(c, "ervshorttermperiod")
	EntityDataStore[key].ERVGrowth.TransitionPeriod = GetInt(c, "ervtransitionperiod")
	EntityDataStore[key].ERVGrowth.LongTermRate = GetFloat(c, "ervlongtermrate") / 100
	EntityDataStore[key].CPIGrowth.ShortTermRate = GetFloat(c, "cpishorttermrate") / 100
	EntityDataStore[key].CPIGrowth.ShortTermPeriod = GetInt(c, "cpishorttermperiod")
	EntityDataStore[key].CPIGrowth.TransitionPeriod = GetInt(c, "cpitransitionperiod")
	EntityDataStore[key].CPIGrowth.LongTermRate = GetFloat(c, "cpilongtermrate") / 100
	EntityDataStore[key].YieldShift = GetFloat(c, "yieldshift")
	EntityDataStore[key].GLA.Void = GetInt(c, "void")
	EntityDataStore[key].GLA.EXTDuration = GetInt(c, "duration")
	EntityDataStore[key].GLA.RentRevisionERV = GetFloat(c, "rentrevision") / 100
	EntityDataStore[key].GLA.Probability = GetFloat(c, "probability") / 100
	EntityDataStore[key].GLA.RentIncentives.Duration = GetInt(c, "incentivemonths")
	EntityDataStore[key].GLA.RentIncentives.PercentOfContractRent = GetFloat(c, "incentivepercent") / 100
	EntityDataStore[key].GLA.FitOutCosts.AmountPerTotalArea = GetFloat(c, "fitoutcosts")
	EntityDataStore[key].OpExpercent = GetFloat(c, "opex") / 100
	EntityDataStore[key].Fees = GetFloat(c, "fees")
	EntityDataStore[key].GLA.Default.Hazard = GetFloat(c, "hazard") / 100
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
	EntityDataStore[key].RETT = GetFloat(c, "rett") / 100
	EntityDataStore[key].Landvalue = GetFloat(c, "landvalue") / 100
	EntityDataStore[key].WOZpercent = GetFloat(c, "minvalue") / 100
	EntityDataStore[key].DeprPeriod = GetInt(c, "usableperiod")
	EntityDataStore[key].VAT = GetFloat(c, "vat") / 100
	EntityDataStore[key].CarryBackYrs = GetInt(c, "carrybackyrs")
	EntityDataStore[key].CarryForwardYrs = GetInt(c, "carryforwardyrs")
	//
	temp := make(map[interface{}]interface{})
	EntityMap[key].Mutex.Lock()
	EntityMap[key].Entity.UpdateEntity(false, EntityDataStore[EntityMap[key].Entity.MasterID])
	EntityMap[key].Entity.MCSetup = mcsetup
	WriteXLSXEntities(EntityMap[key].Entity)
	if EntityMap[key].Entity.Parent != EntityMap[key].Entity && EntityMap[key].Entity.MCSetup.Sims >= 100 {
		EntityMap[key].Entity.MonteCarlo()
	}
	temp["entity"] = EntityMap[key].Entity
	temp["modelslist"] = ModelsList
	temp["fundslist"] = FundsList
	c.TplName = "EntityView.tpl"
	c.Data = temp
	EntityDataStore[key].Mutex.Unlock()
	EntityMap[key].Mutex.Unlock()
	go debug.FreeOSMemory()
}

// MultiplyInputs -
// func (e *Entity) MultiplyInputs() {
// 	e.Valuation.EntryYield = math.Round(e.Valuation.EntryYield*100000) / 1000
// 	e.Valuation.ExitYield = math.Round(e.Valuation.ExitYield*100000) / 1000
// 	e.GLA.DiscountRate = math.Round(e.GLA.DiscountRate*100000) / 1000
// 	e.GLA.PercentSoldRent = math.Round(e.GLA.PercentSoldRent*100000) / 1000
// 	e.BalloonPercent = math.Round(e.BalloonPercent*100000) / 1000
// 	e.GLA.RentIncentives.PercentOfContractRent = math.Round(e.GLA.RentIncentives.PercentOfContractRent*100000) / 1000
// 	temperv := HModel{}
// 	temperv.ShortTermRate = math.Round(e.GrowthInput["ERV"].ShortTermRate*100000) / 1000
// 	temperv.ShortTermPeriod = int(math.Round(float64(e.GrowthInput["ERV"].ShortTermPeriod)))
// 	temperv.TransitionPeriod = int(math.Round(float64(e.GrowthInput["ERV"].TransitionPeriod)))
// 	temperv.LongTermRate = math.Round(e.GrowthInput["ERV"].LongTermRate*100000) / 1000
// 	e.GrowthInput["ERV"] = temperv
// 	tempcpi := HModel{}
// 	tempcpi.ShortTermRate = math.Round(e.GrowthInput["CPI"].ShortTermRate*100000) / 1000
// 	tempcpi.ShortTermPeriod = int(math.Round(float64(e.GrowthInput["CPI"].ShortTermPeriod)))
// 	tempcpi.TransitionPeriod = int(math.Round(float64(e.GrowthInput["CPI"].TransitionPeriod)))
// 	tempcpi.LongTermRate = math.Round(e.GrowthInput["CPI"].LongTermRate*100000) / 1000
// 	e.GrowthInput["CPI"] = tempcpi
// 	e.GLA.RentRevisionERV = math.Round(e.GLA.RentRevisionERV*100000) / 1000
// 	e.GLA.Probability = math.Round(e.GLA.Probability*100000) / 1000
// 	e.OpEx.PercentOfTRI = math.Round(e.OpEx.PercentOfTRI*100000) / 1000
// 	e.DebtInput.LTV = math.Round(e.DebtInput.LTV*100000) / 1000
// 	e.DebtInput.InterestRate = math.Round(e.DebtInput.InterestRate*100000) / 1000
// 	e.GLA.Default.Hazard = math.Round(e.GLA.Default.Hazard*100000) / 1000
// 	tempervmc := HModel{}
// 	tempervmc.ShortTermRate = math.Round(e.MCSetup.ERV.ShortTermRate*100000) / 1000
// 	tempervmc.ShortTermPeriod = int(math.Round(float64(e.MCSetup.ERV.ShortTermPeriod)))
// 	tempervmc.TransitionPeriod = int(math.Round(float64(e.MCSetup.ERV.TransitionPeriod)))
// 	tempervmc.LongTermRate = math.Round(e.MCSetup.ERV.LongTermRate*100000) / 1000
// 	e.MCSetup.ERV = tempervmc
// 	tempcpimc := HModel{}
// 	tempcpimc.ShortTermRate = math.Round(e.MCSetup.CPI.ShortTermRate*100000) / 1000
// 	tempcpimc.ShortTermPeriod = int(math.Round(float64(e.MCSetup.CPI.ShortTermPeriod)))
// 	tempcpimc.TransitionPeriod = int(math.Round(float64(e.MCSetup.CPI.TransitionPeriod)))
// 	tempcpimc.LongTermRate = math.Round(e.MCSetup.CPI.LongTermRate*100000) / 1000
// 	e.MCSetup.CPI = tempcpimc
// 	e.MCSetup.Probability = math.Round(e.MCSetup.Probability*100000) / 1000
// 	e.MCSetup.OpEx = math.Round(e.MCSetup.OpEx*100000) / 1000
// 	e.MCSetup.Hazard = math.Round(e.MCSetup.Hazard*100000) / 1000
// 	// TAX
// 	e.Tax.RETT = math.Round(e.Tax.RETT*100000) / 1000
// 	e.Tax.LandValue = math.Round(e.Tax.LandValue*100000) / 1000
// 	e.Tax.MinValue = math.Round(e.Tax.MinValue*100000) / 1000
// 	e.Tax.UsablePeriod = int(math.Round(float64(e.Tax.UsablePeriod)))
// 	e.Tax.VAT = math.Round(e.Tax.VAT*100000) / 1000
// 	e.Tax.CarryBackYrs = int(math.Round(float64(e.Tax.CarryBackYrs)))
// 	e.Tax.CarryForwardYrs = int(math.Round(float64(e.Tax.CarryForwardYrs)))
// }

// Post -
func (c *ChangeEntityController) Post() {
	// select model
	c.Data["modelname"] = c.GetString("modelname")
	name := c.Data["modelname"].(string)
	key := ModelsList[name]
	//
	c.TplName = "EntityView.tpl"
	temp := make(map[interface{}]interface{})
	temp["entity"] = EntityMap[key].Entity
	temp["modelslist"] = ModelsList
	temp["baseURL"] = BaseURL
	c.Data = temp
}
