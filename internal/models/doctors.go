package models

type Doctors struct {
	Identifier string `json:"identifier" bson:"identifier" validate:"required"`
	Name       string `json:"name" bson:"name" validate:"required" example:"Juan Perez"`
	Specialty  string `json:"specialty" bson:"specialty" validate:"required" example:"Psicologa"`
	OnDoctor   string `json:"on_doctor" bson:"on_doctor" validate:"required" example:"Sobre el doctor"`
}

type DoctorsRequest struct {
	Name      string `json:"name" bson:"name" validate:"required" example:"Juan Perez"`
	Specialty string `json:"specialty" bson:"specialty" validate:"required" example:"Psicologa"`
	OnDoctor  string `json:"on_doctor" bson:"on_doctor" validate:"required" example:"Sobre el doctor"`
}
