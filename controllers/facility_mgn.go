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

type FacilityController struct {
	MainController
}
var err3 = beego.NewFlash()
var err4 = beego.NewFlash()


// EMP registration
func (this *FacilityController) Createemp() {
		this.facility_mgn("facility_mgn")
	if this.Ctx.Input.Method() == "POST" {
		//get the values from the form
		first := this.GetString("first")
		last := this.GetString("last")
		position := this.GetString("position")
		nId := this.GetString("nId")
		o := orm.NewOrm()
		o.Using("default")
		emp := models.Employee{FirstName: first, LastName: last, Position: position, EmpId: nId}
		_, err := o.Insert(&emp)
		if err != nil {
			fmt.Println(err)
			err3.Error(nId + " already registered")
			err3.Store(&this.Controller)
			return
		}
		fmt.Print("successfull registration Employee")
		this.Redirect("/facility_mgn", 302)
	}
}

// employee removal function
func (this *FacilityController) RemoveEmp() {
	this.facility_mgn("facility_mgn")
	if this.Ctx.Input.Method() == "GET" {
		value := this.GetString("id")
		if utf8.RuneCountInString(value) != 0 { //do not process empty values
			o := orm.NewOrm()
			o.Using("default")
			emp := models.Employee{EmpId: value}
			err := o.Read(&emp, "NationalId")
			if err == orm.ErrNoRows {
				fmt.Println(err)
				fmt.Println("incorrect national id")
				err4.Error("You've entered incorrect ID number")
				err4.Store(&this.Controller)
				return
			} else if err != nil {
				fmt.Println("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
				err4.Error("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
				err4.Store(&this.Controller)
				return
			} else if err == nil {
				// DELETE FROM employee WHERE Id = "user input value on the form"
				num, err := o.QueryTable("employee").Filter("national_id", value).Delete()
				fmt.Printf("Affected Num: %s, %s", num, err)
				fmt.Println("Successful removal of the  Employee")
				this.Redirect("/facility_mgn", 302)
			}
		}
	}
}
