package controllers

import (
	"fmt"
	"iReferral/models"

	// "github.com/astaxie/beego"
	"unicode/utf8"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type DocController struct {
	MainController
}
var Huduma string

func (this *DocController) FindPatient() {
	session := this.StartSession()
	userID := session.Get("UserID")

	if userID == nil {
		this.Redirect("/", 302)
		return
	}
	fmt.Println("Session id is:", userID)
	//retrieve the name of the doctor from database using the session id
	user := userID.(string)
	o := orm.NewOrm()
	o.Using("default")
	doc := models.Employee{EmpId: user}
	err := o.Read(&doc, "EmpId")
	if err != nil {
		fmt.Println("Internal Server error")
	}
	this.Data["First"] = doc.FirstName
	this.Data["Last"] = doc.LastName
	this.doctor_portal("doctor")
	if this.Ctx.Input.Method() == "GET" {
		value := this.GetString("huduma")
		//validate the user input
		valid := validation.Validation{}
		valid.Required(value, "Huduma Number")
		valid.Numeric(value, "huduma Number")    //numeric values for huduma are permitted
		valid.Length(value, 11, "huduma Number") //11 digits are permitted
		if valid.HasErrors() {
			errormap := []string{}
			for _, err := range valid.Errors {
				errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")

			}
			this.Data["Errors"] = errormap
			return
		}

		if utf8.RuneCountInString(value) != 0 {
			o := orm.NewOrm()
			o.Using("default")
			hos := models.Patient_account{HudumaNo: value}
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
			} else if err == nil {
				fmt.Println("Successful searching patient")
				Huduma = value
				this.Redirect("/report", 302)

			}

		}

	}
}
