package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	BasicResponse
	Issues RequestIssues `json:"issues,omitempty"`
	Errors []string      `json:"errors,omitempty"`
}

type RequestIssues map[string]string

func SendErrorResponse(w http.ResponseWriter, issues RequestIssues, errors []string, status int) {
	if len(issues) == 0 {
		issues = nil
	}
	if len(errors) == 0 {
		errors = nil
	}
	errorResponse := ErrorResponse{
		BasicResponse: BasicResponse{
			Success: false,
		},
		Issues: issues,
		Errors: errors,
	}
	data, err := json.Marshal(errorResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}
