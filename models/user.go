package models

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
	return nil
}

func validateOM(num int64, pwd string) error {
	return nil
}

func validateIM(num int64, pwd string) error {
	return nil
}

func validateMaster(num int64, pwd string) error {
	return nil
}
