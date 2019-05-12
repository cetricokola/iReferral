package controllers

import (

	//"iReferral/models"
	//"github.com/astaxie/beego"
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Hospitals struct {
	Hos []string
}
type ReferController struct {
	MainController
}

var codes []string //slice of codes having the service name
var names []string
var hospitals Hospitals
var mycode string
var mydistrict string
var myregion string
var count = 1
var myArrayy []string

func (this *ReferController) ListHospitals() {
	this.searchFacility("searchFacility")
	if this.Ctx.Input.Method() == "POST" {
		//Extract the form data
		service := this.GetString("service")
		//return a array of codes of the hospitals offering the service input by the doctor
		o := orm.NewOrm()
		o.Using("default")
		_, err := o.Raw("select code from services where name=?", service).QueryRows(&codes)
		if err != nil {
			fmt.Println("Internal server error")
			flash.Error("The service you are requesting does not exist.")
			flash.Store(&this.Controller)
		}
		
		//querying for array hospital names within the same district offering the service
		if count < 1 {
			for i, _ := range codes {
				o.Raw("select name from hospital_account where code=? and district=?", codes[i], mydistrict).QueryRows(&names)
			}
		} else if count >= 1 && count < 2 {
			//querying for array hospital names within the same region offering the service
			for i, _ := range codes {
				o.Raw("select name from hospital_account where code=? and region=?", codes[i], myregion).QueryRows(&names)
			}
		} else {
			//querying for array hospital names from the rest of the country offering the service
			for i, _ := range codes {
				o.Raw("select name from hospital_account where code=?", codes[i]).QueryRows(&names)
			}
		}
		hospitals.Hos = names
		this.Redirect("/patientreferralform", 302)
	}
}
