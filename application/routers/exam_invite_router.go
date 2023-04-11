package routers

import (
	v1 "english_exam_go/application/api/v1"
	"github.com/gin-gonic/gin"
)

type ExamInviteRouter struct {
	eic *v1.ExamInviteController
}

func (eir ExamInviteRouter) routing(e *gin.RouterGroup) {
	//TODO implement me
	be := e.Group("/exam_invite")
	{
		be.POST("", eir.eic.AddUserToExam())
		be.DELETE("", eir.eic.RemoveUserFromExam())
	}
}

func CreateExamInviteRouter(eic *v1.ExamInviteController) Router {
	return ExamInviteRouter{eic: eic}
}
