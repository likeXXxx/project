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
	utype := c.GetString("type")
	num, _ := c.GetInt64("usr")
	c.SetSession("UID", num)
	c.SetSession("TYPE", utype)
	switch utype {
	case "教师":
		c.TplName = "teacher.html"
	case "学院管理员":

	case "信息化建设管理员":

	case "专家":

	}
}
