package models

import (
	"ProjectManage/db"
	"fmt"
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
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	CreateTime  string `json:"create_time,omitempty"`
	Budget      int    `json:"budget,omitempty"`
	Purpose     string `json:"purpose,omitempty"`
	InviteWay   string `json:"invite_way,omitempty"`
	PFunction   string `json:"p_function,omitempty"`
	Instruction string `json:"instruction,omitempty"`
	Status      string `json:"status,omitempty"`
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
			CreateTime:  projects[i].CreateTime.Format("2006-01-02"),
			Budget:      projects[i].Budget,
			InviteWay:   projects[i].InviteWay,
			Instruction: projects[i].Instruction,
			Status:      projects[i].Status,
			Purpose:     projects[i].Purpose,
			PFunction:   projects[i].PFunction,
		}
		resp = append(resp, tmpProject)
	}
	return resp
}

// GetTempProjects ...
func GetTempProjects(teacherID int64) (*TmpProjectsResp, error) {
	o := db.GetOrmer()

	var projects []db.Project
	_, err := o.QueryTable("project").Exclude("status", StatusRunning).Exclude("status", StatusFinish).Filter("teacher_id", teacherID).All(&projects)
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

// AbolitionProjectResp ...
type AbolitionProjectResp struct {
	ID                    int    `json:"id,omitempty"`
	Name                  string `json:"name,omitempty"`
	CreateTime            string `json:"create_time,omitempty"`
	AbolitionOrganization string `json:"abolition_organization,omitempty"`
	AbolitionInstruction  string `json:"abolition_instruction,omitempty"`
	Operator              string `json:"operator,omitempty"`
	OperatorTel           string `json:"operator_tel,omitempty"`
}

func convertAbolitionProjectsToAbolitionProjectResp(abolitionProjects []db.AbolitionProject) []*AbolitionProjectResp {
	abolitionProjectsResp := make([]*AbolitionProjectResp, 0, len(abolitionProjects))
	for i := 0; i < len(abolitionProjects); i++ {
		abolitionProject := &AbolitionProjectResp{
			ID:                    abolitionProjects[i].ID,
			Name:                  abolitionProjects[i].Name,
			CreateTime:            abolitionProjects[i].CreateTime.Format("2006-01-02"),
			AbolitionOrganization: abolitionProjects[i].AbolitionOrganization,
			AbolitionInstruction:  abolitionProjects[i].AbolitionInstr0uction,
			Operator:              abolitionProjects[i].Operator,
			OperatorTel:           abolitionProjects[i].OperatorTel,
		}
		abolitionProjectsResp = append(abolitionProjectsResp, abolitionProject)
	}

	return abolitionProjectsResp
}

// GetAbolitionProjects ...
func GetAbolitionProjects(teacherID int64) ([]*AbolitionProjectResp, error) {
	o := db.GetOrmer()
	var abolitionProjects []db.AbolitionProject
	_, err := o.QueryTable("abolition_project").Filter("teacher_id", teacherID).All(&abolitionProjects)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}

		logrus.Errorln(err)
		return nil, err
	}

	abolitionProjectsResp := convertAbolitionProjectsToAbolitionProjectResp(abolitionProjects)
	return abolitionProjectsResp, nil
}

// CreateProjectReq ...
type CreateProjectReq struct {
	Name         string `json:"name,omitempty"`
	Organization string `json:"organization,omitempty"`
	Instruction  string `json:"instruction,omitempty"`
	Budget       int    `json:"budget,omitempty"`
	Inviteway    string `json:"inviteway,omitempty"`
	Purpose      string `json:"purpose,omitempty"`
	PFunction    string `json:"p_function,omitempty"`
	Result       string `json:"result,omitempty"`
	TeacherID    int64  `json:"teacher_id,omitempty"`
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
		TeacherID:    projectReq.TeacherID,
		Status:       StatusSchoolVerify,
		Purpose:      projectReq.Purpose,
		PFunction:    projectReq.PFunction,
		ExpectResult: projectReq.Result,
	}

	_, err := o.Insert(&project)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}

//DeleteAbolitionProject ...
func DeleteAbolitionProject(id int) error {
	o := db.GetOrmer()

	if _, err := o.Delete(&db.AbolitionProject{ID: id}); err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}

