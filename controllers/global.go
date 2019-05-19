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

// GetFinishedProjectDetail ...
// @router /finished/detail [get]
func (c *GlobalController) GetFinishedProjectDetail() {
	logrus.Infof("global get finished project detail url: [%s]", c.Ctx.Input.URI())

	id, err := c.GetInt("id")
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusBadRequest, err)
		return
	}

	projectDetail, err := models.GetProjectDetail(id)
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, projectDetail)
}

// GetProjectEventList ...
// @router /project/eventlist [get]
func (c *GlobalController) GetProjectEventList() {
	logrus.Infof("global get project events url: [%s]", c.Ctx.Input.URI())

	id, err := c.GetInt("id")
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusBadRequest, err)
		return
	}

	eventList, err := models.ListRunningProjectEvent(id)
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, eventList)
}
