package controllers

import (
	"encoding/json"
	"fmt"
	"iReferral/models"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type PatientLogInController struct {
	MainController
}

type PatientErrorJson struct {
	Messages        string
	PatientHuduma   string
	PatientPassword string
}

var Pfname, Plname, PiD string

var responsejson PatientErrorJson

func (this *PatientLogInController) Get() {
	this.home("home")
}

func (this *PatientLogInController) Post() {
	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	huduma := dataform["huduma"]
	password := dataform["password"]
	fmt.Println(password)
	fmt.Println(huduma)
	o := orm.NewOrm()
	o.Using("default")
	h := huduma.(string)
	p := password.(string)

	//check if the huduma number exists
	exist := o.QueryTable("patient_account").Filter("HudumaNo", huduma).Exist()
	patient := models.Patient_account{HudumaNo: h}
	err := o.Read(&patient, "HudumaNo")
	err = bcrypt.CompareHashAndPassword([]byte(patient.Password), []byte(p))
	if exist == true && err == nil {
		session := this.StartSession()
		session.Set("UserID", h)
		PiD = h
		//retrieve the patient first name and last name
		Pfname = patient.FirstName
		Plname = patient.LastName
		responsejson.Messages = "/phome" //redirect to patient home page
		obj, _ := json.Marshal(responsejson)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)

	} else {
		//send ajax response that the password entered is invalid
		responsejson.Messages = "valid"
		responsejson.PatientHuduma = "valid"
		responsejson.PatientPassword = "valid"
		if exist == false {
			responsejson.PatientHuduma = "invalid"
		}
		if err != nil {
			responsejson.PatientPassword = "invalid"
		}
		obj, _ := json.Marshal(responsejson)
		this.Ctx.Output.SetStatus(300)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)
	}

}
