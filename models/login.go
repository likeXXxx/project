package models

import (
	"ProjectManage/db"
	"fmt"

	"github.com/sirupsen/logrus"
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
	teacher, err := db.GetTeacherByID(num)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	if teacher.Pwd != pwd {
		return fmt.Errorf("密码错误，请重新输入！")
	}
	println(teacher)

	return nil
}

func validateOM(num int64, pwd string) error {
	om, err := db.GetOMByID(num)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	if om.Pwd != pwd {
		return fmt.Errorf("密码错误，请重新输入！")
	}

	return nil
}

func validateIM(num int64, pwd string) error {
	im, err := db.GetIMByID(num)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	if im.Pwd != pwd {
		return fmt.Errorf("密码错误，请重新输入！")
	}

	return nil
}

func validateMaster(num int64, pwd string) error {
	master, err := db.GetMasterByID(num)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	if master.Pwd != pwd {
		return fmt.Errorf("密码错误，请重新输入！")
	}

	return nil
}
