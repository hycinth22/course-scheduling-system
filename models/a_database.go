package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// get config
	driver := web.AppConfig.DefaultString("sqlDriver", "mysql")
	dsn := web.AppConfig.DefaultString("sqlConn", "")
	if dsn == "" {
		panic("configure dsn in app.conf")
	}
	// set default database
	if err := orm.RegisterDataBase("default", driver, dsn); err != nil {
		panic(err)
	}
	// register model
	orm.RegisterModel(new(Clazz))
	orm.RegisterModel(new(Clazzroom))
	orm.RegisterModel(new(College))
	orm.RegisterModel(new(Course))
	orm.RegisterModel(new(Department))
	orm.RegisterModel(new(Instruct))
	orm.RegisterModel(new(Major))
	orm.RegisterModel(new(Schedule), new(ScheduleItem))
	orm.RegisterModel(new(Semester))
	orm.RegisterModel(new(Student))
	orm.RegisterModel(new(Teacher))
	orm.RegisterModel(new(Timespan))
	orm.BootStrap()
	// create table
	if err := orm.RunSyncdb("default", false, true); err != nil {
		panic(err)
	}
}
