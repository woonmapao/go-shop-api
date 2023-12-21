package models

import (
	"gorm.io/gorm"
)

type UserDetail struct {
	gorm.Model
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
}
