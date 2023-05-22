package dtos

import (
	"english_exam_go/infrastructure/data/entities"
	"english_exam_go/utils/resource"
	"gorm.io/gorm"
	"time"
)

type UpsertExamRequest struct {
	Id              int                     `json:"id"`
	ExamName        string                  `json:"exam_name"`
	ExamDescription string                  `json:"exam_description"`
	ListenFile      string                  `json:"listen_file"`
	ExamStartTime   time.Time               `json:"exam_start_time"  validate:"required,datetime"`
	ExamEndTime     time.Time               `json:"exam_end_time" validate:"required,datetime"`
	CreatorId       int                     `json:"creator_id"`
	ExamQuestions   []UpsertQuestionRequest `json:"exam_questions"`
}

func (cer UpsertExamRequest) CreateExamEntity() entities.Exam {
	var model gorm.Model
	if cer.Id != 0 {
		model = gorm.Model{ID: uint(cer.Id)}
	} else {
		model = gorm.Model{}
	}
	return entities.Exam{
		Model:           model,
		ExamName:        cer.ExamName,
		ExamDescription: cer.ExamDescription,
		ListenFile:      cer.ListenFile,
		ExamStartTime:   cer.ExamStartTime,
		ExamEndTime:     cer.ExamEndTime,
		ExamQuestions:   ListQuestionRequestToListQuestionEntity(cer.ExamQuestions),
		CreatorID:       uint(cer.CreatorId),
	}
}

type UpsertQuestionRequest struct {
	Id           int                   `json:"id"`
	QuestionText string                `json:"question_text"`
	File         string                `json:"file"`
	QuestionCase resource.QuestionCase `json:"question_case,omitempty" binding:"required,questionCase"`
	Answers      []UpsertAnswerRequest `json:"answers"`
}

func (cqr UpsertQuestionRequest) CreateQuestionEntity() entities.ExamQuestion {
	var model gorm.Model
	if cqr.Id != 0 {
		model = gorm.Model{ID: uint(cqr.Id)}
	} else {
		model = gorm.Model{}
	}
	return entities.ExamQuestion{
		Model:        model,
		QuestionText: cqr.QuestionText,
		File:         cqr.File,
		QuestionCase: cqr.QuestionCase,
		Answers:      ListAnswerRequestToListAnswerEntity(cqr.Answers),
	}
}

func ListQuestionRequestToListQuestionEntity(requests []UpsertQuestionRequest) []entities.ExamQuestion {
	var ents []entities.ExamQuestion
	for _, req := range requests {
		ents = append(ents, req.CreateQuestionEntity())
	}
	return ents
}

type UpsertAnswerRequest struct {
	Id        int    `json:"id"`
	Content   string `json:"content"`
	IsCorrect int    `json:"is_correct"`
}

func (car UpsertAnswerRequest) CreateAnswerEntity() entities.QuestionAnswer {
	var model gorm.Model
	if car.Id != 0 {
		model = gorm.Model{ID: uint(car.Id)}
	} else {
		model = gorm.Model{}
	}
	return entities.QuestionAnswer{
		Model:     model,
		Content:   car.Content,
		IsCorrect: car.IsCorrect,
	}
}

func ListAnswerRequestToListAnswerEntity(requests []UpsertAnswerRequest) []entities.QuestionAnswer {
	var ents []entities.QuestionAnswer
	for _, req := range requests {
		ents = append(ents, req.CreateAnswerEntity())
	}
	return ents
}

// response

type ExamListResponse struct {
	Id              uint      `json:"id"`
	ExamName        string    `json:"exam_name"`
	ExamDescription string    `json:"exam_description"`
	ExamStartTime   time.Time `json:"exam_start_time"`
	ExamEndTime     time.Time `json:"exam_end_time"`
	CreatorId       int       `json:"creator_id"`
}

func CreateExamListRes(entity *entities.Exam) *ExamListResponse {
	return &ExamListResponse{
		Id:              entity.ID,
		ExamName:        entity.ExamName,
		ExamDescription: entity.ExamDescription,
		ExamStartTime:   entity.ExamStartTime,
		ExamEndTime:     entity.ExamEndTime,
		CreatorId:       int(entity.CreatorID),
	}
}

