package dtos

import (
	"english_exam_go/infrastructure/data/entities"
	"english_exam_go/utils/resource"
	"gorm.io/gorm"
	"time"
)

type CreateExamRequest struct {
	ExamName        string                  `json:"exam_name"`
	ExamDescription string                  `json:"exam_description"`
	ExamStartTime   time.Time               `json:"exam_start_time"  validate:"required,datetime"`
	ExamEndTime     time.Time               `json:"exam_end_time" validate:"required,datetime"`
	CreatorId       int                     `json:"creator_id"`
	ExamQuestions   []CreateQuestionRequest `json:"exam_questions"`
}

func (cer CreateExamRequest) CreateExamEntity() entities.Exam {
	return entities.Exam{
		Model:           gorm.Model{},
		ExamName:        cer.ExamName,
		ExamDescription: cer.ExamDescription,
		ExamStartTime:   cer.ExamStartTime,
		ExamEndTime:     cer.ExamEndTime,
		ExamQuestions:   ListQuestionRequestToListQuestionEntity(cer.ExamQuestions),
		CreatorID:       uint(cer.CreatorId),
	}
}

type CreateQuestionRequest struct {
	QuestionText string                `json:"question_text"`
	QuestionCase resource.QuestionCase `json:"question_case,omitempty" binding:"required,questionCase"`
	Answers      []CreateAnswerRequest `json:"answers"`
}

func (cqr CreateQuestionRequest) CreateQuestionEntity() entities.ExamQuestion {

	return entities.ExamQuestion{
		Model:        gorm.Model{},
		QuestionText: cqr.QuestionText,
		QuestionCase: cqr.QuestionCase,
		Answers:      ListAnswerRequestToListAnswerEntity(cqr.Answers),
	}
}

func ListQuestionRequestToListQuestionEntity(requests []CreateQuestionRequest) []entities.ExamQuestion {
	var ents []entities.ExamQuestion
	for _, req := range requests {
		ents = append(ents, req.CreateQuestionEntity())
	}
	return ents
}

type CreateAnswerRequest struct {
	Content   string `json:"content"`
	IsCorrect int    `json:"is_correct"`
}

func (car CreateAnswerRequest) CreateAnswerEntity() entities.QuestionAnswer {
	return entities.QuestionAnswer{
		Model:     gorm.Model{},
		Content:   car.Content,
		IsCorrect: car.IsCorrect,
	}
}

func ListAnswerRequestToListAnswerEntity(requests []CreateAnswerRequest) []entities.QuestionAnswer {
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
	ExamEndTime     time.Time          `json:"exam_end_time"`
	CreatorId       int                `json:"creator_id"`
	ExamQuestions   []QuestionResponse `json:"exam_questions"`
}

type QuestionResponse struct {
	ID           int                   `json:"id"`
	QuestionText string                `json:"question_text"`
	QuestionCase resource.QuestionCase `json:"question_case,omitempty"`
	Answers      []AnswerResponse      `json:"answers"`
}

type AnswerResponse struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

func ParseExamDetailRes(entity *entities.Exam) *ExamDetailResponse {
	return &ExamDetailResponse{
		ID:              int(entity.ID),
		ExamName:        entity.ExamName,
		ExamDescription: entity.ExamDescription,
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
