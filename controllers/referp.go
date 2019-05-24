package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/orm"
)

type ReferpController struct {
	MainController
}

type ErrorJson struct {
	Message string
	Huduma string
	Service string
}

var MyHuduma string
var Myservice string
func (this *ReferpController) Get() {
	this.doctor_portal("doctor")
}

func (this *ReferpController) Post() {
	
	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	huduma :=  dataform["huduma"]
	service := dataform["service"]
	fmt.Println(huduma)
	fmt.Println(service)
	o := orm.NewOrm()
	o.Using("default")

	exist := o.QueryTable("patient_account").Filter("HudumaNo", huduma).Exist()
	exist1 := o.QueryTable("services").Filter("Name", service).Exist()
	h := huduma.(string)
	s := service.(string)
	MyHuduma = h
	Myservice = s
	fmt.Println(exist)
	fmt.Println(exist1)
	if exist == true && exist1 == true {
		var responsejson ErrorJson
		//responsejson.Message = "/submitpatient?huduma="+h+"&service="+s
		responsejson.Message = "/services"
		obj, _ := json.Marshal(responsejson)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)
	} else {
		var responsejson ErrorJson
		responsejson.Message = "Invalid details"
		responsejson.Huduma = "valid"
		responsejson.Service = "valid"
		if exist == false {
			responsejson.Huduma = "invalid"
		}
		if exist1 == false {
			responsejson.Service = "invalid"
		}
		obj, _ := json.Marshal(responsejson)
		this.Ctx.Output.SetStatus(300)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)

	}

}
