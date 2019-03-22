package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

// define some constant variable
const (
	ErrKey     = "error"
	SuccessKey = "msg"
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
	resp[SuccessKey] = msg
	if data != nil {
		resp[DataKey] = data
	}

	b.Data["json"] = resp
	b.ServeJSON()
}

// ServeError response code = @status, error msg == @err
func (b *BaseController) ServeError(status int, err error) {
	resp := make(map[string]interface{})
	resp[ErrKey] = err.Error()

	b.Ctx.Output.SetStatus(status)
	b.Data["json"] = resp
	b.ServeJSON()
}
