package controllers

import (
	"encoding/json"
	"fmt"
	"iReferral/models"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type AdLoginController struct {
	MainController
}

type ResJson struct {
	Ress string
	AdUsername string
	AdPassword string
}


var res ResJson

func (this *AdLoginController) Get() {
	this.AdAuth()
}

func (this *AdLoginController) Post() {
	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	username := dataform["user"]
	password := dataform["pass"]

	o := orm.NewOrm()
	o.Using("default")
	h := username.(string)
	p := password.(string)

	//check if theusername exists
	exist := o.QueryTable("admin").Filter("Username", username).Exist()
	admin := models.Admin{Username: h}
	err := o.Read(&admin, "Username")
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(p))
	if exist == true && err == nil {
		session := this.StartSession()
		session.Set("UserID", h)
		res.Ress = "/admin-dash"
		obj, _ := json.Marshal(res)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)

	} else {
			//send ajax response that the password entered is invalid
			if exist == false {
				res.AdUsername = "incorrect"
			}
			if err != nil {
				res.AdPassword = "incorrect"
			}
			obj, _ := json.Marshal(res)
			this.Ctx.Output.SetStatus(300)
			this.Ctx.Output.Header("Content-Type", "application/json")
			this.Ctx.Output.Body(obj)
		}

	}

