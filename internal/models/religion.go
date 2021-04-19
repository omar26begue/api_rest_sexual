package models

type Religion struct {
	Identifier string `json:"identifier" bson:"identifier" validate:"required,uuid"`
	Name       string `json:"name" bson:"name" validate:"required"`
}
