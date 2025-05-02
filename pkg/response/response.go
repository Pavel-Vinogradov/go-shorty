package response

import (
	"encoding/json"
	"net/http"
)

type ErrorsResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Json(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err.Error())
	}

}
