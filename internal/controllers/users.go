package controllers

import (
	"api_rest_sexual/internal/interfaces"
	"api_rest_sexual/internal/models"
	"encoding/json"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// InfoUsers godoc
// @Summary It returns the authenticated user's information
// @Description It returns the authenticated user's information.
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.UsersResponse
// @Failure 400 {object} models.HTTPResponse
// @Failure default {object} models.HTTPResponse
// @Security ApiKeyAuth
// @Router /users/info [get]
func InfoUsers(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	identifier := claims["identifier"].(string)

	dataUsers, err := interfaces.GetUsersIdentifier(identifier)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	dataReligion := new(models.Religion)
	if len(dataUsers.IdReligion) > 0 {
		*dataReligion, err = interfaces.GetReligionsIdentifier(dataUsers.IdReligion)
		if err != nil {
			return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
				Code:    fiber.StatusBadRequest,
				Message: err.Error(),
			})
		}
	}

	dataOrientation := new(models.SexualOrientation)
	if len(dataUsers.SexualOrientation) > 0 {
		*dataOrientation, err = interfaces.GetSexualOrientationIdentifier(dataUsers.SexualOrientation)
		if err != nil {
			return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
				Code:    fiber.StatusBadRequest,
				Message: err.Error(),
			})
		}
	}

	dataResponse := new(models.UsersResponse)
	dataResponse.Identifier = dataUsers.Identifier
	dataResponse.Name = dataUsers.Name
	dataResponse.Email = dataUsers.Email
	dataResponse.Age = dataUsers.Age
	dataResponse.Sex = dataUsers.Sex
	dataResponse.IdReligion = dataUsers.IdReligion
	dataResponse.NameReligion = dataReligion.Name
	dataResponse.SexualOrientation = dataUsers.SexualOrientation
	dataResponse.NameSexualOrientation = dataOrientation.Name
	dataResponse.Coins = dataUsers.Coins
	dataResponse.Image = dataUsers.Image

	return json.NewEncoder(c.Status(fiber.StatusOK).Type("json", "utf-8").Response().BodyWriter()).Encode(dataResponse)
}

// UpdateUsers godoc
// @Summary It upgrades the user's information
// @Description It upgrades the user's information.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param   id      path   string     true  "Users"
// @Param users body models.UsersUpdateRequest true "Register"
// @Success 200 {object} models.HTTPResponse
// @Failure 400 {object} models.HTTPResponse
// @Failure default {object} models.HTTPResponse
// @Security ApiKeyAuth
// @Router /users/{id} [patch]
func UpdateUsers(c *fiber.Ctx) error {
	identifier := c.Params("id")
	users := new(models.UsersUpdateRequest)
	if err := c.BodyParser(&users); err != nil {
		return json.NewEncoder(c.Status(fiber.StatusServiceUnavailable).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code: fiber.StatusServiceUnavailable,
			Message: "error",
		})
	}

	_, err := interfaces.GetUsersIdentifier(identifier)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code: fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err = interfaces.UpdateUsers(identifier, *users)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code: fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return json.NewEncoder(c.Status(fiber.StatusOK).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
		Code: fiber.StatusOK,
		Message: "Usuario actualizado",
	})
}

// UpdateUsersImage godoc
// @Summary Update the user's profile picture
// @Description Update the user's profile picture.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param register body models.UsersImageUpdate true "Register"
// @Success 200 {object} models.AuthResponse
// @Failure 400 {object} models.HTTPResponse
// @Failure default {object} models.HTTPResponse
// @Security ApiKeyAuth
// @Router /users/image [patch]
func UpdateUsersImage(c *fiber.Ctx) error {
	imageRequest := new(models.UsersImageUpdate)
	if err := c.BodyParser(&imageRequest); err != nil {
		return json.NewEncoder(c.Status(fiber.StatusServiceUnavailable).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{Code: fiber.StatusServiceUnavailable, Message: "error"})
	}

	err := interfaces.UpdateImage(imageRequest.Identifier, imageRequest.Image)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return json.NewEncoder(c.Status(fiber.StatusOK).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
		Code:    fiber.StatusOK,
		Message: "Imagen del usuario actualizada",
	})
}
