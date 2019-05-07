package controllers

import (
	"ProjectManage/db"
	"ProjectManage/models"
	"fmt"
	"net/http"
	"regexp"

	"github.com/sirupsen/logrus"
)

// MasterController ...
type MasterController struct {
	BaseController
}

// ListMaster ...
// @router /list [get]
func (c *MasterController) ListMaster() {
	logrus.Infof("list masters url:[%s]", c.Ctx.Input.URI())

	masters, err := models.ListMaster()
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, masters)
}

// Get ...
// @router / [get]
func (c *MasterController) Get() {
	c.TplName = "master.html"
}

// GetInfo ...
// @router /getinfo [get]
func (c *MasterController) GetInfo() {
	logrus.Infof("get user info url: [%s]", c.Ctx.Input.URI())

	name, err := models.GetUserName(c.uID, "专家")
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, name)
}

// ResetPassword ...
// @router /pwd [put]
func (c *MasterController) ResetPassword() {
	logrus.Infof("master[%s] reset pwd url: [%s]", c.uID, c.Ctx.Input.URI())

	oldPwd := c.GetString("old")
	newPwd := c.GetString("new")

	master, err := db.GetMasterByID(c.uID)
	if err != nil {
		logrus.Errorf("get master info fail:[%v]", err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}
	if master.Pwd != oldPwd {
		logrus.Errorf("master old pwd is wrong")
		c.ServeError(http.StatusBadRequest, fmt.Errorf("原密码错误!请输入正确的密码"))
		return
	}

	if ok, _ := regexp.MatchString("^[0-9a-zA-Z_]{8,20}$", newPwd); !ok {
		logrus.Errorf("new password is invalid")
		c.ServeError(http.StatusBadRequest, fmt.Errorf("新密码应由8-20位字母/数字/下划线组成"))
		return
	}

	if err := models.ResetMasterPwd(newPwd, c.uID); err != nil {
		logrus.Errorf("reset master pwd fail: [%v]", err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, nil)
}

// GetApplyProjects ...
// @router /project/apply [get]
func (c *MasterController) GetApplyProjects() {
	logrus.Infof("master get apply projects url: [%s]", c.Ctx.Input.URI())

	applyProjectsResp, err := models.GetMasterApplyProjects(c.uID)
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

// MasterGetProjectDetail ...
// @router /project/detail [get]
func (c *MasterController) MasterGetProjectDetail() {
	logrus.Infof("master get  project detail url: [%s]", c.Ctx.Input.URI())

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

// MasterProjectPass ...
// @router /project/pass [post]
func (c *MasterController) MasterProjectPass() {
	logrus.Infof("manager pass  project url: [%s]", c.Ctx.Input.URI())

	instruction := c.GetString("instruction")
	id, err := c.GetInt("id")
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusBadRequest, err)
		return
	}
	funds, err := c.GetInt("funds")
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusBadRequest, err)
		return
	}

	if err := models.MasterPassProject(id, instruction, funds, c.uID); err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, nil)
}

// MasterProjectFail ...
// @router /project/fail [post]
func (c *MasterController) MasterProjectFail() {
	logrus.Infof("master fail to pass  project url: [%s]", c.Ctx.Input.URI())

	instruction := c.GetString("instruction")
	id, err := c.GetInt("id")
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusBadRequest, err)
		return
	}

	if err := models.MAbolitionProject(id, instruction, c.uID); err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, nil)
}
