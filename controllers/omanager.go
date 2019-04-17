package controllers

import (
	"ProjectManage/db"
	"ProjectManage/models"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

// OManagerController ...
type OManagerController struct {
	BaseController
}

// Get ...
// @router / [get]
func (c *OManagerController) Get() {
	c.TplName = "omanager.html"
}

// GetInfo ...
// @router /getinfo [get]
func (c *OManagerController) GetInfo() {
	logrus.Infof("get user info url: [%s]", c.Ctx.Input.URI())

	uInfo, err := models.GetUserName(c.uID, "学院管理员")
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, strings.Split(uInfo, ","))
}

// Logout ...
// @router /logout [post]
func (c *OManagerController) Logout() {
	logrus.Infof("omanager[%d] logout url: [%s]", c.uID, c.Ctx.Input.URI())

	globalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)
	c.ServeOK(SuccessVal, nil)
}

// ResetPassword ...
// @router /pwd [put]
func (c *OManagerController) ResetPassword() {
	logrus.Infof("omanager[%s] reset pwd url: [%s]", c.uID, c.Ctx.Input.URI())

	oldPwd := c.GetString("old")
	newPwd := c.GetString("new")

	teacher, err := db.GetOMByID(c.uID)
	if err != nil {
		logrus.Errorf("get omanager info fail:[%v]", err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}
	if teacher.Pwd != oldPwd {
		logrus.Errorf("omanager old pwd is wrong")
		c.ServeError(http.StatusBadRequest, fmt.Errorf("原密码错误!请输入正确的密码"))
		return
	}

	if ok, _ := regexp.MatchString("^[0-9a-zA-Z_]{8,20}$", newPwd); !ok {
		logrus.Errorf("new password is invalid")
		c.ServeError(http.StatusBadRequest, fmt.Errorf("新密码应由8-20位字母/数字/下划线组成"))
		return
	}

	if err := models.ResetOManagerPwd(newPwd, c.uID); err != nil {
		logrus.Errorf("reset omanager pwd fail: [%v]", err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, nil)
}

// GetApplyProjects ...
// @router /project/apply [get]
func (c *OManagerController) GetApplyProjects() {
	logrus.Infof("teacher get temp projects url: [%s]", c.uID, c.Ctx.Input.URI())

	applyProjectsResp, err := models.GetApplyProjects(c.uID)
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	resp := make(map[string]interface{})
	resp["rows"] = applyProjectsResp
	resp["total"] = len(applyProjectsResp)
	c.Data["json"] = resp
	c.ServeJSON()
}
