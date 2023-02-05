package main

import (
	"encoding/json"
	_ "encoding/json"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	_ "strconv"
)
import _ "net/http"

func manage() {
	router := gin.Default()

	router.POST("/SignUp", func(context *gin.Context) {
		signUp(context)
	})

	router.GET("/SignIn", func(context *gin.Context) {
		signIn(context)
	})

	router.GET("/UserInfo", func(context *gin.Context) {
		userInfo(context)
	})

	router.POST("/SignOut", func(context *gin.Context) {
		signOut(context)
	})

	router.Run()
}

func initDB() {
	ConnectToDatabase()
	MigrateDatabase()
	initRedis()
}

func signUp(context *gin.Context) {
	jsonData, _ := ioutil.ReadAll(context.Request.Body)
	dataMap := make(map[string]string)
	err := json.Unmarshal(jsonData, &dataMap)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request!",
		})
		return
	}

	if !ckeckDataValidation(context, dataMap) {
		return
	}

	var firstName = dataMap["first_name"]
	var lastName = dataMap["last_name"]
	var email = dataMap["email"]
	var phoneNumber = dataMap["phone_number"]
	var gender = dataMap["gender"]
	var password = dataMap["password"]

	var user = User{
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
		Gender:      gender,
		Password:    HashPassword(password),
	}

	if !SaveUser(context, &user) {
		return
	}

	var token string
	token, err = GenerateJWT(email, user.ID)

	context.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func signIn(context *gin.Context) {
	jsonData, _ := ioutil.ReadAll(context.Request.Body)
	dataMap := make(map[string]string)
	err := json.Unmarshal(jsonData, &dataMap)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request!",
		})
		return
	}
	var email = dataMap["email"]
	var phoneNumber = dataMap["phone_number"]
	var password = dataMap["password"]

	var user User
	if phoneNumber == "" {
		user = loadUserWithEmail(context, email)
		if &user == nil {
			return
		}
	} else {
		user = loadUserWithPhoneNumber(context, phoneNumber)
		if &user == nil {
			return
		}
	}

	err2 := user.ValidatePassword(password)
	if err2 != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong password",
		})
		return
	}

	var token string
	token, err = GenerateJWT(email, user.ID)

	context.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func userInfo(context *gin.Context) {
	jsonData, _ := ioutil.ReadAll(context.Request.Body)
	dataMap := make(map[string]string)
	err := json.Unmarshal(jsonData, &dataMap)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request!",
		})
		return
	}
	var token = dataMap["token"]

	if checkKeyExistanceInRedis(token) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "token expired",
		})
		return
	}

	var tokenObject = loadTokenObjectFromDB(token)

	if &tokenObject != nil && tokenObject.UserId != 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "token expired",
		})
		return
	}

	err2 := ValidateToken(token)
	if err2 != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err2.Error(),
		})
		return
	}

	var userId = getUserIdFromToken(token)

	var user = loadUserWithId(userId)

	context.JSON(http.StatusOK, gin.H{
		"id":          user.ID,
		"first_name":  user.FirstName,
		"last_name":   user.LastName,
		"phoneNumber": user.PhoneNumber,
		"email":       user.Email,
		"gender":      user.Gender,
	})
	return
}

func signOut(context *gin.Context) {
	jsonData, _ := ioutil.ReadAll(context.Request.Body)
	dataMap := make(map[string]string)
	err := json.Unmarshal(jsonData, &dataMap)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request!",
		})
		return
	}
	var token = dataMap["token"]

	if checkKeyExistanceInRedis(token) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "token expired",
		})
		return
	}

	var tokenObject = loadTokenObjectFromDB(token)

	if &tokenObject != nil && tokenObject.UserId != 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "token expired",
		})
		return
	}

	err2 := ValidateToken(token)
	if err2 != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err2.Error(),
		})
		return
	}

	var userId = getUserIdFromToken(token)
	saveExpiredToken(token, userId)
}

func HashPassword(password string) (hashedPassword string) {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(passwordHash)
}

func start() {
	initDB()
	manage()
}
