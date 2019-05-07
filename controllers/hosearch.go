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

type FaController struct {
	MainController
}
var err5 = beego.NewFlash()
var err6= beego.NewFlash()

// create service
func (this *FaController) Createservice() {
	this.facility_mgn("facility_mgn")
	if this.Ctx.Input.Method() == "POST" {
		//get the values from the form
		name := this.GetString("name")
		code := this.GetString("code")
		cost := this.GetString("cost")
		slot := this.GetString("slot")
		dep := this.GetString("dep")
		o := orm.NewOrm()
		o.Using("default")
		service := models.Services{Code: code, Name: name, Cost: cost, Slot: slot, Department: dep}
		_, err := o.Insert(&service)
		if err != nil {
			fmt.Println(err)
			err5.Error(name + " already registered")
			err5.Store(&this.Controller)
			//this.Redirect("/", 302)
			return
		}
		fmt.Print("successfull registration SERVICE")
		this.Redirect("/facility_mgn", 302)
	}
}

// service removal
func (this *FaController) Removeservice() {
	this.facility_mgn("facility_mgn")
	if this.Ctx.Input.Method() == "GET" {
		value := this.GetString("code")
		if utf8.RuneCountInString(value) != 0 { //do not process empty values
			o := orm.NewOrm()
			o.Using("default")
			service := models.Services{Code: value}
			err := o.Read(&service, "Code")
			if err == orm.ErrNoRows {
				fmt.Println(err)
				fmt.Println("invalid service code")
				err6.Error("You've entered incorrect service code-doesn't exist")
				err6.Store(&this.Controller)
				return
			} else if err != nil {
				fmt.Println("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
				err6.Error("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
				err6.Store(&this.Controller)
				return
			} else if err == nil {
				// DELETE FROM employee WHERE Id = "user input value on the form"
				num, err := o.QueryTable("services").Filter("code", value).Delete()
				fmt.Printf("Affected Num: %s, %s", num, err)
				fmt.Println("Successful removal of the  Service")
				this.Redirect("/facility_mgn", 302)
			}
		}
	}
}