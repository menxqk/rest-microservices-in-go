package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/menxqk/rest-microservices-in-go/common/errors"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err *errors.RestError) {
	RespondJson(w, err.Status, err)
}
