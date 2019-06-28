package controllers

import (
	"github.com/astaxie/beego"

)

type MainController struct {
	beego.Controller
}

var flash = beego.NewFlash()
func (this *MainController) Prepare() {

	this.Data["HeadStyles"] = []string{
		"/static/css/mdb/bootstrap.min.css",
		"/static/css/mdb/mdb.min.css",
		"/static/css/mdb/style.min.css",
		"/static/css/custom/main.css",
		"/static/css/custom/search.css",
		"/static/css/custom/login.css",
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
		"/static/js/mdb-js/axios.min.js",
		"/static/js/custom/login.js",
		
	}
}

func (this *MainController) Get() {
	this.Data["Title"] = "iReferral-Welcome"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "home.html"
}

func (this *MainController) AdminAuth() {
	this.Data["Title"] = "iReferral-Welcome"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "adminAuth.html"
}

func (this *MainController) home(view string) {
	this.Data["Title"] = "iReferral-Welcome"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = view + ".html"
}

func (this *MainController) Contactus() {
	this.Data["Title"] = "iReferral-Write to us"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "contactUs.html"
}

func (this *MainController) Patient_Reg_Success() {
	this.Data["Title"] = "Patient account successfully created"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "info/patient_regSuccess.html"
}
func (this *MainController) Admin_Reg_Success() {
	this.Data["Title"] = "Admin account successfully created"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "info/admin_regSuccess.html"
}
func (this *MainController) Emp_Reg_Success() {
	this.Data["Title"] = "Employee account successfully created"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "info/emp_regSuccess.html"
}


func (this *MainController) hospital_reg(view string) {
	//check if the user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")

	if userID == nil {
		this.Redirect("/management_authentication", 302)
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
		this.Redirect("/management_authentication", 302)
		return
	}
	this.Data["Title"] = "Manage facility"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = view + ".html"
}

func (this *MainController) Facility_management() {
	//check if the user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")

	if userID == nil {
		this.Redirect("/management_authentication", 302)
		return
	}
	this.Data["Title"] = "Manage the facility"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "facility_mgn.html"
}

func (this *MainController) EmpUpdate() {
	//check if the user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")

	if userID == nil {
		this.Redirect("/management_authentication", 302)
		return
	}
	this.Data["Title"] = "employee successfully managed"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "info/empreg.html"
}
func (this *MainController) diagnosis(view string) {
	//check if the user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")

	if userID == nil {
		this.Redirect("/", 302)
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
		this.Redirect("/", 302)
		return
	}
	this.Data["Title"] = "Doctors dashboard"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = view + ".html"
}

func (this *MainController) Referral() {
	session := this.StartSession()
	userID := session.Get("UserID")

	if userID == nil {
		this.Redirect("/", 302)
		return
	}
	this.Data["Title"] = "iReferral-updated report"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "healthIssue.html"
}

func (this *MainController) patReferral(view string) {
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID == nil {
		this.Redirect("/", 302)
		return
	}
	this.Data["Title"] = "iReferral-select the facility"
	this.Data["Names"] = Names
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = view + ".html"
}
func (this *MainController) PostDetails() {
	session := this.StartSession()
	userID := session.Get("UserID")

	if userID == nil {
		this.Redirect("/", 302)
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
	this.Data["Name"] = Uname
	this.Data["Len"] = Len
	this.Data["Referral"] = Referral
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = view + ".html"
}

func (this *MainController) Phome() {
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID == nil {
		this.Redirect("/", 302)
		return
	}
	this.Data["Title"] = "iReferral-Welcome to patient portal"
	this.Data["Pfname"] = Pfname
	this.Data["Plname"] = Plname
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "phome.html"
}

func (this *MainController) preferral(view string) {
	this.Data["Title"] = "iReferral-View your referrals"
	this.Data["Plname"] = Plname
	this.Data["Pfname"] = Pfname
	this.Data["MyRef"] = MyRef
	this.Data["Leng"] = Leng
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = view + ".html"
}

func (this *MainController) preport(view string) {
	this.Data["Title"] = "iReferral-View your medical history"
	this.Data["Plname"] = Plname
	this.Data["Pfname"] = Pfname
	this.Data["Repo"] = Repo
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = view + ".html"
}
func (this *MainController) AdminDash() {
	//check if the user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")

	if userID == nil {
		this.Redirect("/management_authentication", 302)
		return
	}
	this.Data["Title"] = "Admin dashbord"
	this.Data["Name"] = Uname
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "admindash.html"
}
func (this *MainController) Unsuccessful() {
	session := this.StartSession()
	userID := session.Get("UserID")

	if userID == nil {
		this.Redirect("/management_authentication", 302)
		return
	}
	this.Data["Title"] = "iReferral-unable to create a facility"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "unsuccessful.html"
}
func (this *MainController) UnsuccessfulHos() {
	session := this.StartSession()
	userID := session.Get("UserID")

	if userID == nil {
		this.Redirect("/management_authentication", 302)
		return
	}
	this.Data["Title"] = "iReferral-unable to manage facility"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "info/reghos.html"
}

func (this *MainController) Conta(){
	session := this.StartSession()
	userID := session.Get("UserID")
	SessionId := userID.(string)
	this.Data["Session"] = SessionId
	this.Data["Pat"] = PiD
	this.Data["Admin"] = ID
	this.Data["Doc"] = Sid
	this.Data["Title"] = "iReferral-contacts"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName = "info/contacts.html"
}
func (this *MainController) AdAuth(){
	
	this.Data["Title"] = "iReferral-Admin auth"
	this.Layout = "layout.tpl"
	
	this.TplName = "admin/login.html"
}

func (this *MainController) AdDash(){
	session := this.StartSession()
	userID := session.Get("UserID")

	if userID == nil {
		this.Redirect("/login-admin", 302)
		return
	}
	this.Data["Title"] = "iReferral-Admin Dash"
	this.Layout = "layout.tpl"
	this.TplName = "admin/adminDash.html"
}

func (this *MainController) Help(){
	this.Data["Title"] = "iReferral-get help"
	this.Layout = "layout.tpl"
	this.TplName = "help.html"
}
func (this *MainController) ViewDoctors(){
	this.Data["Title"] = "iReferral-doctors at your facility"
	this.Data["Emp"] = Emps
	this.Data["Len"] = Len0
	this.Layout = "layout.tpl"
	this.TplName = "viewdoctors.html"
}
func (this *MainController) ViewServices(){
	this.Data["Title"] = "iReferral-services at your facility"
	this.Data["Services"] = Serv
	this.Data["Len"] = Len2
	this.Layout = "layout.tpl"
	this.TplName = "viewservices.html"
}