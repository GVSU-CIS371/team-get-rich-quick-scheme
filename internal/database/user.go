package database

import (
	"errors"
	"invoicegen/internal/security/crypto"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"uniqueIndex"`
	Password  string
}

func (db *Database) CreateUser(firstName, lastName, email, password string) (*User, error) {
	// Hash password
	hashedPassword, err := crypto.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Create user struct
	u := &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  hashedPassword,
	}

	// Attempt to create user
	tx := db.client.Create(u)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return u, nil
}

func (db *Database) GetUserByEmail(email string) *User {
	user := new(User)
	tx := db.client.Where(&User{Email: email}).First(user)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return user
}

func (user *User) ValidatePassword(password string) bool {
	valid, err := crypto.VerifyPassword(password, user.Password)
	return err == nil && valid
}
