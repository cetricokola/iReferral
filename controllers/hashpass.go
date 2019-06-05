package controllers

import (
	"log"
	"regexp"

	"golang.org/x/crypto/bcrypt" //import the hashing algorithm
)

//hash the password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("ERROR:EncryptPassword: %s", err.Error())
	}
	return string(bytes), err
}

//compare the hashed password and the submitted password
func ComparePass(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//check if the password is more than 8 characters, contains a uppercase and lowercase plus a digit 
func IsValid(password string) bool {
	//Check if the password contains a capital letter
	r, _ := regexp.Compile(`[A-Z]`)
	if !r.MatchString(password) {
		return false
	}
	// Check password contain lowercase letter
	r, _ = regexp.Compile(`[a-z]`)
	if !r.MatchString(password) {
		return false
	}
	// Check password contain number
	r, _ = regexp.Compile(`[0-9]`)
	if !r.MatchString(password) {
		return false
	}
	if len(password) < 8{
		return false
	}

	return true
}

//email validation
func EmailValid(email string) bool{
	//check if the email entered is a valid email
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(email){
		return false
	}
	return true
	
}

//check the national id should have a maximum of 8 digits
func NationalIdValid(id string) bool{
	re := regexp.MustCompile("^[0-9]{8}$")
	if !re.MatchString(id){
		return false
	}
	return true
}

//check the employee id should have a maximum of 6 digits
func EmpIdValid(id string) bool{
	re := regexp.MustCompile("^[0-9]{6}$")
	if !re.MatchString(id){
		return false
	}
	return true
}//check the service code should have a maximum of 6 digits
func ServiceValid(serv string) bool{
	re := regexp.MustCompile("^[0-9]{6}$")
	if !re.MatchString(serv){
		return false
	}
	return true
}



//huduma validation
func HudumaValid(id string) bool{
	re := regexp.MustCompile("^[0-9]{11}$")
	if !re.MatchString(id){
		return false
	}
	return true
}

//phone validation
func PhoneValid(id string) bool{
	re := regexp.MustCompile("^[0-9]{10}$")
	if !re.MatchString(id){
		return false
	}
	return true
}