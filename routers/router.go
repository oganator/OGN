package routers

import (
	"OGN/controllers"

	beego "github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ViewEntityController{})
	beego.Router("/ChangeEntity", &controllers.ChangeEntityController{})
	beego.Router("/", &controllers.ViewUnitController{})
}
