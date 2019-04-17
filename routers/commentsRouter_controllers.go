package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["ProjectManage/controllers:LoginController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:LoginController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Validate",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:LoginController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:LoginController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Mapper",
            Router: `/mapper`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:OManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:OManagerController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:OManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:OManagerController"],
        beego.ControllerComments{
            Method: "GetInfo",
            Router: `/getinfo`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:OManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:OManagerController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:OManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:OManagerController"],
        beego.ControllerComments{
            Method: "GetApplyProjects",
            Router: `/project/apply`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:OManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:OManagerController"],
        beego.ControllerComments{
            Method: "ResetPassword",
            Router: `/pwd`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "GetInfo",
            Router: `/getinfo`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "CreateProject",
            Router: `/project`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "GetTempProjects",
            Router: `/project/temp`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "ResetPassword",
            Router: `/pwd`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
