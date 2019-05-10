package db

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// GetTeacherByID ...
func GetTeacherByID(id int64) (*Teacher, error) {
	o := GetOrmer()
	info := o.QueryTable("teacher").Filter("id", id)
	var teacher Teacher
	if err := info.One(&teacher); err != nil {
		if err == orm.ErrNoRows {
			return nil, fmt.Errorf("该帐号不存在！")
		}

		return nil, err
	}

	return &teacher, nil
}

// GetOMByID ...
func GetOMByID(id int64) (*OManager, error) {
	o := GetOrmer()
	info := o.QueryTable("o_manager").Filter("id", id)
	var om OManager
	if err := info.One(&om); err != nil {
		if err == orm.ErrNoRows {
			return nil, fmt.Errorf("该帐号不存在！")
		}

		return nil, err
	}

	return &om, nil
}

// GetIMByID ...
func GetIMByID(id int64) (*IManager, error) {
	o := GetOrmer()
	info := o.QueryTable("i_manager").Filter("id", id)
	var im IManager
	if err := info.One(&im); err != nil {
		if err == orm.ErrNoRows {
			return nil, fmt.Errorf("该帐号不存在！")
		}

		return nil, err
	}

	return &im, nil
}

// GetMasterByID ...
func GetMasterByID(id int64) (*Master, error) {
	o := GetOrmer()
	info := o.QueryTable("master").Filter("id", id)
	var master Master
	if err := info.One(&master); err != nil {
		if err == orm.ErrNoRows {
			return nil, fmt.Errorf("该帐号不存在！")
		}

		return nil, err
	}

	return &master, nil
}

// GetProjectByID ...
func GetProjectByID(id int) (*Project, error) {
	o := GetOrmer()
	info := o.QueryTable("project").Filter("id", id)
	var project Project
	if err := info.One(&project); err != nil {
		if err == orm.ErrNoRows {
			return nil, fmt.Errorf("该项目不存在！")
		}

		return nil, err
	}

	return &project, nil
}

// GetProjectInviteByID ...
func GetProjectInviteByID(id int) (*ProjectInvite, error) {
	o := GetOrmer()
	info := o.QueryTable("project_invite").Filter("id", id)
	var project ProjectInvite
	if err := info.One(&project); err != nil {
		if err == orm.ErrNoRows {
			return nil, fmt.Errorf("该项目不存在！")
		}

		return nil, err
	}

	return &project, nil
}
