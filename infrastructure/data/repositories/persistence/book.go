package persistence

import (
	"context"
	"english_exam_go/infrastructure/data/entities"
	"english_exam_go/infrastructure/data/repositories"
	"english_exam_go/utils/app_logger"
	"fmt"
)

type IBookRepository interface {
	CreateBook(context.Context, *entities.BookEnt) error
	FindBook(context.Context, uint) (*entities.BookEnt, error)
	FindBooks(context.Context) ([]*entities.BookEnt, error)
	UpdateBook(context.Context, *entities.BookEnt) error
	DeleteBook(context.Context, uint) error
}

type BookRepositoryImpl struct {
}

func (br BookRepositoryImpl) CreateBook(ctx context.Context, ent *entities.BookEnt) error {
	//TODO implement me
	app_logger.Logger.Error(ent.Content)
	//db, _ := repositories.GetTx(ctx)
	//if result := db.Create(ent); result.Error != nil {
	//	return &repositories.RdbRuntimeError{
	//		ErrMsg:        fmt.Sprintf("[infrastructure.data.repositories.persistence.CreateBook] fail to insert Book to Database"),
	//		OriginalError: result.Error,
	//	}
	//}
	return nil
}

func (br BookRepositoryImpl) FindBook(ctx context.Context, ID uint) (*entities.BookEnt, error) {
	//TODO implement me
	db := repositories.GetConn()
	bookEnt := &entities.BookEnt{}
	if err := db.First(bookEnt, "id=?", ID); err != nil {
		return nil, &repositories.NotFoundError{
			Msg:           repositories.DefaultNotFoundMsg,
			ErrMsg:        fmt.Sprintf("[infrastructure.data.repositories.persistence.FindBook] failed to find bookEnt from rdb. ID : %d", ID),
			OriginalError: nil,
		}
	}
	return bookEnt, nil
}

func (br BookRepositoryImpl) FindBooks(ctx context.Context) ([]*entities.BookEnt, error) {
	//TODO implement me
	db := repositories.GetConn()
	var books []*entities.BookEnt
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

func (br BookRepositoryImpl) UpdateBook(ctx context.Context, requestBook *entities.BookEnt) error {
	//TODO implement me
	db, _ := repositories.GetTx(ctx)
	updatedAt := requestBook.UpdatedAt
	result := db.Where("updated_at <= ?", updatedAt).Model(&requestBook).Updates(requestBook)
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
	result := db.Where("id = ?", ID).Delete(&entities.BookEnt{})
	if result.Error != nil {
		return &repositories.RdbRuntimeError{
			ErrMsg:        fmt.Sprintf("[infrastructure.persistence.DeleteBook] failed to delete Book in RDB. ID : %d", ID),
			OriginalError: result.Error,
		}
	} else if result.RowsAffected == 0 {
		return &repositories.NotFoundError{
			Msg:           "削除対象の本が見つかりませんでした。その本は既に削除されている可能性があります。",
			ErrMsg:        fmt.Sprintf("[infrastructure.persistence.DeleteBook] failed to delete Book in RDB. Record not found. ID : %d", ID),
			OriginalError: nil,
		}
	}

	return nil
}

func CreateBookRepository() IBookRepository {
	return &BookRepositoryImpl{}
}
