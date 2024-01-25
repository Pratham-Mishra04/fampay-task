package models

import (
	"time"

	"github.com/lib/pq"
)

type Video struct {
	ID          int            `gorm:"autoIncrement;primaryKey" json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Thumbnails  pq.StringArray `gorm:"type:text[]" json:"thumbnails"`
	UploadedAt  time.Time      `gorm:"index:idx_uploaded_at" json:"uploadedAt"`
}
