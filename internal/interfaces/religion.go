package interfaces

import (
	"api_rest_sexual/internal/models"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var colletionReligion *mongo.Collection

func InitInterfaceReligion(c *mongo.Database) {
	colletionReligion = c.Collection("religions")
}

func GetAllReligions() []models.Religion {
	var religions []models.Religion
	religions = []models.Religion{}

	cursor, err := colletionReligion.Find(context.TODO(), bson.M{})
	if err != nil {
		return []models.Religion{}
	}

	for cursor.Next(context.TODO()) {
		var religion models.Religion
		cursor.Decode(&religion)
		religions = append(religions, religion)
	}

	return religions
}

func GetReligionsIdentifier(identifier string) (models.Religion, error) {
	religion := models.Religion{}

	err := colletionReligion.FindOne(context.TODO(), bson.M{"identifier": identifier}).Decode(&religion)
	if err != nil {
		return models.Religion{}, errors.New("No existe el elemento solicitado.")
	}

	return religion, nil
}

func GetReligionsName(name string) (models.Religion, error) {
	religion := models.Religion{}

	err := colletionReligion.FindOne(context.TODO(), bson.M{"name": name}).Decode(&religion)
	if err != nil {
		return models.Religion{}, errors.New("No existe el elemento solicitado.")
	}

	return religion, nil
}

func CreateReligions(religion models.Religion) (models.Religion, error) {
	_, err := GetReligionsName(religion.Name)
	if err == nil {
		return models.Religion{}, errors.New("Ya existe el elemento.")
	}

	var validate = validator.New()
	err = validate.Struct(religion)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return models.Religion{}, errors.New(err.Tag())
		}
	}

	_, err = colletionReligion.InsertOne(context.TODO(), religion)
	if err != nil {
		return models.Religion{}, errors.New("Ha ocurrido un error.")
	}

	return religion, nil
}
