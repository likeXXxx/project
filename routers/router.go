package routers

import (
	"ProjectManage/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/project",
		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
