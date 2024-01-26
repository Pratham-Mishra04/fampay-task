package controllers

import (
	"time"

	"github.com/Pratham-Mishra04/fampay/fampay-backend/config"
	"github.com/Pratham-Mishra04/fampay/fampay-backend/helpers"
	"github.com/Pratham-Mishra04/fampay/fampay-backend/initializers"
	"github.com/Pratham-Mishra04/fampay/fampay-backend/models"
	"github.com/gofiber/fiber/v2"
)

const (
	time_layout = "2006-01-02T15:04:05Z"
)

func FetchLatestVideos() {
	searchResponse, err := helpers.Service.Search.List([]string{"snippet"}).
		Q(config.ServiceQuery).
		Type("video").
		MaxResults(10).
		Order("date").
		Do()

	if err != nil {
		go config.Logger.Errorw("Error while fetching from the service", "Message", err.Error(), "Path", "FetchLatestVideos", "Error", err)

		return
	}

	tx := initializers.DB.Begin()

	defer func() {
		if tx.Error != nil {
			tx.Rollback()
			go config.Logger.Errorw("Transaction rolled back due to error", "Message", tx.Error.Error(), "Path", "FetchLatestVideos", "Error", tx.Error)
		}
	}()

	result := tx.Exec("DELETE FROM videos").Unscoped()
	if result.Error != nil {
		go config.Logger.Errorw("Error while flushing videos", "Message", result.Error.Error(), "Path", "FetchLatestVideos", "Error", result.Error)
	}

	for _, item := range searchResponse.Items {
		video := models.Video{
			Title:        item.Snippet.Title,
			ChannelID:    item.Snippet.ChannelId,
			ChannelTitle: item.Snippet.ChannelTitle,
			Description:  item.Snippet.Description,
			Thumbnail:    item.Snippet.Thumbnails.Default.Url,
		}

		// Parse the time string
		parsedTime, err := time.Parse(time_layout, item.Snippet.PublishedAt)
		if err != nil {
			video.UploadedAt = time.Now()
		} else {
			video.UploadedAt = parsedTime
		}

		result := tx.Create(&video)
		if result.Error != nil {
			go config.Logger.Errorw("Error while adding a video", "Message", result.Error.Error(), "Path", "FetchLatestVideos", "Error", result.Error)
		}
	}

	if err := tx.Commit().Error; err != nil {
		go config.Logger.Errorw("Error while committing a transaction", "Message", err.Error(), "Path", "FetchLatestVideos", "Error", err.Error)
	}
}

func GetVideos(c *fiber.Ctx) error {
	var videos []models.Video
	if err := initializers.DB.
		Order("uploaded_at DESC").
		Find(&videos).Error; err != nil {
		go config.Logger.Errorw("Error while fetching videos", "Message", err.Error(), "Path", "FetchLatestVideos", "Error", err.Error)
		return &fiber.Error{Code: 500, Message: config.SERVER_ERROR}
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Videos fetched",
		"videos":  videos,
	})
}
