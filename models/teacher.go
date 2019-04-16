package models

import (
	"ProjectManage/db"
	"time"

	"github.com/astaxie/beego/orm"

	"github.com/sirupsen/logrus"
)

// TmpProjectsResp ...
type TmpProjectsResp struct {
	Total int           `json:"total,omitempty"`
	Rows  []TmpProjects `json:"rows,omitempty"`
}

// TmpProjects ...
type TmpProjects struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	CreateTime  time.Time `json:"create_time,omitempty"`
	Budget      int       `json:"budget,omitempty"`
	InviteWay   string    `json:"invite_way,omitempty"`
	Instruction string    `json:"instruction,omitempty"`
}

// ResetTeacherPwd ...
func ResetTeacherPwd(newPwd string, id int64) error {
	orm := db.GetOrmer()

	teacher := db.Teacher{ID: id}
	if err := orm.Read(&teacher); err != nil {
		logrus.Errorln(err)
		return err
	}
	teacher.Pwd = newPwd
	if _, err := orm.Update(&teacher, "Pwd"); err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}

func convertProjectToTmpProjects(projects []db.Project) []TmpProjects {
	resp := make([]TmpProjects, 0, len(projects))
	for i := 0; i < len(projects); i++ {
		tmpProject := TmpProjects{
			ID:          projects[i].ID,
			Name:        projects[i].Name,
			CreateTime:  projects[i].CreateTime,
			Budget:      projects[i].Budget,
			InviteWay:   projects[i].InviteWay,
			Instruction: projects[i].Instruction,
		}
		resp = append(resp, tmpProject)
	}
	return resp
}

// GetTempProjects ...
func GetTempProjects(teacherID int64) (*TmpProjectsResp, error) {
	o := db.GetOrmer()

	var projects []db.Project
	_, err := o.QueryTable("project").Exclude("invite_way", StatusRunning).Exclude("invite_way", StatusFinish).Filter("teacher_id", teacherID).All(&projects)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}

		logrus.Errorln(err)
		return nil, err
	}

	tmpProjects := convertProjectToTmpProjects(projects)
	tmpProjectsResp := &TmpProjectsResp{
		Total: len(tmpProjects),
		Rows:  tmpProjects,
	}

	return tmpProjectsResp, nil
}

// CreateProjectReq ...
type CreateProjectReq struct {
	Name         string `json:"name,omitempty"`
	Organization string `json:"organization,omitempty"`
	Instruction  string `json:"instruction,omitempty"`
	Budget       int    `json:"budget,omitempty"`
	Inviteway    string `json:"inviteway,omitempty"`
}

// CreateProject ...
func CreateProject(projectReq *CreateProjectReq) error {
	o := db.GetOrmer()

	project := db.Project{
		Name:         projectReq.Name,
		Organization: projectReq.Organization,
		Instruction:  projectReq.Instruction,
		Budget:       projectReq.Budget,
		InviteWay:    projectReq.Inviteway,
	}

	_, err := o.Insert(&project)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}
