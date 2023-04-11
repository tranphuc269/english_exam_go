package v1

import (
	"english_exam_go/application/http_utils"
	"english_exam_go/application/http_utils/exception"
	dtos "english_exam_go/domain/dtos/exam"
	"english_exam_go/domain/services"
	auth_utils "english_exam_go/utils/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type ExamController struct {
	es services.IExamService
}

func CreateExamController(es services.IExamService) *ExamController {
	return &ExamController{es: es}
}

func (ec *ExamController) CreateExam() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

		claim, _ := auth_utils.ParseToken(tokenString)
		var createExamRequest dtos.CreateExamRequest
		if err := c.ShouldBindJSON(&createExamRequest); err != nil {
			exception.Handle(err, c)
			//return
		}
		createExamRequest.CreatorId = claim.UserID
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
		var params = c.Param("id")
		fmt.Println(params)
		ID, err := strconv.Atoi(params)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		exam, _ := ec.es.GetDetailExams(c, ID)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.SuccessHandle(exam, c)
	}
}

func (ec *ExamController) GetExamsByCreatorID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params = c.Param("id")
		creatorID, err := strconv.Atoi(params)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		exams, err := ec.es.GetExamByCreatorID(c, creatorID)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.SuccessHandle(exams, c)
	}
}

func (ec *ExamController) GetExamByTakerID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params = c.Param("id")
		takerID, err := strconv.Atoi(params)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		exams, err := ec.es.GetExamByTakerID(c, takerID)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.SuccessHandle(exams, c)
	}
}
