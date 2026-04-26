package routes

import (
	"invoicegen/internal/database"
	"invoicegen/internal/server/response"
	"net/http"
)

func GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*database.User)

		response.SendSuccessResponse(w, response.Flex{
			"user": user,
		})
	}
}
