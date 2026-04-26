package auth

import (
	"context"
	"invoicegen/internal/database"
	"net/http"
)

func UserMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sessionSecret := r.Header.Get("Authorization")
			if sessionSecret == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			db := r.Context().Value("db").(*database.Database)
			user, session, err := db.GetUserFromSessionSecret(sessionSecret)
			if err != nil || user == nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			ctx := r.Context()
			ctx = context.WithValue(ctx, "session", session)
			ctx = context.WithValue(ctx, "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ForceUserMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := r.Context().Value("user").(*database.User)
			if user == nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
