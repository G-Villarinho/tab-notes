package models

type SendAuthenticationLinkPayload struct {
	Email string `json:"email" validate:"required,email"`
}

type RegisterPayload struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
