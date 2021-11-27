package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	temp := make(map[interface{}]interface{})
	temp["modelslist"] = ModelsList
	temp["fundslist"] = FundsList
	temp["entity"] = EntityMap[0].Entity
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

// Post - entity settings
func (c *SettingsController) Post() {
	temp := make(map[interface{}]interface{})
	tempkey := EntitiesList[GetStringSettings(c, "entity")]
	temp["tab"] = GetStringSettings(c, "tab")
	temp["entity"] = EntityMap[tempkey].Entity
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

func (c *AppSettingsController) Post() {
	tempmonthly := GetStringAppSettings(c, "monthly")
	Monthly, _ = strconv.ParseBool(tempmonthly)
	fmt.Println(Monthly)
	temp := make(map[interface{}]interface{})
	temp["modelslist"] = ModelsList
	temp["fundslist"] = FundsList
	temp["entity"] = EntityMap[0].Entity
	temp["baseURL"] = BaseURL
	c.TplName = "Home.tpl"
	c.Data = temp
}
