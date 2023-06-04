package services

import (
	"context"
	dtoExam "english_exam_go/domain/dtos/exam"
	dtoUser "english_exam_go/domain/dtos/user"
	"english_exam_go/infrastructure/data/repositories/persistence"
)

type IExamService interface {
	CreateExam(ctx context.Context, exam *dtoExam.UpsertExamRequest) error
	UpdateExam(ctx context.Context, exam *dtoExam.UpsertExamRequest) error
	UpdateExamQuestion(ctx context.Context, question *dtoExam.UpsertQuestionRequest) error
	GetAllExams(ctx context.Context, offset int, limit int) ([]*dtoExam.ExamListResponse, int, error)
	GetExamByCreatorID(ctx context.Context, offset int, limit int, UserID int) ([]*dtoExam.ExamListResponse, int, error)
	GetExamByTakerID(ctx context.Context, offset int, limit int, UserID int) ([]*dtoExam.ExamListResponse, int, error)
	GetDetailExamRoleUser(ctx context.Context, ID int) (*dtoExam.ExamDetailResponse, error)
	GetDetailExamRoleAdmin(ctx context.Context, ID int) (*dtoExam.ExamDetailResponse, error)
	DeleteExam(ctx context.Context, ID int) error
	GetParticipants(ctx context.Context, ID int) ([]*dtoUser.UserResponse, error)
}

type ExamServiceImpl struct {
	er persistence.IExamRepository
}

func (es ExamServiceImpl) GetParticipants(ctx context.Context, ID int) ([]*dtoUser.UserResponse, error) {
	//TODO implement me
	var result []*dtoUser.UserResponse
	userEnts, err := es.er.GetParticipants(ctx, ID)
	for _, ent := range userEnts {
		result = append(result, dtoUser.UserEntToResponse(&ent))
	}
	return result, err
}

func (es ExamServiceImpl) DeleteExam(ctx context.Context, ID int) error {
	//TODO implement me
	return es.er.DeleteExam(ctx, ID)
}

func (es ExamServiceImpl) GetDetailExamRoleAdmin(ctx context.Context, ID int) (*dtoExam.ExamDetailResponse, error) {
	//TODO implement me
	examEnts, err := es.er.FindExamById(ctx, uint(ID))
	if err != nil {
		return nil, err
	}
	return dtoExam.ParseExamDetailAdminRes(examEnts), nil
}

func (es ExamServiceImpl) GetExamByCreatorID(ctx context.Context, offset int, limit int, UserID int) ([]*dtoExam.ExamListResponse, int, error) {
	//TODO implement me
	examEnts, total, err := es.er.FindExamsByCreatorId(ctx, offset, limit, uint(UserID))
	if err != nil {
		return nil, 0, err
	}

	var responseExams []*dtoExam.ExamListResponse
	for _, e := range examEnts {
		responseExams = append(responseExams, dtoExam.CreateExamListRes(e))
	}
	return responseExams, total, nil
}

func (es ExamServiceImpl) GetExamByTakerID(ctx context.Context, offset, limit, UserID int) ([]*dtoExam.ExamListResponse, int, error) {
	//TODO implement me
	examEnts, total, err := es.er.FindExamsByCreatorId(ctx, offset, limit, uint(UserID))
	if err != nil {
		return nil, 0, err
	}

	var responseExams []*dtoExam.ExamListResponse
	for _, e := range examEnts {
		responseExams = append(responseExams, dtoExam.CreateExamListRes(e))
	}
	return responseExams, total, nil
}

func (es ExamServiceImpl) GetDetailExamRoleUser(ctx context.Context, ID int) (*dtoExam.ExamDetailResponse, error) {
	//TODO implement me
	examEnts, err := es.er.FindExamById(ctx, uint(ID))
	if err != nil {
		return nil, err
	}
	return dtoExam.ParseExamDetailRes(examEnts), nil
}

func (es ExamServiceImpl) GetAllExams(ctx context.Context, offset int, limit int) ([]*dtoExam.ExamListResponse, int, error) {
	//TODO implement me
	examEnts, err := es.er.FindAllExams(ctx, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	var responseExams []*dtoExam.ExamListResponse
	for _, e := range examEnts {
		responseExams = append(responseExams, dtoExam.CreateExamListRes(e))
	}
	total := es.er.CountTotal(ctx)
	return responseExams, total, nil
}

func (es ExamServiceImpl) CreateExam(ctx context.Context, exam *dtoExam.UpsertExamRequest) error {
	//TODO implement me
	examEnt := exam.CreateExamEntity()
	err := es.er.CreateExam(ctx, &examEnt)
	if err != nil {
		return err
	}
	return nil
}

func (es ExamServiceImpl) UpdateExam(ctx context.Context, exam *dtoExam.UpsertExamRequest) error {
	//TODO implement me
	examEnt := exam.CreateExamEntity()
	err := es.er.UpdateExam(ctx, &examEnt)
	if err != nil {
		return err
	}
	return nil
}

func (es ExamServiceImpl) UpdateExamQuestion(ctx context.Context, question *dtoExam.UpsertQuestionRequest) error {
	//TODO implement me
	questionEnt := question.CreateQuestionEntity()
	err := es.er.UpdateQuestion(ctx, &questionEnt)
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
