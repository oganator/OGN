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
	beego.Router("/CFTable", &controllers.CFTableController{})
	beego.Router("/MCEndingCash", &controllers.MCEndingCashController{})
	beego.Router("/MCEndingNCF", &controllers.MCEndingNCFController{})
	beego.Router("/MCIRR", &controllers.MCIRRController{})
	beego.Router("/MCEM", &controllers.MCEMController{})
	beego.Router("/MCYTM", &controllers.MCYTMController{})
	beego.Router("/MCDuration", &controllers.MCDurationController{})
	beego.Router("/MCDetails", &controllers.MCDetailsController{})
	beego.Router("/MCIndex", &controllers.MCIndexController{})
}
