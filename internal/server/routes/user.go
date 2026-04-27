package routes

import (
	"invoicegen/internal/database"
	"invoicegen/internal/server/response"
	"net/http"

	"github.com/Oudwins/zog"
)

type passwordChange struct {
	CurrentPassword string
	NewPassword     string
}

var passwordChangeSchema = zog.Struct(zog.Shape{
	"currentPassword": zog.String().Required(zog.Message("Password is required")).Min(11,
		zog.Message("Password is too short")).Max(128, zog.Message("Password is too long")),
	"newPassword": zog.String().Required(zog.Message("Password is required")).Min(11,
		zog.Message("Password is too short")).Max(128, zog.Message("Password is too long")),
})

func GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*database.User)

		response.SendSuccessResponse(w, response.Flex{
			"user": user,
		})
	}
}

func PutPassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		passChange, err := response.ParseJsonRequest[passwordChange](w, r, passwordChangeSchema)
		if err != nil {
			return
		}

		user := r.Context().Value("user").(*database.User)
		db := r.Context().Value("db").(*database.Database)

		if !user.ValidatePassword(passChange.CurrentPassword) {
			response.SendErrorResponse(w, response.RequestIssues{"currentPassword": "Current password is invalid"},
				nil, http.StatusBadRequest)
			return
		}

		err = db.UpdatePassword(user, passChange.NewPassword)
		if err != nil {
			response.SendErrorResponse(w, nil, []string{"unable to update password"},
				http.StatusInternalServerError)
			return
		}

		response.SendSuccessResponse(w, nil)
	}
}
