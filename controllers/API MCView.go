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
		sort.Sort(ByIRR(EntityMap[tempkey].Entity.MCSlice))
	case "irr-r":
		sort.Sort(ByIRRr(EntityMap[tempkey].Entity.MCSlice))
	case "em":
		sort.Sort(ByEM(EntityMap[tempkey].Entity.MCSlice))
	case "em-r":
		sort.Sort(ByEMr(EntityMap[tempkey].Entity.MCSlice))
	case "ytm":
		sort.Sort(ByYTM(EntityMap[tempkey].Entity.MCSlice))
	case "ytm-r":
		sort.Sort(ByYTMr(EntityMap[tempkey].Entity.MCSlice))
	case "duration":
		sort.Sort(ByDuration(EntityMap[tempkey].Entity.MCSlice))
	case "duration-r":
		sort.Sort(ByDurationr(EntityMap[tempkey].Entity.MCSlice))
	case "void":
		sort.Sort(ByVoid(EntityMap[tempkey].Entity.MCSlice))
	case "void-r":
		sort.Sort(ByVoidr(EntityMap[tempkey].Entity.MCSlice))
	case "extdur":
		sort.Sort(ByEXTDUR(EntityMap[tempkey].Entity.MCSlice))
	case "extdur-r":
		sort.Sort(ByEXTDURr(EntityMap[tempkey].Entity.MCSlice))
	case "hazard":
		sort.Sort(ByHazard(EntityMap[tempkey].Entity.MCSlice))
	case "hazard-r":
		sort.Sort(ByHazardr(EntityMap[tempkey].Entity.MCSlice))
	case "opex":
		sort.Sort(ByOpEx(EntityMap[tempkey].Entity.MCSlice))
	case "opex-r":
		sort.Sort(ByOpExr(EntityMap[tempkey].Entity.MCSlice))
	case "cpi":
		sort.Sort(ByCPI(EntityMap[tempkey].Entity.MCSlice))
	case "cpi-r":
		sort.Sort(ByCPIr(EntityMap[tempkey].Entity.MCSlice))
	case "erv":
		sort.Sort(ByERV(EntityMap[tempkey].Entity.MCSlice))
	case "erv-r":
		sort.Sort(ByERVr(EntityMap[tempkey].Entity.MCSlice))
	case "yieldshift":
		sort.Sort(ByYieldShift(EntityMap[tempkey].Entity.MCSlice))
	case "yieldshift-r":
		sort.Sort(ByYieldShiftr(EntityMap[tempkey].Entity.MCSlice))
	case "endcash":
		sort.Sort(ByEndCash(EntityMap[tempkey].Entity.MCSlice))
	case "endcash-r":
		sort.Sort(ByEndCashr(EntityMap[tempkey].Entity.MCSlice))
	case "endncf":
		sort.Sort(ByEndNCF(EntityMap[tempkey].Entity.MCSlice))
	case "endncf-r":
		sort.Sort(ByEndNCFr(EntityMap[tempkey].Entity.MCSlice))
	case "endmarketvalue":
		sort.Sort(ByEndMarketValue(EntityMap[tempkey].Entity.MCSlice))
	case "endmarketvalue-r":
		sort.Sort(ByEndMarketValuer(EntityMap[tempkey].Entity.MCSlice))
	case "prob":
		sort.Sort(ByProb(EntityMap[tempkey].Entity.MCSlice))
	case "prob-r":
		sort.Sort(ByProbr(EntityMap[tempkey].Entity.MCSlice))
	}
	temp["MCSlice"] = EntityMap[tempkey].Entity.MCSlice[startindex:endindex]
	temp["Strategy"] = EntityMap[tempkey].Entity.Strategy
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
		temp["entity"] = EntityMap[tempkey].Entity
		c.TplName = "CFTable.tpl"
	case "endcash":
		temp["data"] = EntityMap[tempkey].Entity.MCResults.EndCash
		temp["varp"] = EntityMap[tempkey].Entity.MCResults.CashBalanceVaR
		c.TplName = "MCHist.tpl"
	case "cashbalance":
		temp["data"] = EntityMap[tempkey].Entity.MCResults.CashBalance
		c.TplName = "3dchart.tpl"
	case "endncf":
		temp["data"] = EntityMap[tempkey].Entity.MCResults.EndNCF
		temp["varp"] = EntityMap[tempkey].Entity.MCResults.NCFVaR
		c.TplName = "MCHist.tpl"
	case "ncf":
		temp["data"] = EntityMap[tempkey].Entity.MCResults.NCF
		c.TplName = "3dchart.tpl"
	case "endmarketvalue":
		temp["data"] = EntityMap[tempkey].Entity.MCResults.EndMarketValue
		temp["varp"] = EntityMap[tempkey].Entity.MCResults.MarketValueVaR
		c.TplName = "MCHist.tpl"
	case "marketvalue":
		temp["data"] = EntityMap[tempkey].Entity.MCResults.MarketValue
		c.TplName = "3dchart.tpl"
	case "irr":
		temp["data"] = EntityMap[tempkey].Entity.MCResults.IRR
		c.TplName = "MCHist.tpl"
	case "em":
		temp["data"] = EntityMap[tempkey].Entity.MCResults.EM
		c.TplName = "MCHist.tpl"
	case "ytm":
		temp["data"] = EntityMap[tempkey].Entity.MCResults.YTM
		c.TplName = "MCHist.tpl"
	case "duration":
		temp["data"] = EntityMap[tempkey].Entity.MCResults.Duration
		c.TplName = "MCHist.tpl"
	case "ytmdur":
		temp["data"] = EntityMap[tempkey].Entity.MCResults.YTMDUR
		c.TplName = "MCHist.tpl"
	case "factoranalysis":
		temp["data"] = EntityMap[tempkey].Entity
		c.TplName = "FactorAnalysis.tpl"
		// for i, v := range EntityMap[tempkey].Entity.FactorAnalysis {
		// 	fmt.Println(i)
		// 	fmt.Printf("%+v\n", v)
		// }

	}
	c.Data = temp
}
