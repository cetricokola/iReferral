package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/orm"
)

type DelServController struct {
	MainController
}

type DelServJson struct {
	Servdel string
	ServCode  string
}

func (this *DelServController) Get() {
	this.Facility_management()
}

func (this *DelServController) Post() {

	session := this.StartSession()
	userID := session.Get("UserID").(string)
	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	code := dataform["code"].(string)

	o := orm.NewOrm()
	o.Using("default")
	var hosCode, hoscode string
	exist := o.QueryTable("services").Filter("service_code", code).Exist()
	o.Raw("select code from services where service_code=?", code).QueryRow(&hoscode)
	o.Raw("select code from hospital_account where mgn_id=?", userID).QueryRow(&hosCode)
	fmt.Println(hoscode)
	fmt.Println(hosCode)
	var responsejson DelServJson
	if exist == true && ServiceValid(code) == true && hoscode == hosCode {
		_, err := o.QueryTable("services").Filter("service_code",code).Delete()
		if err != nil {
			responsejson.Servdel = "/successful_update"
			obj, _ := json.Marshal(responsejson)
			this.Ctx.Output.Header("Content-Type", "application/json")
			this.Ctx.Output.Body(obj)
			panic(err)
		}
		responsejson.Servdel = "/successful_update"
		obj, _ := json.Marshal(responsejson)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)
	} else {

		if exist == false {
			responsejson.ServCode = "invalid"
		}
		if ServiceValid(code) == false {
			responsejson.ServCode = "incorrect"
		}

		if hoscode != hosCode {
			responsejson.ServCode = "notcode"
		}

		if code == "" {
			responsejson.ServCode = "empty"
		}

		obj, _ := json.Marshal(responsejson)
		this.Ctx.Output.SetStatus(300)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)

	}

}
