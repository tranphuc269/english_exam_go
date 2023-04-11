package persistence

import (
	"context"
	"english_exam_go/infrastructure/data/entities"
	"english_exam_go/infrastructure/data/repositories"
)

type IExamInviteRepository interface {
	AddUserToExam(ctx context.Context, entities []*entities.ExamTasker) error
	RemoveUserToExam(ctx context.Context, ExamID int, UserId int) error
}

type ExamInviteRepository struct {
}

func (e ExamInviteRepository) RemoveUserToExam(ctx context.Context, ExamID int, UserId int) error {
	//TODO implement me
	var examTaskerEnt *entities.ExamTasker
	db := repositories.GetConn()
	err := db.Where("exam_id = ? AND user_id = ?", ExamID, UserId).Delete(&examTaskerEnt)
	return err.Error
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
