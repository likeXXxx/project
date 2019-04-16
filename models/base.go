package models

import (
	"ProjectManage/db"
	"fmt"

	"github.com/sirupsen/logrus"
)

// GetUserName ...
func GetUserName(id int64, utype string) (string, error) {
	switch utype {
	case "教师":
		obj, err := db.GetTeacherByID(id)
		if err != nil {
			logrus.Errorln(err)
			return "", err
		}
		return obj.Name + "," + obj.Organization, nil
	case "学院管理员":
		obj, err := db.GetOMByID(id)
		if err != nil {
			logrus.Errorln(err)
			return "", err
		}
		return obj.Name, nil
	case "信息化建设管理员":
		obj, err := db.GetIMByID(id)
		if err != nil {
			logrus.Errorln(err)
			return "", err
		}
		return obj.Name, nil
	case "专家":
		obj, err := db.GetMasterByID(id)
		if err != nil {
			logrus.Errorln(err)
			return "", err
		}
		return obj.Name, nil
	}

	return "", fmt.Errorf("无此类型用户:%s", utype)
}
