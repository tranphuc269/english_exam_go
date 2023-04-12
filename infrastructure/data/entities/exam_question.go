package entities

import (
	"english_exam_go/utils/resource"
	"gorm.io/gorm"
)

type ExamQuestion struct {
	gorm.Model
	ExamId       int                   `json:"exam_id"`
	QuestionText string                `json:"question_text"`
	QuestionCase resource.QuestionCase `json:"question_case"`
	Answers      []QuestionAnswer      `json:"answers" gorm:"ForeignKey:QuestionId;"`
}

func (eq *ExamQuestion) GetCorrectAnswer() int {
	for _, answer := range eq.Answers {
		if answer.IsCorrect == 1 {
			return int(answer.ID)
		}
	}
	return -1
}

type QuestionAnswer struct {
	gorm.Model
	QuestionId int    `json:"question_id"`
	Content    string `json:"content"`
	IsCorrect  int    `json:"is_correct"`
}
