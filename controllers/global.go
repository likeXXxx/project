package controllers

import (
	"ProjectManage/models"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GlobalController ...
type GlobalController struct {
	BaseController
}

// Get ...
// @router / [get]
func (c *GlobalController) Get() {
	c.TplName = "project.html"
}

// GlobalGetPassProject ...
// @router /passproject [get]
func (c *GlobalController) GlobalGetPassProject() {
	logrus.Infof("global get pass projects url: [%s]", c.Ctx.Input.URI())

	passProjectsResp, err := models.GlobalGetPassProject()
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	resp := make(map[string]interface{})
	resp["rows"] = passProjectsResp
	resp["total"] = len(passProjectsResp)
	c.Data["json"] = resp
	c.ServeJSON()
}

// GlobalGetRunningProject ...
// @router /runningproject [get]
func (c *GlobalController) GlobalGetRunningProject() {
	logrus.Infof("global get running projects url: [%s]", c.Ctx.Input.URI())

	runningProjectsResp, err := models.GlobalGetRunningProject()
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	resp := make(map[string]interface{})
	resp["rows"] = runningProjectsResp
	resp["total"] = len(runningProjectsResp)
	c.Data["json"] = resp
	c.ServeJSON()
}

// GlobalGetFinishedProject ...
// @router /finishedproject [get]
func (c *GlobalController) GlobalGetFinishedProject() {
	logrus.Infof("global get finished projects url: [%s]", c.Ctx.Input.URI())

	finishedProjectsResp, err := models.GlobalGetFinishedProject()
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	resp := make(map[string]interface{})
	resp["rows"] = finishedProjectsResp
	resp["total"] = len(finishedProjectsResp)
	c.Data["json"] = resp
	c.ServeJSON()
}
