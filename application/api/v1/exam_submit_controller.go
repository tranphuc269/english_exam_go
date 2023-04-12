package v1

import (
	"english_exam_go/application/http_utils"
	"english_exam_go/application/http_utils/exception"
	dtos "english_exam_go/domain/dtos/exam"
	"english_exam_go/domain/services"
	auth_utils "english_exam_go/utils/auth"
	"github.com/gin-gonic/gin"
	"strings"
)

type ExamSubmitController struct {
	ess services.IExamResultService
}

func (esc ExamSubmitController) SubmitExam() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

		claim, _ := auth_utils.ParseToken(tokenString)
		var createExamSubmit dtos.CreateExamSubmit
		if err := c.ShouldBindJSON(&createExamSubmit); err != nil {
			exception.Handle(err, c)
			return
		}
		createExamSubmit.TakerID = claim.UserID
		result, err := esc.ess.SubmitExam(c, createExamSubmit)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.SuccessHandle(result, c)
	}
}

func CreateExamSubmitController(ess services.IExamResultService) *ExamSubmitController {
	return &ExamSubmitController{ess: ess}
}
