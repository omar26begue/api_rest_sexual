package main

import (
	"api_rest_sexual/internal/database"
	"api_rest_sexual/internal/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

// @title API Rest Sexualidad - Knowdo Studio.
// @version 1.0.0
// @description API Rest Sexualidad.
// @termsOfService http://swagger.io/terms/

// @contact.name Omar Isalgué Begue
// @contact.email omar26begue@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// leyendo las variables de entorno
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	database.GetConnectionMongoDB()

	app := fiber.New()
	app.Use(cors.New())
	file, err := os.OpenFile("./logs/router-" + time.Now().Format("2021-01-06") +".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	app.Use(logger.New(logger.Config{
		Output: file,
	}))

	routers.Routers(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":" + viper.GetString("PORT"))
}
