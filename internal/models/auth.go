package models

type Auth struct {
	Email    string `json:"email" bson:"email" validate:"required,email" example:"user@users.com"`
	Password string `json:"password" bson:"password" validate:"required" minLength:"5"`
}

type AuthResponse struct {
	Identifier string `json:"identifier"`
	Token      string `json:"token"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}
