package server

import (
	"invoicegen/internal/database"
	"invoicegen/internal/security/auth"
	"invoicegen/internal/server/routes"
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
	r.Use(db.Middleware())

	apiRouter := chi.NewRouter()
	apiRouter.Post("/login", routes.PostLogin())
	apiRouter.Post("/register", routes.PostRegister())
	apiRouter.Get("/stats", routes.GetStats())

	apiRouter.Group(func(r chi.Router) {
		r.Use(auth.UserMiddleware())

		r.Get("/user", routes.GetUser())
		r.Get("/organizations", routes.GetOrganizations())
		r.Post("/organizations", routes.PostOrganizations())
		r.Get("/organizations/{orgID}", routes.GetOrganization())
		r.Get("/organizations/{orgID}/invoices", routes.GetInvoices())
		r.Post("/organizations/{orgID}/invoices", routes.PostInvoices())
		r.Post("/organizations/{orgID}/invoices/{invID}/items", routes.PostInvoiceItem())
		r.Get("/organizations/{orgID}/invoices/{invID}", routes.GetInvoice())
	})

	r.Mount("/api/v1", apiRouter)
	setupFrontend(r, config.Dev)

	return r, nil
}
