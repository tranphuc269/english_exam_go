package dtos

import "gorm.io/gorm"

type ExamResultRes struct {
	ID                  int `json:"id"`
	ExamID              int `json:"exam_id"`
	TotalScore          int `json:"total_score"`
	ReadingScore        int `json:"reading_score"`
	ListeningScore      int `json:"listening_score"`
	NumCorrectReading   int `json:"num_correct_reading"`
	NumCorrectListening int `json:"num_correct_listening"`
	TabSwitchCount      int `json:"tab_switch_count"`
}

type ExamResultResponseDetail struct {
	gorm.Model
}
