package controllers

import (
	"fmt"
	"iReferral/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/validation"
	//  "golang.org/x/crypto/bcrypt"
)

type HosregController struct {
	MainController
}

var err1 = beego.NewFlash()
var err2 = beego.NewFlash()

func (this *HosregController) Create() {
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID == nil {
		this.Redirect("/auth/a-login", 302)
		return
	}
	fmt.Println("Logged in user is", userID)
	mgnId := userID.(string)
	this.hospital_reg("hosreg")
	if this.Ctx.Input.Method() == "POST" {
		//get the values from the form
		code := this.GetString("code")
		name := this.GetString("name")
		serial := this.GetString("serial")
		phone := this.GetString("phone")
		email := this.GetString("email")
		country := this.GetString("count")
		region := this.GetString("regio")
		district := this.GetString("dist")
		fmt.Println(region)

		o := orm.NewOrm()
		o.Using("default")

		facility := models.Hospital_account{Code: code, Name: name, SerialNo: serial, PhoneNo: phone, Email: email, Country: country, Region: region, District: district, MgnId: mgnId}
		_, err := o.Insert(&facility)
		if err != nil {
			fmt.Println(err)
			err2.Error(serial + " already registered")
			err2.Store(&this.Controller)
			//this.Redirect("/", 302)
			return
		}
		fmt.Print("successfull registration FACILITY")
		this.Redirect("/facility_mgn", 302)

	}
}

func (this *HosregController) CheckReg() {
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID == nil {
		this.Redirect("/auth/a-login", 302)
		return
	}
	fmt.Println("Logged in user is", userID)
	mgnId := userID.(string)
	o := orm.NewOrm()
	o.Using("default")
	exist := o.QueryTable("hospital_account").Filter("MgnId", mgnId).Exist()
	if exist == true {
		this.Redirect("/unsuccessfulrequest", 302)
	} else {
		this.Redirect("/hosreg", 302)
	}
}
