package main

import (
	_ "iReferral/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/session"
	
)

func init() {
	
	orm.RegisterDriver("mysql", orm.DRMySQL)                                                       //mapping database driver and database name
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/ireferral?charset=utf8") //connecting to the database

}

func main() {
	sessionconf := &session.ManagerConfig{
		CookieName: "sessionID",
		Gclifetime: 3600,
	}
	beego.GlobalSessions, _ = session.NewManager("memory", sessionconf)
	go beego.GlobalSessions.GC()
	beego.Run()
}
