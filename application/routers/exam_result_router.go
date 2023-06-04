package routers

import (
	v1 "english_exam_go/application/api/v1"
	"english_exam_go/application/middleware"
	"github.com/gin-gonic/gin"
)

type ExamResultRouter struct {
	esc *v1.ExamResultController
}

func (esr ExamResultRouter) routing(e *gin.RouterGroup) {
	//TODO implement me
	be := e.Group("/exam-submit").Use(middleware.Authentication())
	{
		be.POST("", esr.esc.SubmitExam())
		be.GET("", esr.esc.GetListExamHaveSubmit())
		be.GET("/admin", esr.esc.GetAllResult())
	}
}

func CreateExamSubmitRouter(esc *v1.ExamResultController) Router {
	return &ExamResultRouter{esc: esc}
}
