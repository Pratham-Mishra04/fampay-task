package routers

import (
	"github.com/Pratham-Mishra04/fampay/fampay-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func VideoRouter(app *fiber.App) {
	app.Get("/", controllers.GetVideos)
}
