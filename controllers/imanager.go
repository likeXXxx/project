package controllers

import (
	"ProjectManage/db"
	"ProjectManage/models"
	"fmt"
	"net/http"
	"regexp"

	"github.com/sirupsen/logrus"
)

// IManagerController ...
type IManagerController struct {
	BaseController
}

// Get ...
// @router / [get]
func (c *IManagerController) Get() {
	c.TplName = "imanager.html"
}

// GetInfo ...
// @router /getinfo [get]
func (c *IManagerController) GetInfo() {
	logrus.Infof("get user info url: [%s]", c.Ctx.Input.URI())

	name, err := models.GetUserName(c.uID, "信息化建设管理员")
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, name)
}

// ResetPassword ...
// @router /pwd [put]
func (c *IManagerController) ResetPassword() {
	logrus.Infof("imanager[%s] reset pwd url: [%s]", c.uID, c.Ctx.Input.URI())

	oldPwd := c.GetString("old")
	newPwd := c.GetString("new")

	imanager, err := db.GetIMByID(c.uID)
	if err != nil {
		logrus.Errorf("get imanager info fail:[%v]", err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}
	if imanager.Pwd != oldPwd {
		logrus.Errorf("imanager old pwd is wrong")
		c.ServeError(http.StatusBadRequest, fmt.Errorf("原密码错误!请输入正确的密码"))
		return
	}

	if ok, _ := regexp.MatchString("^[0-9a-zA-Z_]{8,20}$", newPwd); !ok {
		logrus.Errorf("new password is invalid")
		c.ServeError(http.StatusBadRequest, fmt.Errorf("新密码应由8-20位字母/数字/下划线组成"))
		return
	}

	if err := models.ResetIManagerPwd(newPwd, c.uID); err != nil {
		logrus.Errorf("reset imanager pwd fail: [%v]", err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, nil)
}

// GetApplyProjects ...
// @router /project/apply [get]
func (c *IManagerController) GetApplyProjects() {
	logrus.Infof("imanager get apply projects url: [%s]", c.Ctx.Input.URI())

	applyProjectsResp, err := models.GetOrganizationApplyProjects()
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

// IManagerGetProjectDetail ...
// @router /project/detail [get]
func (c *IManagerController) IManagerGetProjectDetail() {
	logrus.Infof("imanager get  project detail url: [%s]", c.Ctx.Input.URI())

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

// IManagerProjectPass ...
// @router /project/pass [post]
func (c *IManagerController) IManagerProjectPass() {
	logrus.Infof("imanager pass  project url: [%s]", c.Ctx.Input.URI())

	instruction := c.GetString("instruction")
	masterInfo := c.GetString("master")
	id, err := c.GetInt("id")
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusBadRequest, err)
		return
	}

	if err := models.IPassProject(id, instruction, masterInfo); err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, nil)
}

// ImanagerProjectFail ...
// @router /project/fail [post]
func (c *IManagerController) ImanagerProjectFail() {
	logrus.Infof("imanager fail to pass  project url: [%s]", c.Ctx.Input.URI())

	instruction := c.GetString("instruction")
	id, err := c.GetInt("id")
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusBadRequest, err)
		return
	}

	if err := models.IAbolitionProject(id, instruction, c.uID); err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, nil)
}

// GetMasterAuditResult ...
// @router /master/audit [get]
func (c *IManagerController) GetMasterAuditResult() {
	logrus.Infof("imanager get master audit result, url: [%s]", c.Ctx.Input.URI())

	id, err := c.GetInt("id")
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusBadRequest, err)
		return
	}

	result, err := models.GetMasterAuditResult(id)
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, result)
}

// FinAuditPass ...
// @router /finaudit/pass [post]
func (c *IManagerController) FinAuditPass() {
	logrus.Infof("imanager final pass  project url: [%s]", c.Ctx.Input.URI())

	instruction := c.GetString("instruction")
	id, err := c.GetInt("id")
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusBadRequest, err)
		return
	}
	finFunds, err := c.GetInt("funds")
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusBadRequest, err)
		return
	}

	if err := models.FinAuditPass(instruction, finFunds, id); err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, nil)
}

// FinAuditFail ...
// @router /finaudit/fail [post]
func (c *IManagerController) FinAuditFail() {
	logrus.Infof("imanager final fail project url: [%s]", c.Ctx.Input.URI())

	instruction := c.GetString("instruction")
	id, err := c.GetInt("id")
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusBadRequest, err)
		return
	}

	if err := models.FinAuditFail(instruction, id, c.uID); err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, nil)
}

// ApplyChangeInviteProject ...
// @router /project/invite/change [get]
func (c *IManagerController) ApplyChangeInviteProject() {
	logrus.Infof("imanager get apply change invite projects url: [%s]", c.Ctx.Input.URI())

	applyProjectsResp, err := models.GetChangeInviteProjects()
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

// PassChangeProjectApply ...
// @router /project/invite/change/pass [post]
func (c *IManagerController) PassChangeProjectApply() {
	logrus.Infof("imanager pass change invite projects apply url: [%s]", c.Ctx.Input.URI())

	id, err := c.GetInt("id")
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusBadRequest, err)
		return
	}

	if err := models.PassChangeProjectApply(id); err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, nil)
}

// FailChangeProjectApply ...
// @router /project/invite/change/fail [post]
func (c *IManagerController) FailChangeProjectApply() {
	logrus.Infof("imanager fail change invite projects apply url: [%s]", c.Ctx.Input.URI())

	id, err := c.GetInt("id")
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusBadRequest, err)
		return
	}

	if err := models.FailChangeProjectApply(id); err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, nil)
}
