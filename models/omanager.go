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
	ID          int    `json:"id,omitempty"`
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
			ID:          projects[i].ID,
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

// OPassProject ...
func OPassProject(id int, instruction string) error {
	orm := db.GetOrmer()

	project := db.Project{ID: id}
	if err := orm.Read(&project); err != nil {
		logrus.Errorln(err)
		return err
	}
	project.Status = StatusICenterVerify
	project.OAuditInstruction = instruction
	if _, err := orm.Update(&project, "status", "o_audit_instruction"); err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}

func convertProjectToAbolitionProject(project *db.Project, omanager *db.OManager) *db.AbolitionProject {
	return &db.AbolitionProject{
		ID:                    project.ID,
		Name:                  project.Name,
		Organization:          project.Organization,
		TeacherID:             project.TeacherID,
		CreateTime:            project.CreateTime,
		AbolitionOrganization: omanager.Organization,
		Operator:              omanager.Name,
		OperatorTel:           omanager.Tel,
	}
}

// OAbolitionProject ...
func OAbolitionProject(projectID int, instruction string, oID int64) error {
	o := db.GetOrmer()

	project := db.Project{ID: projectID}
	if err := o.Read(&project); err != nil {
		logrus.Errorln(err)
		return err
	}
	omanager, err := db.GetOMByID(oID)
	if err != nil {
		logrus.Errorln(err)
		return err
	}
	abolitionProject := convertProjectToAbolitionProject(&project, omanager)
	abolitionProject.AbolitionInstr0uction = instruction

	if _, err := o.Insert(abolitionProject); err != nil {
		logrus.Errorln(err)
		return err
	}

	if _, err := o.Delete(&project); err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}
