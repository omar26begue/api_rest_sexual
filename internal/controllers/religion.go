package controllers

import (
	"api_rest_sexual/internal/interfaces"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
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
func CreateReligion(c *fiber.Ctx) error {
	return nil
}