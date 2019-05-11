package models

import (
	"ProjectManage/db"
	"fmt"
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

//MasterAuditResultResp ...
type MasterAuditResultResp struct {
	MAuditInstruction string `json:"instruction,omitempty"`
	MasterID          int64  `json:"master_id,omitempty"`
	MasterName        string `json:"master_name,omitempty"`
	Result            string `json:"result,omitempty"`
	FinFunds          int    `json:"fin_funds"`
}

func convertMasterAuditToMasterAuditResultResp(masterAudits []db.MasterAudit) ([]MasterAuditResultResp, error) {
	resp := make([]MasterAuditResultResp, 0, len(masterAudits))
	for i := 0; i < len(masterAudits); i++ {
		master, err := db.GetMasterByID(masterAudits[i].MasterID)
		if err != nil {
			logrus.Errorln(err)
			return nil, err
		}

		obj := MasterAuditResultResp{
			MAuditInstruction: masterAudits[i].MAuditInstruction,
			MasterID:          masterAudits[i].MasterID,
			MasterName:        master.Name,
			Result:            masterAudits[i].Result,
			FinFunds:          masterAudits[i].FinFunds,
		}
		resp = append(resp, obj)
	}
	return resp, nil
}

// GetMasterAuditResult ...
func GetMasterAuditResult(id int) ([]MasterAuditResultResp, error) {
	o := db.GetOrmer()

	var masterAudit []db.MasterAudit
	_, err := o.QueryTable("master_audit").Filter("project_id", id).All(&masterAudit)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	if len(masterAudit) == 0 {
		return nil, fmt.Errorf("不存在此id的审核项目")
	}

	for _, info := range masterAudit {
		if info.Status == StatusWaitToAudit {
			return nil, fmt.Errorf("notReady")
		}
	}

	resp, err := convertMasterAuditToMasterAuditResultResp(masterAudit)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	return resp, nil
}

//FinAuditPass ...
func FinAuditPass(instruction string, finFunds, id int) error {
	orm := db.GetOrmer()

	project := db.Project{ID: id}
	if err := orm.Read(&project); err != nil {
		logrus.Errorln(err)
		return err
	}
	project.Status = StatusVerifyProject
	project.IFinInstruction = instruction
	project.FinFunds = finFunds
	if _, err := orm.Update(&project, "status", "i_fin_instruction", "fin_funds"); err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}

// FinAbolitionName ...
const FinAbolitionName = "信息化建设推进办公室(专家审核)"

func convertProjectToFinAbolitionProject(project *db.Project, imanager *db.IManager) *db.AbolitionProject {
	return &db.AbolitionProject{
		ID:                    project.ID,
		Name:                  project.Name,
		Organization:          project.Organization,
		TeacherID:             project.TeacherID,
		CreateTime:            project.CreateTime,
		AbolitionOrganization: FinAbolitionName,
		Operator:              imanager.Name,
		OperatorTel:           imanager.Tel,
	}
}

//FinAuditFail ...
func FinAuditFail(instruction string, id int, iID int64) error {
	o := db.GetOrmer()

	project := db.Project{ID: id}
	if err := o.Read(&project); err != nil {
		logrus.Errorln(err)
		return err
	}
	imanager, err := db.GetIMByID(iID)
	if err != nil {
		logrus.Errorln(err)
		return err
	}
	abolitionProject := convertProjectToFinAbolitionProject(&project, imanager)
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

//GetChangeInviteProjectResp ...
type GetChangeInviteProjectResp struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Funds        int    `json:"funds,omitempty"`
	InviteWay    string `json:"invite_way,omitempty"`
	ChangeReason string `json:"change_reason"`
	TeacherName  string `json:"teacher_name,omitempty"`
	TeacherTel   string `json:"teacher_tel,omitempty"`
}

func convertInviteProjectToGetChangeInviteProjectResp(inviteProjects []db.ProjectInvite) ([]*GetChangeInviteProjectResp, error) {
	resp := make([]*GetChangeInviteProjectResp, 0, len(inviteProjects))
	for i := 0; i < len(inviteProjects); i++ {
		applyProjectResp := GetChangeInviteProjectResp{
			ID:           inviteProjects[i].ID,
			Name:         inviteProjects[i].Name,
			Funds:        inviteProjects[i].Funds,
			InviteWay:    inviteProjects[i].InviteWay,
			ChangeReason: inviteProjects[i].ChangeReason,
		}
		project, err := db.GetProjectByID(inviteProjects[i].ID)
		if err != nil {
			return nil, err
		}
		teacher, err := db.GetTeacherByID(project.TeacherID)

		applyProjectResp.TeacherName = teacher.Name
		applyProjectResp.TeacherTel = teacher.Tel
		resp = append(resp, &applyProjectResp)
	}
	return resp, nil
}

//GetChangeInviteProjects ...
func GetChangeInviteProjects() ([]*GetChangeInviteProjectResp, error) {
	o := db.GetOrmer()

	var projects []db.ProjectInvite
	_, err := o.QueryTable("project_invite").Filter("change_apply", "true").All(&projects)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}

		logrus.Errorln(err)
		return nil, err
	}

	resp, err := convertInviteProjectToGetChangeInviteProjectResp(projects)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	return resp, nil
}

//PassChangeProjectApply ..
func PassChangeProjectApply(id int) error {
	o := db.GetOrmer()

	if _, err := o.Delete(&db.ProjectInvite{ID: id}); err != nil {
		logrus.Errorln(err)
		return err
	}

	project, err := db.GetProjectByID(id)
	if err != nil {
		logrus.Errorln(err)
		return err
	}
	project.Status = StatusVerifyProject

	if _, err := o.Update(project, "status"); err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}

//FailChangeProjectApply ...
func FailChangeProjectApply(id int) error {
	project, err := db.GetProjectInviteByID(id)
	if err != nil {
		logrus.Errorln(err)
		return err
	}
	project.ChangeReason = ""
	project.ChangeApply = "false"

	o := db.GetOrmer()
	if _, err := o.Update(project, "change_reason", "change_apply"); err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}
