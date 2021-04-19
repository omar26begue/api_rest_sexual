package controllers

import "github.com/gofiber/fiber/v2"

// RegisterApp godoc
// @Summary Registration via email
// @Description Registration via email.
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param auth body models.Auth true "Auth"
// @Success 200 {object} models.AuthResponse
// @Failure 400 {object} models.HTTPResponse
// @Failure default {object} models.HTTPResponse
// @Router /auth/register [post]
func RegisterApp(c *fiber.Ctx) error {
	return nil
}