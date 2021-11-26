package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	temp := make(map[interface{}]interface{})
	temp["modelslist"] = ModelsList
	temp["fundslist"] = FundsList
	temp["entity"] = Entities[0]
	temp["baseURL"] = BaseURL
	c.TplName = "Home.tpl"
	c.Data = temp
}

type SettingsController struct {
	beego.Controller
}

// GetString -
func GetStringSettings(c *SettingsController, field string) string {
	c.Data[field] = c.GetString(field)
	return c.Data[field].(string)
}

func (c *SettingsController) Post() {
	temp := make(map[interface{}]interface{})
	tempkey := EntitiesList[GetStringSettings(c, "entity")]
	temp["tab"] = GetStringSettings(c, "tab")
	temp["entity"] = Entities[tempkey]
	temp["baseURL"] = BaseURL
	c.TplName = "Settings.tpl"
	c.Data = temp
}

type AppSettingsController struct {
	beego.Controller
}

// GetString -
func GetStringAppSettings(c *AppSettingsController, field string) string {
	c.Data[field] = c.GetString(field)
	return c.Data[field].(string)
}

func (c *AppSettingsController) Patch() {
	temp := make(map[interface{}]interface{})
	BaseURL = GetStringAppSettings(c, "baseURL")
	temp["baseURL"] = BaseURL
	c.TplName = ""
	c.Data = temp
}
