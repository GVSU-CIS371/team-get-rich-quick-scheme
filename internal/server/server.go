package server

import (
	"invoicegen/internal/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Config struct {
	Host     string
	Database string
	Dev      bool
}

func Run(config *Config) error {
	db, err := database.Connect(config.Database)
	if err != nil {
		return err
	}

	r, err := setupRoutes(config, db)
	if err != nil {
		return err
	}

	return http.ListenAndServe(config.Host, r)
}

func setupRoutes(config *Config, db *database.Database) (*chi.Mux, error) {
	r := chi.NewRouter()

	// Setup dev middleware
	if config.Dev {
		r.Use(middleware.Logger)
	}

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	apiRouter := chi.NewRouter()

	r.Mount("/api/v1", apiRouter)
	setupFrontend(r, config.Dev)

	return r, nil
}
