package db

import (
	"fmt"
	"net/url"
	"strconv"
	"sync"

	"github.com/astaxie/beego/orm"
)

// MySQL mysql
type MySQL struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

var mysql = MySQL{
	Host:     "172.17.0.1",
	Port:     33306,
	Username: "root",
	Password: "password",
	Database: "project",
}

// InitMysql 初始化mysql
func InitMysql() {
	sql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=%s", mysql.Username,
		mysql.Password, mysql.Host, strconv.Itoa(mysql.Port), mysql.Database, url.QueryEscape("Asia/Shanghai"))

	orm.RegisterDataBase("default", "mysql", sql)
	orm.RegisterModel(new(Teacher), new(OManager), new(Organization), new(IManager), new(Master), new(Project), new(AbolitionProject), new(ProjectInvite), new(MasterAudit), new(ProjectEvent))
	orm.RunSyncdb("default", false, true)
}

var (
	globalOrm orm.Ormer
	once      sync.Once
)

// GetOrmer ...
func GetOrmer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}
