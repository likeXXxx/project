package models

import (
	"ProjectManage/db"

	"github.com/sirupsen/logrus"
)

// ResetOManagerPwd ...
func ResetOManagerPwd(newPwd string, id int64) error {
	orm := db.GetOrmer()

	omanager := db.OManager{ID: id}
	if err := orm.Read(&omanager); err != nil {
		logrus.Errorln(err)
		return err
	}
	omanager.Pwd = newPwd
	if _, err := orm.Update(&omanager, "Pwd"); err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}
