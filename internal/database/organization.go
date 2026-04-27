package database

import (
	"errors"

	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	Name        string
	Description string
	UserID      uint  `json:"-"`
	User        *User `json:"-"`
}

func (db *Database) GetOrganizations(user *User) []Organization {
	var organizations []Organization
	db.client.Where(&Organization{UserID: user.ID}).Find(&organizations)

	return organizations
}

func (db *Database) GetOrganization(user *User, id uint) (*Organization, error) {
	org := new(Organization)
	tx := db.client.Where("id = ? AND user_id = ?", id, user.ID).First(org)

	if tx.Error != nil {
		return nil, errors.New("unable to find organization")
	}

	return org, nil
}

func (db *Database) CreateOrganization(user *User, name, description string) *Organization {
	org := &Organization{
		Name:        name,
		Description: description,
		UserID:      user.ID,
	}

	db.client.Create(org)

	return org
}
