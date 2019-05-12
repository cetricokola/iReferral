package controllers

import (
	"fmt"
	"iReferral/models"

	"time"

	"github.com/astaxie/beego/orm"

	//"github.com/astaxie/beego/validation"
	"github.com/twinj/uuid"
)

type DiagnosisController struct {
	MainController
}

func (this *DiagnosisController) UpdateReport() {
	//check if the user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")

	if userID == nil {
		this.Redirect("/auth/s_login", 302)
		return
	}
	this.Data["Huduma"] = Huduma
	fmt.Println("The huduma is:", Huduma)
	o := orm.NewOrm()
	o.Using("default")
	patient := models.Patient_account{HudumaNo: Huduma}
	err := o.Read(&patient, "hudumaNo")
	if err != nil {
		fmt.Println("Internal Server error")
		flash.Error("Internal Server error")
		flash.Store(&this.Controller)
	}
	this.Data["FName"] = patient.FirstName
	this.Data["LName"] = patient.LastName

	this.diagnosis("report")
	if this.Ctx.Input.Method() == "POST" {
		weight := this.GetString("weight")
		temp := this.GetString("temp")
		pressure := this.GetString("blood")
		diagnosis := this.GetString("diagnosis")
		response := this.GetString("response")
		prescription := this.GetString("prescription")
		currentTime := time.Now()
		today := currentTime.Format("2006-01-02")
		o := orm.NewOrm()
		o.Using("default")
		diag := models.Patient_diagnosis{
			HudumaNo:      Huduma,
			Weight:        weight,
			Temperature:   temp,
			BloodPressure: pressure,
			Diagnosis:     diagnosis,
			Response:      response,
			Prescription:  prescription,
			Reg_date:      today}
		u := uuid.NewV4()
		diag.Id = u.String()
		_, err := o.Insert(&diag)
		if err != nil {
			fmt.Println(err)
			flash.Error(diag.Id + " already registered")
			flash.Store(&this.Controller)
			return
		}
		fmt.Print("successfull registration update")
		this.Redirect("/healthIssue", 302)
	}
}
