package controllers

import (
	// "github.com/astaxie/beego/orm"
	"github.com/jung-kurt/gofpdf"
	"fmt"
	// "iReferral/models"
)

type ReportController struct {
	MainController
}
// type Services struct{
// 	ServiceCode string
// 	Name string
// 	Department string
// }
// type Employee struct {
// 	FirstName string
// 	LastName  string
// 	EmpId     string 
// 	// Code      string
// }
// var services []Services
// var emp []Employee

func (this *ReportController) ViewReferrals() {

	// o := orm.NewOrm()
	// o.Using("default")
		pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetHeaderFuncMode(func() {
		// pdf.Image(example.ImageFile("logo.png"), 10, 6, 30, 0, false, "", 0, "")
		pdf.SetY(5)
		pdf.SetFont("Arial", "B", 20)
		pdf.Cell(80, 0, "")
		pdf.SetTextColor(66, 133, 244)
		pdf.CellFormat(30, 10, "iReferral.", "", 0, "C", false, 0, "")
		pdf.Ln(20)
	}, true)
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 8)
		pdf.CellFormat(0, 10, fmt.Sprintf("iReferral-Your full circle of care."),
			"", 0, "C", false, 0, "")
			pdf.Ln(5)
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d/{nb}", pdf.PageNo()),
			"", 0, "C", false, 0, "")
	})
	pdf.AliasNbPages("")
	pdf.AddPage()
	// pdf.PageNo()
	pdf.SetFont("Arial", "B", 16)
	// pdf.SetTextColor(66, 133, 244)
	// pdf.CellFormat(40, 10, "iReferral.", "", 0, "", false, 0, "")
	// pdf.Ln(10)
	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "B", 10)
	pdf.Ln(10)
	pdf.Cell(40, 10, "Dates:")
	pdf.Cell(20, 10, Start)
	pdf.Cell(3, 10, "--")
	pdf.Cell(20, 10, End)
	pdf.Ln(10)
	pdf.Line(10, 30, 200, 30)
	pdf.Cell(40, 10, "Patient Huduma No")
	pdf.Cell(30, 10, "Service")
	pdf.Cell(30, 10, "PhoneNumber")
	pdf.Cell(30, 10, "Referred from")
	pdf.Cell(30, 10, "Referral Date")
	pdf.Cell(30, 10, "Referral Time")
	pdf.Ln(10)
	pdf.Line(10, 40, 200, 40)
	pdf.SetFont("Arial", "", 10)
	for i, _ := range Referral {
		pdf.Cell(40, 10, Referral[i].HudumaNo)
		pdf.Cell(30, 10, Referral[i].Service)
		pdf.Cell(30, 10, Referral[i].Phone)
		pdf.Cell(30, 10, Referral[i].Refer_hos)
		pdf.Cell(30, 10, Referral[i].RDate)
		pdf.Cell(30, 10, Referral[i].RTime)
		pdf.Ln(10)
	}

	err := pdf.OutputFileAndClose("downloads/viewreferrals.pdf")
		if err != nil {
		panic(err)
	}

	this.Redirect("/viewreferrals", 302)
}

func (this *ReportController) PrintServices() {
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID == nil {
		this.Redirect("/management_authentication", 302)
		return
	}
	// myId := userID.(string)

	// o := orm.NewOrm()
	// o.Using("default")
	// mgn := models.Hospital_account{MgnId: myId}
	// err := o.Read(&mgn, "MgnId")
	// if err != nil {
	// 	panic(err)
	// }

	// code := mgn.Code
	// o.Raw("SELECT service_code,name,department FROM services WHERE code=?", code).QueryRows(&services)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetHeaderFuncMode(func() {
		// pdf.Image(example.ImageFile("logo.png"), 10, 6, 30, 0, false, "", 0, "")
		pdf.SetY(5)
		pdf.SetFont("Arial", "B", 20)
		pdf.Cell(80, 0, "")
		pdf.SetTextColor(66, 133, 244)
		pdf.CellFormat(30, 10, "iReferral.", "", 0, "C", false, 0, "")
		pdf.Ln(20)
	}, true)
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 8)
		pdf.CellFormat(0, 10, fmt.Sprintf("iReferral-Your full circle of care."),
			"", 0, "C", false, 0, "")
			pdf.Ln(5)
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d/{nb}", pdf.PageNo()),
			"", 0, "C", false, 0, "")
	})
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(40, 10, "Services")
	pdf.Ln(10)
	
	pdf.Line(10, 30, 200, 30)
	pdf.Cell(60, 10, "Service code")
	pdf.Cell(60, 10, "Service Name")
	pdf.Cell(60, 10, "Department")
	pdf.Ln(10)
	pdf.Line(10, 20, 200, 20)
	pdf.SetFont("Arial", "", 10)
	for i, _ := range Serv {
		fmt.Println("Services are as follows", Serv[i])
		pdf.Cell(60, 10, Serv[i].ServiceCode)
		pdf.Cell(60, 10, Serv[i].Name)
		pdf.Cell(60, 10, Serv[i].Department)
		pdf.Ln(10)
	}

	err := pdf.OutputFileAndClose("downloads/services.pdf")
		if err != nil {
		panic(err)
	}

	this.Redirect("/adminhome", 302)
}

func (this *ReportController) PrintDoctors() {
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID == nil {
		this.Redirect("/management_authentication", 302)
		return
	}
	// myId := userID.(string)

	// o := orm.NewOrm()
	// o.Using("default")
	// // mgn := models.Hospital_account{MgnId: myId}
	// err := o.Read(&mgn, "MgnId")
	// if err != nil {
	// 	panic(err)
	// }

	// code := mgn.Code
	// o.Raw("SELECT emp_id,first_name,last_name FROM employee WHERE code=?", code).QueryRows(&emp)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetHeaderFuncMode(func() {
		// pdf.Image(example.ImageFile("logo.png"), 10, 6, 30, 0, false, "", 0, "")
		pdf.SetY(5)
		pdf.SetFont("Arial", "B", 20)
		pdf.Cell(80, 0, "")
		pdf.SetTextColor(66, 133, 244)
		pdf.CellFormat(30, 10, "iReferral.", "", 0, "C", false, 0, "")
		pdf.Ln(20)
	}, true)
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 8)
		pdf.CellFormat(0, 10, fmt.Sprintf("iReferral-Your full circle of care."),
			"", 0, "C", false, 0, "")
			pdf.Ln(5)
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d/{nb}", pdf.PageNo()),
			"", 0, "C", false, 0, "")
	})
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(40, 10, "Employees")
	pdf.Ln(10)
	
	pdf.Line(10, 30, 200, 30)
	pdf.Cell(60, 10, "Employee Id")
	pdf.Cell(60, 10, "First Name")
	pdf.Cell(60, 10, "Last Name")
	pdf.Ln(10)
	pdf.Line(10, 20, 200, 20)
	pdf.SetFont("Arial", "", 10)
	for i, _ := range Emps {
		fmt.Println("Services are as follows", Emps[i])
			pdf.Cell(60, 10, Emps[i].EmpId)
		pdf.Cell(60, 10,Emps[i].FirstName)
		pdf.Cell(60, 10, Emps[i].LastName)
		pdf.Ln(10)
	}
	err := pdf.OutputFileAndClose("downloads/doctors.pdf")
		if err != nil {
		panic(err)
	}
	this.Redirect("/adminhome", 302)
}
