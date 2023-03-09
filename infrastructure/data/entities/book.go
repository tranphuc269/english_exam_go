package entities

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}
