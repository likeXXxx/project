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
		beego.NSNamespace("/teacher",
			beego.NSInclude(
				&controllers.TeacherController{},
			),
		),
		beego.NSNamespace("/omanager",
			beego.NSInclude(
				&controllers.OManagerController{},
			),
		),
		beego.NSNamespace("/imanager",
			beego.NSInclude(
				&controllers.IManagerController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
