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
		if len(MyHuduma) == 0 {
			mhuduma = Huduma
		} else {
			mhuduma = MyHuduma
		}
		o := orm.NewOrm()
		o.Using("default")
		pat := models.Patient_account{HudumaNo: mhuduma}
		err := o.Read(&pat, "HudumaNo")
		if err != nil {
			fmt.Println("Internal Server error")
			flash.Error("Internal Server error")
			flash.Store(&this.Controller)
		}
		Fname = pat.FirstName
		Lname = pat.LastName
		fmt.Println(pat.HudumaNo)
		fmt.Println(Sid)
		//find the hospital code of the employee
		emp := models.Employee{EmpId: Sid}
		err = o.Read(&emp, "EmpId")
		if err != nil {
			fmt.Println("Internal Server error")
			flash.Error("Internal Server error")
			flash.Store(&this.Controller)
		}
		fmt.Print(emp.Code)
		//read the name, email and phone of the referring hospital
		hos := models.Hospital_account{Code: emp.Code}
		err = o.Read(&hos, "Code")
		if err != nil {
			fmt.Println("Internal Server error")
			flash.Error("Internal Server error")
			flash.Store(&this.Controller)
		}

		rhos := hos.Name
		rby := Sid
		email := hos.Email
		phone := hos.PhoneNo
		fmt.Println(rhos)
		referral := models.Referrals{Service: Service, HosName: name, RDate: date, RTime: time, Refer_hos: rhos, Refer_by: rby, Email: email, Phone: phone}
		u := uuid.NewV4()
		referral.HudumaNo = mhuduma
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
