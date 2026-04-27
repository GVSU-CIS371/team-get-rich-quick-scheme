package database

import (
	"errors"
	"invoicegen/internal/security/crypto"
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

var errSessionInvalid = errors.New("session is invalid")

func (db *Database) GetUserFromSessionSecret(secret string) (*User, *Session, error) {
	session := new(Session)
	result := db.client.Preload("User").Where("secret = ?", secret).First(session)

	if result.Error != nil || session == nil {
		return nil, nil, errSessionInvalid
	}

	if session.Expiration.Before(time.Now()) {
		db.client.Delete(session)
		return nil, nil, errSessionInvalid
	}

	return session.User, session, nil
}

func (db *Database) CreateSession(user *User) (*Session, error) {
	session := &Session{
		Secret:     crypto.GenerateRandomString(32),
		Expiration: time.Now().UTC().Add(24 * time.Hour),
		UserID:     user.ID,
	}

	_ = db.client.Create(session).Error
	return session, nil
}
