package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

var globalSessions *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	globalSessions, _ = session.NewManager("memory", sessionConfig)
	go globalSessions.GC()
}

// define some constant variable
const (
	MessageKey = "msg"
	DataKey    = "data"
	SuccessVal = "success"
)

// BaseController base class for other all controller
type BaseController struct {
	beego.Controller
}

// ServeOK response code = 200, and carry a data field that store a json object
func (b *BaseController) ServeOK(msg string, data interface{}) {
	b.Ctx.Output.SetStatus(http.StatusOK)

	resp := make(map[string]interface{})
	resp[MessageKey] = msg
	if data != nil {
		resp[DataKey] = data
	}

	b.Data["json"] = resp
	b.ServeJSON()
}

// ServeError response code = @status, error msg == @err
func (b *BaseController) ServeError(status int, err error) {
	resp := make(map[string]interface{})
	resp[MessageKey] = err.Error()

	b.Ctx.Output.SetStatus(status)
	b.Data["json"] = resp
	b.ServeJSON()
}
