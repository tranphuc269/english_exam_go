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
	CreateAuthRouter Router `name:"auth"`
}

func NewService(allRouter AllRouter) *Service {
	routers := []Router{allRouter.BookRouter}
	return &Service{
		routers: routers,
	}
}

func (rs *Service) RouterInit(e *gin.RouterGroup) {

	for _, r := range rs.routers {
		r.routing(e)
	}

}
