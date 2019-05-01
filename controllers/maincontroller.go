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
	}
	this.Data["HeadScripts"] = []string{
		"/static/js/mdb-js/jquery-3.3.1.min.js",
		"/static/js/mdb-js/mdb.min.js",
		"/static/js/mdb-js/bootstrap.min.js",
		"/static/js/mdb-js/popper.min.js",
		"/static/js/custom/jquery.plugin.min.js",
		"/static/js/custom/jquery.datepick.js",
		"/static/js/custom/date.js",
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

// func (this *BaseController) reset_password(view string) {
// 	this.Data["Title"] = "Reset your cetride. password"
// 	this.Layout = "layout.tpl"
// 	this.LayoutSections = make(map[string]string)
// 	this.LayoutSections["Footer"] = "footer.html"
// 	this.TplName = "profile/"+ view + ".html"
// }

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
func (this *MainController) Hospital_reg(view string) {
	this.Data["Title"] = "Create facility"
	this.Layout = "layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Footer"] = "footer.html"
	this.TplName ="hosreg.html"
}



