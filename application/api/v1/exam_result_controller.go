package v1

import (
	"english_exam_go/application/http_utils"
	"english_exam_go/application/http_utils/exception"
	dtos "english_exam_go/domain/dtos/exam"
	"english_exam_go/domain/services"
	auth_utils "english_exam_go/utils/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type ExamResultController struct {
	ess services.IExamResultService
}

func (esc ExamResultController) SubmitExam() gin.HandlerFunc {
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

func (esc ExamResultController) GetListExam() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, _ := auth_utils.GetUserIdFromToken(c)
		//if err != nil {
		//	domain_exception.Handle(err, c)
		//	return
		//}
		exams := esc.ess.GetYourExamDone(c, userId)
		fmt.Println(exams)
		http_utils.SuccessHandle(exams, c)
	}
}

func CreateExamSubmitController(ess services.IExamResultService) *ExamResultController {
	return &ExamResultController{ess: ess}
}
