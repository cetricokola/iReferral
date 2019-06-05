package controllers

import (
	"encoding/json"
	"fmt"
	"iReferral/models"

	"github.com/astaxie/beego/orm"
)

type CreateServController struct {
	MainController
}

type CreateServJson struct {
	ServCreate string
	ServName   string
	ServCode   string
	ServDept   string
}

func (this *CreateServController) Get() {
	this.Facility_management()
}

func (this *CreateServController) Post() {

	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	name := dataform["name"].(string)
	code := dataform["code"].(string)
	dep := dataform["dep"].(string)

	o := orm.NewOrm()
	o.Using("default")

	exist := o.QueryTable("services").Filter("service_code", code).Exist()
	var responsejson CreateServJson
	if exist == false && ServiceValid(code) == true && dep != "Select department" {
		session := this.StartSession()
		userID := session.Get("UserID")
		var hosCode string
		o.Raw("select code from hospital_account where mgn_id=?", userID.(string)).QueryRow(&hosCode)
		service := models.Services{ServiceCode: code, Code: hosCode, Name: name, Department: dep}
		_, err := o.Insert(&service)
		if err != nil {
			responsejson.ServCreate = "/successful_update"
			obj, _ := json.Marshal(responsejson)
			this.Ctx.Output.Header("Content-Type", "application/json")
			this.Ctx.Output.Body(obj)
			panic(err)
		}
		responsejson.ServCreate = "/successful_update"
		obj, _ := json.Marshal(responsejson)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)
	} else {

		if exist == true {
			responsejson.ServCode = "invalid"
		}
		if ServiceValid(code) == false {
			responsejson.ServCode = "incorrect"
		}

		if name == "" {
			responsejson.ServName = "invalid"
		}

		if dep == "Select department" {
			responsejson.ServDept = "noselection"
		}

		
		obj, _ := json.Marshal(responsejson)
		this.Ctx.Output.SetStatus(300)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)

	}

}
