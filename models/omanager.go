package models

import (
	"ProjectManage/db"

	"github.com/astaxie/beego/orm"
	"github.com/sirupsen/logrus"
)

// ResetOManagerPwd ...
func ResetOManagerPwd(newPwd string, id int64) error {
	orm := db.GetOrmer()

	omanager := db.OManager{ID: id}
	if err := orm.Read(&omanager); err != nil {
		logrus.Errorln(err)
		return err
	}
	omanager.Pwd = newPwd
	if _, err := orm.Update(&omanager, "Pwd"); err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}

// ApplyProjectResp ...
type ApplyProjectResp struct {
	Name        string `json:"name,omitempty"`
	CreateTime  string `json:"create_time,omitempty"`
	Budget      int    `json:"budget,omitempty"`
	InviteWay   string `json:"invite_way,omitempty"`
	Instruction string `json:"instruction,omitempty"`
	TeacherName string `json:"teacher_name,omitempty"`
	TeacherTel  string `json:"teacher_tel,omitempty"`
}

func convertProjectToApplyProjects(projects []db.Project) ([]*ApplyProjectResp, error) {
	resp := make([]*ApplyProjectResp, 0, len(projects))
	for i := 0; i < len(projects); i++ {
		applyProjectResp := ApplyProjectResp{
			Name:        projects[i].Name,
			CreateTime:  projects[i].CreateTime.Format("2006-01-02"),
			Budget:      projects[i].Budget,
			InviteWay:   projects[i].InviteWay,
			Instruction: projects[i].Instruction,
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

// GetApplyProjects ...
func GetApplyProjects(ID int64) ([]*ApplyProjectResp, error) {
	o := db.GetOrmer()

	omanager, err := db.GetOMByID(ID)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	var projects []db.Project
	_, err = o.QueryTable("project").Filter("organization", omanager.Organization).Filter("status", StatusSchoolVerify).All(&projects)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}

		logrus.Errorln(err)
		return nil, err
	}

	tmpProjects, err := convertProjectToApplyProjects(projects)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	return tmpProjects, nil
}
