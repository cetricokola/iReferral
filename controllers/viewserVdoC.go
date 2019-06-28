package controllers

import (
	"github.com/astaxie/beego/orm"
	// "fmt"
	"iReferral/models"
)

type ViewservdocController struct {
	MaController
}
type Services struct {
	ServiceCode string
	Name        string
	Department  string
}
type Employee struct {
	FirstName string
	LastName  string
	EmpId     string
	// Code      string
}

var Len2, Len0 int
var Serv []Services
var Emps []Employee

func (this *ViewservdocController) ViewDoc() {
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID == nil {
		this.Redirect("/management_authentication", 302)
		return
	}
	myId := userID.(string)
	// this.Viewservices()
	Emps = nil
	o := orm.NewOrm()
	o.Using("default")
	mgn := models.Hospital_account{MgnId: myId}
	err := o.Read(&mgn, "MgnId")
	if err != nil {
		panic(err)
	}
	code := mgn.Code
	_, err = o.Raw("SELECT emp_id,first_name,last_name FROM employee WHERE code=?", code).QueryRows(&Emps)
	if err != nil {
		panic(err)
		return
	}
	Len0 = len(Emps)
	this.Redirect("/view_doctors", 302)
}

func (this *ViewservdocController) ViewServ() {
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID == nil {
		this.Redirect("/management_authentication", 302)
		return
	}
	myId := userID.(string)
	// this.Viewdoctors()

	Serv = nil
	o := orm.NewOrm()
	o.Using("default")
	mgn := models.Hospital_account{MgnId: myId}
	err := o.Read(&mgn, "MgnId")
	if err != nil {
		panic(err)
	}
	code := mgn.Code
	_, err = o.Raw("SELECT service_code,name,department FROM services WHERE code=?", code).QueryRows(&Serv)
	if err != nil {
		panic(err)
		return
	}
Len2 = len(Serv)
	this.Redirect("/view_services", 302)

}