type ExamDetailResponse struct {
	ID              int                `json:"id"`
	ExamName        string             `json:"exam_name"`
	ExamDescription string             `json:"exam_description"`
	ExamStartTime   time.Time          `json:"exam_start_time"`
	ListenFile      string             `json:"listen_file"`
	ExamEndTime     time.Time          `json:"exam_end_time"`
	CreatorId       int                `json:"creator_id"`
	ExamQuestions   []QuestionResponse `json:"exam_questions"`
}

type QuestionResponse struct {
	ID           int                   `json:"id"`
	QuestionText string                `json:"question_text"`
	File         string                `json:"file"`
	QuestionCase resource.QuestionCase `json:"question_case,omitempty"`
	Answers      []AnswerResponse      `json:"answers"`
}

type AnswerResponse struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	IsCorrect int    `json:"is_correct"`
}

func ParseExamDetailRes(entity *entities.Exam) *ExamDetailResponse {
	return &ExamDetailResponse{
		ID:              int(entity.ID),
		ExamName:        entity.ExamName,
		ExamDescription: entity.ExamDescription,
		ListenFile:      entity.ListenFile,
		ExamStartTime:   entity.ExamStartTime,
		ExamEndTime:     entity.ExamEndTime,
		CreatorId:       int(entity.CreatorID),
		ExamQuestions:   ParseListQuestionResponse(entity.ExamQuestions),
	}
}

func ParseQuestionResponse(question *entities.ExamQuestion) *QuestionResponse {
	return &QuestionResponse{
		ID:           int(question.ID),
		QuestionText: question.QuestionText,
		File:         question.File,
		QuestionCase: question.QuestionCase,
		Answers:      ParseListAnswerResponse(question.Answers),
	}
}

func ParseListQuestionResponse(questions []entities.ExamQuestion) []QuestionResponse {
	var questionResponse []QuestionResponse
	for _, ques := range questions {
		questionResponse = append(questionResponse, *ParseQuestionResponse(&ques))
	}
	return questionResponse
}

func ParseAnswerResponse(answer *entities.QuestionAnswer) *AnswerResponse {
	return &AnswerResponse{
		ID:      int(answer.ID),
		Content: answer.Content,
	}
}

func ParseListAnswerResponse(answers []entities.QuestionAnswer) []AnswerResponse {
	var answerResponse []AnswerResponse
	for _, ans := range answers {
		answerResponse = append(answerResponse, *ParseAnswerResponse(&ans))
	}
	return answerResponse
}

// ParseExamDetailAdminRes for admin
func ParseExamDetailAdminRes(entity *entities.Exam) *ExamDetailResponse {
	return &ExamDetailResponse{
		ID:              int(entity.ID),
		ExamName:        entity.ExamName,
		ExamDescription: entity.ExamDescription,
		ListenFile:      entity.ListenFile,
		ExamStartTime:   entity.ExamStartTime,
		ExamEndTime:     entity.ExamEndTime,
		CreatorId:       int(entity.CreatorID),
		ExamQuestions:   ParseListQuestionAdminResponse(entity.ExamQuestions),
	}
}

func ParseQuestionAdminResponse(question *entities.ExamQuestion) *QuestionResponse {
	return &QuestionResponse{
		ID:           int(question.ID),
		QuestionText: question.QuestionText,
		File:         question.File,
		QuestionCase: question.QuestionCase,
		Answers:      ParseListAnswerAdminResponse(question.Answers),
	}
}

func ParseListQuestionAdminResponse(questions []entities.ExamQuestion) []QuestionResponse {
	var questionResponse []QuestionResponse
	for _, ques := range questions {
		questionResponse = append(questionResponse, *ParseQuestionAdminResponse(&ques))
	}
	return questionResponse
}

func ParseAnswerAdminResponse(answer *entities.QuestionAnswer) *AnswerResponse {
	return &AnswerResponse{
		ID:        int(answer.ID),
		Content:   answer.Content,
		IsCorrect: answer.IsCorrect,
	}
}

func ParseListAnswerAdminResponse(answers []entities.QuestionAnswer) []AnswerResponse {
	var answerResponse []AnswerResponse
	for _, ans := range answers {
		answerResponse = append(answerResponse, *ParseAnswerAdminResponse(&ans))
	}
	return answerResponse
}
