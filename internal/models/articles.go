package models

import "time"

type Articles struct {
	Identifier    string    `json:"identifier" bson:"identifier" validate:"required,uuid"`
	DateArticle   time.Time `json:"date_article" bson:"date_article"`
	Category      string    `json:"category" bson:"category" validate:"required"`
	Title         string    `json:"title" bson:"title" validate:"required"`
	SubTitle      string    `json:"sub_title" bson:"subtitle" validate:"required"`
	ImageSubTitle string    `json:"image_sub_title,omitempty" bson:"image_sub_title"`
	ImageArticle  string    `json:"image_article,omitempty" bson:"image_article"`
}

type ArticleRequest struct {
	Category string `json:"category" bson:"category" validate:"required" example:"Hombre"`
	SubTitle string `json:"sub_title" bson:"subtitle" validate:"required" example:"Subtitle"`
	Title    string `json:"title" bson:"title" validate:"required" example:"Title"`
}
