package mode

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    *string
	Age      uint8
	Birthday *time.Time
}
