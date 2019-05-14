package controllers

import (
	"iReferral/models"
	//"github.com/astaxie/beego"
	"fmt"

	"github.com/astaxie/beego/orm"
)

type ReferController struct {
	MainController
}

//variable declaration
var Names []string
func (this *ReferController) ListHospitals() {
	Names = nil
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID == nil {
		this.Redirect("/auth/s_login", 302)
		return
	}

	myId := userID.(string)
	this.searchFacility("searchFacility")

	if this.Ctx.Input.Method() == "POST" {

		//Extract the form data
		service := this.GetString("service")
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

		//searches if the services are available in the database
		serv := models.Services{Name: service}
		err := o.Read(&serv, "Name")
		if err == orm.ErrNoRows {
			fmt.Println("Service not found")
			flash.Error("Invalid service name or service not found.Try again!!")
			flash.Store(&this.Controller)
			return

		} else if err != nil {
			fmt.Println("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
			flash.Error("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
			flash.Store(&this.Controller)
			return
		}

		//Returns the number of hospital codes offerring the service
		var count int
		o.Raw("select COUNT(*) from services where name = ?", service).QueryRow(&count)
		fmt.Println("The number of codes offerring the service", count)

		//returns the list of hospital codes offering the service
		var lists []orm.ParamsList
		o.Raw("select code from services where name=?", service).ValuesList(&lists)

		//return hospital names offering the service requested in the same district
		fmt.Println(len(Names))
		control := false //control variable
		if control == false {
			///var Names []string
			for i := 0; i < len(lists); i++ {
				var name string
				fmt.Println(lists[i])
				o.Raw("select name from hospital_account where code=? AND district=?", lists[i], mydistrict).QueryRow(&name)
				fmt.Println(name)
				Names = append(Names, name)
			}
			if len(Names) == 0 {
				control = false
			} else {
				control = true
			}
		}
		fmt.Println("After a check fro mthe same district:=", Names)
		fmt.Println("the contro:", control)
		//return hospital names offering the service requested in the same region

		if control == false {
			for i := 0; i < len(lists); i++ {
				var name string
				fmt.Println(lists[i])
				o.Raw("SELECT name FROM hospital_account WHERE code=? AND region=? EXCEPT SELECT name FROM hospital_account WHERE code=? AND district=?", lists[i], myregion, lists[i], mydistrict).QueryRow(&name)
				fmt.Println(name)
				Names = append(Names, name)
			}
			if len(Names) == 0 {
				control = false
			} else {
				control = true
			}
		}

		if control == false {
			//return hospital names offering the service requested in the rest of the country parts
			for i := 0; i < len(lists); i++ {
				var name string
				fmt.Println(lists[i])
				o.Raw("SELECT name FROM hospital_account WHERE code=? AND country=? EXCEPT SELECT name FROM hospital_account WHERE code=? AND region=?", lists[i], mycountry, lists[i], myregion).QueryRow(&name)
				fmt.Println(name)
				Names = append(Names, name)
			}
		}

		fmt.Println(Names)
		fmt.Println(len(Names))
		this.Redirect("/patientreferralform", 302)
	
	}
}
