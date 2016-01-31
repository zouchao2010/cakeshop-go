package routers

import (
	"cakeshop-go/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//    beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.MainController{})
}
