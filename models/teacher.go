package models

import (
	"ProjectManage/db"

	"github.com/sirupsen/logrus"
)

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
