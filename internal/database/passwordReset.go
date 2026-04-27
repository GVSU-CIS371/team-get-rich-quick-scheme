package database

import (
	"time"

	"gorm.io/gorm"
)

type PasswordReset struct {
	gorm.Model
	Secret     string `gorm:"uniqueIndex"`
	Expiration time.Time
	UserID     uint
	User       *User
}
