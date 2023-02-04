package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

var database *gorm.DB

func ConnectToDatabase() {
	var err error
	host := "localhost"
	username := "postgres"
	password := "123456"
	databaseName := "postgres"
	port := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, username, password, databaseName, port)
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}

func MigrateDatabase() {
	database.AutoMigrate(&User{})
	err := database.AutoMigrate(&DBToken{})
	if err != nil {
		println(err.Error())
		return
	}
}

func SaveUser(context *gin.Context, user *User) bool {
	err := database.Create(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "can't save data",
		})
		return false
	}
	return true
}

func loadUserWithEmail(context *gin.Context, email string) (user User) {
	err := database.Where("email=?", email).Find(&user).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "can't load user",
		})
		return
	}
	return user
}

func loadUserWithPhoneNumber(context *gin.Context, phoneNumber string) (user User) {
	err := database.Where("phoneNumber = ?", phoneNumber).Find(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "can't load user",
		})
		return
	}
	return user
}

func loadUserWithId(id uint) (user User) {
	database.Where("id = ?", id).Find(&user)
	return user
}

func saveExpiredToken(token string, userId uint) {
	var dbToken = DBToken{
		Token:          token,
		UserId:         userId,
		ExpirationTime: getTokenExpiration(token),
	}

	err := database.Create(&dbToken)
	println(err.Error)
}

func loadTokenObject(token string) (tokenObj DBToken) {
	err := database.Where("token=?", token).Find(&tokenObj).Error
	if err != nil {
		println(err.Error())
		return
	}
	return tokenObj
}
