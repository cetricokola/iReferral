package controllers

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"iReferral/models"
)

type ServiceController struct {
	MainController
}

var Service string
var Names orm.ParamsList
var service string

func (this *ServiceController) Listservices() {
	Names = nil
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID == nil {
		this.Redirect("/", 302)
		return
	}
	myId := userID.(string)
	if len(Myservice) == 0 {
		service = Myservice2
	} else {
		service = Myservice
	}
	
	// get the hospital code of the doctor
	o := orm.NewOrm()
	o.Using("default")
	emp := models.Employee{EmpId: myId}
	err := o.Read(&emp)
	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
	}
	code := emp.Code
	
	//get the district and region of the doctors hospital
	hos := models.Hospital_account{Code: code}
	err = o.Read(&hos)
	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
	}
	mydistrict := hos.District
	myregion := hos.Region
	
	// //returns the list of hospital codes offering the service requested
	var lists []orm.ParamsList
	o.Raw("select code from services where name=?", service).ValuesList(&lists)
	
	// get the hospital names offering the service in the same district
	var mylists orm.ParamsList
	_, err = o.QueryTable("hospital_account").Filter("code__in", lists).Filter("district", mydistrict).ValuesFlat(&mylists, "name")
	if err != nil {
		// No result
		fmt.Printf("Not row found")
		return
	}
	if err == nil && len(mylists) > 0 {

		Names = mylists
		fmt.Println("First query  %% my list are^^^^^^^^^^", Names)
		this.Redirect("/patientreferralform", 302)
		return
	}

	//return hospital names offering the service requested in the same region
	_, err = o.QueryTable("hospital_account").Exclude("district", mydistrict).Filter("code__in", lists).Filter("region", myregion).ValuesFlat(&mylists, "name")
	if err != nil {
		// No result
		fmt.Printf("Not row found")
		return
	}
	if err == nil && len(mylists) > 0 {
		Names = mylists
		this.Redirect("/patientreferralform", 302)
		return
	}

		//return hospital names offering the service requested in the diferrent region
	_, err = o.QueryTable("hospital_account").Exclude("region", myregion).Filter("code__in", lists).ValuesFlat(&mylists, "name")
	if err != nil {
		// No result
		fmt.Printf("Not row found")
		return
	}
	if err == nil && len(mylists) > 0 {
		Names = mylists
		this.Redirect("/patientreferralform", 302)
		Service = service
	}
}
