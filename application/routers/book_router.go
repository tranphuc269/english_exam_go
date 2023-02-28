package routers

import (
	"english_exam_go/application/api/v1"
	"github.com/gin-gonic/gin"
)

type BookRouter struct {
	bc *v1.BookController
}

func (b BookRouter) routing(e *gin.RouterGroup) {
	//TODO implement me
	be := e.Group("/books")
	{
		be.POST("", b.bc.CreateBook())
		be.GET("/:id", b.bc.FindBook())
		be.GET("", b.bc.FindBooks())
		be.PUT("/:id", b.bc.UpdateBook())
		be.DELETE("/:id", b.bc.DeleteBook())
	}
}

func CreateBookRouter(bc *v1.BookController) Router {
	return &BookRouter{bc: bc}
}
