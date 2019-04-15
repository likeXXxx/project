package db

import "time"

// Teacher ...
type Teacher struct {
	ID           int64     `orm:"column(id);pk"`
	Name         string    `orm:"column(name)"`
	Organization string    `orm:"column(organization)"`
	Sex          string    `orm:"column(sex)"`
	Birth        time.Time `orm:"column(birth)"`
	Tel          string    `orm:"column(tel)"`
	Pwd          string    `orm:"column(pwd)"`
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
}

// IManager ...
type IManager struct {
	ID   int64  `orm:"column(id);pk"`
	Name string `orm:"column(name)"`
	Pwd  string `orm:"column(pwd)"`
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
	ID           int       `orm:"column(id);pk"`
	Name         string    `orm:"column(name)"`
	Organization string    `orm:"column(organization)"`
	TeacherID    int64     `orm:"column(teacher_id)"`
	CreateTime   time.Time `orm:"column(create_time);auto_now_add;type(date)"`
	Budget       int       `orm:"column(budget)"`
	Status       string    `orm:"column(status)"`
	FinTime      time.Time `orm:"column(fin_time);type(date)"`
	InviteWay    string    `orm:"column(invite_way)"`
	Instruction  string    `orm:"column(instruction);type(text)"`
}
