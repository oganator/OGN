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
	temp := make(map[interface{}]interface{})
	switch order {
	case "irr":
		sort.Sort(ByIRR(Entities[tempkey].MCSlice))
	case "irr-r":
		sort.Sort(ByIRRr(Entities[tempkey].MCSlice))
	case "em":
		sort.Sort(ByEM(Entities[tempkey].MCSlice))
	case "em-r":
		sort.Sort(ByEMr(Entities[tempkey].MCSlice))
	case "ytm":
		sort.Sort(ByYTM(Entities[tempkey].MCSlice))
	case "ytm-r":
		sort.Sort(ByYTMr(Entities[tempkey].MCSlice))
	case "duration":
		sort.Sort(ByDuration(Entities[tempkey].MCSlice))
	case "duration-r":
		sort.Sort(ByDurationr(Entities[tempkey].MCSlice))
	case "void":
		sort.Sort(ByVoid(Entities[tempkey].MCSlice))
	case "void-r":
		sort.Sort(ByVoidr(Entities[tempkey].MCSlice))
	case "extdur":
		sort.Sort(ByEXTDUR(Entities[tempkey].MCSlice))
	case "extdur-r":
		sort.Sort(ByEXTDURr(Entities[tempkey].MCSlice))
	case "hazard":
		sort.Sort(ByHazard(Entities[tempkey].MCSlice))
	case "hazard-r":
		sort.Sort(ByHazardr(Entities[tempkey].MCSlice))
	case "opex":
		sort.Sort(ByOpEx(Entities[tempkey].MCSlice))
	case "opex-r":
		sort.Sort(ByOpExr(Entities[tempkey].MCSlice))
	case "cpi":
		sort.Sort(ByCPI(Entities[tempkey].MCSlice))
	case "cpi-r":
		sort.Sort(ByCPIr(Entities[tempkey].MCSlice))
	case "erv":
		sort.Sort(ByERV(Entities[tempkey].MCSlice))
	case "erv-r":
		sort.Sort(ByERVr(Entities[tempkey].MCSlice))
	case "yieldshift":
		sort.Sort(ByYieldShift(Entities[tempkey].MCSlice))
	case "yieldshift-r":
		sort.Sort(ByYieldShiftr(Entities[tempkey].MCSlice))
	case "endcash":
		sort.Sort(ByEndCash(Entities[tempkey].MCSlice))
	case "endcash-r":
		sort.Sort(ByEndCashr(Entities[tempkey].MCSlice))
	case "endncf":
		sort.Sort(ByEndNCF(Entities[tempkey].MCSlice))
	case "endncf-r":
		sort.Sort(ByEndNCFr(Entities[tempkey].MCSlice))
	case "endmarketvalue":
		sort.Sort(ByEndMarketValue(Entities[tempkey].MCSlice))
	case "endmarketvalue-r":
		sort.Sort(ByEndMarketValuer(Entities[tempkey].MCSlice))
	case "prob":
		sort.Sort(ByProb(Entities[tempkey].MCSlice))
	case "prob-r":
		sort.Sort(ByProbr(Entities[tempkey].MCSlice))
	}
	temp["MCSlice"] = Entities[tempkey].MCSlice[startindex:endindex]
	temp["Strategy"] = Entities[tempkey].Strategy
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
		temp["entity"] = Entities[tempkey]
		c.TplName = "CFTable.tpl"
	case "endcash":
		temp["data"] = Entities[tempkey].MCResults.EndCash
		temp["varp"] = Entities[tempkey].MCResults.CashBalanceVaR
		c.TplName = "MCHist.tpl"
	case "cashbalance":
		temp["data"] = Entities[tempkey].MCResults.CashBalance
		c.TplName = "3dchart.tpl"
	case "endncf":
		temp["data"] = Entities[tempkey].MCResults.EndNCF
		temp["varp"] = Entities[tempkey].MCResults.NCFVaR
		c.TplName = "MCHist.tpl"
	case "ncf":
		temp["data"] = Entities[tempkey].MCResults.NCF
		c.TplName = "3dchart.tpl"
	case "endmarketvalue":
		temp["data"] = Entities[tempkey].MCResults.EndMarketValue
		temp["varp"] = Entities[tempkey].MCResults.MarketValueVaR
		c.TplName = "MCHist.tpl"
	case "marketvalue":
		temp["data"] = Entities[tempkey].MCResults.MarketValue
		c.TplName = "3dchart.tpl"
	case "irr":
		temp["data"] = Entities[tempkey].MCResults.IRR
		c.TplName = "MCHist.tpl"
	case "em":
		temp["data"] = Entities[tempkey].MCResults.EM
		c.TplName = "MCHist.tpl"
	case "ytm":
		temp["data"] = Entities[tempkey].MCResults.YTM
		c.TplName = "MCHist.tpl"
	case "duration":
		temp["data"] = Entities[tempkey].MCResults.Duration
		c.TplName = "MCHist.tpl"
	}
	c.Data = temp
}

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	c.Data["json"] = Entities[1].MCResults
	c.ServeJSON()
}
