package dtos

import (
	dtos "english_exam_go/domain/dtos/user"
	"gorm.io/gorm"
	"time"
)

type ExamResultRes struct {
	ID                  int               `json:"id"`
	ExamID              int               `json:"exam_id"`
	TotalScore          int               `json:"total_score"`
	User                dtos.UserResponse `json:"user"`
	ReadingScore        int               `json:"reading_score"`
	ListeningScore      int               `json:"listening_score"`
	NumCorrectReading   int               `json:"num_correct_reading"`
	NumCorrectListening int               `json:"num_correct_listening"`
	TabSwitchCount      int               `json:"tab_switch_count"`
	CreatedAt           time.Time         `json:"created_at"`
	UpdatedAt           time.Time         `json:"updated_at"`
}

type ExamResultResponseDetail struct {
	gorm.Model
}
