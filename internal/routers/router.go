package routers

import (
	_ "api_rest_sexual/docs"
	"api_rest_sexual/internal/controllers"
	"api_rest_sexual/internal/middlewares"
	websocket2 "api_rest_sexual/internal/websocket"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Routers(app *fiber.App) {
	// swagger
	app.Get("/docs/*", swagger.Handler)

	// monitor
	app.Get("/monitor", monitor.New())

	wbRouter := app.Group("/ws", logger.New())
	{
		websocketRouter(wbRouter)
	}

	apiRouter := app.Group("/api", logger.New())
	{
		authRouter := apiRouter.Group("/auth")
		{
			routerAuth(authRouter)
		}

		v1NoProte := apiRouter.Group("/")
		{
			routerReligion(v1NoProte)
			routerSexualOrientation(v1NoProte)
		}

		images := apiRouter.Group("/images")
		{
			routerImages(images)
		}

		v1Proteted := apiRouter.Group("/v1", middlewares.Protected())
		{
			routerArticles(v1Proteted)
		}
	}
}

func websocketRouter(router fiber.Router) {
	router.Use("/chats", websocket2.ChatsWS)
}

func routerAuth(router fiber.Router) {
	router.Post("/register", controllers.RegisterApp)
	router.Post("/login", controllers.LoginEmail)
	router.Post("/facebook", controllers.LoginFacebook)
}

func routerReligion(router fiber.Router) {
	router.Get("religions", controllers.GetAllReligion)
	router.Post("religions", /*middlewares.Protected(),*/ controllers.CreateReligion)
}

func routerSexualOrientation(router fiber.Router) {
	router.Get("sexual", controllers.GetAllSexualOrientation)
	router.Post("sexual", /*middlewares.Protected(),*/ controllers.CreateSexualOrientation)
}

func routerImages(router fiber.Router) {
	router.Get("articles/subtitle/:id/photo", controllers.GetImageArticleSubtitle)
}

func routerArticles(router fiber.Router) {
	router.Get("/articles", controllers.GetArticle)
	router.Get("/articles", controllers.GetImageArticleSubtitle)
	router.Post("/articles", controllers.CreateArticle)
	router.Post("/articles/subtitle/:id/upload", controllers.UploadImageSubtitle)
}
