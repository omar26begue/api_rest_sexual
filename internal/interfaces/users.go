package interfaces

import (
	"api_rest_sexual/internal/models"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var colletionUser *mongo.Collection

func InitInterfaceUsers(c *mongo.Database) {
	colletionUser = c.Collection("users")
}

func GetAllUsers() []models.Users {
	var users []models.Users
	users = []models.Users{}

	cursor, err := colletionUser.Find(context.TODO(), bson.M{})
	if err != nil {
		return []models.Users{}
	}

	for cursor.Next(context.TODO()) {
		var user models.Users
		cursor.Decode(&user)
		users = append(users, user)
	}

	return users
}

func GetUsersEmail(email string) (models.Users, error) {
	users := models.Users{}

	err := colletionUser.FindOne(context.TODO(), bson.M{"email": email}).Decode(&users)
	if err != nil {
		return models.Users{}, errors.New("No existe el elemento solicitado.")
	}

	return users, nil
}

func GetUsersIdentifier(identifier string) (models.Users, error) {
	users := models.Users{}

	err := colletionUser.FindOne(context.TODO(), bson.M{"identifier": identifier}).Decode(&users)
	if err != nil {
		return models.Users{}, errors.New("No existe el elemento solicitado.")
	}

	return users, nil
}

func CreateUsers(users models.Users) (models.Users, error) {
	_, err := GetUsersEmail(users.Email)
	if err == nil {
		return models.Users{}, errors.New("Ya existe un usuario autenticado con ese correo electrónico " + users.Email + " en el sistema.")
	}

	var validate = validator.New()
	err = validate.Struct(users)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return models.Users{}, errors.New(err.StructNamespace())
		}
	}

	_, err = colletionUser.InsertOne(context.TODO(), users)
	if err != nil {
		return models.Users{}, errors.New("Ha ocurrido un error.")
	}

	return users, nil
}

func UpdateUsers(identifier string, users models.UsersUpdateRequest) error {
	_, err := colletionUser.UpdateOne(context.TODO(), bson.M{"identifier": identifier}, bson.D{{"$set", bson.D{{"name", users.Name},
		{"email", users.Email}, {"age", users.Age}, {"sex", users.Sex}, {"id_religion", users.IdReligion},
		{"sexual_orientation", users.SexualOrientation}}}})
	if err != nil {
		return errors.New("No se pudo actualizar el elemento solicitado.")
	}

	return nil
}

func UpdateToken(email string, token string) error {
	user, err := GetUsersEmail(email)
	if err != nil {
		return err
	}

	user.Token = token

	_, err = colletionUser.UpdateOne(context.TODO(), bson.M{"email": email}, bson.D{{"$set", bson.D{{"token", token}}}})
	if err != nil {
		return errors.New("No existe el elemento solicitado.")
	}

	return nil
}

func UpdateImage(identifier string, image string) error {
	_, err := GetUsersIdentifier(identifier)
	if err != nil {
		return err
	}

	_, err = colletionUser.UpdateOne(context.TODO(), bson.M{"identifier": identifier}, bson.D{{"$set", bson.D{{"image", image}}}})
	if err != nil {
		return errors.New("No existe el elemento solicitado.")
	}

	return nil
}
