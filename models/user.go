package models

import (
	"github.com/astaxie/beego/orm"
)

type Patient_account struct {
	FirstName string
	LastName  string
	HudumaNo  string `orm:"pk"`
	DoB       string
	PhoneNo   string `orm:"unique"`
	Password  string
	Sex       string
}

type Employee_account struct {
	EmpId    string `orm:"pk"`
	Email    string `orm:"unique"`
	PhoneNo  string
	Password string
}

type Admin_account struct {
	Id       string `orm:"pk"`
	Email    string `orm: "unique"`
	Password string
}

type Hospital_account struct {
	Name     string
	SerialNo string `orm:"pk"`
	PhoneNo  string
	Email    string `orm:"unique"`
	Country  string
	Region   string
	District string
}

type Employee struct {
	FirstName  string
	LastName   string
	Position   string
	NationalId string `orm:"pk"`
}

type Services struct{
	Code string `orm:"pk"`
	Name string
	Cost string
	Slot string
	Department string
}

type Patient_diagnosis struct {
	Id string `orm:"pk"`
	HudumaNo string
	Weight string
	Temperature string
	BloodPressure string
	Diagnosis string 
	Response string
	Prescription string
	Reg_date string
}
/********/
func init() {
	orm.RegisterModel(new(Patient_account), new(Employee_account),
	 new(Admin_account), new(Hospital_account), new(Services), new(Employee), new(Patient_diagnosis))
}
