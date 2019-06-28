package routers

import (
	"iReferral/controllers"
	"github.com/astaxie/beego"
)
func init() {
	beego.Router("/create-admin", &controllers.AdminSUPController{},"get,post:CreateAdmin")
	beego.Router("/login-admin", &controllers.AdLoginController{})
	beego.Router("/admin-dash", &controllers.MainController{},"get,post:AdDash")
}