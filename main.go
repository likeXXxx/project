package main

import (
	"ProjectManage/db"
	_ "ProjectManage/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.SetStaticPath("/bootstrap-4.3.1-dist", "bootstrap-4.3.1-dist")
	beego.SetStaticPath("/jquery", "jquery")
	db.InitMysql()
	beego.Run()
}
