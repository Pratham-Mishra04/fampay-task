package utils

import (
	"strconv"

	"github.com/Pratham-Mishra04/fampay/fampay-backend/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginator(c *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageStr := c.Query("page", "1")
		limitStr := c.Query("limit", "10")

		page, err := strconv.Atoi(pageStr)
		if err != nil {
			config.Logger.Warnw("Failed to Paginate due to integer conversion.", "Error", err)
			return db
		}

		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			config.Logger.Warnw("Failed to Paginate due to integer conversion.", "Error", err)
			return db
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}
