package controllers

import (
	"iReferral/models"
	// //"github.com/astaxie/beego"
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Referrals struct {
	HudumaNo string
	Service  string
	RDate    string
	RTime    string
}

var Referral []Referrals
var Hdm []string
var Serv []string
var Dat []string
var Tim []string
var Len int

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
	/*Hdm = nil
	Serv = nil
	Dat = nil
	Tim = nil
	Len = 0*/
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
		o.Raw("SELECT huduma_no, service, r_date, r_time FROM referrals WHERE hos_name=? AND r_date BETWEEN ? AND ? ORDER BY r_date", mycode, start, end).QueryRows(&Referral)
		len := len(Referral)
		fmt.Println("the length of the struct", len)
		fmt.Println("patient", Referral)
		fmt.Println("My ids ", myId)
		fmt.Println("My names ", mycode)
		fmt.Println(end)
		if len == 0 {
			Referral = nil
			Len = len
		} else {
		
		for i := 0; i < len; i++ {
			hdm := Referral[i].HudumaNo
			dat := Referral[i].RDate
			tim := Referral[i].RTime
			serv := Referral[i].Service
			Hdm = append(Hdm, hdm)
			Serv = append(Serv, serv)
			Dat = append(Dat, dat)
			Tim = append(Tim, tim)

		}
		
			Len = len
		}
		fmt.Println("Hudumas are the following", Serv, Dat, Tim, Hdm)
		this.Redirect("/viewreferrals", 302)

	}
}
