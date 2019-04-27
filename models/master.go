package models

import (
	"ProjectManage/db"

	"github.com/astaxie/beego/orm"
	"github.com/sirupsen/logrus"
)

// ListMasterResp ...
type ListMasterResp struct {
	Name string `json:"name,omitempty"`
	ID   int64  `json:"id,omitempty"`
}

// ListMaster ...
func ListMaster() ([]*ListMasterResp, error) {
	o := db.GetOrmer()

	var masters []db.Master
	_, err := o.QueryTable("master").All(&masters)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}

		logrus.Errorln(err)
		return nil, err
	}

	mastersResp := make([]*ListMasterResp, 0, len(masters))
	for i := 0; i < len(masters); i++ {
		masterResp := &ListMasterResp{
			Name: masters[i].Name,
			ID:   masters[i].ID,
		}
		mastersResp = append(mastersResp, masterResp)
	}

	return mastersResp, nil
}
