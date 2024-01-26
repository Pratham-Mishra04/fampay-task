package models

import (
	"time"
)

type Video struct {
	ID           int       `gorm:"autoIncrement;primaryKey" json:"id"`
	Title        string    `json:"title"`
	ChannelID    string    `json:"channelID"`
	ChannelTitle string    `json:"channelTitle"`
	Description  string    `json:"description"`
	Thumbnail    string    `json:"thumbnail"`
	UploadedAt   time.Time `gorm:"index:idx_uploaded_at" json:"uploadedAt"`
}
