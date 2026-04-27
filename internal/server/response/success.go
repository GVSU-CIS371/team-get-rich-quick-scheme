package response

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	BasicResponse
	Data any `json:"data,omitempty"`
}

type Flex map[string]any

func SendSuccessResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	res := SuccessResponse{
		BasicResponse: BasicResponse{
			Success: true,
		},
		Data: data,
	}
	jsonRes, _ := json.Marshal(res)
	_, _ = w.Write(jsonRes)
}
