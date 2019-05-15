package controllers

import (
	"iReferral/models"
	// //"github.com/astaxie/beego"
	"fmt"

	"github.com/astaxie/beego/orm"
)

var Hdm []string
var Serv []string
var Dat []string
var Tim []string

type ViewReferralsController struct {
	MainController
}

func (this *ViewReferralsController) MyReferrals() {

	//get the id of the logged in user
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID == nil {
		this.Redirect("/auth/a-login", 302)
		return
	}

	myId := userID.(string)
	this.viewReferrals("viewreferrals")
	if this.Ctx.Input.Method() == "POST" {
		start := this.GetString("start")
		end := this.GetString("end")
		o := orm.NewOrm()
		o.Using("default")
		me := models.Hospital_account{MgnId: myId}
		err := o.Read(&me, "MgnId")
		if err != nil {
			fmt.Println("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
			flash.Error("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
			flash.Store(&this.Controller)
			return
		}
		mycode := me.Name
		var obj string
		
		o.Raw("SELECT huduma_no FROM referrals WHERE hos_name=? AND r_date >=? AND r_date <=?", mycode, start, end).QueryRow(&obj)
		fmt.Println("patient", obj)
		fmt.Println("My ids ", myId)
		fmt.Println("My names ", mycode)
		fmt.Println(start)
		fmt.Println(end)
		this.Redirect("/viewreferrals", 302)

	}
}
