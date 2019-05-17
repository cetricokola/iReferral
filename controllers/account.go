package controllers

import (
	"fmt"
	"iReferral/models"
	"regexp"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"golang.org/x/crypto/bcrypt"
)

type AccountController struct {
	MainController
}

var flash = beego.NewFlash()
var valid = validation.Validation{}

var Name string
var Pfname string
var Plname string
var Nam string
var PiD string

//*PATIENT PORTAL LOG IN*//
func (this *AccountController) Patients_login() {
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID != nil {
		// User is logged in already, display another page
		this.Redirect("/phome", 302)
		return
	}
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
			flash.Error("You've entered incorrect huduma number or password.Try again!")
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
			flash.Error("You've entered incorrect password.Try again!")
			flash.Store(&this.Controller)
			return
		}
		//set session for patient log in
		session.Set("UserID", hudumaNo)
		Pfname = patient.FirstName
		Plname = patient.LastName
		PiD = hudumaNo
		fmt.Println(patient.HudumaNo, ":successful log in ")
		this.Redirect("/phome", 302)
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
			flash.Error("You've entered incorrect national id number or password.Try again!")
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
			flash.Error("You've entered incorrect password.Try again")
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
			flash.Error("Incorrect national id or password.Try again!")
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
			flash.Error("You've entered incorrect password.Try again!")
			flash.Store(&this.Controller)
			return
		}
		//set session
		this.SetSession("UserID", iD)
		Name = admin.Email
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

		//input validation
		valid.Required(first, "first name") //No null values are accepted
		valid.Required(last, "last name")
		valid.Required(huduma, "huduma no")   //non empty values are permitted
		valid.Numeric(huduma, "huduma no")    //numeric values for huduma are permitted
		valid.Length(huduma, 11, "huduma no") //11 digits are permitted
		valid.Required(dob, "date of birth")  //non empty values are permitted
		valid.Required(phone, "phone number") // non empty values are permitted
		valid.Required(sex, "sex")            // non empty values are permitted
		valid.Required(confirmPass, "confirm password")
		valid.MinSize(submittedpassword, 8, "password") // minimum size for the password is 8 characters
		valid.Length(phone, 10, "phone number")
		valid.Numeric(phone, "phone number")

		if phone != "" {
			matched, err := regexp.MatchString("^(07)([0-9]{8})$", phone)
			if err != nil {
				valid.SetError("Phone", "Something wierd is going on here.")
			} else if !matched {
				valid.SetError("Phone", "Phone number should be something like: 07123145678")
			}
		} // end Phone2 validation if block

		if valid.HasErrors() {
			errormap := []string{}
			for _, err := range valid.Errors {
				errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")

			}
			this.Data["Errors"] = errormap
			return
		}
		//Check if the password contains a capital letter
		r, _ := regexp.Compile(`[A-Z]`)
		if !r.MatchString(submittedpassword) {
			flash.Error("Password must contain at least one capital letter")
			flash.Store(&this.Controller)
			return
		}
		// Check password contain lowercase letter
		r, _ = regexp.Compile(`[a-z]`)
		if !r.MatchString(submittedpassword) {
			flash.Error("Password must contain at least one lower case letter")
			flash.Store(&this.Controller)
			return
		}
		// Check password contain number
		r, _ = regexp.Compile(`[0-9]`)
		if !r.MatchString(submittedpassword) {
			flash.Error("Password must contain at least one number")
			flash.Store(&this.Controller)
			return
		}
		//match the submitted password with the confirm password
		if submittedpassword != confirmPass {
			fmt.Println("The password submitted does no match the confirm password")
			flash.Error("The password submitted does no match the confirm password!")
			flash.Store(&this.Controller)
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
			flash.Error(phone + " already registered!")
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
		//input validation
		valid.Required(empId, "National id") //No null values are accepted
		valid.Required(email, "email")
		valid.Email(email, "email")           //input a valid email address
		valid.Required(phone, "phone number") // non empty values are permitted
		valid.Required(confirmPass, "confirm password")
		valid.MinSize(submittedpassword, 8, "password") // minimum size for the password is 8 characters
		valid.Length(phone, 10, "phone number")
		valid.Numeric(phone, "phone number")

		if phone != "" {
			matched, err := regexp.MatchString("^(07)([0-9]{8})$", phone)
			if err != nil {
				valid.SetError("Phone", "Something wierd is going on here.")
			} else if !matched {
				valid.SetError("Phone", "Phone number should be something like: 07123145678")
			}
		} // end Phone2 validation if block

		if valid.HasErrors() {
			errormap := []string{}
			for _, err := range valid.Errors {
				errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
			}
			this.Data["Errors"] = errormap
			return
		}
		//Check if the password contains a capital letter
		r, _ := regexp.Compile(`[A-Z]`)
		if !r.MatchString(submittedpassword) {
			flash.Error("Password must contain at least one capital letter")
			flash.Store(&this.Controller)
			return
		}
		// Check password contain lowercase letter
		r, _ = regexp.Compile(`[a-z]`)
		if !r.MatchString(submittedpassword) {
			flash.Error("Password must contain at least one lower case letter")
			flash.Store(&this.Controller)
			return
		}
		// Check password contain number
		r, _ = regexp.Compile(`[0-9]`)
		if !r.MatchString(submittedpassword) {
			flash.Error("Password must contain at least one number")
			flash.Store(&this.Controller)
			return
		}

		//match the submitted password with the confirm password
		if submittedpassword != confirmPass {
			fmt.Println("The password submitted does no match the confirm password")
			flash.Error("The password submitted does no match the confirm password!")
			flash.Store(&this.Controller)
			return
		}
		//hash the password
		password, _ := HashPassword(submittedpassword)
		//******** Save user info to database

		o := orm.NewOrm()
		o.Using("default")

		//check if the emp id is valid.The one admin assigned to the employer and the emp id is already used by the employee to sign up
		exist := o.QueryTable("employee").Filter("EmpId", empId).Exist()
		exist1 := o.QueryTable("employee_account").Filter("EmpId", empId).Exist()
		if exist == true && exist1 == false {
			staff := models.Employee_account{EmpId: empId, Email: email, PhoneNo: phone, Password: password}
			_, err := o.Insert(&staff)
			if err != nil {
				fmt.Println(err)
				flash.Error(phone + " already registered!")
				flash.Store(&this.Controller)
				return
			}
		} else {
			fmt.Println("Invalid emp id")
			flash.Error("Invalid employee id.Try again")
			flash.Store(&this.Controller)
			return
		}
		fmt.Println("Successfull sign up")
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
		//input validation

		valid.Required(id, "National id") //No null values are accepted
		valid.Required(email, "email")
		valid.Email(email, "email") //input a valid email address
		valid.Required(confirmPass, "confirm password")
		valid.MinSize(submittedpassword, 8, "password") // minimum size for the password is 8 characters

		if valid.HasErrors() {
			errormap := []string{}
			for _, err := range valid.Errors {
				errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")

			}
			this.Data["Errors"] = errormap
			return
		}
		//Check if the password contains a capital letter
		r, _ := regexp.Compile(`[A-Z]`)
		if !r.MatchString(submittedpassword) {
			flash.Error("Password must contain at least one capital letter")
			flash.Store(&this.Controller)
			return
		}
		// Check password contain lowercase letter
		r, _ = regexp.Compile(`[a-z]`)
		if !r.MatchString(submittedpassword) {
			flash.Error("Password must contain at least one lower case letter")
			flash.Store(&this.Controller)
			return
		}
		// Check password contain number
		r, _ = regexp.Compile(`[0-9]`)
		if !r.MatchString(submittedpassword) {
			flash.Error("Password must contain at least one number")
			flash.Store(&this.Controller)
			return
		}

		//match the submitted password with the confirm password
		if submittedpassword != confirmPass {
			fmt.Println("The password submitted does no match the confirm password")
			flash.Error("The password submitted does no match the confirm password!")
			flash.Store(&this.Controller)
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
			flash.Error(email + " already registered!")
			flash.Store(&this.Controller)
			return
		}

		this.Redirect("../info/admin_regSuccess", 302)
	}
}

func (this *AccountController) Logout() {
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
