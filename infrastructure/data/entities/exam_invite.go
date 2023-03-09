package entities

import (
	"gorm.io/gorm"
	"time"
)

type ExamInvite struct {
	gorm.Model
	ExamId         int       `json:"exam_id"`
	InviterId      int       `json:"inviter_id"`
	InviteeId      int       `json:"invitee_id"`
	InviteeEmail   string    `json:"invitee_email"`
	InvitationTime time.Time `json:"invitation_time"`
}
