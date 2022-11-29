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
	tempkey := EntitiesList[GetStringMCDetails(c, "name")]
	page := GetIntMCDetails(c, "page")
	order := GetStringMCDetails(c, "order")
	startindex := (page - 1) * 10
	endindex := startindex + 10
	//
	// startindex = 0
	// endindex = EntityMap[tempkey].Entity.MCSetup.Sims
	//
	temp := make(map[interface{}]interface{})
	switch order {
	case "irr":
		sort.Sort(ByIRR(EntityMap[tempkey].EntityModel.MCSlice))
	case "irr-r":
		sort.Sort(ByIRRr(EntityMap[tempkey].EntityModel.MCSlice))
	case "em":
		sort.Sort(ByEM(EntityMap[tempkey].EntityModel.MCSlice))
	case "em-r":
		sort.Sort(ByEMr(EntityMap[tempkey].EntityModel.MCSlice))
	case "ytm":
		sort.Sort(ByYTM(EntityMap[tempkey].EntityModel.MCSlice))
	case "ytm-r":
		sort.Sort(ByYTMr(EntityMap[tempkey].EntityModel.MCSlice))
	case "duration":
		sort.Sort(ByDuration(EntityMap[tempkey].EntityModel.MCSlice))
	case "duration-r":
		sort.Sort(ByDurationr(EntityMap[tempkey].EntityModel.MCSlice))
	case "void":
		sort.Sort(ByVoid(EntityMap[tempkey].EntityModel.MCSlice))
	case "void-r":
		sort.Sort(ByVoidr(EntityMap[tempkey].EntityModel.MCSlice))
	case "extdur":
		sort.Sort(ByEXTDUR(EntityMap[tempkey].EntityModel.MCSlice))
	case "extdur-r":
		sort.Sort(ByEXTDURr(EntityMap[tempkey].EntityModel.MCSlice))
	case "hazard":
		sort.Sort(ByHazard(EntityMap[tempkey].EntityModel.MCSlice))
	case "hazard-r":
		sort.Sort(ByHazardr(EntityMap[tempkey].EntityModel.MCSlice))
	case "opex":
		sort.Sort(ByOpEx(EntityMap[tempkey].EntityModel.MCSlice))
	case "opex-r":
		sort.Sort(ByOpExr(EntityMap[tempkey].EntityModel.MCSlice))
	case "cpi":
		sort.Sort(ByCPI(EntityMap[tempkey].EntityModel.MCSlice))
	case "cpi-r":
		sort.Sort(ByCPIr(EntityMap[tempkey].EntityModel.MCSlice))
	case "erv":
		sort.Sort(ByERV(EntityMap[tempkey].EntityModel.MCSlice))
	case "erv-r":
		sort.Sort(ByERVr(EntityMap[tempkey].EntityModel.MCSlice))
	case "yieldshift":
		sort.Sort(ByYieldShift(EntityMap[tempkey].EntityModel.MCSlice))
	case "yieldshift-r":
		sort.Sort(ByYieldShiftr(EntityMap[tempkey].EntityModel.MCSlice))
	case "endcash":
		sort.Sort(ByEndCash(EntityMap[tempkey].EntityModel.MCSlice))
	case "endcash-r":
		sort.Sort(ByEndCashr(EntityMap[tempkey].EntityModel.MCSlice))
	case "endncf":
		sort.Sort(ByEndNCF(EntityMap[tempkey].EntityModel.MCSlice))
	case "endncf-r":
		sort.Sort(ByEndNCFr(EntityMap[tempkey].EntityModel.MCSlice))
	case "endmarketvalue":
		sort.Sort(ByEndMarketValue(EntityMap[tempkey].EntityModel.MCSlice))
	case "endmarketvalue-r":
		sort.Sort(ByEndMarketValuer(EntityMap[tempkey].EntityModel.MCSlice))
	case "prob":
		sort.Sort(ByProb(EntityMap[tempkey].EntityModel.MCSlice))
	case "prob-r":
		sort.Sort(ByProbr(EntityMap[tempkey].EntityModel.MCSlice))
	}
	temp["MCSlice"] = EntityMap[tempkey].EntityModel.MCSlice[startindex:endindex]
	temp["Strategy"] = EntityMap[tempkey].EntityModel.Strategy
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
	tempkey := EntitiesList[GetStringMCTabs(c, "name")]
	tab := c.Data["tab"].(string)
	switch tab {
	case "cf":
		temp["entity"] = EntityMap[tempkey].EntityModel
		c.TplName = "CFTable.tpl"
	case "endcash":
		temp["data"] = EntityMap[tempkey].EntityModel.MCResults.EndCash
		temp["varp"] = EntityMap[tempkey].EntityModel.MCResults.CashBalanceVaR
		c.TplName = "MCHist.tpl"
	case "cashbalance":
		temp["data"] = EntityMap[tempkey].EntityModel.MCResults.CashBalance
		c.TplName = "3dchart.tpl"
	case "endncf":
		temp["data"] = EntityMap[tempkey].EntityModel.MCResults.EndNCF
		temp["varp"] = EntityMap[tempkey].EntityModel.MCResults.NCFVaR
		c.TplName = "MCHist.tpl"
	case "ncf":
		temp["data"] = EntityMap[tempkey].EntityModel.MCResults.NCF
		c.TplName = "3dchart.tpl"
	case "endmarketvalue":
		temp["data"] = EntityMap[tempkey].EntityModel.MCResults.EndMarketValue
		temp["varp"] = EntityMap[tempkey].EntityModel.MCResults.MarketValueVaR
		c.TplName = "MCHist.tpl"
	case "marketvalue":
		temp["data"] = EntityMap[tempkey].EntityModel.MCResults.MarketValue
		c.TplName = "3dchart.tpl"
	case "irr":
		temp["data"] = EntityMap[tempkey].EntityModel.MCResults.IRR
		c.TplName = "MCHist.tpl"
	case "em":
		temp["data"] = EntityMap[tempkey].EntityModel.MCResults.EM
		c.TplName = "MCHist.tpl"
	case "ytm":
		temp["data"] = EntityMap[tempkey].EntityModel.MCResults.YTM
		c.TplName = "MCHist.tpl"
	case "duration":
		temp["data"] = EntityMap[tempkey].EntityModel.MCResults.Duration
		c.TplName = "MCHist.tpl"
	case "ytmdur":
		temp["data"] = EntityMap[tempkey].EntityModel.MCResults.YTMDUR
		c.TplName = "MCHist.tpl"
	case "factoranalysis":
		temp["data"] = EntityMap[tempkey].EntityModel
		c.TplName = "FactorAnalysis.tpl"
		// for i, v := range EntityMap[tempkey].Entity.FactorAnalysis {
		// 	fmt.Println(i)
		// 	fmt.Printf("%+v\n", v)
		// }

	}
	c.Data = temp
}
