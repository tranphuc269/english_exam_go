package entities

import (
	"gorm.io/gorm"
	"time"
)

type Exam struct {
	gorm.Model
	ExamName        string         `json:"exam_name"`
	ExamDescription string         `json:"exam_description"`
	ListenFile      string         `json:"listen_file"`
	ExamStartTime   time.Time      `json:"exam_start_time"`
	ExamEndTime     time.Time      `json:"exam_end_time"`
	ExamQuestions   []ExamQuestion `json:"exam_questions" gorm:"ForeignKey:ExamId;"`
	CreatorID       uint           `json:"creator_id"`
	ExamCreator     User           `json:"exam_creator" gorm:"ForeignKey:CreatorID;"`
	ExamTakers      []User         `json:"exam_takers" gorm:"many2many:exam_takers"`
}

func (Exam) TableName() string {
	return "exams"
}

type ExamTasker struct {
	ExamID int `json:"exam_id"`
	UserID int `json:"user_id"`
}

func (ExamTasker) TableName() string {
	return "exam_takers"
}