//DeleteProject ...
func DeleteProject(id int) error {
	o := db.GetOrmer()

	if _, err := o.Delete(&db.Project{ID: id}); err == nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}

//TeacherVerifyProject ...
func TeacherVerifyProject(id int, instruction string, fileName string) error {
	o := db.GetOrmer()

	project := db.Project{ID: id}
	if err := o.Read(&project); err != nil {
		logrus.Errorln(err)
		return err
	}
	project.Budget = project.FinFunds
	project.Status = StatusBudget
	if _, err := o.Update(&project, "status", "budget"); err != nil {
		logrus.Errorln(err)
		return err
	}

	inviteProject := db.ProjectInvite{
		ID:             project.ID,
		Funds:          project.FinFunds,
		InviteWay:      project.InviteWay,
		Instruction:    instruction,
		InviteFileName: fileName,
		Name:           project.Name,
		ChangeApply:    "false",
	}

	_, err := o.Insert(&inviteProject)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}

//GetInviteProject ...
func GetInviteProject(id int) (*db.ProjectInvite, error) {
	project, err := db.GetProjectInviteByID(id)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	return project, nil
}

//ChangeInviteProject ...
func ChangeInviteProject(id int, instruction string) error {
	projectInvite, err := db.GetProjectInviteByID(id)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	if projectInvite.ChangeApply == "true" {
		return fmt.Errorf("该招标项目已申请过修改，请耐心等待审批，务重复申请")
	}

	o := db.GetOrmer()
	projectInvite.ChangeApply = "true"
	projectInvite.ChangeReason = instruction
	if _, err := o.Update(projectInvite, "change_reason", "change_apply"); err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}

//ProjectRun ...
func ProjectRun(company string, id, funds int) error {
	o := db.GetOrmer()
	projectInvite, err := db.GetProjectInviteByID(id)
	if err != nil {
		logrus.Errorln(err)
		return err
	}
	projectInvite.CompanyName = company
	projectInvite.FinFunds = funds
	projectInvite.FinTime = time.Now()
	if _, err := o.Update(projectInvite, "company_name", "fin_funds", "fin_time"); err != nil {
		logrus.Errorln(err)
		return err
	}

	project, err := db.GetProjectByID(id)
	project.RunTime = time.Now()
	project.Status = StatusRunning
	if _, err := o.Update(project, "run_time", "status"); err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}

//RunningProjectResp ..
type RunningProjectResp struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	RunTime       string `json:"run_time,omitempty"`
	CompanyName   string `json:"company_name,omitempty"`
	FinFunds      int    `json:"fin_funds,omitempty"`
	LeftoverFunds int    `json:"leftover_funds,omitempty"`
}

func getProjectLeftFunds(id int, totalFunds int) (int, error) {
	o := db.GetOrmer()

	var projects []db.ProjectEvent
	_, err := o.QueryTable("project_event").Filter("project_id", id).All(&projects)
	if err != nil {
		if err == orm.ErrNoRows {
			return totalFunds, nil
		}

		logrus.Errorln(err)
		return totalFunds, err
	}
	if len(projects) == 0 {
		return totalFunds, nil
	}

	result := totalFunds
	for i := 0; i < len(projects); i++ {
		result = result - projects[i].UseFunds
	}
	return result, nil
}

func convertProjectToRunningProjectResp(projects []db.Project) ([]RunningProjectResp, error) {
	if len(projects) == 0 {
		return make([]RunningProjectResp, 0, 0), nil
	}

	projectResp := make([]RunningProjectResp, 0, len(projects))
	for i := 0; i < len(projects); i++ {
		runningProject := RunningProjectResp{
			ID:       projects[i].ID,
			Name:     projects[i].Name,
			RunTime:  projects[i].RunTime.Format("2006-01-02"),
			FinFunds: projects[i].FinFunds,
		}

		projectInvite, err := db.GetProjectInviteByID(projects[i].ID)
		if err != nil {
			logrus.Errorln(err)
			return nil, err
		}
		runningProject.CompanyName = projectInvite.CompanyName

		runningProject.LeftoverFunds, err = getProjectLeftFunds(projects[i].ID, projects[i].FinFunds)
		if err != nil {
			logrus.Errorln(err)
			return nil, err
		}

		projectResp = append(projectResp, runningProject)
	}

	return projectResp, nil
}

