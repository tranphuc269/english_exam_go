package entities

import "gorm.io/gorm"

type ExamResult struct {
	gorm.Model
	ExamId              int  `json:"exam_id"`
	UserId              int  `json:"user_id"`
	TotalScore          int  `json:"total_score"`
	ReadingScore        int  `json:"reading_score"`
	ListeningScore      int  `json:"listening_score"`
	NumCorrectReading   int  `json:"num_correct_reading"`
	NumCorrectListening int  `json:"num_correct_listening"`
	TabSwitchCount      int  `json:"tab_switch_count"`
	UserOwner           User `json:"user_owner" gorm:"ForeignKey:UserId;"`
}
