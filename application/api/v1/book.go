package v1

import (
	"english_exam_go/application/http_utils"
	"english_exam_go/application/http_utils/exception"
	"english_exam_go/domain/dtos/book"
	"english_exam_go/domain/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BookController struct {
	bs services.IBookService
}

func CreateBookController(bs services.IBookService) *BookController {
	return &BookController{bs: bs}
}

func (bc *BookController) CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var bookRequest book.RequestBook
		if err := c.ShouldBind(&bookRequest); err != nil {
			bindingError := &exception.BindingError{
				ErrMsg:        "[application.api.v1.book.CreateBook] failed to bind Book from request.",
				OriginalError: err,
			}
			exception.Handle(bindingError, c)
			return
		}
		err := bc.bs.CreateBook(c, &bookRequest)
		if err != nil {
			fmt.Printf("Error database : %s", err.Error())
			exception.Handle(err, c)
			return
		}
		http_utils.CreatedHandle(c)
	}
}

func (bc *BookController) FindBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		ID, err := strconv.Atoi(idStr)
		if err != nil {
			exception.Handle(err, c)
		}
		responseBook, err := bc.bs.FindBook(c, uint(ID))
		if err != nil {
			exception.Handle(err, c)
		}
		http_utils.SuccessHandle(*responseBook, c)
	}
}

func (bc *BookController) FindBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := bc.bs.FindBooks(c)
		if err != nil {
			exception.Handle(err, c)
		}
		http_utils.SuccessHandle(books, c)
	}
}

func (bc *BookController) UpdateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		ID, err := strconv.Atoi(idStr)
		if err != nil {
			exception.Handle(err, c)
		}
		var bookRequest book.RequestBook
		if err := c.ShouldBind(&bookRequest); err != nil {
			bindingError := &exception.BindingError{
				ErrMsg:        "[application.api.v1.book.UpdatedBook] failed to bind Book from request",
				OriginalError: err,
			}
			exception.Handle(bindingError, c)
			return
		}
		if err := bc.bs.UpdateBook(c, &bookRequest, uint(ID)); err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.SuccessHandle(nil, c)
	}
}

func (bc *BookController) DeleteBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		ID, err := strconv.Atoi(idStr)
		if err != nil {
			exception.Handle(err, c)
		}
		if err := bc.bs.DeleteBook(c, uint(ID)); err != nil {
			exception.Handle(err, c)
			return
		}
		http_utils.NoContentHandle(c)
	}
}
