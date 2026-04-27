package routes

import (
	"invoicegen/internal/database"
	"invoicegen/internal/server/response"
	"net/http"
	"strconv"

	"github.com/Oudwins/zog"
	"github.com/go-chi/chi/v5"
)

type organizationData struct {
	Name        string
	Description string
}

var organizationSchema = zog.Struct(zog.Shape{
	"name": zog.String().Required(zog.Message("Organization name is required")).
		Min(2, zog.Message("Organization name is too short")).
		Max(80, zog.Message("Organization name is too long")),
	"description": zog.String().Required(zog.Message("Organization description is required")).
		Max(255, zog.Message("Organization description is too long")),
})

func GetOrganizations() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*database.User)
		db := r.Context().Value("db").(*database.Database)

		response.SendSuccessResponse(w, response.Flex{
			"organizations": db.GetOrganizations(user),
		})
	}
}

func GetOrganization() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*database.User)
		db := r.Context().Value("db").(*database.Database)

		ordIdStr := chi.URLParam(r, "orgID")
		ordId, err := strconv.Atoi(ordIdStr)
		if err != nil {
			response.SendErrorResponse(w, nil, []string{"organization not found"}, http.StatusNotFound)
			return
		}

		org, err := db.GetOrganization(user, uint(ordId))
		if err != nil {
			response.SendErrorResponse(w, nil, []string{"organization not found"}, http.StatusNotFound)
			return
		}

		response.SendSuccessResponse(w, response.Flex{
			"organization": org,
		})
	}
}

func PostOrganizations() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orgData, err := response.ParseJsonRequest[organizationData](w, r, organizationSchema)
		if err != nil {
			return
		}

		user := r.Context().Value("user").(*database.User)
		db := r.Context().Value("db").(*database.Database)

		org := db.CreateOrganization(user, orgData.Name, orgData.Description)
		response.SendSuccessResponse(w, response.Flex{
			"organization": org,
		})
	}
}
