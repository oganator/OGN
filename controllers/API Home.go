package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func GetStringHome(c *HomeController, field string) string {
	c.Data[field] = c.GetString(field)
	return c.Data[field].(string)
}

func (c *HomeController) Get() {
	temp := make(map[interface{}]interface{})
	temp["modelslist"] = AssetModelsList
	temp["fundslist"] = FundModelsList
	temp["entity"] = EntityModelsMap[0].EntityModel
	temp["entityMap"] = EntityMap
	switch InitVersion {
	case "v1":
		temp["path"] = "/ViewEntity"
	case "v2":
		temp["path"] = "/ViewEntity2"
	}
	temp["Monthly"] = Monthly
	temp["AzureURL"] = AzureURL
	temp["Compute"] = Compute
	temp["MCActive"] = MCActive
	temp["Debug"] = Debug
	// temp["baseURL"] = BaseURL
	c.TplName = "Home.tpl"
	c.Data = temp
}

// Post - retrieves entity model table
func (c *HomeController) Post() {
	temp := make(map[interface{}]interface{})
	entity := GetStringHome(c, "entity")
	entityInt := EntitiesList[entity]
	temp["models"] = EntityMap[entityInt].Models //GetModels(entityInt)
	c.TplName = "EntityModelTable.tpl"
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

// Post - entity settings
func (c *SettingsController) Post() {
	temp := make(map[interface{}]interface{})
	tempkey := EntityModelsList[GetStringSettings(c, "entity")]
	temp["tab"] = GetStringSettings(c, "tab")
	temp["MCActive"] = MCActive
	temp["entity"] = EntityModelsMap[tempkey].EntityModel
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

func (c *AppSettingsController) Post() {
	tempmonthly := GetStringAppSettings(c, "monthly")
	Monthly, _ = strconv.ParseBool(tempmonthly)
	Compute = GetStringAppSettings(c, "compute")
	AzureURL = GetStringAppSettings(c, "azureurl")
	tempMCActive := GetStringAppSettings(c, "mcActive")
	MCActive, _ = strconv.ParseBool(tempMCActive)
	tempDebug := GetStringAppSettings(c, "debug")
	Debug, _ = strconv.ParseBool(tempDebug)
	temp := make(map[interface{}]interface{})
	temp["modelslist"] = AssetModelsList
	temp["fundslist"] = FundModelsList
	temp["entity"] = EntityModelsMap[0].EntityModel
	// temp["baseURL"] = BaseURL
	c.TplName = "Home.tpl"
	c.Data = temp
}
