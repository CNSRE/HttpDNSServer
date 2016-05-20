package controllers

import (
	"github.com/astaxie/beego"
	"../dns/config"
	"fmt"
)

type ConfigController struct {
	beego.Controller
}

func (c *ConfigController) Get() {

	appkey := c.GetString("appkey")
	if len(appkey) < 32{
		// 错误处理
	}

	conModel , err := config.Find(appkey)
	if err != nil{
		// 错误处理
	}

	fmt.Println(appkey)
	fmt.Println(c.GetString("version"))
	fmt.Println(conModel)

	c.Data["json"] = conModel
	c.ServeJSON()
}
