package handler

import (
	"fmt"
	"main/configs"
	"main/internal/middleware"
	"net/http"
)

type AuthHandlerImpl struct {
	*configs.Config
}
type AuthHandler struct {
	*configs.Config
}

// Registration godoc
// @Summary Регистрация пользователя
// @Description Создаёт нового пользователя
// @Tags auth
// @Accept  json
// @Produce  json
// @Param request body RegistrationRequest true "Данные для регистрации"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 405 {object} map[string]string
// @Router /auth/registration [post]
func (h *AuthHandler) Registration(w http.ResponseWriter, r *http.Request) {
	if !middleware.AllowMethod(w, r, http.MethodPost) {
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintln(w, "Registration OK")

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if !middleware.AllowMethod(w, r, http.MethodPost) {
		return
	}
	fmt.Println(h.Auth.Secret)
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintln(w, "Login OK")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func NewAuthHandler(router *http.ServeMux, impl AuthHandlerImpl) {
	handler := &AuthHandler{
		Config: impl.Config,
	}
	router.HandleFunc("/auth/registration", handler.Registration)
	router.HandleFunc("/auth/login", handler.Login)
}
