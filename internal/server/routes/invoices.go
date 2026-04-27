package routes

import (
	"invoicegen/internal/database"
	"invoicegen/internal/server/response"
	"net/http"
	"strconv"

	"github.com/Oudwins/zog"
	"github.com/go-chi/chi/v5"
)

type invoiceData struct {
	Note string
}

var invoiceSchema = zog.Struct(zog.Shape{
	"note": zog.String().Required(zog.Message("Note is required")).
		Max(1000, zog.Message("Note is too long")),
})

type invoiceItemData struct {
	Description string
	Quantity    int
	Price       int
}

var invoiceItemSchema = zog.Struct(zog.Shape{
	"description": zog.String().Required(zog.Message("Description is required")),
	"quantity":    zog.Int().Required(zog.Message("Quantity is required")),
	"price":       zog.Int().Required(zog.Message("Price is required")),
})

func PostInvoices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		invData, err := response.ParseJsonRequest[invoiceData](w, r, invoiceSchema)
		if err != nil {
			return
		}

		user := r.Context().Value("user").(*database.User)
		db := r.Context().Value("db").(*database.Database)

		orgIdStr := chi.URLParam(r, "orgID")

		orgId, err := strconv.Atoi(orgIdStr)
		if err != nil {
			response.SendErrorResponse(w, nil, []string{"organization not found"}, http.StatusNotFound)
			return
		}

		org, err := db.GetOrganization(user, uint(orgId))
		if org == nil || err != nil {
			response.SendErrorResponse(w, nil, []string{"organization not found"}, http.StatusNotFound)
			return
		}

		inv := db.CreateInvoice(org, invData.Note)
		response.SendSuccessResponse(w, response.Flex{
			"invoice": inv,
		})
	}
}

func PostInvoiceItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		invItemData, err := response.ParseJsonRequest[invoiceItemData](w, r, invoiceItemSchema)
		if err != nil {
			return
		}

		user := r.Context().Value("user").(*database.User)
		db := r.Context().Value("db").(*database.Database)

		orgIdStr := chi.URLParam(r, "orgID")
		invIdStr := chi.URLParam(r, "invID")

		orgId, err := strconv.Atoi(orgIdStr)
		if err != nil {
			response.SendErrorResponse(w, nil, []string{"invoice not found"}, http.StatusNotFound)
			return
		}

		invId, err := strconv.Atoi(invIdStr)
		if err != nil {
			response.SendErrorResponse(w, nil, []string{"invoice not found"}, http.StatusNotFound)
			return
		}

		org, err := db.GetOrganization(user, uint(orgId))
		if org == nil || err != nil {
			response.SendErrorResponse(w, nil, []string{"invoice not found"}, http.StatusNotFound)
			return
		}

		invoice, err := db.GetInvoice(org, uint(invId))
		if invoice == nil || err != nil {
			response.SendErrorResponse(w, nil, []string{"invoice not found"}, http.StatusNotFound)
			return
		}

		invItem := db.CreateInvoiceItem(invoice, invItemData.Description,
			uint(invItemData.Quantity), uint(invItemData.Price))

		response.SendSuccessResponse(w, response.Flex{
			"invoiceItem": invItem,
		})
	}
}

func GetInvoices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*database.User)
		db := r.Context().Value("db").(*database.Database)

		orgIdStr := chi.URLParam(r, "orgID")

		orgId, err := strconv.Atoi(orgIdStr)
		if err != nil {
			response.SendErrorResponse(w, nil, []string{"organization not found"}, http.StatusNotFound)
			return
		}

		org, err := db.GetOrganization(user, uint(orgId))
		if org == nil || err != nil {
			response.SendErrorResponse(w, nil, []string{"organization not found"}, http.StatusNotFound)
			return
		}

		invoices := db.GetInvoices(org)

		response.SendSuccessResponse(w, response.Flex{
			"invoices": invoices,
		})
	}
}

func GetInvoice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*database.User)
		db := r.Context().Value("db").(*database.Database)

		orgIdStr := chi.URLParam(r, "orgID")
		invIdStr := chi.URLParam(r, "invID")

		orgId, err := strconv.Atoi(orgIdStr)
		if err != nil {
			response.SendErrorResponse(w, nil, []string{"invoice not found"}, http.StatusNotFound)
			return
		}

		invId, err := strconv.Atoi(invIdStr)
		if err != nil {
			response.SendErrorResponse(w, nil, []string{"invoice not found"}, http.StatusNotFound)
			return
		}

		org, err := db.GetOrganization(user, uint(orgId))
		if org == nil || err != nil {
			response.SendErrorResponse(w, nil, []string{"invoice not found"}, http.StatusNotFound)
			return
		}

		invoice, err := db.GetInvoice(org, uint(invId))
		if invoice == nil || err != nil {
			response.SendErrorResponse(w, nil, []string{"invoice not found"}, http.StatusNotFound)
			return
		}

		response.SendSuccessResponse(w, response.Flex{
			"invoice": invoice,
		})
	}
}
