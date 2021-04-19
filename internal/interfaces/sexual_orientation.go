package interfaces

import (
	"api_rest_sexual/internal/models"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var colletionSexualOrientation *mongo.Collection

func InitInterfaceSexualOrientation(c *mongo.Database) {
	colletionSexualOrientation = c.Collection("sexual_orientation")
}

func GetAllSexualOrientation() []models.Religion {
	var religions []models.Religion
	religions = []models.Religion{}

	cursor, err := colletionSexualOrientation.Find(context.TODO(), bson.M{})
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

func GetSexualOrientation(name string) (models.Religion, error) {
	religion := models.Religion{}

	err := colletionSexualOrientation.FindOne(context.TODO(), bson.M{"name": name}).Decode(&religion)
	if err != nil {
		return models.Religion{}, errors.New("No existe el elemento solicitado.")
	}

	return religion, nil
}

func CreateSexualOrientation(sexual models.SexualOrientation) (models.SexualOrientation, error) {
	_, err := GetSexualOrientation(sexual.Name)
	if err == nil {
		return models.SexualOrientation{}, errors.New("Ya existe el elemento.")
	}

	var validate = validator.New()
	err = validate.Struct(sexual)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return models.SexualOrientation{}, errors.New(err.Tag())
		}
	}

	_, err = colletionSexualOrientation.InsertOne(context.TODO(), sexual)
	if err != nil {
		return models.SexualOrientation{}, errors.New("Ha ocurrido un error.")
	}

	return sexual, nil
}
