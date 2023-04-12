package persistence

import (
	"context"
	dtos "english_exam_go/domain/dtos/exam"
	"english_exam_go/infrastructure/data/entities"
	"english_exam_go/infrastructure/data/repositories"
)

type IExamResultRepository interface {
	SubmitExam(ctx context.Context, result *entities.ExamResult) (*dtos.ExamResultRes, error)
}

type ExamResultRepositoryImpl struct {
}

func (e ExamResultRepositoryImpl) SubmitExam(ctx context.Context, result *entities.ExamResult) (*dtos.ExamResultRes, error) {
	//TODO implement me
	db := repositories.GetConn()
	err := db.Create(result)
	return &dtos.ExamResultRes{
		ID:                  int(result.ID),
		ExamID:              result.ExamId,
		TotalScore:          result.TotalScore,
		ReadingScore:        result.ReadingScore,
		ListeningScore:      result.ListeningScore,
		NumCorrectReading:   result.NumCorrectReading,
		NumCorrectListening: result.NumCorrectListening,
		TabSwitchCount:      result.TabSwitchCount,
	}, err.Error
}

func CreateExamResultRepository() IExamResultRepository {
	return &ExamResultRepositoryImpl{}
}
