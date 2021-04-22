package models

type Facebook struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required"`
	Token string `json:"token"`
}
