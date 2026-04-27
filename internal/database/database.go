package database

import (
	"context"
	"net/http"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	client *gorm.DB
}

func Connect(uri string) (*Database, error) {
	db, err := gorm.Open(sqlite.Open(uri), &gorm.Config{
		TranslateError: true,
	})
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

	err = db.AutoMigrate(&Invoice{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&InvoiceItem{})
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "db", db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
