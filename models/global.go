package models

import (
	"ProjectManage/db"

	"github.com/astaxie/beego/orm"
	"github.com/sirupsen/logrus"
)

//GlobalPassProject ...
type GlobalPassProject struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	CreateTime   string `json:"create_time,omitempty"`
	Organization string `json:"organization"`
	TeacherName  string `json:"teacher"`
	TeacherTel   string `json:"teacher_tel"`
}

//GlobalGetPassProject ...
func GlobalGetPassProject() ([]GlobalPassProject, error) {
	o := db.GetOrmer()

	var projects []db.Project
	_, err := o.QueryTable("project").Filter("status__in", StatusBudget, StatusVerifyProject).All(&projects)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}

		logrus.Errorln(err)
		return nil, err
	}

	resp := make([]GlobalPassProject, 0, len(projects))
	for i := 0; i < len(projects); i++ {
		teacher, err := db.GetTeacherByID(projects[i].TeacherID)
		if err != nil {
			logrus.Errorln(err)
			return nil, err
		}
		projectResp := GlobalPassProject{
			ID:           projects[i].ID,
			Name:         projects[i].Name,
			CreateTime:   projects[i].CreateTime.Format("2006-01-02"),
			Organization: projects[i].Organization,
			TeacherName:  teacher.Name,
			TeacherTel:   teacher.Tel,
		}
		resp = append(resp, projectResp)
	}

	return resp, nil
}

//GlobalRunningProject ...
type GlobalRunningProject struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	RunTime      string `json:"run_time,omitempty"`
	Organization string `json:"organization"`
	TeacherName  string `json:"teacher"`
	TeacherTel   string `json:"teacher_tel"`
}

//GlobalGetRunningProject ...
func GlobalGetRunningProject() ([]GlobalRunningProject, error) {
	o := db.GetOrmer()

	var projects []db.Project
	_, err := o.QueryTable("project").Filter("status", StatusRunning).All(&projects)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}

		logrus.Errorln(err)
		return nil, err
	}

	resp := make([]GlobalRunningProject, 0, len(projects))
	for i := 0; i < len(projects); i++ {
		teacher, err := db.GetTeacherByID(projects[i].TeacherID)
		if err != nil {
			logrus.Errorln(err)
			return nil, err
		}
		projectResp := GlobalRunningProject{
			ID:           projects[i].ID,
			Name:         projects[i].Name,
			RunTime:      projects[i].RunTime.Format("2006-01-02"),
			Organization: projects[i].Organization,
			TeacherName:  teacher.Name,
			TeacherTel:   teacher.Tel,
		}
		resp = append(resp, projectResp)
	}

	return resp, nil
}

//GlobalFinishedProject ...
type GlobalFinishedProject struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	FinishTime   string `json:"finish_time,omitempty"`
	Organization string `json:"organization"`
	TeacherName  string `json:"teacher"`
	TeacherTel   string `json:"teacher_tel"`
}

//GlobalGetFinishedProject ...
func GlobalGetFinishedProject() ([]GlobalFinishedProject, error) {
	o := db.GetOrmer()

	var projects []db.Project
	_, err := o.QueryTable("project").Filter("status", StatusFinish).All(&projects)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}

		logrus.Errorln(err)
		return nil, err
	}

	resp := make([]GlobalFinishedProject, 0, len(projects))
	for i := 0; i < len(projects); i++ {
		teacher, err := db.GetTeacherByID(projects[i].TeacherID)
		if err != nil {
			logrus.Errorln(err)
			return nil, err
		}
		projectResp := GlobalFinishedProject{
			ID:           projects[i].ID,
			Name:         projects[i].Name,
			FinishTime:   projects[i].FinTime.Format("2006-01-02"),
			Organization: projects[i].Organization,
			TeacherName:  teacher.Name,
			TeacherTel:   teacher.Tel,
		}
		resp = append(resp, projectResp)
	}

	return resp, nil
}
