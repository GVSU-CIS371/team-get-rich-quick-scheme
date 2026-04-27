package routes

import (
	"invoicegen/internal/database"
	"invoicegen/internal/server/response"
	"net/http"
)

func GetStats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := r.Context().Value("db").(*database.Database)

		response.SendSuccessResponse(w, response.Flex{
			"userCount":    db.GetUserCount(),
			"invoiceCount": db.GetInvoiceCount(),
		})
	}
}
