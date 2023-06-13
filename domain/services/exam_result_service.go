package services

import (
	"context"
	dtosExam "english_exam_go/domain/dtos/exam"
	dtosUser "english_exam_go/domain/dtos/user"
	"english_exam_go/infrastructure/data/entities"
	"english_exam_go/infrastructure/data/repositories/persistence"
	"english_exam_go/utils/resource"
	"gorm.io/gorm"
)

type IExamResultService interface {
	SubmitExam(ctx context.Context, submitRequest dtosExam.CreateExamSubmit) (*dtosExam.ExamResultRes, error)
	GetYourExamDone(ctx context.Context, UserID int) []*dtosExam.ExamResultRes
	GetAllResult(ctx context.Context, examId int, score int) []*dtosExam.ExamResultRes
}

type ExamResultServiceImpl struct {
	err persistence.IExamResultRepository
	er  persistence.IExamRepository
	ar  persistence.IAuthRepository
}

func (ers ExamResultServiceImpl) GetAllResult(ctx context.Context, examId int, score int) []*dtosExam.ExamResultRes {
	//TODO implement me
	examDoneEnts := ers.err.GetAllResult(ctx, examId, score)
	var results []*dtosExam.ExamResultRes
	for _, ent := range examDoneEnts {
		userEnt, _ := ers.ar.FindById(ctx, ent.UserId)
		userRes := dtosUser.UserEntToResponse(userEnt)
		results = append(results, &dtosExam.ExamResultRes{
			ID:                  int(ent.ID),
			CreatedAt:           ent.CreatedAt,
			UpdatedAt:           ent.UpdatedAt,
			ExamID:              ent.ExamId,
			User:                *userRes,
			TotalScore:          ent.TotalScore,
			ReadingScore:        ent.ReadingScore,
			ListeningScore:      ent.ListeningScore,
			NumCorrectReading:   ent.NumCorrectReading,
			NumCorrectListening: ent.NumCorrectListening,
			TabSwitchCount:      ent.TabSwitchCount,
		})
	}
	return results
}

func (ers ExamResultServiceImpl) GetYourExamDone(ctx context.Context,
	UserID int) []*dtosExam.ExamResultRes {
	//TODO implement me
	examDoneEnts := ers.err.GetListExamByTakerID(ctx, UserID)
	var results []*dtosExam.ExamResultRes
	for _, ent := range examDoneEnts {
		results = append(results, &dtosExam.ExamResultRes{
			ID:                  int(ent.ID),
			CreatedAt:           ent.CreatedAt,
			UpdatedAt:           ent.UpdatedAt,
			ExamID:              ent.ExamId,
			TotalScore:          ent.TotalScore,
			ReadingScore:        ent.ReadingScore,
			ListeningScore:      ent.ListeningScore,
			NumCorrectReading:   ent.NumCorrectReading,
			NumCorrectListening: ent.NumCorrectListening,
			TabSwitchCount:      ent.TabSwitchCount,
		})
	}
	return results
}

func (ers ExamResultServiceImpl) SubmitExam(ctx context.Context,
	submitRequest dtosExam.CreateExamSubmit) (*dtosExam.ExamResultRes, error) {
	//TODO implement me

	numCorrectReading := 0
	numCorrectListening := 0

	examDetail, err := ers.er.FindExamById(ctx, uint(submitRequest.ExamId))
	///  tạo một map, có key là id của question, value là kết quả answer của question gửi lên
	confRequestQuestion := map[int]int{}
	for _, reqQues := range submitRequest.SubmissionResults {
		confRequestQuestion[reqQues.QuestionId] = reqQues.AnswerId
	}
	for _, ques := range examDetail.ExamQuestions {
		if ques.QuestionCase == resource.QuestionReading {
			if confRequestQuestion[int(ques.ID)] == ques.GetCorrectAnswer() {
				numCorrectReading++
			}
		} else {
			if confRequestQuestion[int(ques.ID)] == ques.GetCorrectAnswer() {
				numCorrectListening++
			}
		}
	}
	readingScore := resource.WritingToeicResult[numCorrectReading]
	listeningScore := resource.ReadingToeicResult[numCorrectListening]
	totalScore := readingScore + listeningScore
	examResultEnt := entities.ExamResult{
		Model:               gorm.Model{},
		ExamId:              submitRequest.ExamId,
		UserId:              submitRequest.TakerID,
		TotalScore:          totalScore,
		ReadingScore:        readingScore,
		ListeningScore:      listeningScore,
		NumCorrectReading:   numCorrectReading,
		NumCorrectListening: numCorrectListening,
		TabSwitchCount:      submitRequest.TabSwitchCount,
	}
	result, err := ers.err.SubmitExam(ctx, &examResultEnt)
	return result, err
}

func CreateExamResultService(err persistence.IExamResultRepository,
	er persistence.IExamRepository, ar persistence.IAuthRepository) IExamResultService {
	return &ExamResultServiceImpl{
		err: err,
		er:  er,
		ar:  ar,
	}
}
