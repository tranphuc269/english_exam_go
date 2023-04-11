package v1

import (
	"english_exam_go/application/http_utils"
	"english_exam_go/application/http_utils/exception"
	dtos "english_exam_go/domain/dtos/exam"
	"english_exam_go/domain/services"
	"github.com/gin-gonic/gin"
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
