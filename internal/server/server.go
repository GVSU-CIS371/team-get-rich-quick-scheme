package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Config struct {
	Host string
	Dev  bool
}

func Run(config *Config) error {
	r, err := setupRoutes(config)
	if err != nil {
		return err
	}

	return http.ListenAndServe(config.Host, r)
}

func setupRoutes(config *Config) (*chi.Mux, error) {
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
