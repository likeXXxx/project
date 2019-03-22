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

	c.ServeOK(SuccessVal, "true")
}
