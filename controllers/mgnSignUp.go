package controllers

import (
	"encoding/json"
	"fmt"
	"iReferral/models"

	"github.com/astaxie/beego/orm"
	//"golang.org/x/crypto/bcrypt"
)

type MgnSignUpController struct {
	MainController
}

type MgnSErrorJson struct {
	Feedback  string
	Username  string
	Id        string
	Email     string
	Mpassword string
	Cpassword string
}

var mgnresponsejson MgnSErrorJson

func (this *MgnSignUpController) Get() {
	this.AdminAuth()
}

func (this *MgnSignUpController) Post() {
	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	username := dataform["username"].(string)
	id := dataform["id"].(string)
	email := dataform["email"].(string)
	password := dataform["password"].(string)
	cpassword := dataform["cpassword"].(string)

	o := orm.NewOrm()
	o.Using("default")

	//check if the huduma number exists
	exist := o.QueryTable("admin_account").Filter("Username", username).Exist()
	idexist := o.QueryTable("admin_account").Filter("Id", id).Exist()
	if exist == false && idexist == false && password == cpassword && IsValid(password) == true && EmailValid(email) == true && NationalIdValid(id) == true {
		password, _ := HashPassword(password) //hash the submitted password
		//insert into the database the details
		admin := models.Admin_account{Username: username, Id: id, Email: email, Password: password}
		_, err := o.Insert(&admin)
		if err != nil {
			fmt.Println("Internal server error")
		}
		mgnresponsejson.Feedback = "/management_authentication" //redirect to log in page
		obj, _ := json.Marshal(mgnresponsejson)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)

	} else {
		//send ajax response that the password entered is invalid
		if exist == true {
			mgnresponsejson.Username = "invalid"
		}
		if idexist == true {
			mgnresponsejson.Id = "invalid"
		}
		if password != cpassword {
			mgnresponsejson.Mpassword = "invalid"
			mgnresponsejson.Cpassword = "invalid"
		}

		if IsValid(password) == false {
			mgnresponsejson.Mpassword = "incorrect"
		}
		if EmailValid(email) == false {
			mgnresponsejson.Email = "invalid"
		}

		if NationalIdValid(id) == false {
			mgnresponsejson.Id = "wrong"
		}

		obj, _ := json.Marshal(mgnresponsejson)
		this.Ctx.Output.SetStatus(300)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)
	}

}
