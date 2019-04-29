package controllers

import (
	"golang.org/x/crypto/bcrypt" //import the hashing algorithm
	"log"
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
