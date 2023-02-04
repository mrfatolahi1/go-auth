package main

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email       string `gorm:"size:255;not null;unique" json:"content" binding:"required"`
	PhoneNumber string `gorm:"size:255;not null;unique" json:"content" binding:"required"`
	Gender      string `gorm:"size:8" json:"content"`
	FirstName   string `gorm:"size:255" json:"content"`
	LastName    string `gorm:"size:255" json:"content"`
	Password    string `json:"password" `
}

func (user User) toString() (toString string) {
	return "\nEmail: " + user.Email +
		"\nPhoneNumber: " + user.PhoneNumber +
		"\nGender: " + user.Gender +
		"\nFirstName: " + user.FirstName +
		"\nLastName: " + user.LastName +
		"\nPassword: " + user.Password
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

type DBToken struct {
	gorm.Model
	UserId         uint      `json:"content"`
	Token          string    `gorm:"size:255;not null;unique" json:"content" binding:"required"`
	ExpirationTime time.Time `json:"content"`
}
