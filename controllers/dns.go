package controllers


import (
	"github.com/astaxie/beego"
	"../models/dns"
)

type DnsController struct {
	beego.Controller
}

func (c *DnsController) Get() {

	data , err := dns.Find( c.GetString("domain") , c.GetString("ip") , c.Ctx.Input.IP() )
	if err != nil {
		//终止程序, 记录log. 返回对应的错误信息
	}

	c.Data["json"] = &data
	c.ServeJSON()

}