package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/mail"
	"unicode"
)

func ckeckDataValidation(context *gin.Context, dataMap map[string]string) bool {
	var firstName = dataMap["first_name"]
	var lastName = dataMap["last_name"]
	var email = dataMap["email"]
	var phoneNumber = dataMap["phone_number"]
	var gender = dataMap["gender"]
	var password = dataMap["password"]

	if !checkEmailValidation(context, email) {
		return false
	}

	if !ckeckNameValidation(context, firstName) {
		return false
	}

	if !ckeckNameValidation(context, lastName) {
		return false
	}

	if !checkPhoneNumberValidation(context, phoneNumber) {
		return false
	}

	if !checkGenderValidation(context, gender) {
		return false
	}

	if !checkPasswordValidation(context, password) {
		return false
	}

	return true
}

func checkEmailValidation(context *gin.Context, email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email!",
		})
		return false
	}
	if len(email) > 255 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "too long email, email should have less than 255 characters",
		})
		return false
	}

	return true
}

func ckeckNameValidation(context *gin.Context, name string) bool {
	if !isAlphabetical(name) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "firstname and lastname should contain only alphabetical characters",
		})
		return false
	}
	if len(name) > 255 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "too long name, firstname and lastname should have less than 255 characters",
		})
		return false
	}

	return true
}

func checkPhoneNumberValidation(context *gin.Context, phoneNumber string) bool {
	if phoneNumber[0] != '0' {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "phonenumber should begin with 0",
		})
		return false
	}

	if len(phoneNumber) != 11 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "phonenumber should have exactly 11 characters",
		})
		return false
	}

	return true
}

func checkGenderValidation(context *gin.Context, gender string) bool {
	if len(gender) == 1 {
		if gender[0] == 'M' || gender[0] == 'F' {
			return true
		} else {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "gender value should be M or F",
			})
			return false
		}
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "gender length should be 1",
		})
		return false
	}

}

func checkPasswordValidation(context *gin.Context, password string) bool {
	if len(password) < 8 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "password should have at least 8 characters",
		})
		return false
	}

	return true
}

func isAlphabetical(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
