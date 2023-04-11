package v1

import (
	"english_exam_go/application/http_utils"
	"english_exam_go/application/http_utils/exception"
	dtos "english_exam_go/domain/dtos/exam"
	"english_exam_go/domain/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ExamInviteController struct {
	eis services.IExamInviteService
}

func CreateExamInviteController(eis services.IExamInviteService) *ExamInviteController {
	return &ExamInviteController{eis: eis}
}

func (eic ExamInviteController) AddUserToExam() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dtos.AddTakerToExam
		if err := c.ShouldBindJSON(&request); err != nil {
			exception.Handle(err, c)
			return
		}
		err := eic.eis.AddUserToExam(c, request)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.CreatedHandle(c)
	}
}

func (eic ExamInviteController) RemoveUserFromExam() gin.HandlerFunc {
	return func(c *gin.Context) {
		ExamID, err := strconv.Atoi(c.Query("exam_id"))
		if err != nil {
			exception.Handle(err, c)
			return
		}
		UserID, err := strconv.Atoi(c.Query("user_id"))
		if err != nil {
			exception.Handle(err, c)
			return
		}
		err = eic.eis.RemoveUserToExam(c, ExamID, UserID)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.NoContentHandle(c)
	}
}
