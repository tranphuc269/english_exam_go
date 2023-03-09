package persistence

import (
	"context"
	"encoding/json"
	"english_exam_go/infrastructure/data/entities"
	"english_exam_go/infrastructure/data/repositories"
	"fmt"
)

const bookTables = "book_ents"

type IBookRepository interface {
	CreateBook(context.Context, *entities.Book) error
	FindBook(context.Context, uint) (*entities.Book, error)
	FindBooks(context.Context) ([]*entities.Book, error)
	UpdateBook(context.Context, *entities.Book) error
	DeleteBook(context.Context, uint) error
}

type BookRepositoryImpl struct {
}

func (br BookRepositoryImpl) CreateBook(ctx context.Context, ent *entities.Book) error {
	//TODO implement me
	db, _ := repositories.GetTx(ctx)
	//db := repositories.GetConn()
	result := db.Create(ent)
	if result.Error != nil {
		return &repositories.RdbRuntimeError{
			ErrMsg:        fmt.Sprintf("[infrastructure.data.repositories.persistence.CreateBook] fail to insert Book to Database"),
			OriginalError: result.Error,
		}
	}
	return nil
}

func (br BookRepositoryImpl) FindBook(ctx context.Context, ID uint) (*entities.Book, error) {
	//TODO implement me
	db := repositories.GetConn().Table(bookTables)
	bookEnt := &entities.Book{}

	err := db.Take(&bookEnt, "id=?", ID)
	if err != nil {
		return nil, &repositories.NotFoundError{
			Msg:           repositories.DefaultNotFoundMsg,
			ErrMsg:        fmt.Sprintf("[infrastructure.data.repositories.persistence.FindBook] failed to find bookEnt from rdb. ID : %d", ID),
			OriginalError: err.Error,
		}
	}
	return bookEnt, err.Error
}

func (br BookRepositoryImpl) FindBooks(ctx context.Context) ([]*entities.Book, error) {
	//TODO implement me
	db := repositories.GetConn()
	var books []*entities.Book
	result := db.Order("title").Limit(51).Find(&books)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, &repositories.NotFoundError{
			Msg:           repositories.DefaultNotFoundMsg,
			ErrMsg:        fmt.Sprintf("[infrastructure.data.repositories.persistence.FindBooks] failed to find books from rdb"),
			OriginalError: nil,
		}
	}
	return books, nil
}

func (br BookRepositoryImpl) UpdateBook(ctx context.Context, requestBook *entities.Book) error {
	//TODO implement me
	db := repositories.GetConn()
	b, _ := json.Marshal(requestBook)
	fmt.Printf("entity update : %s", b)
	//db, _ := repositories.GetTx(ctx)

	result := db.Model(&requestBook).Updates(requestBook)
	if result.Error != nil {
		return &repositories.RdbRuntimeError{
			ErrMsg:        fmt.Sprintf("[[infrastructure.data.repositories.persistence.UpdateBook] failed to update Book in RDB. ID : %d", requestBook.ID),
			OriginalError: result.Error,
		}
	} else if result.RowsAffected == 0 {
		return &repositories.OptimisticLockError{
			Msg:    "Cập nhật không thành công. Sách có thể đã bị xóa hoặc cập nhật bởi người dùng khác.",
			ErrMsg: fmt.Sprintf("[infrastructure.persistence.UpdateBook] failed to update Book in RDB. May be optimistic lock. ID : %d", requestBook.ID),
		}
	}
	return nil
}

func (br BookRepositoryImpl) DeleteBook(ctx context.Context, ID uint) error {
	//TODO implement me
	db, _ := repositories.GetTx(ctx)
	result := db.Where("id = ?", ID).Delete(&entities.Book{})
	if result.Error != nil {
		return &repositories.RdbRuntimeError{
			ErrMsg:        fmt.Sprintf("[infrastructure.persistence.DeleteBook] failed to delete Book in RDB. ID : %d", ID),
			OriginalError: result.Error,
		}
	} else if result.RowsAffected == 0 {
		return &repositories.NotFoundError{
			Msg:           "The book to be deleted was not found. The book may have already been deleted.",
			ErrMsg:        fmt.Sprintf("[infrastructure.persistence.DeleteBook] failed to delete Book in RDB. Record not found. ID : %d", ID),
			OriginalError: nil,
		}
	}

	return nil
}

func CreateBookRepository() IBookRepository {
	return &BookRepositoryImpl{}
}
