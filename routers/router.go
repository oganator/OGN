package routers

import (
	"OGN/controllers"

	beego "github.com/astaxie/beego"
)

func init() {
	// HOME
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/GetSettings", &controllers.SettingsController{})
	// ENTITY
	beego.Router("/ViewEntity", &controllers.ViewEntityController{})
	beego.Router("/ChangeEntity", &controllers.ChangeEntityController{})
	// UNIT
	beego.Router("/ViewUnitCF", &controllers.ViewUnitCFController{})
	beego.Router("/ViewUnitTable", &controllers.ViewUnitTableController{})
	beego.Router("/AddChildUnit", &controllers.AddChildUnitController{})
	beego.Router("/ViewRentSchedule", &controllers.ViewRentScheduleController{})
	// MONTECARLO
	beego.Router("/MCTabs", &controllers.MCTabsController{})
	beego.Router("/MCDetails", &controllers.MCDetailsController{})
	beego.Router("/MCIndex", &controllers.MCIndexController{})
	beego.Router("/test", &controllers.TestController{})
	// FUND
	beego.Router("/Fund", &controllers.FundController{})
}
