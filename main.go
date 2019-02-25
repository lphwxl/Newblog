package main

import (
	_ "shop/routers"
	"github.com/astaxie/beego"
	_ "shop/models"
	"github.com/astaxie/beego/orm"
)

func main() {
	orm.Debug = true
	beego.Run()
}

