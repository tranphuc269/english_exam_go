package routers

import (
	v1 "english_exam_go/application/api/v1"
	"english_exam_go/application/middleware"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	ac *v1.AuthController
}

func (ar AuthRouter) routing(e *gin.RouterGroup) {
	//TODO implement me
	be := e.Group("/auth")
	{
		be.POST("/login", ar.ac.Login())
		be.POST("/register", ar.ac.Register())
		be.GET("/teachers", ar.ac.GetTeachers())
		be.GET("/students", ar.ac.GetStudents())
		secured := be.Group("/secured").Use(middleware.Authentication())
		{
			secured.GET("/me", ar.ac.Me())
			secured.PUT("/update", ar.ac.Update())
			secured.DELETE("/:id", ar.ac.DeleteUser())
			secured.GET("/user/:id", ar.ac.UserDetail())
		}
	}
}

func CreateAuthRouter(ac *v1.AuthController) Router {
	return AuthRouter{ac: ac}
}
