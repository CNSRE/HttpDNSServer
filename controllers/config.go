package controllers

import (
	"github.com/astaxie/beego"
	"../models/config"
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

	c.Data["json"] = conModel
	c.ServeJSON()

}
