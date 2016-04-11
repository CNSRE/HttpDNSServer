package controllers


import (
	"github.com/astaxie/beego"
)

type DnsController struct {
	beego.Controller
}

func (c *DnsController) Get() {

	ip := c.GetString("ip")
	if( len(ip) == 0 ){
		ip = c.Ctx.Input.IP()
	}

	domain := c.GetString("dn")
	d := domainModel{ domain, ip }

	c.Data["json"] = &d
	c.ServeJSON()
}

type domainModel struct  {
	Domain string
	Ip string
}