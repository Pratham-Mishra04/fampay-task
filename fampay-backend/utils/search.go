package utils

import (
	"time"

	"github.com/Pratham-Mishra04/fampay/fampay-backend/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func genericSearch(db *gorm.DB, field, value string) *gorm.DB {
	if value != "" {
		return db.Where(field+" ILIKE ?", "%"+value+"%")
	}
	return db
}

func timestampSearch(db *gorm.DB, start, end string) *gorm.DB {
	if start != "" && end != "" {
		startTime, err := time.Parse(time.RFC3339, start)
		if err != nil {
			config.Logger.Warnw("Error parsing start timestamp", "Error", err)
			return db
		}

		endTime, err := time.Parse(time.RFC3339, end)
		if err != nil {
			config.Logger.Warnw("Error parsing end timestamp", "Error", err)
			return db
		}

		return db.Where("uploaded_at BETWEEN ? AND ?", startTime, endTime)
	} else if start != "" {
		startTime, err := time.Parse(time.RFC3339, start)
		if err != nil {
			config.Logger.Warnw("Error parsing start timestamp", "Error", err)
			return db
		}

		return db.Where("uploaded_at >= ?", startTime)
	} else if end != "" {
		endTime, err := time.Parse(time.RFC3339, end)
		if err != nil {
			config.Logger.Warnw("Error parsing end timestamp", "Error", err)
			return db
		}

		return db.Where("uploaded_at <= ?", endTime)
	}
	return db
}

func Search(c *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		fields := []string{"title", "channel_title"}

		for _, field := range fields {
			value := c.Query(field, "")
			db = genericSearch(db, field, value)
		}

		startTime := c.Query("start", "")
		endTime := c.Query("end", "")
		db = timestampSearch(db, startTime, endTime)

		return db
	}
}
