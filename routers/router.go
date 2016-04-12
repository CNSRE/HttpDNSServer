package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/dns", &controllers.DnsController{})
}
