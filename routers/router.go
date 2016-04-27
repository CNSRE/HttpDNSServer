package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/dns", &controllers.DnsController{})
	beego.Router("/config", &controllers.ConfigController{})
	beego.Router("/qps", &controllers.QpsController{})
}
