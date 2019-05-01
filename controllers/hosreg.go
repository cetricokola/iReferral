package controllers

import (
	"fmt"
	"iReferral/models"

	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/validation"
	// "golang.org/x/crypto/bcrypt"
)

type HosregController struct {
	MainController
}

func (this *HosregController) Post() {
	this.hospital_reg("hosreg")
	// if this.Ctx.Input.Method() == "POST" {
	//get the values from the form
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

	facility := models.Hospital_account{Name: name, SerialNo: serial, PhoneNo: phone, Email: email, Country: country, Region: region, District: district}

	_, err := o.Insert(&facility)
	if err != nil {
		fmt.Println(err)
		flash.Error(serial + " already registered")
		flash.Store(&this.Controller)
		//this.Redirect("/", 302)
		return
	}
	fmt.Print("successfull registration")
	this.Redirect("/", 302)
	// }
}

func (this *HosregController) Get() {
	this.hospital_reg("hosreg")

	value := this.GetString("name")
	o := orm.NewOrm()
	o.Using("default")
	hos := models.Hospital_account{Email: value}
	err := o.Read(&hos, "Email")

	if err == orm.ErrNoRows{
		fmt.Println(err)
		fmt.Println("incorrect email")
		flash.Error("You've entered incorrect email address")
		flash.Store(&this.Controller)
		return

	} else if err != nil {
		fmt.Println("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
		flash.Error("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
		flash.Store(&this.Controller)
		return
	}

	fmt.Println("Successful searching")
	this.Redirect("/", 302)
}
