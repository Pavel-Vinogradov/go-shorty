package request

import (
	"main/pkg"
	"main/pkg/response"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := pkg.Decode[T](r.Body)
	if err != nil {
		errData := response.ErrorsResponse{Code: http.StatusBadRequest, Message: err.Error()}
		response.Json(w, errData, http.StatusBadRequest)
		return nil, err
	}
	err = pkg.IsValid(body)
	if err != nil {
		errData := response.ErrorsResponse{Code: http.StatusBadRequest, Message: err.Error()}
		response.Json(w, errData, http.StatusBadRequest)
		return nil, err
	}
	return &body, nil
}
