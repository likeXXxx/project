package controllers

import (
	"ProjectManage/db"
	"ProjectManage/models"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

// TeacherController ..
type TeacherController struct {
	BaseController
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

	uInfo, err := models.GetUserName(c.uID, "教师")
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, strings.Split(uInfo, ","))
}

// ResetPassword ...
// @router /pwd [put]
func (c *TeacherController) ResetPassword() {
	logrus.Infof("teacher reset pwd url: [%s]", c.uID, c.Ctx.Input.URI())

	oldPwd := c.GetString("old")
	newPwd := c.GetString("new")

	teacher, err := db.GetTeacherByID(c.uID)
	if err != nil {
		logrus.Errorf("get teacher info fail:[%v]", err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}
	if teacher.Pwd != oldPwd {
		logrus.Errorf("teacher old pwd is wrong")
		c.ServeError(http.StatusBadRequest, fmt.Errorf("原密码错误!请输入正确的密码"))
		return
	}

	if ok, _ := regexp.MatchString("^[0-9a-zA-Z_]{8,20}$", newPwd); !ok {
		logrus.Errorf("new password is invalid")
		c.ServeError(http.StatusBadRequest, fmt.Errorf("新密码应由8-20位字母/数字/下划线组成"))
		return
	}

	if err := models.ResetTeacherPwd(newPwd, c.uID); err != nil {
		logrus.Errorf("reset teacher pwd fail: [%v]", err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, nil)
}

// GetTempProjects ...
// @router /project/temp [get]
func (c *TeacherController) GetTempProjects() {
	logrus.Infof("teacher[%s] get temp projects url: [%s]", c.uID, c.Ctx.Input.URI())

	tmpProjectsResp, err := models.GetTempProjects(c.uID)
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	resp := make(map[string]interface{})
	resp["rows"] = tmpProjectsResp.Rows
	resp["total"] = tmpProjectsResp.Total
	c.Data["json"] = resp
	c.ServeJSON()
}

// GetAbolitionProjects ...
// @router /project/abolition [get]
func (c *TeacherController) GetAbolitionProjects() {
	logrus.Infof("teacher[%s] get abolition projects url: [%s]", c.uID, c.Ctx.Input.URI())

	abolitionProjectsResp, err := models.GetAbolitionProjects(c.uID)
	if err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	resp := make(map[string]interface{})
	resp["rows"] = abolitionProjectsResp
	resp["total"] = len(abolitionProjectsResp)
	c.Data["json"] = resp
	c.ServeJSON()
}

// CreateProject ...
// @router /project [post]
func (c *TeacherController) CreateProject() {
	logrus.Infof("teacher create project, url:[%s], body:[%s]", c.Ctx.Input.URI(), string(c.Ctx.Input.RequestBody))

	var project models.CreateProjectReq
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &project); err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusBadRequest, err)
		return
	}
	project.TeacherID = c.uID

	if err := models.CreateProject(&project); err != nil {
		logrus.Errorln(err)
		c.ServeError(http.StatusInternalServerError, err)
		return
	}

	c.ServeOK(SuccessVal, nil)
}
