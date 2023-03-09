package dtos

import "time"

type CreateExamInviteRequest struct {
	ExamId         int       `json:"exam_id"`
	InviterId      int       `json:"inviter_id"`
	InviteeEmail   int       `json:"invitee_email"`
	InvitationTime time.Time `json:"invitation_time" validate:"required,datetime"`
}

type ExamInviteResponse struct {
}
