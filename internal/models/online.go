package models

import "time"

type Online struct {
	Identifier string    `json:"identifier" bson:"identifier" validate:"required"`
	Uuid       string    `json:"uuid" bson:"uuid" validate:"required"`
	Time       time.Time `json:"time" bson:"time" validate:"required"`
}
