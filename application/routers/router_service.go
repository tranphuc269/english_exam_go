package routers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type Router interface {
	routing(e *gin.RouterGroup)
}

type Service struct {
	routers []Router
}

type AllRouter struct {
	dig.In
	BookRouter       Router `name:"book"`
	AuthRouter       Router `name:"auth"`
	ExamRouter       Router `name:"exam"`
	ExamSubmitRouter Router `name:"exam_submit"`
	ExamInviteRouter Router `name:"exam_invite"`
}

func NewService(allRouter AllRouter) *Service {
	routers := []Router{allRouter.BookRouter,
		allRouter.AuthRouter,
		allRouter.ExamRouter,
		allRouter.ExamSubmitRouter,
		allRouter.ExamInviteRouter}
	return &Service{
		routers: routers,
	}
}

func (rs *Service) RouterInit(e *gin.RouterGroup) {
	for _, r := range rs.routers {
		r.routing(e)
	}
}
