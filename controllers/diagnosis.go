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
	this.diagnosis("report")
	if this.Ctx.Input.Method() == "POST" {
		huduma := this.GetString("huduma")
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
			HudumaNo:      huduma,
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
