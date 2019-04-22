package models

import (
	"ProjectManage/db"

	"github.com/sirupsen/logrus"
)

// ProjectDetail ...
type ProjectDetail struct {
	ID                int    `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	Organization      string `json:"organization,omitempty"`
	TeacherID         int64  `json:"teacher_id,omitempty"`
	CreateTime        string `json:"create_time,omitempty"`
	Budget            int    `json:"budget,omitempty"`
	Status            string `json:"status,omitempty"`
	FinTime           string `json:"fin_time,omitempty"`
	InviteWay         string `json:"invite_way,omitempty"`
	Instruction       string `json:"instruction,omitempty"`
	RunTime           string `json:"run_time,omitempty"`
	OAuditInstruction string `json:"o_audit_instruction,omitempty"`
	IAuditInstruction string `json:"i_audit_instruction,omitempty"`
	MAuditInstruction string `json:"m_audit_instruction,omitempty"`
}

// ProjectDetailResp ...
type ProjectDetailResp struct {
	Project *ProjectDetail `json:"project,omitempty"`
	Teacher *db.Teacher    `json:"teacher,omitempty"`
}

func convertProjectToProjectDetail(project *db.Project) *ProjectDetail {
	projectDetail := &ProjectDetail{
		ID:                project.ID,
		Name:              project.Name,
		Organization:      project.Organization,
		CreateTime:        project.CreateTime.Format("2006-01-02"),
		Budget:            project.Budget,
		InviteWay:         project.InviteWay,
		Instruction:       project.Instruction,
		TeacherID:         project.TeacherID,
		Status:            project.Status,
		RunTime:           project.RunTime.Format("2006-01-02"),
		FinTime:           project.FinTime.Format("2006-01-02"),
		OAuditInstruction: project.OAuditInstruction,
		IAuditInstruction: project.IAuditInstruction,
		MAuditInstruction: project.MAuditInstruction,
	}
	return projectDetail
}

// GetProjectDetail ...
func GetProjectDetail(id int) (*ProjectDetailResp, error) {
	project, err := db.GetProjectByID(id)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	projectDetail := convertProjectToProjectDetail(project)

	teacher, err := db.GetTeacherByID(project.TeacherID)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	projectDetailResp := &ProjectDetailResp{
		Project: projectDetail,
		Teacher: teacher,
	}
	return projectDetailResp, nil
}
