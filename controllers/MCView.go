package controllers

import (
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

// Get -
func (c *MCDetailsController) Get() {
	temp := make(map[interface{}]interface{})
	temp["MCSlice"] = Models[Key].MCSlice
	c.TplName = "MCDetails.tpl"
	c.Data = temp
}
