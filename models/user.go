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

/********/
func init() {
	orm.RegisterModel(new(Patient_account), new(Employee_account), new(Admin_account), new(Hospital_account))
}
