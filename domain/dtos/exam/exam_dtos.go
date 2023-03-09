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
	ExamStartTime   time.Time               `json:"exam_start_time"`
	ExamEndTime     time.Time               `json:"exam_end_time" validate:"required,datetime"`
	CreatorId       int                     `json:"creator_id" validate:"required,datetime"`
	ExamQuestions   []CreateQuestionRequest `json:"exam_questions"`
}

func (cer CreateExamRequest) CreateExamEntity() entities.Exam {
	return entities.Exam{
		Model:           gorm.Model{},
		ExamName:        cer.ExamName,
		ExamDescription: cer.ExamDescription,
		ExamStartTime:   time.Time{},
		ExamEndTime:     time.Time{},
		ExamQuestions:   ListQuestionRequestToListQuestionEntity(cer.ExamQuestions),
		CreatorID:       1,
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
	IsCorrect string `json:"is_correct"`
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
