package models

import (
	"ProjectManage/db"

	"github.com/astaxie/beego/orm"
	"github.com/sirupsen/logrus"
)

// ResetIManagerPwd ...
func ResetIManagerPwd(newPwd string, id int64) error {
	o := db.GetOrmer()

	imanager := db.IManager{ID: id}
	if err := o.Read(&imanager); err != nil {
		logrus.Errorln(err)
		return err
	}
	imanager.Pwd = newPwd
	if _, err := o.Update(&imanager, "Pwd"); err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}

// OrganizationApplyProjectResp ...
type OrganizationApplyProjectResp struct {
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

func convertProjectToOrganizationApplyProjectResp(projects []db.Project) ([]*OrganizationApplyProjectResp, error) {
	resp := make([]*OrganizationApplyProjectResp, 0, len(projects))
	for i := 0; i < len(projects); i++ {
		applyProjectResp := OrganizationApplyProjectResp{
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

//GetOrganizationApplyProjects ...
func GetOrganizationApplyProjects() ([]*OrganizationApplyProjectResp, error) {
	o := db.GetOrmer()

	var projects []db.Project
	_, err := o.QueryTable("project").Filter("status", StatusICenterVerify).All(&projects)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}

		logrus.Errorln(err)
		return nil, err
	}

	tmpProjects, err := convertProjectToOrganizationApplyProjectResp(projects)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	return tmpProjects, nil
}
