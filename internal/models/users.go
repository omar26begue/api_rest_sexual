package models

type Users struct {
	Identifier        string `json:"identifier" bson:"identifier" validate:"required"`
	Name              string `json:"name" bson:"name" validate:"required" example:"Pepe Gonzalez"`
	Email             string `json:"email" bson:"email" validate:"required,email" example:"pepe@gmail.com"`
	Password          string `json:"password" bson:"password" validate:"required"`
	Age               int    `json:"age,omitempty" bson:"age"`
	Sex               string `json:"sex,omitempty" bson:"sex" example:"Masculino"`
	IdReligion        string `json:"id_religion,omitempty" bson:"id_religion"`
	SexualOrientation string `json:"sexual_orientation,omitempty" bson:"sexual_orientation"`
	Active            bool   `json:"active" bson:"active"`
	Token             string `json:"token,omitempty" bson:"token"`
	Coins             int    `json:"coins" bson:"coins" example:"0"`
}

type UsersRequest struct {
	Name              string `json:"name" bson:"name" validate:"required" example:"Pepe Gonzalez"`
	Email             string `json:"email" bson:"email" validate:"required,email" example:"pepe@gmail.com"`
	Password          string `json:"password" bson:"password" validate:"required"`
	Age               int    `json:"age,omitempty" bson:"age" validate:"min=12,max=100"`
	Sex               string `json:"sex,omitempty" bson:"sex" example:"Masculino"`
	IdReligion        string `json:"id_religion,omitempty" bson:"id_religion"`
	SexualOrientation string `json:"sexual_orientation,omitempty" bson:"sexual_orientation"`
	Active            bool   `json:"active" bson:"active"`
}
