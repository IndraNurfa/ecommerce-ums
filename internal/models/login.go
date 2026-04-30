package models

import "github.com/go-playground/validator/v10"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l LoginRequest) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type LoginResponse struct {
	UserID       int    `json:"user_id"`
	Username     string `json:"username"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
