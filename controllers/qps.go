package controllers

import (
	"github.com/astaxie/beego"
)

type QpsController struct {
	beego.Controller
}

func (c *QpsController) Get() {
	c.Data["json"] = "qps"
	c.ServeJSON()
}
