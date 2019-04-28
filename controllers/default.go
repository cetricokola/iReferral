package controllers

import (
	"github.com/astaxie/beego"
)

type MaController struct {
	beego.Controller
}

func (this *MaController) sessionOff(view string) {
	this.Data["Title"] = "cetride.hub for rides"
	this.Layout = "basic-layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Sidebar"] = "sidebar.tpl"
	this.LayoutSections["Footer"] = "footer.tpl"
	this.TplName = view + ".tpl"
}
