package routes

import (
	"errors"
	"invoicegen/internal/database"
	"invoicegen/internal/server/response"
	"net/http"

	"github.com/Oudwins/zog"
	"gorm.io/gorm"
)

type registerData struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

var registerSchema = zog.Struct(zog.Shape{
	"firstName": zog.String().Required(zog.Message("First name is required")).
		Min(2, zog.Message("First name is too short")).Max(20, zog.Message("First name is too long")),
	"lastName": zog.String().Required(zog.Message("Last name is required")).
		Min(2, zog.Message("Last name is too short")).Max(20, zog.Message("Last name is too long")),
	"email": zog.String().Required(zog.Message("Email is required")).Email(zog.Message("Email is invalid")),
	"password": zog.String().Required(zog.Message("Password is required")).Min(11,
		zog.Message("Password is too short")).Max(128, zog.Message("Password is too long")),
})

func PostRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		register, err := response.ParseJsonRequest[registerData](w, r, registerSchema)
		if err != nil {
			return
		}

		db := r.Context().Value("db").(*database.Database)

		user, err := db.CreateUser(register.FirstName, register.LastName, register.Email, register.Password)
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			response.SendErrorResponse(w, nil, []string{"You already have an account"}, http.StatusBadRequest)
			return
		}

		if err != nil {
			response.SendErrorResponse(w, nil, nil, http.StatusInternalServerError)
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
