package models

import (
	"github.com/astaxie/beego/orm"
)

type Patient_account struct {
	FirstName string
	LastName  string
	HudumaNo  string `orm:"pk"`
	DoB       string `valid:"Required"`
	PhoneNo   string `orm:"unique"`
	Password  string `valid:"Required"`
	Sex       string `valid:"Required"`
}

type Employee_account struct {
	EmpId    string `orm:"pk"`
	Email    string `orm:"unique"`
	PhoneNo  string `valid:"Required"`
	Password string `valid:"Required"`
}

type Admin_account struct {
	Id       string `orm:"pk"`
	Email    string `orm: "unique"`
	Password string `valid:"Required"`
}

type Hospital_account struct {
	Code     string `orm:"unique"`
	Name     string
	SerialNo string `orm:"pk"`
	PhoneNo  string `orm:"unique"`
	Email    string `orm:"unique"`
	Country  string
	Region   string
	District string
	MgnId    string `orm:"unique"`
}

type Employee struct {
	FirstName string
	LastName  string
	Position  string
	EmpId     string `orm:"pk"`
	Code      string
}

type Services struct {
	ServiceCode string `orm:"pk"`
	Code        string
	Name        string
	Cost        string
	Slot        string
	Department  string
}

type Patient_diagnosis struct {
	Id            string `orm:"pk"`
	HudumaNo      string
	Weight        string
	Temperature   string
	BloodPressure string
	Diagnosis     string
	Response      string
	Prescription  string
	Reg_date      string
}

type Referrals struct {
	Id       string `orm:"pk"`
	HudumaNo string
	Service string
	HosName  string
	RDate    string
	RTime    string
}

/********/
func init() {
	orm.RegisterModel(new(Patient_account), new(Employee_account),
		new(Admin_account), new(Hospital_account), new(Services), new(Employee), new(Patient_diagnosis), new(Referrals))
}
