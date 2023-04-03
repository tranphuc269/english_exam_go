package v1

import (
	"english_exam_go/application/http_utils"
	"english_exam_go/application/http_utils/exception"
	dtos "english_exam_go/domain/dtos/exam"
	"english_exam_go/domain/services"
	"github.com/gin-gonic/gin"
)

type ExamController struct {
	es services.IExamService
}

func CreateExamController(es services.IExamService) *ExamController {
	return &ExamController{es: es}
}

func (ec *ExamController) CreateExam() gin.HandlerFunc {
	return func(c *gin.Context) {
		var createExamRequest dtos.CreateExamRequest
		if err := c.ShouldBindJSON(&createExamRequest); err != nil {
			exception.Handle(err, c)
			//return
		}
		err := ec.es.CreateExam(c, &createExamRequest)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.SuccessHandle(createExamRequest, c)
	}
}

func (ec *ExamController) GetExams() gin.HandlerFunc {
	return func(c *gin.Context) {
		exams, err := ec.es.GetAllExams(c)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.SuccessHandle(exams, c)
	}
}

func (ec *ExamController) GetDetailExam() gin.HandlerFunc {
	return func(c *gin.Context) {
		http_utils.SuccessHandle("Get detail exams", c)
	}
}

func (ec *ExamController) UpdateExams() gin.HandlerFunc {
	return func(c *gin.Context) {
		http_utils.SuccessHandle("Update Exam", c)
	}
}

func (ec *ExamController) DeleteExam() gin.HandlerFunc {
	return func(c *gin.Context) {
		http_utils.SuccessHandle("Delete exam", c)
	}
}
