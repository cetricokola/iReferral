package controllers

import (
	"encoding/json"
	"fmt"
	"iReferral/models"

	"github.com/astaxie/beego/orm"
)

type CreateEmpController struct {
	MainController
}

type CreateEmpJson struct {
	EmpCreate string
	EmpFirst  string
	EmpLast   string
	EmpNid    string
}

func (this *CreateEmpController) Get() {
	this.Facility_management()
}

func (this *CreateEmpController) Post() {

	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	first := dataform["first"]
	last := dataform["last"]
	nId := dataform["nId"]

	o := orm.NewOrm()
	o.Using("default")

	exist := o.QueryTable("employee").Filter("EmpId", nId).Exist()
	var responsejson CreateEmpJson
	if exist == false && EmpIdValid(nId.(string)) == true {
		session := this.StartSession()
		userID := session.Get("UserID")
		var hosCode string
		o.Raw("select code from hospital_account where mgn_id=?", userID.(string)).QueryRow(&hosCode)
		emp := models.Employee{FirstName: first.(string), LastName: last.(string), EmpId: nId.(string), Code: hosCode}
		_, err := o.Insert(&emp)
		if err != nil {
			responsejson.EmpCreate = "/successful_update"
			obj, _ := json.Marshal(responsejson)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)
			panic(err)
		}
		responsejson.EmpCreate = "/successful_update"
		obj, _ := json.Marshal(responsejson)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)
	} else {

		if exist == true {
			responsejson.EmpNid = "invalid"
		}
		if EmpIdValid(nId.(string)) == false {
			responsejson.EmpNid = "incorrect"
		}

		if first == "" {
			responsejson.EmpFirst = "invalid"
		}

		if last == "" {
			responsejson.EmpLast = "invalid"
		}

		if nId == "" {
			responsejson.EmpNid = "empty"
		}
		obj, _ := json.Marshal(responsejson)
		this.Ctx.Output.SetStatus(300)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)

	}

}
