package controllers

import (
	"github.com/twinj/uuid"
	"fmt"
	"iReferral/models"
	
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	
)

type ReferralsController struct {
	MainController
}

func (this *ReferralsController) PatReferral(){
	this.patReferral("patientreferral")
	if this.Ctx.Input.Method() == "POST" {
		name := this.GetString("name")
		date := this.GetString("date")
		time := this.GetString("time")

		//save the details to the database
		o := orm.NewOrm()
		o.Using("default")
		referral := models.Referrals{HudumaNo: Huduma, HosName:name, RDate: date, RTime: time}
		u := uuid.NewV4()
		referral.Id = u.String()
		_, err := o.Insert(&referral)
		if err != nil {
			fmt.Println(err)
			flash.Error(HudumaNo+ " already refered")
			flash.Store(&this.Controller)
			return
		}

		this.Redirect("/doctor", 302)
	}

}