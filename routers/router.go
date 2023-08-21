package routers

import (
	"OGN/controllers"

	beego "github.com/astaxie/beego"
)

func init() {
	// HOME
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/GetSettings", &controllers.SettingsController{})
	beego.Router("/AppSettings", &controllers.AppSettingsController{})
	// ENTITY
	beego.Router("/ViewEntity", &controllers.ViewEntityController{})
	beego.Router("/ViewEntity2", &controllers.ViewEntity2Controller{})
	beego.Router("/ChangeEntity", &controllers.ChangeEntityController{})
	beego.Router("/Chart", &controllers.ChartController{})
	// PROPERTY
	beego.Router("/AddCost", &controllers.AddCostController{})
	// UNIT
	beego.Router("/ViewUnitCF", &controllers.ViewUnitCFController{})
	beego.Router("/UpdateUnit", &controllers.UpdateUnitController{})
	beego.Router("/ViewUnitTable", &controllers.ViewUnitTableController{})
	beego.Router("/AddChildUnit", &controllers.AddChildUnitController{})
	beego.Router("/ViewRentSchedule", &controllers.ViewRentScheduleController{})
	// MONTECARLO & SENSITIVITY
	beego.Router("/MCTabs", &controllers.MCTabsController{})
	beego.Router("/MCDetails", &controllers.MCDetailsController{})
	beego.Router("/MCIndex", &controllers.MCIndexController{})
	beego.Router("/Sensitivity", &controllers.SensitivityController{})
	// FUND
	beego.Router("/Fund", &controllers.FundController{})
	// AZURE
	beego.Router("/Function", &controllers.FunctionController{})
	// FLUTTER
	beego.Router("/test", &controllers.TestController{})
}
