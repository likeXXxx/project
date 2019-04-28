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

// ResetMasterPwd ...
func ResetMasterPwd(newPwd string, id int64) error {
	o := db.GetOrmer()

	master := db.Master{ID: id}
	if err := o.Read(&master); err != nil {
		logrus.Errorln(err)
		return err
	}
	master.Pwd = newPwd
	if _, err := o.Update(&master, "Pwd"); err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}

// MasterApplyProjectResp ...
type MasterApplyProjectResp struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Organization string `json:"organization,omitempty"`
	CreateTime   string `json:"create_time,omitempty"`
	Budget       int    `json:"budget,omitempty"`
	InviteWay    string `json:"invite_way,omitempty"`
	Instruction  string `json:"instruction,omitempty"`
	TeacherName  string `json:"teacher_name,omitempty"`
	TeacherTel   string `json:"teacher_tel,omitempty"`
}

func convertProjectToMasterApplyProjectResp(projects []db.Project) ([]*MasterApplyProjectResp, error) {
	resp := make([]*MasterApplyProjectResp, 0, len(projects))
	for i := 0; i < len(projects); i++ {
		applyProjectResp := MasterApplyProjectResp{
			ID:           projects[i].ID,
			Name:         projects[i].Name,
			Organization: projects[i].Organization,
			CreateTime:   projects[i].CreateTime.Format("2006-01-02"),
			Budget:       projects[i].Budget,
			InviteWay:    projects[i].InviteWay,
			Instruction:  projects[i].Instruction,
		}
		teacher, err := db.GetTeacherByID(projects[i].TeacherID)
		if err != nil {
			return nil, err
		}
		applyProjectResp.TeacherName = teacher.Name
		applyProjectResp.TeacherTel = teacher.Tel
		resp = append(resp, &applyProjectResp)
	}
	return resp, nil
}

//GetMasterApplyProjects ...
func GetMasterApplyProjects(id int64) ([]*MasterApplyProjectResp, error) {
	o := db.GetOrmer()

	var projects []db.Project
	_, err := o.QueryTable("project").Filter("status", StatusMasterVerify).Filter("master_id", id).All(&projects)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}

		logrus.Errorln(err)
		return nil, err
	}

	tmpProjects, err := convertProjectToMasterApplyProjectResp(projects)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	return tmpProjects, nil
}
