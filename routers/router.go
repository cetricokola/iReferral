package routers

import (
	"iReferral/controllers"
	"github.com/astaxie/beego"
)
func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("healthIssue", &controllers.MainController{}, "get,post:Referral")
	beego.Router("searchfacility", &controllers.ReferController{}, "get,post:ListHospitals")
	beego.Router("patientreferralform", &controllers.ReferralsController{}, "get,post:PatReferral")
	beego.Router("auth/p_login", &controllers.AccountController{}, "get,post:Patients_login") //patient log in url
	beego.Router("auth/s_login", &controllers.AccountController{}, "get,post:Staff_login")
	beego.Router("auth/a-login", &controllers.AccountController{}, "get,post:Admin_login")
	beego.Router("registration/p_signup", &controllers.AccountController{}, "get,post:Patient_reg")
	beego.Router("info/admin_regSuccess", &controllers.MainController{}, "get,post:Admin_Reg_Success")
	beego.Router("info/emp_regSuccess", &controllers.MainController{}, "get,post:Emp_Reg_Success")
	beego.Router("info/patient_regSuccess", &controllers.MainController{}, "get,post:Patient_Reg_Success")
	beego.Router("registration/s-signup", &controllers.AccountController{}, "get,post:Staff_reg")
	beego.Router("registration/a-signup", &controllers.AccountController{}, "get,post:Admin_reg")
	beego.Router("logout", &controllers.AccountController{}, "get,post:Logout")
	beego.Router("auth/myadmin", &controllers.MainController{}, "get,post:AdminHome")
	beego.Router("hosreg", &controllers.HosregController{}, "post:Create")
	beego.Router("hosreg", &controllers.HosregController{}, "get:Search")
	beego.Router("facility_mgn", &controllers.FacilityController{}, "post:Createemp")
	beego.Router("facility_mgn", &controllers.FacilityController{}, "get:RemoveEmp")
	beego.Router("hosearch", &controllers.FaController{}, "post:Createservice")
	beego.Router("hosearch", &controllers.FaController{}, "get:Removeservice")
	beego.Router("doctor", &controllers.DocController{}, "get:FindPatient")
	beego.Router("/report", &controllers.DiagnosisController{}, "get,post:UpdateReport")
	beego.Router("/submitpatient", &controllers.SubmitPatientController{}, "get,post:SubmitPatient")
}
