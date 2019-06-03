package controllers

import (
	"fmt"
	"iReferral/models"

	"github.com/astaxie/beego/orm"
)

type ServiceController struct {
	MainController
}

//variable declaration
var hos models.Hospital_account
var Service string
var Names []orm.ParamsList
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
	fmt.Println("The service is ****", service)
	o := orm.NewOrm()
	o.Using("default")

	//find the hospital code of the employee
	var hosCode string
	o.Raw("select code from employee where emp_id=?", myId).QueryRow(&hosCode)
	fmt.Println(hosCode)

	//find the district of the hospital
	var mydistrict string
	o.Raw("select district from hospital_account where code=?", hosCode).QueryRow(&mydistrict)
	fmt.Println(mydistrict)

	//find the region of the hospital
	var myregion string
	o.Raw("select region from hospital_account where code=?", hosCode).QueryRow(&myregion)
	fmt.Println(myregion)

	//find the country of the hospital
	var mycountry string
	o.Raw("select country from hospital_account where code=?", hosCode).QueryRow(&mycountry)
	fmt.Println(mycountry)

	// //returns the list of hospital codes offering the service
	var lists []orm.ParamsList
	o.Raw("select code from services where name=?", service).ValuesList(&lists)
	fmt.Println("The length before the list having the codes@@@", lists)
	fmt.Println("The length before the list having the codes", len(lists))

	var mylists []orm.ParamsList
	_, err := o.QueryTable("hospital_account").Filter("code__in", lists).Filter("district", mydistrict).ValuesList(&mylists, "name")
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

	fmt.Println("the length of the list is", len(Names))
	fmt.Println("The second query is now executing")

	//return hospital names offering the service requested in the same region
	_, err = o.QueryTable("hospital_account").Exclude("district", mydistrict).Filter("code__in", lists).Filter("region", myregion).ValuesList(&mylists, "name")
	if err != nil {
		// No result
		fmt.Printf("Not row found")
		return
	}
	if err == nil && len(mylists) > 0 {
		Names = mylists
		fmt.Println("Second query  %% my list are^^^^^^^^^^", mylists)
		this.Redirect("/patientreferralform", 302)
		return
	}

	_, err = o.QueryTable("hospital_account").Exclude("region", myregion).Filter("code__in", lists).ValuesList(&mylists, "name")
	if err != nil {
		// No result
		fmt.Printf("Not row found")
		return
	}
	if err == nil && len(mylists) > 0 {
		Names = mylists
		fmt.Println("Third query  %% my list are^^^^^^^^^^", mylists)
		this.Redirect("/patientreferralform", 302)
		Service = service
	}
}
