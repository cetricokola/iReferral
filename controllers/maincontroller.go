package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Prepare() {

	this.Data["HeadStyles"] = []string{
		"/static/css/mdb/bootstrap.min.css",
		"/static/css/mdb/mdb.min.css",
		"/static/css/mdb/style.min.css",
		"/static/css/custom/main.css",
		"/static/css/custom/search.css",
		"/static/css/custom/jquery.datepick.css",
		"/static/css/custom/chung-timepicker.css",
			}

	this.Data["HeadScripts"] = []string{
		"/static/js/mdb-js/jquery-3.3.1.min.js",
		"/static/js/mdb-js/mdb.min.js",
		"/static/js/mdb-js/bootstrap.min.js",
		"/static/js/mdb-js/popper.min.js",
		"/static/js/custom/jquery.plugin.min.js",
		"/static/js/custom/jquery.datepick.js",
		"/static/js/custom/date.js",
		"/static/js/custom/date2.js",
		"/static/js/custom/date3.js",
		"/static/js/custom/date4.js",
		"/static/js/custom/time.js",
		"/static/js/custom/chung-timepicker.js",
		
	}
}

func (this *MainController) Get(){
	this.Data["Title"] = "iReferral-Welcome"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "home.html"
}
func (this *MainController) Patient_Reg_Success(){
	this.Data["Title"] = "Patient account successfully created"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "info/patient_regSuccess.html"
}
func (this *MainController) Admin_Reg_Success(){
	this.Data["Title"] = "Admin account successfully created"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "info/admin_regSuccess.html"
}
func (this *MainController) Emp_Reg_Success(){
	this.Data["Title"] = "Employee account successfully created"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "info/emp_regSuccess.html"
}

func (this *MainController) patient_logIn(view string) {
	this.Data["Title"] = "Log in to iReferral"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "auth/"+ view + ".html"
}

func (this *MainController) staff_logIn(view string) {
	this.Data["Title"] = "Log in to iReferral"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "auth/"+ view + ".html"
}

func (this *MainController) admin_logIn(view string) {
	this.Data["Title"] = "Log in to iReferral"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "auth/"+ view + ".html"
}

func (this *MainController) patient_signUp(view string) {
	this.Data["Title"] = "Sign up to iReferral"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName ="registration/"+ view + ".html"
}

func (this *MainController) staff_signUp(view string) {
	this.Data["Title"] = "Sign up to iReferral"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName ="registration/"+ view + ".html"
}

func (this *MainController) admin_signUp(view string) {
	this.Data["Title"] = "Sign up to iReferral"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName ="registration/"+ view + ".html"
}

func (this *MainController) AdminHome(){
	this.Data["Title"] = "iReferral-myAdmin portal"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName ="auth/myadmin.html"

}
func (this *MainController) hospital_reg(view string) {
	//check if the user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")
	
	if userID == nil {
		this.Redirect("/auth/a-login", 302)
		return
	}
		this.Data["Title"] = "Create facility"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = view + ".html"
}
func (this *MainController) facility_mgn(view string) {
	//check if the user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")
	
	if userID == nil {
		this.Redirect("/auth/a-login", 302)
		return
	}
	this.Data["Title"] = "Manage facility"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = view + ".html"
}
func (this *MainController) diagnosis(view string) {
	//check if the user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")
	
	if userID == nil {
		this.Redirect("/auth/s_login", 302)
		return
	}
	this.Data["Title"] = "Patient medical reports update"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = view + ".html"
}
func (this *MainController) doctor_portal(view string) {
	session := this.StartSession()
	userID := session.Get("UserID")
	
	if userID == nil {
		this.Redirect("/auth/s_login", 302)
		return
	}
	this.Data["Title"] = "Doctors dashboard"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = view + ".html"
}

func (this *MainController) Referral(){
	session := this.StartSession()
	userID := session.Get("UserID")
	
	if userID == nil {
		this.Redirect("/auth/s_login", 302)
		return
	}
	this.Data["Title"] = "iReferral-updated report"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "healthIssue.html"
}

func (this *MainController) searchFacility( view string){
	session := this.StartSession()
	userID := session.Get("UserID")
		if userID == nil {
		this.Redirect("/auth/s_login", 302)
		return
	}
	this.Data["Title"] = "iReferral-search for facilities"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = view + ".html"
}

func (this *MainController) submitPatient( view string){
	session := this.StartSession()
	userID := session.Get("UserID")
		if userID == nil {
		this.Redirect("/auth/s_login", 302)
		return
	}
	this.Data["Title"] = "iReferral-input your patient"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = view + ".html"
}
func (this *MainController) patReferral(view string){
	session := this.StartSession()
	userID := session.Get("UserID")
		if userID == nil {
		this.Redirect("/auth/s_login", 302)
		return
	}
	this.Data["Title"] = "iReferral-select the facility"
	this.Data["Names"] = Names
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = view+".html"
}
func (this *MainController) PostDetails(){
	session := this.StartSession()
	userID := session.Get("UserID")
	
	if userID == nil {
		this.Redirect("/auth/s_login", 302)
		return
	}
	this.Data["Title"] = "iReferral-Confirm that you want to refer the patient"
	this.Data["Fname"] = Fname
	this.Data["Lname"] = Lname
	this.Data["HudumaNumber"] = Myhuduma
	this.Data["Servicename"] = Service
	this.Data["Thisdate"] = Date
	this.Data["Thistime"] = Time
	this.Data["Hosname"] = Hosname
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "success.html"
}
func (this *MainController) viewReferrals(view string) {
	//check if the user is logged in
	
		this.Data["Title"] = "View referrals at your facility"
		this.Data["Name"] = Name
		this.Data["Hdm"] = Hdm
		this.Data["Serv"] = Serv
		this.Data["Dat"] = Dat
		this.Data["Tim"] = Tim
		this.Data["Len"] = Len
		this.Data["Referral"] = Referral
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = view+".html"
}
