package handler

import (
	"fmt"
	"main/configs"
	"main/internal/auth/dto"
	"main/internal/middleware"
	"main/pkg/request"
	"main/pkg/response"
	"net/http"
)

type AuthHandlerImpl struct {
	*configs.Config
}
type AuthHandler struct {
	*configs.Config
}

func (h *AuthHandler) Registration(w http.ResponseWriter, r *http.Request) {
	if !middleware.AllowMethod(w, r, http.MethodPost) {
		return
	}
	_, err := request.HandleBody[dto.LoginRequest](w, r)
	if err != nil {
		return
	}
	var responseRegistration dto.RegisterResponse
	response.Json(w, responseRegistration, http.StatusCreated)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if !middleware.AllowMethod(w, r, http.MethodPost) {
		return
	}
	body, err := request.HandleBody[dto.LoginRequest](w, r)
	if err != nil {
		return
	}
	fmt.Println(body)
}

func NewAuthHandler(router *http.ServeMux, impl AuthHandlerImpl) {
	handler := &AuthHandler{
		Config: impl.Config,
	}
	router.HandleFunc("/auth/registration", handler.Registration)
	router.HandleFunc("/auth/login", handler.Login)
}
