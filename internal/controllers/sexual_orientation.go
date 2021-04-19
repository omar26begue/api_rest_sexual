package controllers

import (
	"api_rest_sexual/internal/interfaces"
	"api_rest_sexual/internal/models"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetAllSexualOrientation godoc
// @Summary Returns all religions
// @Description Returns all religions.
// @Tags Sexual Orientation
// @Accept  json
// @Produce  json
// @Success 200 {array} models.SexualOrientation
// @Failure 400 {object} models.HTTPResponse
// @Failure default {object} models.HTTPResponse
// @Router /sexual [get]
func GetAllSexualOrientation(c *fiber.Ctx) error {
	sexuals := interfaces.GetAllSexualOrientation()

	return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(sexuals)
}

// CreateSexualOrientation godoc
// @Summary Returns all religions
// @Description Returns all religions.
// @Tags Sexual Orientation
// @Accept  json
// @Produce  json
// @Param sexual body models.SexualOrientationRequest true "Sexual Orientation"
// @Success 200 {object} models.SexualOrientation
// @Failure 400 {object} models.HTTPResponse
// @Failure default {object} models.HTTPResponse
// @Security ApiKeyAuth
// @Router /sexual [post]
func CreateSexualOrientation(c *fiber.Ctx) error {
	sexualRequest := new(models.SexualOrientationRequest)
	if err := c.BodyParser(&sexualRequest); err != nil {
		return json.NewEncoder(c.Status(fiber.StatusServiceUnavailable).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{Code: fiber.StatusServiceUnavailable, Message: "error"})
	}

	sexual := new(models.SexualOrientation)
	sexual.Identifier = uuid.New().String()
	sexual.Name = sexualRequest.Name

	_, err := interfaces.CreateSexualOrientation(*sexual)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return json.NewEncoder(c.Status(fiber.StatusCreated).Type("json", "utf-8").Response().BodyWriter()).Encode(sexual)
}

// UpdateSexualOrientation godoc
// @Summary Returns all religions
// @Description Returns all religions.
// @Tags Sexual Orientation
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Religion
// @Failure 400 {object} models.HTTPResponse
// @Failure default {object} models.HTTPResponse
// @Security ApiKeyAuth
// @Router /sexual/{id} [patch]
func UpdateSexualOrientation(c *fiber.Ctx) error {
	return nil
}

// DeleteSexualOrientation godoc
// @Summary Returns all religions
// @Description Returns all religions.
// @Tags Sexual Orientation
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Religion
// @Failure 400 {object} models.HTTPResponse
// @Failure default {object} models.HTTPResponse
// @Security ApiKeyAuth
// @Router /sexual/{id} [delete]
func DeleteSexualOrientation(c *fiber.Ctx) error {
	return nil
}
