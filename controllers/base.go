package controllers

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/sirupsen/logrus"
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
	ServerHost = "http://localhost:8080"
)

// BaseController base class for other all controller
type BaseController struct {
	beego.Controller
	uID   int64
	uType string
}

// Prepare ...
func (b *BaseController) Prepare() {
	if b.Ctx.Input.URI() == "/project/login" || b.Ctx.Input.URI() == "/project/login/mapper" {
		return
	}

	sess, _ := globalSessions.SessionStart(b.Ctx.ResponseWriter, b.Ctx.Request)

	uType := sess.Get("type")
	uNum := sess.Get("user")
	if uType == nil || uNum == nil {
		b.Ctx.Redirect(302, ServerHost+"/project/login")
		return
	}

	if t, ok := uType.(string); ok {
		b.uType = t
	} else {
		logrus.Errorf("获取uType失败")
		b.ServeError(http.StatusInternalServerError, fmt.Errorf("获取uType失败"))
		return
	}

	if id, ok := uNum.(int64); ok {
		b.uID = id
	} else {
		logrus.Errorf("获取uID失败")
		b.ServeError(http.StatusInternalServerError, fmt.Errorf("获取uID失败"))
		return
	}
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
