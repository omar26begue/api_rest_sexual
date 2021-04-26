package models

import "time"

type Articles struct {
	Identifier    string    `json:"identifier" bson:"identifier" validate:"required,uuid"`
	DateArticle   time.Time `json:"date_article" bson:"date_article"`
	DoctorArticle string    `json:"doctor_article,omitempty" bson:"doctor_article" validate:"required"`
	Category      string    `json:"category" bson:"category" validate:"required"`
	Title         string    `json:"title" bson:"title" validate:"required"`
	SubTitle      string    `json:"sub_title" bson:"subtitle" validate:"required"`
	ImageSubTitle string    `json:"image_sub_title,omitempty" bson:"image_sub_title"`
	ImageArticle  string    `json:"image_article,omitempty" bson:"image_article"`
	TextArticle   string    `json:"text_article" bson:"text_article" validate:"required"`
}

type ArticleRequest struct {
	DoctorArticle string `json:"doctor_article" bson:"doctor_article" validate:"required" example:"Omar Isalgue Begue"`
	Category      string `json:"category" bson:"category" validate:"required" example:"man"`
	SubTitle      string `json:"sub_title" bson:"subtitle" validate:"required" example:"Subtitle"`
	Title         string `json:"title" bson:"title" validate:"required" example:"Title"`
	TextArticle   string `json:"text_article" bson:"text_article" validate:"required"`
}
