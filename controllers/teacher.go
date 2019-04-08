package controllers

import (
	"ProjectManage/models"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

// TeacherController ..
type TeacherController struct {
	BaseController
	uID   int64
	uType string
}

// Prepare ...
func (c *TeacherController) Prepare() {
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)

	uType := sess.Get("type")
	uNum := sess.Get("user")
	if uType == nil || uNum == nil {
		c.Ctx.Redirect(302, "http://localhost:8080/project/login")
		return
	}

	if t, ok := uType.(string); ok {
		c.uType = t
	} else {
		logrus.Errorf("获取uType失败")
		c.ServeError(http.StatusInternalServerError, fmt.Errorf("获取uType失败"))
		return
	}

	if id, ok := uNum.(int64); ok {
		c.uID = id
	} else {
		logrus.Errorf("获取uID失败")
		c.ServeError(http.StatusInternalServerError, fmt.Errorf("获取uID失败"))
		return
	}
}

// Get ...
// @router / [get]
func (c *TeacherController) Get() {
	c.TplName = "teacher.html"
}

// GetInfo ...
// @router /getinfo [get]
func (c *TeacherController) GetInfo() {
	logrus.Infof("get user info url: [%s]", c.Ctx.Input.URI())

	uName, err := models.GetUserName(c.uID, "教师")
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, uName)
}
