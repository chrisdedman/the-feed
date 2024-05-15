package models

import "gorm.io/gorm"

type Feed struct {
	gorm.Model
	URL    string  `gorm:"size:255;not null;" json:"url"`
	UserID float64 `gorm:"not null;" json:"user_id"`
}

type FeedInfo struct {
	FeedID uint   `json:"feed_id"`
	URL    string `json:"url"`
}
