package services

import (
	"context"
	dtos "english_exam_go/domain/dtos/exam"
	"english_exam_go/infrastructure/data/repositories/persistence"
)

type IExamService interface {
	CreateExam(ctx context.Context, exam *dtos.CreateExamRequest) error
	GetAllExams(ctx context.Context) ([]*dtos.ExamListResponse, error)
	GetDetailExams(ctx context.Context, ID int) (*dtos.ExamDetailResponse, error)
}

type ExamServiceImpl struct {
	er persistence.IExamRepository
}

func (es ExamServiceImpl) GetDetailExams(ctx context.Context, ID int) (*dtos.ExamDetailResponse, error) {
	//TODO implement me
	examEnts, err := es.er.FindExamById(ctx, uint(ID))
	if err != nil {
		return nil, err
	}
	return dtos.ParseExamDetailRes(examEnts), nil
}

func (es ExamServiceImpl) GetAllExams(ctx context.Context) ([]*dtos.ExamListResponse, error) {
	//TODO implement me
	examEnts, err := es.er.FindAllExams(ctx)
	if err != nil {
		return nil, err
	}

	var responseExams []*dtos.ExamListResponse
	for _, e := range examEnts {
		responseExams = append(responseExams, dtos.CreateExamListRes(e))
	}
	return responseExams, nil
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
