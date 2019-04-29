package db

import "time"

// Teacher ...
type Teacher struct {
	ID                int64     `orm:"column(id);pk" json:"id,omitempty"`
	Name              string    `orm:"column(name)" json:"name,omitempty"`
	Organization      string    `orm:"column(organization)" json:"organization,omitempty"`
	Sex               string    `orm:"column(sex)" json:"sex,omitempty"`
	Birth             time.Time `orm:"column(birth)" json:"birth,omitempty"`
	Tel               string    `orm:"column(tel)" json:"tel,omitempty"`
	Pwd               string    `orm:"column(pwd)" json:"pwd,omitempty"`
	ProfessionalTitle string    `orm:"column(professional_title)" json:"professional_title,omitempty"`
}

// Organization ...
type Organization struct {
	Name string `orm:"column(name);pk"`
}

// OManager ...
type OManager struct {
	ID           int64  `orm:"column(id);pk"`
	Organization string `orm:"column(organization)"`
	Pwd          string `orm:"column(pwd)"`
	Name         string `orm:"column(name)"`
	Tel          string `orm:"column(tel)"`
}

// IManager ...
type IManager struct {
	ID   int64  `orm:"column(id);pk"`
	Name string `orm:"column(name)"`
	Pwd  string `orm:"column(pwd)"`
	Tel  string `orm:"column(tel)"`
}

// Master ...
type Master struct {
	ID   int64  `orm:"column(id);pk"`
	Name string `orm:"column(name)"`
	Pwd  string `orm:"column(pwd)"`
	Tel  string `orm:"column(tel)"`
}

// Project ...
type Project struct {
	ID                int       `orm:"column(id);pk" json:"id,omitempty"`
	Name              string    `orm:"column(name)" json:"name,omitempty"`
	Organization      string    `orm:"column(organization)" json:"organization,omitempty"`
	TeacherID         int64     `orm:"column(teacher_id)" json:"teacher_id,omitempty"`
	CreateTime        time.Time `orm:"column(create_time);auto_now_add;type(date)" json:"create_time,omitempty"`
	Budget            int       `orm:"column(budget)" json:"budget,omitempty"`
	Status            string    `orm:"column(status)" json:"status,omitempty"`
	FinTime           time.Time `orm:"column(fin_time);type(date)" json:"fin_time,omitempty"`
	InviteWay         string    `orm:"column(invite_way)" json:"invite_way,omitempty"`
	Instruction       string    `orm:"column(instruction);type(text)" json:"instruction,omitempty"`
	RunTime           time.Time `orm:"column(run_time);type(date)" json:"run_time,omitempty"`
	OAuditInstruction string    `orm:"column(o_audit_instruction)" json:"o_audit_instruction,omitempty"`
	IAuditInstruction string    `orm:"column(i_audit_instruction)" json:"i_audit_instruction,omitempty"`
	MAuditInstruction string    `orm:"column(m_audit_instruction)" json:"m_audit_instruction,omitempty"`
	MasterID          int64     `orm:"column(master_id)" json:"master_id,omitempty"`
	FinFunds          int       `orm:"column(fin_funds)" json:"fin_funds,omitempty"`
}

//AbolitionProject ...
type AbolitionProject struct {
	ID                    int       `orm:"column(id);pk" json:"id,omitempty"`
	Name                  string    `orm:"column(name)" json:"name,omitempty"`
	Organization          string    `orm:"column(organization)" json:"organization,omitempty"`
	TeacherID             int64     `orm:"column(teacher_id)" json:"teacher_id,omitempty"`
	CreateTime            time.Time `orm:"column(create_time);auto_now_add;type(date)" json:"create_time,omitempty"`
	AbolitionOrganization string    `orm:"column(abolition_organization)"`
	AbolitionInstr0uction string    `orm:"column(abolition_instruction)"`
	Operator              string    `orm:"column(operator)"`
	OperatorTel           string    `orm:"column(operator_tel)"`
}

//ProjectInvite ...
type ProjectInvite struct {
	ID             int       `orm:"column(id);pk" json:"id,omitempty"`
	BeginTime      time.Time `orm:"column(begin_time);auto_now_add;type(date)" json:"begin_time,omitempty"`
	FinTime        time.Time `orm:"column(fin_time);type(date)" json:"fin_time,omitempty"`
	Funds          int       `orm:"column(funds)" json:"funds,omitempty"`
	FinFunds       int       `orm:"column(fin_funds)" json:"fin_funds,omitempty"`
	CompanyName    string    `orm:"column(company_name)" json:"company_name,omitempty"`
	InviteWay      string    `orm:"column(invite_way)" json:"invite_way,omitempty"`
	Instruction    string    `orm:"column(instruction)" json:"instruction,omitempty"`
	InviteFileName string    `orm:"column(invite_file_name)" json:"invite_file_name,omitempty"`
}
