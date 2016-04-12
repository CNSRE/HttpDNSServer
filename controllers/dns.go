package controllers


import (
	"github.com/astaxie/beego"
	"../util"
	"../models"
)

type DnsController struct {
	beego.Controller
}

func (c *DnsController) Get() {

	domain := c.GetString("domain")
	if len(domain) == 0 {
		c.Ctx.WriteString("domain err")
	}

	ip := c.GetString("ip")
	if( len(ip) == 0 ){
		ip = c.Ctx.Input.IP()
	}

	longIp := util.IpStr2Long(ip)
	regIp := models.NewRegionIP(domain ,longIp)
	domainDns := models.NewDomainDNS( domain , regIp.ID )
	ipinfo := models.NewIpInfo(regIp.ID)
	data := models.NewdataModel( domain , ip , ipinfo.Isp , domainDns.A )

	c.Data["json"] = &data
	c.ServeJSON()

}