package dtos

import (
	"english_exam_go/utils/resource"
	"time"
)

type CreateExamRequest struct {
	ExamName        string    `json:"exam_name"`
	ExamDescription string    `json:"exam_description"`
	ExamStartTime   time.Time `json:"exam_start_time"`
	ExamEndTime     time.Time `json:"exam_end_time" validate:"required,datetime"`
	CreatorId       int       `json:"creator_id" validate:"required,datetime"`
}

type CreateQuestionRequest struct {
	QuestionText string                `json:"question_text"`
	QuestionCase resource.QuestionCase `json:"question_case,omitempty" binding:"required,questionCase"`
}

type CreateAnswerRequest struct {
	Content   string `json:"content"`
	IsCorrect string `json:"is_correct"`
}
