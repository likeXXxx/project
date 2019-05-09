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

	var masterAudits []db.MasterAudit
	_, err := o.QueryTable("master_audit").Filter("status", StatusWaitToAudit).Filter("master_id", id).All(&masterAudits)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}

		logrus.Errorln(err)
		return nil, err
	}
	if len(masterAudits) == 0 {
		return nil, nil
	}

	ids := make([]int, len(masterAudits))
	for i := 0; i < len(masterAudits); i++ {
		ids = append(ids, masterAudits[i].ProjectID)
	}
	var projects []db.Project
	_, err = o.QueryTable("project").Filter("id__in", ids).All(&projects)
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

// MasterPassProject ...
func MasterPassProject(id int, instruction string, funds int, mID int64) error {
	o := db.GetOrmer()

	masterAudit := db.MasterAudit{ProjectID: id, MasterID: mID}
	if err := o.Read(&masterAudit); err != nil {
		logrus.Errorln(err)
		return err
	}
	masterAudit.MAuditInstruction = instruction
	masterAudit.FinFunds = funds
	masterAudit.Result = ResultPass
	masterAudit.Status = StatusFinish
	if _, err := o.QueryTable("master_audit").Filter("project_id", id).Filter("master_id", mID).Update(orm.Params{
		"m_audit_instruction": instruction, "result": ResultPass, "status": StatusFinish, "fin_funds": funds,
	}); err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}

//MAbolitionProject ...
func MAbolitionProject(projectID int, instruction string, mID int64) error {
	o := db.GetOrmer()

	masterAudit := db.MasterAudit{ProjectID: projectID, MasterID: mID}
	if err := o.Read(&masterAudit); err != nil {
		logrus.Errorln(err)
		return err
	}
	masterAudit.MAuditInstruction = instruction
	masterAudit.Result = ResultFail
	masterAudit.Status = StatusFinish
	if _, err := o.QueryTable("master_audit").Filter("project_id", projectID).Filter("master_id", mID).Update(orm.Params{
		"m_audit_instruction": instruction, "result": ResultFail, "status": StatusFinish,
	}); err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}
