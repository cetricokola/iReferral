package controllers

import (
	"crypto/tls"
	"encoding/json"
	"fmt"

	"github.com/gomail"
)

type EmailController struct {
	MainController
}

type EmailJson struct {
	Contact        string
	ContactName    string
	ContactEmail   string
	ContactSubject string
	ContactMessage string
}

var emailRes EmailJson

func (this *EmailController) Get() {
	this.Contactus()
}

func (this *EmailController) Post() {

	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	name := dataform["name"].(string)
	email := dataform["email"].(string)
	subject := dataform["subject"].(string)
	message := dataform["message"].(string)

	fmt.Println(dataform)

	if EmailValid(email) == true {

		m := gomail.NewMessage()
		m.SetHeader("From", email)
		m.SetHeader("To", "cetokola2015@gmail.com")
		m.SetAddressHeader("Cc", "okolacetric@gmail.com", name)
		m.SetHeader("Subject", subject)
		m.SetBody("text/html", message)
		// m.Attach("/home/Alex/lolcat.jpg")
		d := gomail.NewPlainDialer("smtp.gmail.com", 587, "cetokola2015@gmail.com", "password2015Gmail")
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		err := d.DialAndSend(m)
		if err != nil {
			fmt.Println("the error is", err)
			emailRes.Contact = "/"
			obj, _ := json.Marshal(emailRes)
			this.Ctx.Output.Header("Content-Type", "application/json")
			this.Ctx.Output.Body(obj)
			panic(err)
		}

		if err == nil {
			emailRes.Contact = "/"
			obj, _ := json.Marshal(emailRes)
			this.Ctx.Output.Header("Content-Type", "application/json")
			this.Ctx.Output.Body(obj)
		}
	}
	if EmailValid(email) == false {
		emailRes.ContactEmail = "incorrect"
	}

	if email == "" {
		emailRes.ContactEmail = "invalid"
	}

	if subject == "" {
		emailRes.ContactSubject = "invalid"
	}

	if message == "" {
		emailRes.ContactMessage = "invalid"
	}

	if name == "" {
		emailRes.ContactName = "invalid"
	}

	obj, _ := json.Marshal(emailRes)
	this.Ctx.Output.SetStatus(300)
	this.Ctx.Output.Header("Content-Type", "application/json")
	this.Ctx.Output.Body(obj)

}
