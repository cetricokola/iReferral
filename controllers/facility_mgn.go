package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"

	
)

type FacilityController struct {
	MainController
}


func(this *FacilityController) Iscreated(){
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID == nil {
		this.Redirect("/management_authentication", 302)
		return
	}
	fmt.Println("Logged in user is", userID)
	mgnId := userID.(string)
	o := orm.NewOrm()
	o.Using("default")
	exist := o.QueryTable("hospital_account").Filter("MgnId",mgnId).Exist()
	if exist == true {
		this.Redirect("/facility_mgn", 302)
	} else {
		this.Redirect("/unsuccessfulservicerequest", 302)
	}
}
