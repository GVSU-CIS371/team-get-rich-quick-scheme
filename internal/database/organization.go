package database

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name   string
	UserID uint
	User   *User
}
