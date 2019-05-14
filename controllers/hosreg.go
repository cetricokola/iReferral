package controllers

import (
	"fmt"
	"iReferral/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/validation"
	//  "golang.org/x/crypto/bcrypt"
	"unicode/utf8"
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
		facility := models.Hospital_account{Code:code, Name: name, SerialNo: serial, PhoneNo: phone, Email: email, Country: country, Region: region, District: district, MgnId:mgnId}
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

func (this *HosregController) Search() {
		this.hospital_reg("hosreg")
	if this.Ctx.Input.Method() == "GET"{
		value := this.GetString("name")
		o := orm.NewOrm()
		o.Using("default")
		hos := models.Hospital_account{Email: value}
		err := o.Read(&hos, "Email")
		if err == orm.ErrNoRows {
			fmt.Println(err)
			fmt.Println("incorrect email")
			err1.Error("You've entered incorrect email address")
			err1.Store(&this.Controller)
			return
		} else if err != nil {
			fmt.Println("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
			err1.Error("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
			err1.Store(&this.Controller)
			return
		} else if utf8.RuneCountInString(value) != 0 {
			fmt.Println("Successful searching FACILITY")
			this.Redirect("/facility_mgn", 302)
		}
		
	}
	
}
