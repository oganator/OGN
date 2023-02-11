package controllers

import (
	"sort"
	"strconv"

	beego "github.com/astaxie/beego"
)

// MCDetailsController -
type MCDetailsController struct {
	beego.Controller
}

// Post -
func (c *MCDetailsController) Post() {
	tempkey := EntityModelsList[GetStringMCDetails(c, "name")]
	page := GetIntMCDetails(c, "page")
	order := GetStringMCDetails(c, "order")
	startindex := (page - 1) * 10
	endindex := startindex + 10
	//
	// startindex = 0
	// endindex = EntityModelsMap[tempkey].Entity.MCSetup.Sims
	//
	temp := make(map[interface{}]interface{})
	switch order {
	case "irr":
		sort.Sort(ByIRR(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "irr-r":
		sort.Sort(ByIRRr(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "em":
		sort.Sort(ByEM(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "em-r":
		sort.Sort(ByEMr(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "ytm":
		sort.Sort(ByYTM(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "ytm-r":
		sort.Sort(ByYTMr(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "duration":
		sort.Sort(ByDuration(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "duration-r":
		sort.Sort(ByDurationr(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "void":
		sort.Sort(ByVoid(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "void-r":
		sort.Sort(ByVoidr(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "extdur":
		sort.Sort(ByEXTDUR(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "extdur-r":
		sort.Sort(ByEXTDURr(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "hazard":
		sort.Sort(ByHazard(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "hazard-r":
		sort.Sort(ByHazardr(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "opex":
		sort.Sort(ByOpEx(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "opex-r":
		sort.Sort(ByOpExr(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "cpi":
		sort.Sort(ByCPI(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "cpi-r":
		sort.Sort(ByCPIr(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "erv":
		sort.Sort(ByERV(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "erv-r":
		sort.Sort(ByERVr(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "yieldshift":
		sort.Sort(ByYieldShift(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "yieldshift-r":
		sort.Sort(ByYieldShiftr(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "endcash":
		sort.Sort(ByEndCash(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "endcash-r":
		sort.Sort(ByEndCashr(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "endncf":
		sort.Sort(ByEndNCF(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "endncf-r":
		sort.Sort(ByEndNCFr(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "endmarketvalue":
		sort.Sort(ByEndMarketValue(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "endmarketvalue-r":
		sort.Sort(ByEndMarketValuer(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "prob":
		sort.Sort(ByProb(EntityModelsMap[tempkey].EntityModel.MCSlice))
	case "prob-r":
		sort.Sort(ByProbr(EntityModelsMap[tempkey].EntityModel.MCSlice))
	}
	temp["MCSlice"] = EntityModelsMap[tempkey].EntityModel.MCSlice[startindex:endindex]
	temp["Strategy"] = EntityModelsMap[tempkey].EntityModel.Strategy
	// temp["baseURL"] = BaseURL
	c.TplName = "MCDetails.tpl"
	c.Data = temp
}

// GetStringMCDetails -
func GetStringMCDetails(c *MCDetailsController, field string) string {
	c.Data[field] = c.GetString(field)
	return c.Data[field].(string)
}

// GetIntMCDetails -
func GetIntMCDetails(c *MCDetailsController, field string) (result int) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.Atoi(temp)
	return result
}

type MCTabsController struct {
	beego.Controller
}

// GetStringMCDetails -
func GetStringMCTabs(c *MCTabsController, field string) string {
	c.Data[field] = c.GetString(field)
	return c.Data[field].(string)
}

// Post -
func (c *MCTabsController) Post() {
	temp := make(map[interface{}]interface{})
	c.Data["tab"] = c.GetString("tab")
	tempkey := EntityModelsList[GetStringMCTabs(c, "name")]
	tab := c.Data["tab"].(string)
	switch tab {
	case "cf":
		temp["entity"] = EntityModelsMap[tempkey].EntityModel
		c.TplName = "CFTable.tpl"
	case "endcash":
		temp["data"] = EntityModelsMap[tempkey].EntityModel.MCResults.EndCash
		temp["varp"] = EntityModelsMap[tempkey].EntityModel.MCResults.CashBalanceVaR
		c.TplName = "MCHist.tpl"
	case "cashbalance":
		temp["data"] = EntityModelsMap[tempkey].EntityModel.MCResults.CashBalance
		c.TplName = "3dchart.tpl"
	case "endncf":
		temp["data"] = EntityModelsMap[tempkey].EntityModel.MCResults.EndNCF
		temp["varp"] = EntityModelsMap[tempkey].EntityModel.MCResults.NCFVaR
		c.TplName = "MCHist.tpl"
	case "ncf":
		temp["data"] = EntityModelsMap[tempkey].EntityModel.MCResults.NCF
		c.TplName = "3dchart.tpl"
	case "endmarketvalue":
		temp["data"] = EntityModelsMap[tempkey].EntityModel.MCResults.EndMarketValue
		temp["varp"] = EntityModelsMap[tempkey].EntityModel.MCResults.MarketValueVaR
		c.TplName = "MCHist.tpl"
	case "marketvalue":
		temp["data"] = EntityModelsMap[tempkey].EntityModel.MCResults.MarketValue
		c.TplName = "3dchart.tpl"
	case "irr":
		temp["data"] = EntityModelsMap[tempkey].EntityModel.MCResults.IRR
		c.TplName = "MCHist.tpl"
	case "em":
		temp["data"] = EntityModelsMap[tempkey].EntityModel.MCResults.EM
		c.TplName = "MCHist.tpl"
	case "ytm":
		temp["data"] = EntityModelsMap[tempkey].EntityModel.MCResults.YTM
		c.TplName = "MCHist.tpl"
	case "duration":
		temp["data"] = EntityModelsMap[tempkey].EntityModel.MCResults.Duration
		c.TplName = "MCHist.tpl"
	case "ytmdur":
		temp["data"] = EntityModelsMap[tempkey].EntityModel.MCResults.YTMDUR
		c.TplName = "MCHist.tpl"
	case "factoranalysis":
		temp["data"] = EntityModelsMap[tempkey].EntityModel
		c.TplName = "FactorAnalysis.tpl"
		// for i, v := range EntityModelsMap[tempkey].Entity.FactorAnalysis {
		// 	fmt.Println(i)
		// 	fmt.Printf("%+v\n", v)
		// }

	}
	c.Data = temp
}
