package routes

import (
	"invoicegen/internal/database"
	"invoicegen/internal/server/response"
	"net/http"

	"github.com/Oudwins/zog"
)

type loginData struct {
	Email    string
	Password string
}

var loginSchema = zog.Struct(zog.Shape{
	"email": zog.String().Required(zog.Message("Email is required")).Email(zog.Message("Email is invalid")),
	"password": zog.String().Required(zog.Message("Password is required")).Min(11,
		zog.Message("Password is too short")).Max(128, zog.Message("Password is too long")),
})

func PostLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login, err := response.ParseJsonRequest[loginData](w, r, loginSchema)
		if err != nil {
			return
		}

		db := r.Context().Value("db").(*database.Database)

		// Get the user
		user := db.GetUserByEmail(login.Email)
		if user == nil {
			response.SendErrorResponse(w, nil, []string{"Invalid Email or Password"}, http.StatusBadRequest)
			return
		}

		if !user.ValidatePassword(login.Password) {
			response.SendErrorResponse(w, nil, []string{"Invalid Email or Password"}, http.StatusBadRequest)
			return
		}

		session, err := db.CreateSession(user)
		if err != nil {
			response.SendErrorResponse(w, nil, nil, http.StatusInternalServerError)
			return
		}

		response.SendSuccessResponse(w, response.Flex{
			"session":    session.Secret,
			"expiration": session.Expiration.UnixMilli(),
		})
	}
}
