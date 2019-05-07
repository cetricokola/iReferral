package controllers

import (
	"fmt"
	"iReferral/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	//"github.com/astaxie/beego/validation"

	"golang.org/x/crypto/bcrypt"
)

type AccountController struct {
	MainController
}

var flash = beego.NewFlash()

//*PATIENT PORTAL LOG IN*//
func (this *AccountController) Patients_login() {
	this.patient_logIn("p_login")
	if this.Ctx.Input.Method() == "POST" {
		//Extract the form data
		hudumaNo := this.GetString("hudumaNo")
		submittedPassword := this.GetString("pss")

		//Read user account from database
		o := orm.NewOrm()
		o.Using("default")
		patient := models.Patient_account{HudumaNo: hudumaNo}
		err := o.Read(&patient, "hudumaNo")

		if err == orm.ErrNoRows || err == orm.ErrMissPK {
			fmt.Println("incorrect huduma number or password")
			flash.Error("You've entered incorrect huduma number or password")
			flash.Store(&this.Controller)
			return

		} else if err != nil {
			fmt.Println("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
			flash.Error("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
			flash.Store(&this.Controller)
			return
		}
		//******** Compare submitted password with the saved hash
		err = bcrypt.CompareHashAndPassword([]byte(patient.Password), []byte(submittedPassword))
		if err != nil {
			fmt.Println(err)
			fmt.Println("Incorrect password")
			flash.Error("You've entered incorrect password")
			flash.Store(&this.Controller)
			return
		}
		//set session for patient log in
		this.SetSession("patient.HudumaNo", 50)
		fmt.Println(patient.HudumaNo, ":successful log in ")
		this.Redirect("/", 302)
	}
}

//*STAFF PORTAL LOG IN*//
func (this *AccountController) Staff_login() {
	// Check if user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID != nil {
		// User is logged in already, display another page
		this.Redirect("/doctor", 302)
		return
	}
	this.staff_logIn("s_login")
	if this.Ctx.Input.Method() == "POST" {
		//Extract the form data
		empId := this.GetString("empId")
		submittedPassword := this.GetString("pass")
		fmt.Println(submittedPassword)
		//Read user account from database
		o := orm.NewOrm()
		o.Using("default")
		staff := models.Employee_account{EmpId: empId}
		err := o.Read(&staff, "empId")

		if err == orm.ErrNoRows || err == orm.ErrMissPK {
			fmt.Println("incorrect employee id or password")
			flash.Error("You've entered incorrect employmee id or password")
			flash.Store(&this.Controller)
			return

		} else if err != nil {
			fmt.Println("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
			flash.Error("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
			flash.Store(&this.Controller)
			return
		}
		//******** Compare submitted password with the saved hash
		err = bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(submittedPassword))
		if err != nil {
			fmt.Println(err)
			fmt.Println("Incorrect password")
			flash.Error("You've entered incorrect password")
			flash.Store(&this.Controller)
			return
		}
		// Set the UserID if everything is ok
		session.Set("UserID", empId)
		fmt.Println(staff.EmpId, ":successful log in ")
		this.Redirect("/doctor", 302)
	}

}

//*ADMIN PORTAL LOG IN*//
func (this *AccountController) Admin_login() {
	this.admin_logIn("a-login")
	if this.Ctx.Input.Method() == "POST" {
		//Extract the form data
		iD := this.GetString("nationalId")
		submittedPassword := this.GetString("password")

		//Read user account from database
		o := orm.NewOrm()
		o.Using("default")
		admin := models.Admin_account{Id: iD}
		err := o.Read(&admin, "Id")

		if err == orm.ErrNoRows || err == orm.ErrMissPK {
			fmt.Print(err)
			fmt.Println("incorrect national id or password")
			flash.Error("You've entered incorrect national id or password")
			flash.Store(&this.Controller)
			return

		} else if err != nil {
			fmt.Println("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
			flash.Error("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
			flash.Store(&this.Controller)
			return
		}
		//******** Compare submitted password with the saved hash
		err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(submittedPassword))
		if err != nil {
			fmt.Println("Incorrect password")
			flash.Error("You've entered incorrect password")
			flash.Store(&this.Controller)
			return
		}
		//set session
		this.SetSession("UserID", iD)
		fmt.Println(admin.Id, ":successful log in ")
		this.Redirect("../hosreg", 302)

	}

}

//*PATIENT SIGN UP*//
func (this *AccountController) Patient_reg() {
	this.patient_signUp("p_signup")
	if this.Ctx.Input.Method() == "POST" {

		//get the values from the form
		first := this.GetString("first")
		last := this.GetString("last")
		huduma := this.GetString("huduma")
		dob := this.GetString("dob")
		phone := this.GetString("phone")
		sex := this.GetString("sex")
		submittedpassword := this.GetString("pass")
		confirmPass := this.GetString("copass")

		fmt.Println(first)
		//validation of user input
		//match the submitted password with the confirm password
		if submittedpassword != confirmPass {
			fmt.Println("The password submitted does no match the confirm password")
			flash.Error("The password failed to match")
			return
		}
		//hash the password
		password, _ := HashPassword(submittedpassword)

		//******** Save user info to database
		o := orm.NewOrm()
		o.Using("default")

		patient := models.Patient_account{FirstName: first, LastName: last, HudumaNo: huduma, DoB: dob, PhoneNo: phone, Password: password, Sex: sex}

		_, err := o.Insert(&patient)
		if err != nil {
			fmt.Println(err)
			flash.Error(phone + " already registered")
			flash.Store(&this.Controller)
			//this.Redirect("/", 302)
			return
		}
		fmt.Print("successfull registration")
		this.Redirect("../info/patient_regSuccess", 302)
	}
}

//*STAFF SIGN UP*//
func (this *AccountController) Staff_reg() {
	this.staff_signUp("s-signup")
	if this.Ctx.Input.Method() == "POST" {

		//get the values from the form
		empId := this.GetString("empId")
		email := this.GetString("email")
		phone := this.GetString("phone")
		submittedpassword := this.GetString("pass")
		confirmPass := this.GetString("copass")

		//validation of user input
		//match the submitted password with the confirm password
		if submittedpassword != confirmPass {
			fmt.Println("The password submitted does no match the confirm password")
			flash.Error("The password failed to match")
			return
		}
		//hash the password
		password, _ := HashPassword(submittedpassword)
		//******** Save user info to database
		o := orm.NewOrm()
		o.Using("default")

		staff := models.Employee_account{EmpId: empId, Email: email, PhoneNo: phone, Password: password}

		_, err := o.Insert(&staff)
		if err != nil {
			fmt.Println(err)
			flash.Error(phone + " already registered")
			flash.Store(&this.Controller)
			return
		}

		this.Redirect("../info/emp_regSuccess", 302)
	}
}

//*Admin SIGN UP*//
func (this *AccountController) Admin_reg() {
	this.admin_signUp("a-signup")
	if this.Ctx.Input.Method() == "POST" {

		//get the values from the form
		id := this.GetString("id")
		email := this.GetString("email")
		submittedpassword := this.GetString("pass")
		confirmPass := this.GetString("copass")

		//validation of user input
		//match the submitted password with the confirm password
		if submittedpassword != confirmPass {
			fmt.Println("The password submitted does no match the confirm password")
			flash.Error("The password failed to match")
			return
		}
		//hash the password
		password, _ := HashPassword(submittedpassword)
		//******** Save user info to database
		o := orm.NewOrm()
		o.Using("default")

		admin := models.Admin_account{Id: id, Email: email, Password: password}

		_, err := o.Insert(&admin)
		if err != nil {
			flash.Error(email + " already registered")
			flash.Store(&this.Controller)
			return
		}

		this.Redirect("../info/admin_regSuccess", 302)
	}
}

func (this *AccountController) Logout(){
	// Check if user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID != nil {
		// UserID is set and can be deleted
		session.Delete("UserID")
	}
	fmt.Println("user successfully logged out", userID)
	this.Redirect("/", 302)
}