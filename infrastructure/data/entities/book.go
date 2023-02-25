package entities

import (
	"gorm.io/gorm"
)

type BookEnt struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}
