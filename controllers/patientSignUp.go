package controllers

import (
	"encoding/json"
	"fmt"
	"iReferral/models"

	"github.com/astaxie/beego/orm"
	//"golang.org/x/crypto/bcrypt"
)

type PatSignUpController struct {
	MainController
}

type PatErrorJson struct {
	Res            string
	PatFirstName   string
	PatLastName    string
	PatHudumaNo    string
	PatDoB         string
	PatPhoneNo     string
	PatPatPassword string
	Copass         string
	PatSex         string
}

var patjson PatErrorJson

func (this *PatSignUpController) Get() {
	this.home("home")
}

func (this *PatSignUpController) Post() {
	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	first := dataform["first"].(string)
	last := dataform["last"].(string)
	huduma := dataform["huduma"].(string) //
	phone := dataform["phone"].(string)   //
	dob := dataform["dob"].(string)
	sex := dataform["sex"].(string)
	pass := dataform["pass"].(string)     //
	copass := dataform["copass"].(string) //

	o := orm.NewOrm()
	o.Using("default")
	//check if the huduma number exists
	exist := o.QueryTable("patient_account").Filter("HudumaNo", huduma).Exist()
	if exist == false && pass == copass && IsValid(pass) == true && PhoneValid(phone) == true && HudumaValid(huduma) == true {
		password, _ := HashPassword(pass) //hash the submitted password
		//insert into the database the details
		patient := models.Patient_account{FirstName: first, LastName: last, HudumaNo: huduma, DoB: dob, PhoneNo: phone, Password: password, Sex: sex}
		_, err := o.Insert(&patient)
		if err != nil {
			fmt.Println("Internal server error")
		}
		patjson.Res = "/" //redirect to log in page
		obj, _ := json.Marshal(patjson)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)

	} else {
		//send ajax response that the password entered is invalid
		if exist == true {
			patjson.PatHudumaNo = "invalid"
		}
		
		if first == ""{
			patjson.PatFirstName = "empty"
		}

		if last == ""{
			patjson.PatLastName = "empty"
		}

		if dob == ""{
			patjson.PatDoB = "empty"
		}

		if pass != copass {
			patjson.PatPatPassword = "invalid"
			patjson.Copass = "invalid"
		}

		if IsValid(pass) == false {
			patjson.PatPatPassword = "incorrect"
		}

		if HudumaValid(huduma) == false {
			patjson.PatHudumaNo = "wrong"
		}

		if PhoneValid(phone) == false {
			patjson.PatPhoneNo = "wrong"
		}

		obj, _ := json.Marshal(patjson)
		this.Ctx.Output.SetStatus(300)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)
	}

}
