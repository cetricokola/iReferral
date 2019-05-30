package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type ServiceController struct {
	MainController
}

//variable declaration
var Service string
var Names []string
var service string
func (this *ServiceController) Listservices(){
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

		//Returns the number of hospital codes offerring the service
		var count int
		o.Raw("select COUNT(*) from services where name = ?", service).QueryRow(&count)
		fmt.Println("The number of codes offerring the service", count)

		//returns the list of hospital codes offering the service
		var lists []orm.ParamsList
		o.Raw("select code from services where name=?", service).ValuesList(&lists)

		//return hospital names offering the service requested in the same district
		fmt.Println(len(Names))

		var control bool
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

			//return hospital names offering the service requested in the same region

		} else if control == false {
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
		} else if control == false {
			//return hospital names offering the service requested in the rest of the country parts
			for i := 0; i < len(lists); i++ {
				var name string
				fmt.Println(lists[i])
				o.Raw("SELECT name FROM hospital_account WHERE code=? AND country=? EXCEPT SELECT name FROM hospital_account WHERE code=? AND region=?", lists[i], mycountry, lists[i], myregion).QueryRow(&name)
				fmt.Println(name)
				Names = append(Names, name)
			}
		}
		fmt.Println("After a check fro mthe same district:=", Names)
		fmt.Println("the contro:", control)
		fmt.Println(Names)
		fmt.Println(len(Names))
		Service = service
		this.Redirect("/patientreferralform", 302)

	}
