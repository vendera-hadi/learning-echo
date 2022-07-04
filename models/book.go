package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
	UserID   int    `json:"user_id"`
	// belongs to user
	User User `gorm:"foreignKey:UserID;references:ID"`
}
