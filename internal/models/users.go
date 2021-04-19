package models

type Users struct {
	Identifier        string `json:"identifier" bson:"identifier" validate:"required"`
	Name              string `json:"name" bson:"name" validate:"required" example:"Pepe Gonzalez"`
	Email             string `json:"email" bson:"email" validate:"required,email" example:"pepe@gmail.com"`
	Password          string `json:"password" bson:"password" validate:"required"`
	Age               int    `json:"age" bson:"age" validate:"required"`
	Sex               string `json:"sex" bson:"sex" validate:"required" example:"Masculino"`
	IdReligion        string `json:"id_religion" bson:"id_religion" validate:"required,uuid"`
	SexualOrientation string `json:"sexual_orientation" bson:"sexual_orientation" validate:"required,uuid"`
	Active            bool   `json:"active" bson:"active" validate:"required"`
}
