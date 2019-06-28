package controllers

import (
	"fmt"
	"iReferral/models"

	"github.com/astaxie/beego/orm"
	//"golang.org/x/crypto/bcrypt"
)

type AdminSUPController struct {
	MainController
}

func (this *AdminSUPController) CreateAdmin() {
	username := "cetric"
	password := "Cet12345"
	o := orm.NewOrm()
	o.Using("default")

	pass, _ := HashPassword(password) //hash the submitted password
	//insert into the database the details
	admin := models.Admin{Username: username, Password: pass}
	_, err := o.Insert(&admin)
	if err != nil {
		fmt.Println("Internal server error")
	}
this.Redirect("/", 302)
}
