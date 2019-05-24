package controllers

import (
	"fmt"
	"iReferral/models"
	"github.com/astaxie/beego/orm"
	"github.com/twinj/uuid"
	"strings"
)

type ReferralsController struct {
	MainController
}

var Hosname, Date, Time, Myhuduma, Fname, Lname string

func (this *ReferralsController) PatReferral() {
	
	this.patReferral("patientreferral")
	if this.Ctx.Input.Method() == "POST" {
		name := this.GetString("name")
		date := this.GetString("date")
		time := this.GetString("time")

		//input validation== you cannot input empty values
		d := strings.TrimSpace(date)
		t := strings.TrimSpace(time)
		if t =="" && d ==""{
			flash.Error("Date and time fields cannot be left empty")
			flash.Store(&this.Controller)
			return
		}else if d ==""{
			flash.Error("Date field cannot be left empty")
			flash.Store(&this.Controller)
			return
		}else if t ==""{
			flash.Error("Time field cannot be left empty")
			flash.Store(&this.Controller)
			return
		}

		var mhuduma string
		if len(Myhuduma) == 0 {
			mhuduma = Huduma
		} else {
			mhuduma = Myhuduma
		}
		o := orm.NewOrm()
		o.Using("default")
		pat := models.Patient_account{HudumaNo: mhuduma}
		err := o.Read(&pat, "hudumaNo")
		if err != nil {
			fmt.Println("Internal Server error")
			flash.Error("Internal Server error")
			flash.Store(&this.Controller)
		}
		Fname = pat.FirstName
		Lname = pat.LastName
		referral := models.Referrals{HudumaNo: Myhuduma, Service: Service, HosName: name, RDate: date, RTime: time}
		u := uuid.NewV4()
		referral.Id = u.String()
		_, err = o.Insert(&referral)
		if err != nil {
			fmt.Println(err)
			flash.Error(Myhuduma + " already refered")
			flash.Store(&this.Controller)
			return
		}
		Myhuduma = mhuduma
		Hosname = name
		Date = date
		Time = time
		this.Redirect("/confirmreferral", 302)

	}

}
