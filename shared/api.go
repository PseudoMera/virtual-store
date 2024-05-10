package shared

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func WriteErrorResponse(w http.ResponseWriter, e error, statusCode int) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(ErrorResponse{
		Error:   true,
		Message: e.Error(),
	}); err != nil {
		http.Error(w, e.Error(), statusCode)
	}
}

func WriteResponse(statusCode int, data any, w http.ResponseWriter) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		WriteErrorResponse(w, err, http.StatusInternalServerError)
	}
}
