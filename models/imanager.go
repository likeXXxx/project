package models

import (
	"ProjectManage/db"
	"strconv"
	"strings"

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
	Status       string `json:"status"`
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
			Status:       projects[i].Status,
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
	_, err := o.QueryTable("project").Filter("status__in", StatusICenterVerify, StatusMasterVerify).All(&projects)
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

// IPassProject ...
func IPassProject(id int, instruction string, masterInfo string) error {
	o := db.GetOrmer()
	var err error

	project := db.Project{ID: id}
	if err := o.Read(&project); err != nil {
		logrus.Errorln(err)
		return err
	}
	project.Status = StatusMasterVerify
	project.IAuditInstruction = instruction
	if err != nil {
		logrus.Errorln(err)
		return err
	}
	if _, err := o.Update(&project, "status", "i_audit_instruction"); err != nil {
		logrus.Errorln(err)
		return err
	}

	masters := strings.Split(masterInfo, ",")
	for i := 0; i < len(masters); i++ {
		masterID, _ := strconv.ParseInt(strings.Split(masters[i], " ")[0], 10, 64)
		masterAudit := db.MasterAudit{
			ProjectID: id,
			MasterID:  masterID,
			Status:    StatusWaitToAudit,
		}
		if _, err := o.Insert(&masterAudit); err != nil {
			logrus.Errorln(err)
			return err
		}
	}

	return nil
}

// OrganizationName ...
const OrganizationName = "信息化建设推进办公室"

func convertProjectToIManagerAbolitionProject(project *db.Project, imanager *db.IManager) *db.AbolitionProject {
	return &db.AbolitionProject{
		ID:                    project.ID,
		Name:                  project.Name,
		Organization:          project.Organization,
		TeacherID:             project.TeacherID,
		CreateTime:            project.CreateTime,
		AbolitionOrganization: OrganizationName,
		Operator:              imanager.Name,
		OperatorTel:           imanager.Tel,
	}
}

//IAbolitionProject ...
func IAbolitionProject(projectID int, instruction string, iID int64) error {
	o := db.GetOrmer()

	project := db.Project{ID: projectID}
	if err := o.Read(&project); err != nil {
		logrus.Errorln(err)
		return err
	}
	imanager, err := db.GetIMByID(iID)
	if err != nil {
		logrus.Errorln(err)
		return err
	}
	abolitionProject := convertProjectToIManagerAbolitionProject(&project, imanager)
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
