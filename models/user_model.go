package models

import (
	"gorm.io/gorm"
)

type UserDetail struct {
	gorm.Model
	Name  string
	Email *string
	Age   uint8
}