//GetRunningProjects ...
func GetRunningProjects(teacherID int64) ([]RunningProjectResp, error) {
	o := db.GetOrmer()

	var projects []db.Project
	_, err := o.QueryTable("project").Filter("teacher_id", teacherID).Filter("status", StatusRunning).All(&projects)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}
		logrus.Errorln(err)
		return nil, err
	}

	return convertProjectToRunningProjectResp(projects)
}

//RunningProjectAddEvent ...
func RunningProjectAddEvent(instruction string, funds, id int) error {
	project, err := db.GetProjectByID(id)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	projectEvent := db.ProjectEvent{
		ProjectID:   id,
		Name:        project.Name,
		Time:        time.Now(),
		UseFunds:    funds,
		Instruction: instruction,
	}
	o := db.GetOrmer()
	if _, err := o.Insert(&projectEvent); err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}

//RunningProjectEvent ...
type RunningProjectEvent struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Time        string `json:"time"`
	UseFunds    int    `json:"use_funds"`
	Instruction string `json:"instruction"`
}

//ListRunningProjectEvent ...
func ListRunningProjectEvent(id int) ([]RunningProjectEvent, error) {
	o := db.GetOrmer()

	var projectEvents []db.ProjectEvent
	if _, err := o.QueryTable("project_event").Filter("project_id", id).All(&projectEvents); err != nil {
		if err == orm.ErrNoRows {
			return make([]RunningProjectEvent, 0, 0), nil
		}

		return nil, err
	}

	resp := make([]RunningProjectEvent, 0, len(projectEvents))
	for i := 0; i < len(projectEvents); i++ {
		runningProjectEvent := RunningProjectEvent{
			ID:          projectEvents[i].ProjectID,
			Name:        projectEvents[i].Name,
			Time:        projectEvents[i].Time.Format("2006-01-02"),
			UseFunds:    projectEvents[i].UseFunds,
			Instruction: projectEvents[i].Instruction,
		}
		resp = append(resp, runningProjectEvent)
	}

	return resp, nil
}

//RunningProjectFinish ...
func RunningProjectFinish(id int, selfEvaluation, completionStatus string) error {
	o := db.GetOrmer()

	project, err := db.GetProjectByID(id)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	leftFunds, err := getProjectLeftFunds(id, project.FinFunds)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	project.Status = StatusFinish
	project.UsedFunds = project.FinFunds - leftFunds
	project.FinTime = time.Now()
	project.SelfEvaluation = selfEvaluation
	project.CompletionStatus = completionStatus
	if _, err := o.Update(project, "used_funds", "status", "fin_time", "completion_status", "self_evaluation"); err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}

//FinishedProjectResp ..
type FinishedProjectResp struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	CreateTime string `json:"create_time"`
	RunTime    string `json:"run_time,omitempty"`
	FinishTime string `json:"fin_time,omitempty"`
	FinFunds   int    `json:"fin_funds,omitempty"`
	UsedFunds  int    `json:"used_funds,omitempty"`
}

//GetFinishedProjects ...
func GetFinishedProjects(id int64) ([]FinishedProjectResp, error) {
	o := db.GetOrmer()

	var projects []db.Project
	_, err := o.QueryTable("project").Filter("teacher_id", id).Filter("status", StatusFinish).All(&projects)
	if err != nil {
		if err == orm.ErrNoRows {
			return make([]FinishedProjectResp, 0, 0), nil
		}
		logrus.Errorln(err)
		return nil, err
	}

	resp := make([]FinishedProjectResp, 0, len(projects))
	for i := 0; i < len(projects); i++ {
		finProject := FinishedProjectResp{
			ID:         projects[i].ID,
			Name:       projects[i].Name,
			RunTime:    projects[i].RunTime.Format("2006-01-02"),
			CreateTime: projects[i].CreateTime.Format("2006-01-02"),
			FinishTime: projects[i].FinTime.Format("2006-01-02"),
			FinFunds:   projects[i].FinFunds,
			UsedFunds:  projects[i].UsedFunds,
		}
		resp = append(resp, finProject)
	}

	return resp, nil
}
