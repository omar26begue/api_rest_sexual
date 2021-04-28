package routers

import (
	_ "api_rest_sexual/docs"
	"api_rest_sexual/internal/controllers"
	"api_rest_sexual/internal/middlewares"
	"api_rest_sexual/internal/models"
	"api_rest_sexual/internal/websocket"
	"encoding/json"
	"fmt"
	"github.com/antoniodipinto/ikisocket"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"time"
)

type MessageObject struct {
	Data string `json:"data"`
	From string `json:"from"`
	To   string `json:"to"`
}

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
			routerUsers(v1Proteted)
		}
	}
}

func websocketRouter(router fiber.Router) {
	router.Use("/online", ikisocket.New(func(kws *ikisocket.Websocket) {
		websocket.OnlineUsers(kws)
	}))

	router.Use("/chats", ikisocket.New(func(kws *ikisocket.Websocket) {
		kws.SetAttribute("type", "chats")
		fmt.Println(kws.UUID)
	}))

	ikisocket.On(ikisocket.EventConnect, func(payload *ikisocket.EventPayload) {
		fmt.Println("connected online" + payload.Kws.GetStringAttribute("type"))
	})

	ikisocket.On(ikisocket.EventDisconnect, func(payload *ikisocket.EventPayload) {
		fmt.Println("closed online")
		fmt.Println(payload.SocketUUID)
	})

	ikisocket.On(ikisocket.EventMessage, func(payload *ikisocket.EventPayload) {
		switch payload.Kws.GetStringAttribute("type") {
		case "online":
			online := models.Online{}
			err := json.Unmarshal(payload.Data, &online)
			if err != nil {
				fmt.Println(err)
				return
			} else {
				online.Uuid = payload.SocketUUID
				online.Time = time.Now()

				websocket.ImplementOnlineUsers(online)
			}

		case "chats":
			message := MessageObject{}
			err := json.Unmarshal(payload.Data, &message)
			if err != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Println(message)
			}
		}
	})
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

func routerUsers(router fiber.Router) {
	router.Patch("/users/:id", controllers.UpdateUsers)
	router.Patch("/users/image", controllers.UpdateUsersImage)
	router.Get("/users/info", controllers.InfoUsers)
}
