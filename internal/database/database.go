package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	client *gorm.DB
}

func Connect(uri string) (*Database, error) {
	db, err := gorm.Open(sqlite.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = migrate(db)
	if err != nil {
		return nil, err
	}

	return &Database{
		client: db,
	}, nil
}

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Session{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&PasswordReset{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Organization{})
	if err != nil {
		return err
	}

	return nil
}
