package models

import "gorm.io/gorm"

type Feed struct {
	gorm.Model
	URL    string  `gorm:"size:255;not null;" json:"url"`
	UserID float64 `gorm:"not null;" json:"user_id"`
}

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type LoginInput struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type UpdateAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
