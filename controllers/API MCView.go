package controllers

import (
	"sort"
	"strconv"

	beego "github.com/astaxie/beego"
)

// CFTableController -
type CFTableController struct {
	beego.Controller
}

// Get -
func (c *CFTableController) Get() {
	temp := make(map[interface{}]interface{})
	temp["entity"] = Models[Key]
	c.TplName = "CFTable.tpl"
	c.Data = temp
}

// MCEndingCashController -
type MCEndingCashController struct {
	beego.Controller
}

// Get -
func (c *MCEndingCashController) Get() {
	temp := make(map[interface{}]interface{})
	temp["data"] = Models[Key].MCResults.EndCash
	c.TplName = "MCHist.tpl"
	c.Data = temp
}

// MCEndingNCFController -
type MCEndingNCFController struct {
	beego.Controller
}

// Get -
func (c *MCEndingNCFController) Get() {
	temp := make(map[interface{}]interface{})
	temp["data"] = Models[Key].MCResults.EndNCF
	c.TplName = "MCHist.tpl"
	c.Data = temp
}

// MCIRRController -
type MCIRRController struct {
	beego.Controller
}

// Get -
func (c *MCIRRController) Get() {
	temp := make(map[interface{}]interface{})
	temp["data"] = Models[Key].MCResults.IRR
	c.TplName = "MCHist.tpl"
	c.Data = temp
}

// MCEMController -
type MCEMController struct {
	beego.Controller
}

// Get -
func (c *MCEMController) Get() {
	temp := make(map[interface{}]interface{})
	temp["data"] = Models[Key].MCResults.EM
	c.TplName = "MCHist.tpl"
	c.Data = temp
}

// MCYTMController -
type MCYTMController struct {
	beego.Controller
}

// Get -
func (c *MCYTMController) Get() {
	temp := make(map[interface{}]interface{})
	temp["data"] = Models[Key].MCResults.YTM
	c.TplName = "MCHist.tpl"
	c.Data = temp
}

// MCDurationController -
type MCDurationController struct {
	beego.Controller
}

// Get -
func (c *MCDurationController) Get() {
	temp := make(map[interface{}]interface{})
	temp["data"] = Models[Key].MCResults.Duration
	c.TplName = "MCHist.tpl"
	c.Data = temp
}

// MCDetailsController -
type MCDetailsController struct {
	beego.Controller
}

// Post -
func (c *MCDetailsController) Post() {
	tempkey := ModelsList[GetStringMCDetails(c, "name")]
	page := GetIntMCDetails(c, "page")
	order := GetStringMCDetails(c, "order")
	startindex := (page - 1) * 10
	endindex := startindex + 10
	temp := make(map[interface{}]interface{})
	switch order {
	case "irr":
		sort.Sort(ByIRR(Models[tempkey].MCSlice))
	case "irr-r":
		sort.Sort(ByIRRr(Models[tempkey].MCSlice))
	case "em":
		sort.Sort(ByEM(Models[tempkey].MCSlice))
	case "em-r":
		sort.Sort(ByEMr(Models[tempkey].MCSlice))
	case "ytm":
		sort.Sort(ByYTM(Models[tempkey].MCSlice))
	case "ytm-r":
		sort.Sort(ByYTMr(Models[tempkey].MCSlice))
	case "duration":
		sort.Sort(ByDuration(Models[tempkey].MCSlice))
	case "duration-r":
		sort.Sort(ByDurationr(Models[tempkey].MCSlice))
	case "void":
		sort.Sort(ByVoid(Models[tempkey].MCSlice))
	case "void-r":
		sort.Sort(ByVoidr(Models[tempkey].MCSlice))
	case "extdur":
		sort.Sort(ByEXTDUR(Models[tempkey].MCSlice))
	case "extdur-r":
		sort.Sort(ByEXTDURr(Models[tempkey].MCSlice))
	case "hazard":
		sort.Sort(ByHazard(Models[tempkey].MCSlice))
	case "hazard-r":
		sort.Sort(ByHazardr(Models[tempkey].MCSlice))
	case "opex":
		sort.Sort(ByOpEx(Models[tempkey].MCSlice))
	case "opex-r":
		sort.Sort(ByOpExr(Models[tempkey].MCSlice))
	case "cpi":
		sort.Sort(ByCPI(Models[tempkey].MCSlice))
	case "cpi-r":
		sort.Sort(ByCPIr(Models[tempkey].MCSlice))
	case "erv":
		sort.Sort(ByERV(Models[tempkey].MCSlice))
	case "erv-r":
		sort.Sort(ByERVr(Models[tempkey].MCSlice))
	case "yieldshift":
		sort.Sort(ByYieldShift(Models[tempkey].MCSlice))
	case "yieldshift-r":
		sort.Sort(ByYieldShiftr(Models[tempkey].MCSlice))
	case "endcash":
		sort.Sort(ByEndCash(Models[tempkey].MCSlice))
	case "endcash-r":
		sort.Sort(ByEndCashr(Models[tempkey].MCSlice))
	case "endncf":
		sort.Sort(ByEndCash(Models[tempkey].MCSlice))
	case "endncf-r":
		sort.Sort(ByEndCashr(Models[tempkey].MCSlice))
	}
	temp["MCSlice"] = Models[tempkey].MCSlice[startindex:endindex]
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

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	c.Data["json"] = Models[1].MCResults
	c.ServeJSON()
}
