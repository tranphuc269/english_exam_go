package routers

import (
	v1 "english_exam_go/application/api/v1"
	"english_exam_go/application/middleware"
	"github.com/gin-gonic/gin"
)

type ExamSubmitRouter struct {
	esc *v1.ExamSubmitController
}

func (esr ExamSubmitRouter) routing(e *gin.RouterGroup) {
	//TODO implement me
	be := e.Group("/exam-submit").Use(middleware.Authentication())
	{
		be.POST("", esr.esc.SubmitExam())
	}
}

func CreateExamSubmitRouter(esc *v1.ExamSubmitController) Router {
	return &ExamSubmitRouter{esc: esc}
}
