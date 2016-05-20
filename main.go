package main

import (
	_ "./routers"
	"github.com/astaxie/beego"
	"./dns/db"
)

func main() {
	db.Init()
	beego.Run()
}
