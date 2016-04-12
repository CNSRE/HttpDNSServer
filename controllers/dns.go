package controllers


import (
	"github.com/astaxie/beego"
	"../models/dns"
	"time"
	"fmt"
)

type DnsController struct {
	beego.Controller
}

func (c *DnsController) Get() {

//	t1 := time.Now().UnixNano();
	data , err := dns.Find( c.GetString("domain") , c.GetString("ip") , c.Ctx.Input.IP() )
//	t2 := time.Now().UnixNano();
//	fmt.Println(t2-t1);

	if err != nil {
		//终止程序, 记录log. 返回对应的错误信息
	}

	c.Data["json"] = &data
	c.ServeJSON()

}