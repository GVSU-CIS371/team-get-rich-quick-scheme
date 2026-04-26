package database

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Secret     string `gorm:"uniqueIndex"`
	Expiration time.Time
	UserID     uint
	User       *User
}
