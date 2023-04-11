package persistence

import (
	"context"
	"english_exam_go/infrastructure/data/entities"
	"english_exam_go/infrastructure/data/repositories"
)

type IExamInviteRepository interface {
	AddUserToExam(ctx context.Context, entities []*entities.ExamTasker) error
}

type ExamInviteRepository struct {
}

func (e ExamInviteRepository) AddUserToExam(ctx context.Context, entities []*entities.ExamTasker) error {
	//TODO implement me
	db := repositories.GetConn()
	err := db.Create(&entities)
	return err.Error
}

func CreateExamInviteRepository() IExamInviteRepository {
	return &ExamInviteRepository{}
}
