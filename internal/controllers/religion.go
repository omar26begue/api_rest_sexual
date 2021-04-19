package controllers

import (
	"api_rest_sexual/internal/interfaces"
	"api_rest_sexual/internal/models"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetAllReligion godoc
// @Summary Returns all religions
// @Description Returns all religions.
// @Tags Religions
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Religion
// @Failure 400 {object} models.HTTPResponse
// @Failure default {object} models.HTTPResponse
// @Router /religions [get]
func GetAllReligion(c *fiber.Ctx) error {
	religions := interfaces.GetAllReligions()

	return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(religions)
}

// CreateReligion godoc
// @Summary Returns all religions
// @Description Returns all religions.
// @Tags Religions
// @Accept  json
// @Produce  json
// @Param religion body models.ReligionRequest true "Religion"
// @Success 200 {object} models.Religion
// @Failure 400 {object} models.HTTPResponse
// @Failure default {object} models.HTTPResponse
// @Security ApiKeyAuth
// @Router /religions [post]
func CreateReligion(c *fiber.Ctx) error {
	religionRequest := new(models.ReligionRequest)
	if err := c.BodyParser(&religionRequest); err != nil {
		return json.NewEncoder(c.Status(fiber.StatusServiceUnavailable).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{Code: fiber.StatusServiceUnavailable, Message: "error"})
	}

	religion := new(models.Religion)
	religion.Identifier = uuid.New().String()
	religion.Name = religionRequest.Name

	_, err := interfaces.CreateReligions(*religion)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return json.NewEncoder(c.Status(fiber.StatusCreated).Type("json", "utf-8").Response().BodyWriter()).Encode(religion)
}

// UpdateReligion godoc
// @Summary Returns all religions
// @Description Returns all religions.
// @Tags Religions
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Religion
// @Failure 400 {object} models.HTTPResponse
// @Failure default {object} models.HTTPResponse
// @Security ApiKeyAuth
// @Router /religions/{id} [patch]
func UpdateReligion(c *fiber.Ctx) error {
	return nil
}

// DeleteReligion godoc
// @Summary Returns all religions
// @Description Returns all religions.
// @Tags Religions
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Religion
// @Failure 400 {object} models.HTTPResponse
// @Failure default {object} models.HTTPResponse
// @Security ApiKeyAuth
// @Router /religions/{id} [delete]
func DeleteReligion(c *fiber.Ctx) error {
	return nil
}
