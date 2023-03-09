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

type QuestionAnswer struct {
	gorm.Model
	QuestionId int    `json:"question_id"`
	Content    string `json:"content"`
	IsCorrect  string `json:"is_correct"`
}
