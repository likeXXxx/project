package models

import (
	"ProjectManage/db"
	"fmt"

	"github.com/astaxie/beego/orm"
)

// ValidateUser ...
func ValidateUser(utype string, num int64, pwd string) error {
	var err error
	switch utype {
	case "教师":
		err = validateTeacher(num, pwd)
	case "学院管理员":
		err = validateOM(num, pwd)
	case "信息化建设管理员":
		err = validateIM(num, pwd)
	case "专家":
		err = validateMaster(num, pwd)
	}

	return err
}

func validateTeacher(num int64, pwd string) error {
	o := db.GetOrmer()
	info := o.QueryTable("teacher").Filter("id", num)
	var teacher db.Teacher
	if err := info.One(&teacher); err != nil {
		if err == orm.ErrNoRows {
			return fmt.Errorf("该帐号不存在！")
		}

		return err
	}

	if teacher.Pwd != pwd {
		return fmt.Errorf("密码错误，请重新输入！")
	}

	return nil
}

func validateOM(num int64, pwd string) error {
	o := db.GetOrmer()
	info := o.QueryTable("o_manager").Filter("id", num)
	var om db.OManager
	if err := info.One(&om); err != nil {
		if err == orm.ErrNoRows {
			return fmt.Errorf("该帐号不存在！")
		}

		return err
	}

	if om.Pwd != pwd {
		return fmt.Errorf("密码错误，请重新输入！")
	}

	return nil
}

func validateIM(num int64, pwd string) error {
	o := db.GetOrmer()
	info := o.QueryTable("i_manager").Filter("id", num)
	var im db.IManager
	if err := info.One(&im); err != nil {
		if err == orm.ErrNoRows {
			return fmt.Errorf("该帐号不存在！")
		}

		return err
	}

	if im.Pwd != pwd {
		return fmt.Errorf("密码错误，请重新输入！")
	}

	return nil
}

func validateMaster(num int64, pwd string) error {
	o := db.GetOrmer()
	info := o.QueryTable("master").Filter("id", num)
	var master db.Master
	if err := info.One(&master); err != nil {
		if err == orm.ErrNoRows {
			return fmt.Errorf("该帐号不存在！")
		}

		return err
	}

	if master.Pwd != pwd {
		return fmt.Errorf("密码错误，请重新输入！")
	}

	return nil
}
