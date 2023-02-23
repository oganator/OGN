package controllers

import (
	"strconv"

	beego "github.com/astaxie/beego"
)

// ViewEntityController - default route - "/". Executed from submit button
type ViewEntity2Controller struct {
	beego.Controller
}

// GetString2 -
func GetString2(c *ViewEntity2Controller, field string) string {
	c.Data[field] = c.GetString(field)
	return c.Data[field].(string)
}

// GetFloat2 -
func GetFloat2(c *ViewEntity2Controller, field string) (result float64) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.ParseFloat(temp, 64)
	return result
}

// GetInt -
func GetInt2(c *ViewEntity2Controller, field string) (result int) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.Atoi(temp)
	return result
}

// Get -
func (c *ViewEntity2Controller) Get() {
	key := AssetModelsList[GetString2(c, "name")]
	temp := make(map[interface{}]interface{})
	temp["entity"] = EntityModelsMap[key].EntityModel
	temp["modelslist"] = AssetModelsList
	// temp["baseURL"] = BaseURL
	c.TplName = "EntityView.tpl"
	c.Data = temp
}

// Post -
func (c *ViewEntity2Controller) Post() {
	key := EntityModelsList[GetString2(c, "name")]
	EntityModelsMap[key].EntityModel.StartDate = Datetype{Month: ReturnMonth(GetString2(c, "startmonth")), Year: GetInt2(c, "startyear")}
	EntityModelsMap[key].EntityModel.StartDate.Add(0)
	EntityModelsMap[key].EntityModel.SalesDate = Datetype{Month: ReturnMonth(GetString2(c, "salesmonth")), Year: GetInt2(c, "salesyear")}
	EntityModelsMap[key].EntityModel.SalesDate.Add(0)
	EntityModelsMap[key].EntityModel.HoldPeriod = dateintdiff(EntityModelsMap[key].EntityModel.SalesDate.Dateint, EntityModelsMap[key].EntityModel.StartDate.Dateint) - 1
	EntityModelsMap[key].EntityModel.Valuation.EntryYield = GetFloat2(c, "entryyield") / 100
	EntityModelsMap[key].EntityModel.Valuation.DiscountRate = GetFloat2(c, "discountrate") / 100
	EntityModelsMap[key].EntityModel.Valuation.Method = GetString2(c, "valuationmethod")
	EntityModelsMap[key].EntityModel.DebtInput.LTV = GetFloat2(c, "ltv") / 100
	EntityModelsMap[key].EntityModel.DebtInput.InterestRate = GetFloat2(c, "rate") / 100
	EntityModelsMap[key].EntityModel.GLA.DiscountRate = GetFloat2(c, "discount") / 100
	EntityModelsMap[key].EntityModel.GLA.PercentSoldRent = GetFloat2(c, "soldrent") / 100
	EntityModelsMap[key].EntityModel.Strategy = GetString2(c, "strategy")
	EntityModelsMap[key].EntityModel.BalloonPercent = GetFloat2(c, "balpercent") / 100
	growth := map[string]HModel{}
	cpi := HModel{}
	erv := HModel{}
	erv.ShortTermRate = GetFloat2(c, "ervshorttermrate") / 100
	erv.ShortTermPeriod = GetInt2(c, "ervshorttermperiod")
	erv.TransitionPeriod = GetInt2(c, "ervtransitionperiod")
	erv.LongTermRate = GetFloat2(c, "ervlongtermrate") / 100
	cpi.ShortTermRate = GetFloat2(c, "cpishorttermrate") / 100
	cpi.ShortTermPeriod = GetInt2(c, "cpishorttermperiod")
	cpi.TransitionPeriod = GetInt2(c, "cpitransitionperiod")
	cpi.LongTermRate = GetFloat2(c, "cpilongtermrate") / 100
	growth["CPI"] = cpi
	growth["ERV"] = erv
	EntityModelsMap[key].EntityModel.GrowthInput = growth
	EntityModelsMap[key].EntityModel.Valuation.YieldShift = GetFloat2(c, "yieldshift")
	EntityModelsMap[key].EntityModel.GLA.Void = GetInt2(c, "void")
	EntityModelsMap[key].EntityModel.GLA.EXTDuration = GetInt2(c, "duration")
	EntityModelsMap[key].EntityModel.GLA.RentRevisionERV = GetFloat2(c, "rentrevision") / 100
	EntityModelsMap[key].EntityModel.GLA.Probability = GetFloat2(c, "probability") / 100
	EntityModelsMap[key].EntityModel.GLA.RentIncentives.Duration = GetInt2(c, "incentivemonths")
	EntityModelsMap[key].EntityModel.GLA.RentIncentives.PercentOfContractRent = GetFloat2(c, "incentivepercent") / 100
	EntityModelsMap[key].EntityModel.GLA.FitOutCosts.AmountPerTotalArea = GetFloat2(c, "fitoutcosts")
	EntityModelsMap[key].EntityModel.OpEx.PercentOfTRI = GetFloat2(c, "opex") / 100
	EntityModelsMap[key].EntityModel.Fees.PercentOfGAV = GetFloat2(c, "fees")
	EntityModelsMap[key].EntityModel.GLA.Default.Hazard = GetFloat2(c, "hazard") / 100
	// TAX
	EntityModelsMap[key].EntityModel.Tax.RETT = GetFloat2(c, "rett") / 100
	EntityModelsMap[key].EntityModel.Tax.LandValue = GetFloat2(c, "landvalue") / 100
	EntityModelsMap[key].EntityModel.Tax.MinValue = GetFloat2(c, "minvalue") / 100
	EntityModelsMap[key].EntityModel.Tax.UsablePeriod = GetInt2(c, "usableperiod")
	EntityModelsMap[key].EntityModel.Tax.VAT = GetFloat2(c, "vat") / 100
	EntityModelsMap[key].EntityModel.Tax.CarryBackYrs = GetInt2(c, "carrybackyrs")
	EntityModelsMap[key].EntityModel.Tax.CarryForwardYrs = GetInt2(c, "carryforwardyrs")
	mcsetup := MCSetup{
		Sims: GetInt2(c, "sims"),
		ERV: HModel{
			ShortTermRate:    GetFloat2(c, "ervshorttermratesigma") / 100,
			ShortTermPeriod:  GetInt2(c, "ervshorttermperiodsigma"),
			TransitionPeriod: GetInt2(c, "ervtransitionperiodsigma"),
			LongTermRate:     GetFloat2(c, "ervlongtermratesigma") / 100,
		},
		CPI: HModel{
			ShortTermRate:    GetFloat2(c, "cpishorttermratesigma") / 100,
			ShortTermPeriod:  GetInt2(c, "cpishorttermperiodsigma"),
			TransitionPeriod: GetInt2(c, "cpitransitionperiodsigma"),
			LongTermRate:     GetFloat2(c, "cpilongtermratesigma") / 100,
		},
		YieldShift:  GetFloat2(c, "yieldshiftsigma"),
		Void:        GetFloat2(c, "voidsigma"),
		Probability: GetFloat2(c, "probabilitysigma") / 100,
		OpEx:        GetFloat2(c, "opexsigma") / 100,
		Hazard:      GetFloat2(c, "hazardsigma") / 100,
	}
	EntityModelsMap[key].EntityModel.MCSetup = mcsetup
	// TODO: write updated values to DB, Monte Carlo
	EntityModelsMap[key].EntityModel.SalesDate = Dateadd(EntityModelsMap[key].EntityModel.StartDate, EntityModelsMap[key].EntityModel.HoldPeriod)
	switch EntityModelsMap[key].EntityModel.Entity.EntityType {
	case "Asset":
		EntityModelsMap[key].EntityModel.UpdateEntityModel()
	case "Fund":
		EntityModelsMap[key].EntityModel.UpdateFundModel()
	}
	// EntityModelsMap[key].EntityModel.MonteCarlo2("Internal")
	// temp := make(map[interface{}]interface{})
	// temp["entity"] = EntityModelsMap[key].EntityModel
	// temp["modelslist"] = AssetModelsList
	// temp["fundslist"] = FundModelsList
	// c.TplName = "EntityView.tpl"
	c.TplName = "test.tpl"
	// c.Data = temp
}

func ReturnMonth(month string) (output int) {
	switch month {
	case "Jan":
		output = 1
	case "Feb":
		output = 2
	case "Mar":
		output = 3
	case "Apr":
		output = 4
	case "May":
		output = 5
	case "Jun":
		output = 6
	case "Jul":
		output = 7
	case "Aug":
		output = 8
	case "Sep":
		output = 9
	case "Oct":
		output = 10
	case "Nov":
		output = 11
	case "Dec":
		output = 12
	}
	return output
}
