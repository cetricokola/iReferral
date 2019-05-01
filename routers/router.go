package routers

import (
	"iReferral/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("auth/p_login", &controllers.AccountController{}, "get,post:Patients_login")
	beego.Router("auth/s_login", &controllers.AccountController{}, "get,post:Staff_login")
	beego.Router("auth/a-login", &controllers.AccountController{}, "get,post:Admin_login")
	beego.Router("registration/p_signup", &controllers.AccountController{}, "get,post:Patient_reg")
	beego.Router("registration/s-signup", &controllers.AccountController{}, "get,post:Staff_reg")
	beego.Router("registration/a-signup", &controllers.AccountController{}, "get,post:Admin_reg")
	beego.Router("auth/myadmin", &controllers.MainController{}, "get,post:AdminHome")
	beego.Router("hosreg", &controllers.HosregController{})
	

}
