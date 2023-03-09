package book

import (
	"english_exam_go/infrastructure/data/entities"
	"time"
)

type RequestBook struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ResponseBook struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateResponseBook(ent entities.Book) *ResponseBook {
	return &ResponseBook{
		ID:        ent.ID,
		Title:     ent.Title,
		Content:   ent.Content,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.UpdatedAt,
	}
}

func NewBookEnt(request *RequestBook) (entities.Book, error) {
	// validate
	return entities.Book{
		Title:   request.Title,
		Content: request.Content,
	}, nil
}
