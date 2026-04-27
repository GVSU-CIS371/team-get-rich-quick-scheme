package response

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Oudwins/zog"
)

var (
	ErrBadRequest = errors.New("bad request data")
)

func ParseJsonRequest[T any](w http.ResponseWriter, r *http.Request, schema *zog.StructSchema) (*T, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil, ErrBadRequest
	}
	m := new(map[string]any)
	err = json.Unmarshal(data, m)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil, ErrBadRequest
	}
	t := new(T)
	issues := schema.Parse(*m, t)
	if len(issues) == 0 {
		return t, nil
	}
	reqIssues := RequestIssues{}

	for _, issue := range issues {
		if v, ok := reqIssues[issue.Path[0]]; ok {
			reqIssues[issue.Path[0]] = v + " and " + issue.Message
			continue
		}

		reqIssues[issue.Path[0]] = issue.Message
	}

	SendErrorResponse(w, reqIssues, nil, http.StatusBadRequest)
	return nil, ErrBadRequest
}
