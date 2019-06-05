package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/orm"
)

type DelEmpController struct {
	MainController
}

type DelEmpJson struct {
	Empdel string
	Empid  string
}

func (this *DelEmpController) Get() {
	this.Facility_management()
}

func (this *DelEmpController) Post() {

	session := this.StartSession()
	userID := session.Get("UserID").(string)
	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	empid := dataform["empid"].(string)

	o := orm.NewOrm()
	o.Using("default")
	var hosCode, code string
	exist := o.QueryTable("employee").Filter("EmpId", empid).Exist()
	o.Raw("select code from employee where emp_id=?", empid).QueryRow(&code)
	o.Raw("select code from hospital_account where mgn_id=?", userID).QueryRow(&hosCode)
	var responsejson DelEmpJson
	if exist == true && EmpIdValid(empid) == true && code == hosCode {
		_, err := o.QueryTable("employee").Filter("emp_id", empid).Delete()
		if err != nil {
			responsejson.Empdel = "/successful_update"
			obj, _ := json.Marshal(responsejson)
			this.Ctx.Output.Header("Content-Type", "application/json")
			this.Ctx.Output.Body(obj)
			panic(err)
		}
		responsejson.Empdel = "/successful_update"
		obj, _ := json.Marshal(responsejson)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)
	} else {

		if exist == false {
			responsejson.Empid = "invalid"
		}
		if EmpIdValid(empid) == false {
			responsejson.Empid = "incorrect"
		}

		if code != hosCode {
			responsejson.Empid = "notstaff"
		}

		if empid == "" {
			responsejson.Empid = "empty"
		}

		obj, _ := json.Marshal(responsejson)
		this.Ctx.Output.SetStatus(300)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)

	}

}
