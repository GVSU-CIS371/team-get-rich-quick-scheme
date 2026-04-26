package database

import "gorm.io/gorm"

type Invoice struct {
	gorm.Model
	Note           string
	OrganizationID uint
	Organization   *Organization
}
