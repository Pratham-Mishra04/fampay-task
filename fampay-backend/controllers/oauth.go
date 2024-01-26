package controllers

import (
	"context"
	"net/http"

	"github.com/Pratham-Mishra04/fampay/fampay-backend/helpers"
	"github.com/Pratham-Mishra04/fampay/fampay-backend/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

func LogIn(c *fiber.Ctx) error {
	// Redirect to Google's OAuth consent screen
	url := helpers.OAuthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(url, http.StatusSeeOther)
}

func OAuthCallback(c *fiber.Ctx) error {
	// Handle the callback from Google after the user grants/denies permission
	code := c.Query("code")
	token, err := helpers.OAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	// Store the token securely (e.g., in a database)
	utils.SaveToken(token)

	// Redirect to the dashboard or a success page
	return c.Redirect("/dashboard", http.StatusSeeOther)
}
