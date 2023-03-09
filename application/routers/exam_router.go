package routers

import (
	v1 "english_exam_go/application/api/v1"
	"github.com/gin-gonic/gin"
)

type ExamRouter struct {
	ec *v1.ExamController
}

func (er ExamRouter) routing(e *gin.RouterGroup) {
	//TODO implement me
	be := e.Group("/exams")
	{
		be.POST("/", er.ec.CreateExam())
	}
}

func CreateExamRouter(ec *v1.ExamController) Router {
	return ExamRouter{ec: ec}
}
