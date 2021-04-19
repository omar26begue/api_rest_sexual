package controllers

import (
	"api_rest_sexual/internal/helpers"
	"api_rest_sexual/internal/interfaces"
	"api_rest_sexual/internal/models"
	"encoding/json"
	"github.com/form3tech-oss/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"time"
)

// RegisterApp godoc
// @Summary Registration via email
// @Description Registration via email.
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param register body models.UsersRequest true "Register"
// @Success 201 {object} models.AuthResponse
// @Failure 400 {object} models.HTTPResponse
// @Failure default {object} models.HTTPResponse
// @Router /auth/register [post]
func RegisterApp(c *fiber.Ctx) error {
	usersRequest := new(models.UsersRequest)
	if err := c.BodyParser(&usersRequest); err != nil {
		return json.NewEncoder(c.Status(fiber.StatusServiceUnavailable).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{Code: fiber.StatusServiceUnavailable, Message: "error"})
	}

	password, err := helpers.HashPassword(usersRequest.Password)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code: fiber.StatusBadRequest,
			Message: "Ha ocurrido un error inesperado en el sistema.",
		})
	}

	users := new(models.Users)
	users.Identifier = uuid.New().String()
	users.Name = usersRequest.Name
	users.Email = usersRequest.Email
	users.Password = password
	users.Age = usersRequest.Age
	users.Sex = usersRequest.Sex
	users.IdReligion = usersRequest.IdReligion
	users.SexualOrientation = usersRequest.SexualOrientation
	users.Active = false

	_, err = interfaces.CreateUsers(*users)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return json.NewEncoder(c.Status(fiber.StatusCreated).Type("json", "utf-8").Response().BodyWriter()).Encode(users)
}

// LoginEmail godoc
// @Summary Login email
// @Description Login email.
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param register body models.Auth true "Register"
// @Success 200 {object} models.AuthResponse
// @Failure 400 {object} models.HTTPResponse
// @Failure default {object} models.HTTPResponse
// @Router /auth/login [post]
func LoginEmail(c *fiber.Ctx) error {
	login := new(models.Auth)

	// variable del formulario
	if err := c.BodyParser(&login); err != nil {
		return json.NewEncoder(c.Status(fiber.StatusServiceUnavailable).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{Code: fiber.StatusServiceUnavailable, Message: err.Error()})
	}

	// validacion
	var validate = validator.New()
	err := validate.Struct(login)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
				Code:    fiber.StatusBadRequest,
				Message: err.Tag(),
			})
		}
	}

	resultUser, err := interfaces.GetUsersEmail(login.Email)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// comprobar la contraseña
	result := helpers.CheckPasswordHash(login.Password, resultUser.Password)
	if result == false {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Su contraseña no coincide.",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["identifier"] = resultUser.Identifier
	claims["type_login"] = "email"
	claims["email"] = resultUser.Email
	claims["expire"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(viper.GetString("SECRET_JWT")))
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	interfaces.UpdateToken(login.Email, tokenString)

	return json.NewEncoder(c.Status(fiber.StatusOK).Type("json", "utf-8").Response().BodyWriter()).Encode(models.AuthResponse{
		Token:      "Bearer " + tokenString,
		Identifier: resultUser.Identifier,
	})
}
