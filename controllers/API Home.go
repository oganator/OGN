package controllers

import (
	"fmt"

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
	temp["entity"] = Entities[tempkey]
	temp["strategy"] = Entities[tempkey].Strategy
	c.TplName = "Settings.tpl"
	fmt.Println(Entities[tempkey].Strategy, Entities[tempkey].Name)
	c.Data = temp
}
