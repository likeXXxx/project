package controllers

import (
	"ProjectManage/models"
	"net/http"

	"github.com/sirupsen/logrus"
)

// LoginController ..
type LoginController struct {
	BaseController
}

// Login ...
// @router / [get]
func (c *LoginController) Login() {
	c.TplName = "login.html"
}

// Validate ...
// @router / [post]
func (c *LoginController) Validate() {
	utype := c.GetString("type")
	pwd := c.GetString("password")
	num, err := c.GetInt64("num")
	if err != nil {
		logrus.Errorf("get id fail:[%v]", err)
		c.ServeError(http.StatusBadRequest, err)
		return
	}

	if err := models.ValidateUser(utype, num, pwd); err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, nil)
}

// Mapper ...
// @router /mapper [post]
func (c *LoginController) Mapper() {
	w := c.Ctx.ResponseWriter
	sess, _ := globalSessions.SessionStart(w, c.Ctx.Request)
	defer sess.SessionRelease(w)
	utype := c.GetString("type")
	num, _ := c.GetInt64("user")

	if err := sess.Set("type", utype); err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}
	if err := sess.Set("user", num); err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}
	logrus.Infof("[%s]:[%d] login", utype, num)

	switch utype {
	case "教师":
		c.Ctx.Redirect(302, "http://localhost:8080/project/teacher")
	case "学院管理员":

	case "信息化建设管理员":

	case "专家":

	}
}
