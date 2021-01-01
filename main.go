package main

import (
	"fmt"
	_ "go-api/routers"
	"register"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		mysqluser := beego.AppConfig.String("mysqluser")
		fmt.Println("11111:"+mysqluser)
	}
	register.
	beego.Run()
}
