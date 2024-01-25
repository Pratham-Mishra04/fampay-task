package main

import (
	"github.com/Pratham-Mishra04/fampay/fampay-backend/config"
	"github.com/Pratham-Mishra04/fampay/fampay-backend/helpers"
	"github.com/Pratham-Mishra04/fampay/fampay-backend/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.ConnectToCache()
	initializers.AutoMigrate()

	config.AddLogger()
	helpers.InitializeYoutubeService()
}

func main() {
	defer config.LoggerCleanUp()
	app := fiber.New(fiber.Config{
		ErrorHandler: fiber.DefaultErrorHandler,
	})

	app.Use(helmet.New())
	app.Use(logger.New())
	app.Use(config.CORS())

	app.Listen(":" + initializers.CONFIG.PORT)
}
