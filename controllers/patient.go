package controllers

import (
	//"iReferral/models"
	// //"github.com/astaxie/beego"
	"fmt"

	"github.com/astaxie/beego/orm"
)

type PatientController struct {
	MainController
}

type Report struct {
	Weight        string
	Temperature   string
	BloodPressure string
	Diagnosis     string
	Response      string
	Prescription  string
	Reg_date      string
}

type Myreferrals struct {
	HosName string
	Service  string
	RDate    string
	RTime    string
}

var Repo []Report
var MyRef []Myreferrals
var Leng int

// function for printing the medical report
func (this *PatientController) Preport() {
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID == nil {
		this.Redirect("/auth/p_login", 302)
		return
	}
	myHd := userID.(string)
		if this.Ctx.Input.Method() == "GET" {
		o := orm.NewOrm()
		o.Using("default")
		o.Raw("SELECT reg_date, temperature, weight, blood_pressure, diagnosis, prescription FROM patient_diagnosis WHERE huduma_no=? ORDER BY reg_date", myHd).QueryRows(&Repo)
		fmt.Println("The user id is", myHd)
		fmt.Println("This is my med history:=", Repo)
		this.preport("preport")
			}
}

// function for printing the referrals

func (this *PatientController) Preferral() {
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID == nil {
		this.Redirect("/auth/p_login", 302)
		return
	}
	myHd := userID.(string)
	this.preferral("preferral")
	if this.Ctx.Input.Method() == "POST" {
		start := this.GetString("start")
		end := this.GetString("end")
		o := orm.NewOrm()
		o.Using("default")
		if Leng > 0 {
			MyRef = nil
		}
		o.Raw("SELECT hos_name, service, r_date, r_time FROM referrals WHERE huduma_no=? AND r_date BETWEEN ? AND ? ORDER BY r_date", myHd, start, end).QueryRows(&MyRef)
		len := len(MyRef)
		Leng = len
		this.Redirect("/myreferrals", 302)
	}
}