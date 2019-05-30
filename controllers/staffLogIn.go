package controllers

import (
	"encoding/json"
	"iReferral/models"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type StaffLogInController struct {
	MainController
}

//create a structure to hold contents of response back to ajax
type StaffErrorJson struct {
	Success       string
	StaffPassword string
	EmpId         string
}

var Sid string

//render the home page
func (this *StaffLogInController) Get() {
	this.home("home")
}

//perform the authentication
func (this *StaffLogInController) Post() {
	//create a map for the ajax request
	var dataform map[string]interface{}
	//decode the ajax request
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	//get the inputs from the form
	empid := dataform["empid"]
	password := dataform["password"]
	//create an orm object
	o := orm.NewOrm()
	o.Using("default")
	//type case the inputs to the strings
	h := empid.(string)
	p := password.(string)

	//check if the huduma number exists
	exist := o.QueryTable("employee_account").Filter("EmpId", empid).Exist()
	//read the employee account
	staff := models.Employee_account{EmpId: h}
	err := o.Read(&staff, "EmpId")

	//check the password in the database with the submitted password
	err = bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(p))
	// declare a variable of type StaffErrorJson
	var responsejson StaffErrorJson
	//check the conditions necessary for the log in
	if exist == true && err == nil {
		//set sessions
		session := this.StartSession()
		session.Set("UserID", h)
		Sid = h
		responsejson.Success = "/doctor" //redirect to patient home page
		obj, _ := json.Marshal(responsejson)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj) //send response to ajax

		// log in fails
	} else {
		responsejson.Success = "valid"
		responsejson.EmpId = "valid"
		responsejson.StaffPassword = "valid"
		if exist == false { //the employee id is incorrect
			responsejson.EmpId = "invalid"
		}
		if err != nil { //the passwords does not match
			responsejson.StaffPassword = "invalid"
		}
		//send the response to ajax
		obj, _ := json.Marshal(responsejson)
		this.Ctx.Output.SetStatus(300)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)
	}

}
