package routers

import (
	"OGN/controllers"

	beego "github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ViewEntityController{})
	beego.Router("/ChangeEntity", &controllers.ChangeEntityController{})
	beego.Router("/", &controllers.ViewUnitController{})
	// MONTECARLO
	beego.Router("/MCTabs", &controllers.MCTabsController{})
	beego.Router("/MCDetails", &controllers.MCDetailsController{})
	beego.Router("/MCIndex", &controllers.MCIndexController{})
	beego.Router("/test", &controllers.TestController{})
}
