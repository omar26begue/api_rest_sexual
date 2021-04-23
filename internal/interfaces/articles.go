package interfaces

import (
	"api_rest_sexual/internal/models"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var colletionArticles *mongo.Collection

func InitInterfaceArticles(c *mongo.Database) {
	colletionArticles = c.Collection("articles")
}

func GetAllArticles() []models.Articles  {
	var articles []models.Articles
	articles = []models.Articles{}

	cursor, err := colletionArticles.Find(context.TODO(), bson.M{})
	if err != nil {
		return []models.Articles{}
	}

	for cursor.Next(context.TODO()) {
		var article models.Articles
		cursor.Decode(&article)
		articles = append(articles, article)
	}

	return articles
}

func GetArticles(identifier string) (models.Articles, error) {
	articles := models.Articles{}

	err := colletionArticles.FindOne(context.TODO(), bson.M{"identifier": identifier}).Decode(&articles)
	if err != nil {
		return models.Articles{}, errors.New("No existe el elemento solicitado.")
	}

	return articles, nil
}

func CreateArticles(article models.Articles) (models.Articles, error) {
	var validate = validator.New()
	err := validate.Struct(article)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return models.Articles{}, errors.New(err.Tag())
		}
	}

	_, err = colletionArticles.InsertOne(context.TODO(), article)
	if err != nil {
		return models.Articles{}, errors.New("Ha ocurrido un error.")
	}

	return article, nil
}

func UpdateArticlePhotoSubtitle(identifier string, photo string) error {
	_, err := colletionArticles.UpdateOne(context.TODO(), bson.M{"identifier": identifier}, bson.D{{"$set",
		bson.D{{"image_sub_title", photo}}}})
	if err != nil {
		return errors.New("No se pudo actualizar la información.")
	}

	return nil
}
