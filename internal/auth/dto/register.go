package dto

type RegisterRequest struct {
	Name string `json:"name"`
	LoginRequest
}

type RegisterResponse struct {
	Token string `json:"token"`
}
