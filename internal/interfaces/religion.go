package interfaces

import (
	"api_rest_sexual/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var colletionReligion *mongo.Collection

func InitInterfaceReligion(c *mongo.Database) {
	colletionReligion = c.Collection("religions")
}

func GetAllReligions() ([]models.Religion) {
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
