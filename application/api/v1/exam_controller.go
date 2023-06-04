package v1

import (
	"english_exam_go/application/http_utils"
	"english_exam_go/application/http_utils/exception"
	dtos "english_exam_go/domain/dtos/exam"
	"english_exam_go/domain/services"
	auth_utils "english_exam_go/utils/auth"
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
		var updateRequest dtos.UpsertExamRequest
		if err := c.ShouldBindJSON(&updateRequest); err != nil {
			exception.Handle(err, c)
			//return
		}
		updateRequest.CreatorId = claim.UserID
		err := ec.es.CreateExam(c, &updateRequest)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.SuccessHandle(updateRequest, c)
	}
}

func (ec *ExamController) UpdateExam() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

		claim, _ := auth_utils.ParseToken(tokenString)
		var updateExamRequest dtos.UpsertExamRequest
		var params = c.Param("id")
		ID, err := strconv.Atoi(params)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		if err := c.ShouldBindJSON(&updateExamRequest); err != nil {
			exception.Handle(err, c)
			//return
		}
		updateExamRequest.CreatorId = claim.UserID
		updateExamRequest.Id = ID
		err = ec.es.UpdateExam(c, &updateExamRequest)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.SuccessHandle(updateExamRequest, c)
	}
}

func (ec *ExamController) UpdateQuestion() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

		//claim, _ := auth_utils.ParseToken(tokenString)
		var questionRequest dtos.UpsertQuestionRequest
		var params = c.Param("id")
		ID, err := strconv.Atoi(params)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		if err := c.ShouldBindJSON(&questionRequest); err != nil {
			exception.Handle(err, c)
			//return
		}
		questionRequest.Id = ID
		err = ec.es.UpdateExamQuestion(c, &questionRequest)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.SuccessHandle(questionRequest, c)
	}
}

func (ec *ExamController) GetExams() gin.HandlerFunc {
	return func(c *gin.Context) {
		var offsetStr, _ = c.GetQuery("offset")
		var limitStr, _ = c.GetQuery("limit")
		offsetInt, err := strconv.Atoi(offsetStr)
		if err != nil {
			offsetInt = 0
		}
		limitInt, err := strconv.Atoi(limitStr)
		if err != nil {
			limitInt = 10
		}
		exams, total, err := ec.es.GetAllExams(c, offsetInt, limitInt)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.SuccessHandlePaginate(exams, total, c)
	}
}

func (ec *ExamController) GetDetailExamRoleUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params = c.Param("id")
		ID, err := strconv.Atoi(params)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		exam, _ := ec.es.GetDetailExamRoleUser(c, ID)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.SuccessHandle(exam, c)
	}
}

func (ec *ExamController) GetDetailExamRoleAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params = c.Param("id")
		ID, err := strconv.Atoi(params)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		exam, _ := ec.es.GetDetailExamRoleAdmin(c, ID)
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
		var offsetStr, _ = c.GetQuery("offset")
		var limitStr, _ = c.GetQuery("limit")
		offsetInt, err := strconv.Atoi(offsetStr)
		if err != nil {
			offsetInt = 0
		}
		limitInt, err := strconv.Atoi(limitStr)
		if err != nil {
			limitInt = 10
		}
		exams, total, err := ec.es.GetExamByCreatorID(c, offsetInt, limitInt, creatorID)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.SuccessHandlePaginate(exams, total, c)
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
		var offsetStr, _ = c.GetQuery("offset")
		var limitStr, _ = c.GetQuery("limit")
		offsetInt, err := strconv.Atoi(offsetStr)
		if err != nil {
			offsetInt = 0
		}
		limitInt, err := strconv.Atoi(limitStr)
		if err != nil {
			limitInt = 10
		}
		exams, total, err := ec.es.GetExamByTakerID(c, offsetInt, limitInt, takerID)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.SuccessHandlePaginate(exams, total, c)
	}
}

func (ec *ExamController) DeleteExam() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params = c.Param("id")
		id, err := strconv.Atoi(params)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		err = ec.es.DeleteExam(c, id)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.SuccessHandle(nil, c)
	}
}
func (ec *ExamController) GetParticipants() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params = c.Param("id")
		id, err := strconv.Atoi(params)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		users, err := ec.es.GetParticipants(c, id)
		if err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.SuccessHandle(users, c)
	}
}
