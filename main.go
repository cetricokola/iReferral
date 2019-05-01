package main

import (
	_ "iReferral/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	
)

func init() {
	
	orm.RegisterDriver("mysql", orm.DRMySQL)                                                       //mapping database driver and database name
	orm.RegisterDataBase("default", "mysql", "cetride:9988@tcp(127.0.0.1:3306)/ireferral?charset=utf8") //connecting to the database

}

func main() {
	beego.Run()
}
