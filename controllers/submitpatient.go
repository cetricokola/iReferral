package controllers
import(
	"fmt"
	"iReferral/models"
	"github.com/astaxie/beego/orm"
)

type SubmitPatientController struct {
	MainController
}

var HudumaNo string
func (this *SubmitPatientController) SubmitPatient(){
	this.submitPatient("submitpatient")
	if this.Ctx.Input.Method() == "POST" {
		huduma :=this.GetString("huduma")
		o := orm.NewOrm()
			o.Using("default")
			hos := models.Patient_account{HudumaNo: huduma}
			err := o.Read(&hos, "HudumaNo")
			if err == orm.ErrNoRows {
				fmt.Println(err)
				fmt.Println("incorrect huduma number")
				flash.Error("Incorrect Huduma No.Try again!")
				flash.Store(&this.Controller)
				return
			} else if err != nil {
				fmt.Println("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
				flash.Error("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
				flash.Store(&this.Controller)
				return
			} else if err == nil {
				fmt.Println("Successful searching patient")
				HudumaNo = huduma
				this.Redirect("/searchfacility", 302)

			}
	}
}