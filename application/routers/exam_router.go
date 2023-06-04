package routers

import (
	v1 "english_exam_go/application/api/v1"
	"english_exam_go/application/middleware"
	"english_exam_go/utils/file"
	"github.com/gin-gonic/gin"
)

type ExamRouter struct {
	ec *v1.ExamController
}

func (er ExamRouter) routing(e *gin.RouterGroup) {
	//TODO implement me
	be := e.Group("/exams")
	{
		be.GET("/", er.ec.GetExams())
		be.GET("/:id", er.ec.GetDetailExamRoleUser())
		be.GET("/admin/:id", er.ec.GetDetailExamRoleAdmin())
		be.GET("/creators/:id", er.ec.GetExamsByCreatorID())
		be.GET("/takers/:id", er.ec.GetExamsByCreatorID())
		be.GET("/participants/:id", er.ec.GetParticipants())
		be.POST("/invite/:id", er.ec.Invite())
		be.POST("/file", file.UploadFile)

		secured := be.Group("/secured").Use(middleware.Authentication())
		{
			secured.POST("/", er.ec.CreateExam())
			secured.DELETE("/:id", er.ec.DeleteExam())
			secured.PUT("/:id", er.ec.UpdateExam())
			secured.PUT("/question/:id", er.ec.UpdateQuestion())
		}
	}
}

func CreateExamRouter(ec *v1.ExamController) Router {
	return ExamRouter{ec: ec}
}
