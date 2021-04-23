package controllers

import (
	"api_rest_sexual/internal/interfaces"
	"api_rest_sexual/internal/models"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

// GetArticle godoc
// @Summary Returns a list of items
// @Description Returns a list of items.
// @Tags Articles
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Articles
// @Failure 400 {object} models.HTTPResponse
// @Failure default {object} models.HTTPResponse
// @Security ApiKeyAuth
// @Router /v1/articles [get]
func GetArticle(c *fiber.Ctx) error {
	articles := interfaces.GetAllArticles()

	return json.NewEncoder(c.Status(fiber.StatusOK).Type("json", "utf-8").Response().BodyWriter()).Encode(articles)
}

// GetImageArticleSubtitle godoc
// @Summary Download photo server
// @Description Download photo server
// @Tags Images
// @Accept  json
// @Produce  json
// @Param   id      path   string     true  "Article"
// @Success 200 {object} models.Image
// @Failure 400 {object} models.HTTPResponse
// @Failure default {object} models.HTTPResponse
// @Router /images/articles/subtitle/{id}/photo [get]
func GetImageArticleSubtitle(c *fiber.Ctx) error {
	identifier := c.Params("id")

	dataArticle, err := interfaces.GetArticles(identifier)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    fiber.StatusBadRequest,
			Message: "No se encuentra la información.",
		})
	}

	if len(dataArticle.ImageSubTitle) == 0 {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Foto de la aplicación.",
		})
	}

	err = c.SendFile(dataArticle.ImageSubTitle, true)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return nil
}

// CreateArticle godoc
// @Summary Create article
// @Description Create article.
// @Tags Articles
// @Accept  json
// @Produce  json
// @Param religion body models.ArticleRequest true "Article"
// @Success 201 {object} models.Articles
// @Failure 400 {object} models.HTTPResponse
// @Failure default {object} models.HTTPResponse
// @Security ApiKeyAuth
// @Router /v1/articles [post]
func CreateArticle(c *fiber.Ctx) error {
	articleRequest := new(models.ArticleRequest)
	if err := c.BodyParser(&articleRequest); err != nil {
		return json.NewEncoder(c.Status(fiber.StatusServiceUnavailable).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code: fiber.StatusServiceUnavailable,
			Message: "error",
		})
	}

	article := new(models.Articles)
	article.Identifier = uuid.New().String()
	article.Title = articleRequest.Title
	article.SubTitle = articleRequest.SubTitle
	article.Category = articleRequest.Category
	article.DateArticle = time.Now()

	_, err := interfaces.CreateArticles(*article)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return json.NewEncoder(c.Status(fiber.StatusCreated).Type("json", "utf-8").Response().BodyWriter()).Encode(article)
}

// UploadImageSubtitle godoc
// @Summary Upload an image to the server
// @Description Upload an image to the server. The supported extensions are [png, jpeg]
// @Tags Articles
// @Accept  json
// @Produce  json
// @Param   id      path   string     true  "Article"
// @Param   subtitle formData file true  "Image Subtitle"
// @Success 200 {object} models.Image
// @Failure 400 {object} models.HTTPResponse
// @Failure default {object} models.HTTPResponse
// @Security ApiKeyAuth
// @Router /v1/articles/subtitle/{id}/upload [post]
func UploadImageSubtitle(c *fiber.Ctx) error {
	identifier := c.Params("id")
	file, err := c.FormFile("subtitle")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	_, err = interfaces.GetArticles(identifier)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png" {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Error:   false,
			Code:    fiber.StatusBadRequest,
			Message: "Tipo de archivo no sportado.",
		})
	}

	// extension file
	ext := "jpeg"
	if file.Header.Get("Content-Type") == "image/png" {
		ext = "png"
	}

	image := uuid.New().String() + "." + ext
	path := "public/subtitle/" + image
	err = c.SaveFile(file, path)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err = interfaces.UpdateArticlePhotoSubtitle(identifier, path)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return json.NewEncoder(c.Status(fiber.StatusOK).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
		Code:    fiber.StatusOK,
		Message: image,
	})
}
