package controllers

import (
	"ProjectManage/models"
	"net/http"

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
