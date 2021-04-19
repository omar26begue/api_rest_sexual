package routers

import (
	_ "api_rest_sexual/docs"
	"api_rest_sexual/internal/controllers"
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

	apiRouter := app.Group("/api", logger.New())
	{
		authRouter := apiRouter.Group("/auth")
		{
			routerAuth(authRouter)
		}

		v1NoProte := apiRouter.Group("/")
		{
			routerReligion(v1NoProte)
		}
	}
}

func routerAuth(router fiber.Router) {
	router.Post("/register", controllers.RegisterApp)
}

func routerReligion(router fiber.Router) {
	router.Get("religions", controllers.GetAllReligion)
}

