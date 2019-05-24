package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type ServicesController struct {
	MainController
}

type ResponseJson struct {
	Mymessage string
	Services  string
}

var Myservice2 string
func (this *ServicesController) Get() {
	this.preport("report")
}

func (this *ServicesController) Post() {

	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	service := dataform["service"]
	fmt.Println(service)
	o := orm.NewOrm()
	o.Using("default")

	exist := o.QueryTable("services").Filter("Name", service).Exist()
	s := service.(string)
	Myservice2 = s
	fmt.Println(exist)
	if exist == true {
		var responsejson ResponseJson
		//responsejson.Message = "/submitpatient?huduma="+h+"&service="+s
		responsejson.Mymessage = "/services"
		obj, _ := json.Marshal(responsejson)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)
	} else {
		var responsejson ResponseJson
		responsejson.Mymessage = "Invalid details"
		responsejson.Services = "valid"
		if exist == false {
			responsejson.Services = "invalid"
		}
		obj, _ := json.Marshal(responsejson)
		this.Ctx.Output.SetStatus(300)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)

	}

}
