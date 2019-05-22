package controllers

import (
	"fmt"
	"iReferral/models"

	"github.com/astaxie/beego/orm"
)

type ReferpController struct {
	MainController
}

func(this *ReferpController) Get(){
	this.doctor_portal("doctor")
}
func (this *ReferpController) Post() {
	huduma := this.GetString("huduma")
	service := this.GetString("service")
	o := orm.NewOrm()
	o.Using("default")
	hos := models.Patient_account{HudumaNo: huduma}
	err := o.Read(&hos, "HudumaNo")
	if err == orm.ErrNoRows {
		fmt.Println(err)
		fmt.Println("incorrect huduma number")
		flash.Error("Incorrect Huduma No.Try again!")
		flash.Store(&this.Controller)
		return
	} else if err != nil {
		fmt.Println("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
		flash.Error("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
		flash.Store(&this.Controller)
		return
	}

	//searches if the services are available in the database
	serv := models.Services{Name: service}
	err = o.Read(&serv, "Name")
	if err == orm.ErrNoRows {
		fmt.Println("Service not found")
		flash.Error("Invalid service name or service not found.Try again!!")
		flash.Store(&this.Controller)
		return

	} else if err != nil {
		fmt.Println("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
		flash.Error("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
		flash.Store(&this.Controller)
		return
	}

	this.Redirect("/patientreferralform", 302)

}
