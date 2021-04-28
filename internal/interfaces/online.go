package interfaces

import (
	"api_rest_sexual/internal/models"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var colletionOnline *mongo.Collection

func InitInterfaceOnline(c *mongo.Database) {
	colletionOnline = c.Collection("online")
}

func GetOnline(identifier string) (models.Online, error) {
	online := models.Online{}

	err := colletionOnline.FindOne(context.TODO(), bson.M{"identifier": identifier}).Decode(&online)
	if err != nil {
		return models.Online{}, errors.New("No existe el elemento solicitado.")
	}

	return online, nil
}

func CreateOnline(online models.Online) (models.Online, error) {
	var validate = validator.New()
	err := validate.Struct(online)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return models.Online{}, errors.New(err.Tag())
		}
	}

	_, err = colletionOnline.InsertOne(context.TODO(), online)
	if err != nil {
		return models.Online{}, errors.New("Ha ocurrido un error.")
	}

	return online, nil
}
