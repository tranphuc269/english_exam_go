package services

import (
	"context"
	dtos "english_exam_go/domain/dtos/exam"
	"english_exam_go/infrastructure/data/repositories/persistence"
)

type IExamService interface {
	CreateExam(ctx context.Context, exam *dtos.CreateExamRequest) error
}

type ExamServiceImpl struct {
	er persistence.IExamRepository
}

func (es ExamServiceImpl) CreateExam(ctx context.Context, exam *dtos.CreateExamRequest) error {
	//TODO implement me
	examEnt := exam.CreateExamEntity()
	err := es.er.CreateExam(ctx, &examEnt)
	if err != nil {
		return err
	}
	return nil
}

func CreateExamService(er persistence.IExamRepository) IExamService {
	return &ExamServiceImpl{
		er: er,
	}
}
