package routers

import (
	"github.com/Pratham-Mishra04/fampay/fampay-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app *fiber.App) {
	app.Post("/login", controllers.LogIn)
	app.Post("/oauth2callback", controllers.OAuthCallback)
}
