package main

import (
	_ "./routers"
	"github.com/astaxie/beego"
	"./models/db"
)

func main() {

	db.Init()

	beego.Run()
}
