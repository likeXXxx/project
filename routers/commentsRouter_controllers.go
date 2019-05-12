package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"],
        beego.ControllerComments{
            Method: "FinAuditFail",
            Router: `/finaudit/fail`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"],
        beego.ControllerComments{
            Method: "FinAuditPass",
            Router: `/finaudit/pass`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"],
        beego.ControllerComments{
            Method: "GetInfo",
            Router: `/getinfo`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"],
        beego.ControllerComments{
            Method: "GetMasterAuditResult",
            Router: `/master/audit`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"],
        beego.ControllerComments{
            Method: "GetApplyProjects",
            Router: `/project/apply`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"],
        beego.ControllerComments{
            Method: "IManagerGetProjectDetail",
            Router: `/project/detail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"],
        beego.ControllerComments{
            Method: "ImanagerProjectFail",
            Router: `/project/fail`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"],
        beego.ControllerComments{
            Method: "ApplyChangeInviteProject",
            Router: `/project/invite/change`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"],
        beego.ControllerComments{
            Method: "FailChangeProjectApply",
            Router: `/project/invite/change/fail`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"],
        beego.ControllerComments{
            Method: "PassChangeProjectApply",
            Router: `/project/invite/change/pass`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"],
        beego.ControllerComments{
            Method: "IManagerProjectPass",
            Router: `/project/pass`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:IManagerController"],
        beego.ControllerComments{
            Method: "ResetPassword",
            Router: `/pwd`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

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

    beego.GlobalControllerRouter["ProjectManage/controllers:MasterController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:MasterController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:MasterController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:MasterController"],
        beego.ControllerComments{
            Method: "GetInfo",
            Router: `/getinfo`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:MasterController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:MasterController"],
        beego.ControllerComments{
            Method: "ListMaster",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:MasterController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:MasterController"],
        beego.ControllerComments{
            Method: "GetApplyProjects",
            Router: `/project/apply`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:MasterController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:MasterController"],
        beego.ControllerComments{
            Method: "MasterGetProjectDetail",
            Router: `/project/detail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:MasterController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:MasterController"],
        beego.ControllerComments{
            Method: "MasterProjectFail",
            Router: `/project/fail`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:MasterController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:MasterController"],
        beego.ControllerComments{
            Method: "MasterProjectPass",
            Router: `/project/pass`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:MasterController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:MasterController"],
        beego.ControllerComments{
            Method: "ResetPassword",
            Router: `/pwd`,
            AllowHTTPMethods: []string{"put"},
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
            Method: "GetProjectDetail",
            Router: `/project/detail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:OManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:OManagerController"],
        beego.ControllerComments{
            Method: "ApplyProjectFail",
            Router: `/project/fail`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:OManagerController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:OManagerController"],
        beego.ControllerComments{
            Method: "ApplyProjectPass",
            Router: `/project/pass`,
            AllowHTTPMethods: []string{"post"},
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
            Method: "DeleteProject",
            Router: `/project`,
            AllowHTTPMethods: []string{"delete"},
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
            Method: "GetAbolitionProjects",
            Router: `/project/abolition`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "DeleteAbolitionProject",
            Router: `/project/abolition`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "GetProjectDetail",
            Router: `/project/detail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "GetFinishedProjects",
            Router: `/project/finished`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "GetInviteProject",
            Router: `/project/invite`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "ApplyChangeInviteProject",
            Router: `/project/invite/change`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "ProjectRun",
            Router: `/project/run`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "GetRunningProjects",
            Router: `/project/run`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "RunningProjectAddEvent",
            Router: `/project/run/addevent`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "RunningProjectFinish",
            Router: `/project/run/finish`,
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
            Method: "VerifyProject",
            Router: `/project/verify`,
            AllowHTTPMethods: []string{"post"},
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

    beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"] = append(beego.GlobalControllerRouter["ProjectManage/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "ListRunningProjectEvent",
            Router: `/run/eventlist`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
