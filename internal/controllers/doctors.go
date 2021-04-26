package controllers

import "github.com/gofiber/fiber/v2"

// GetDoctors godoc
// @Summary It returns Doctor's information
// @Description It returns Doctor's information.
// @Tags Doctors
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Articles
// @Failure 400 {object} models.HTTPResponse
// @Failure default {object} models.HTTPResponse
// @Security ApiKeyAuth
// @Router /v1/doctors/{id} [get]
func GetDoctors(c *fiber.Ctx) error {
	return nil
}
