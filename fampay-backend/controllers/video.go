package controllers

import (
	"net/http"

	"github.com/Pratham-Mishra04/fampay/fampay-backend/helpers"
	"github.com/gofiber/fiber/v2"
)

const (
	QUERY = "cricket"
)

func FetchLatestVideos(c *fiber.Ctx) error {
	// Make API requests using youtubeService...
	// For example, retrieve the user's channel information
	searchResponse, err := helpers.Service.Search.List([]string{"snippet"}).
		Q(QUERY).
		Type("video").
		MaxResults(10). // Adjust the number of videos you want to retrieve
		Order("date").  // Order by date to get the latest videos
		Do()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	// Display the retrieved data on the dashboard
	return c.JSON(fiber.Map{
		"data": searchResponse,
	})
}
