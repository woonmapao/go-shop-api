package modes

import (
	"time"

	"gorm.io/gorm"
)

type UserDetail struct {
	gorm.Model
	Name     string
	Email    *string
	Age      uint8
	Birthday *time.Time
}
