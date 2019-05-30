package controllers

import (
	"encoding/json"
	"fmt"
	"iReferral/models"

	"github.com/astaxie/beego/orm"
	//"golang.org/x/crypto/bcrypt"
)

type StaffSignUpController struct {
	MainController
}

type StaffError struct {
	Resp string
	StaffEmpId    string
	StaffEmail    string
	StaffPhoneNo  string
	StaffPassword string
	StaffCopass   string
}

var staffjson StaffError

func (this *StaffSignUpController) Get() {
	this.home("home")
}

func (this *StaffSignUpController) Post() {
	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	empid := dataform["emp"].(string)
	email := dataform["em"].(string)
	phone := dataform["phone"].(string) //
	pass := dataform["mypass"].(string)     //
	copass := dataform["comypass"].(string) //

	o := orm.NewOrm()
	o.Using("default")
	//check if the huduma number exists
	exist := o.QueryTable("employee_account").Filter("EmpId", empid).Exist()
	exist1 := o.QueryTable("employee").Filter("EmpId", empid).Exist()
	if exist == false && exist1 == true && pass == copass && IsValid(pass) == true && PhoneValid(phone) == true && EmailValid(email) == true {
		password, _ := HashPassword(pass) //hash the submitted password
		//insert into the database the details
		staff := models.Employee_account{EmpId: empid, Email: email, PhoneNo: phone, Password: password}
			_, err := o.Insert(&staff)
			if err != nil {
			fmt.Println("Internal server error")
		}
		staffjson.Resp = "/" //redirect to log in page
		obj, _ := json.Marshal(staffjson)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)

	} else {
		//send ajax response that the password entered is invalid
		if exist == true {
			staffjson.StaffEmpId = "invalid"
		}

		if exist1 == false {
			staffjson.StaffEmpId = "incorrect"
		}

		if pass != copass {
			staffjson.StaffPassword = "invalid"
			staffjson.StaffCopass = "invalid"
		}

		if IsValid(pass) == false {
			staffjson.StaffPassword = "incorrect"
		}

		if EmailValid(email) == false {
			staffjson.StaffEmail = "wrong"
		}

		if PhoneValid(phone) == false {
			staffjson.StaffPhoneNo = "wrong"
		}

		obj, _ := json.Marshal(staffjson)
		this.Ctx.Output.SetStatus(300)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)
	}

}
