package controllers

import (
	"encoding/json"
	"fmt"
	"iReferral/models"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type MgnLogInController struct {
	MainController
}

type MgnErrorJson struct {
	Response    string
	NationalId  string
	MgnPassword string
}
var ID, Uname string
var myresponsejson MgnErrorJson

func (this *MgnLogInController) Get() {
	this.AdminAuth()
}

func (this *MgnLogInController) Post() {
	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	nationalId := dataform["nationalId"]
	password := dataform["password"]

	o := orm.NewOrm()
	o.Using("default")
	n := nationalId.(string)
	p := password.(string)

	//check if the huduma number exists
	exist := o.QueryTable("admin_account").Filter("Id", nationalId).Exist()
	mgn := models.Admin_account{Id: n}
	err := o.Read(&mgn, "Id")
	err = bcrypt.CompareHashAndPassword([]byte(mgn.Password), []byte(p))
	if exist == true && err == nil {
		session := this.StartSession()
		session.Set("UserID", n)
		ID = n
		Uname = mgn.Username
		//retrieve the patient first name and last name
		myresponsejson.Response = "/adminhome" //redirect to patient home page
		obj, _ := json.Marshal(myresponsejson)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)

	} else {
		//send ajax response that the password entered is invalid
		myresponsejson.NationalId = "valid"
		myresponsejson.MgnPassword = "valid"
		if exist == false {
			myresponsejson.NationalId = "invalid"
		}
		if err != nil {
			myresponsejson.MgnPassword = "invalid"
		}
		obj, _ := json.Marshal(myresponsejson)
		this.Ctx.Output.SetStatus(300)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)
	}

}
