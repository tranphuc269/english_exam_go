package services

import (
	"context"
	"english_exam_go/domain/dtos/book"
	"english_exam_go/infrastructure/data/repositories/persistence"
	"english_exam_go/utils/transaction"
)

type IBookService interface {
	CreateBook(ctx context.Context, book *book.RequestBook) error
	FindBook(context.Context, uint) (*book.ResponseBook, error)
	FindBooks(ctx context.Context) ([]*book.ResponseBook, error)
	UpdateBook(context.Context, *book.RequestBook, uint) error
	DeleteBook(context.Context, uint) error
}

type BookServiceImpl struct {
	br persistence.IBookRepository
	tx transaction.Transaction
}

func (bs BookServiceImpl) CreateBook(ctx context.Context, request *book.RequestBook) error {
	//TODO implement me
	//id, err := bs.tx.Required(ctx, func(ctx context.Context) (interface{}, error) {
	//	newBookEnt, _ := entities.NewBookEnt(request)
	//})
	newBookEnt, _ := book.NewBookEnt(request)
	err := bs.br.CreateBook(ctx, &newBookEnt)
	return err
}

func (bs BookServiceImpl) FindBook(ctx context.Context, ID uint) (*book.ResponseBook, error) {
	//TODO implement me
	ebook, err := bs.br.FindBook(ctx, ID)
	if err != nil {
		return nil, err
	}
	return book.CreateResponseBook(*ebook), nil
}

func (bs BookServiceImpl) FindBooks(ctx context.Context) ([]*book.ResponseBook, error) {
	//TODO implement me
	ebooks, err := bs.br.FindBooks(ctx)
	if err != nil {
		return nil, err
	}
	var responseBooks []*book.ResponseBook
	for _, b := range ebooks {
		responseBooks = append(responseBooks, book.CreateResponseBook(*b))
	}
	return responseBooks, nil
}

func (bs BookServiceImpl) UpdateBook(ctx context.Context, request *book.RequestBook, ID uint) error {
	//TODO implement me
	ebook, err := bs.br.FindBook(ctx, ID)
	if err != nil {
		return err
	}
	ebook.Content = request.Content
	ebook.Title = request.Title
	err = bs.br.UpdateBook(ctx, ebook)
	return err
}

func (bs BookServiceImpl) DeleteBook(ctx context.Context, ID uint) error {
	//TODO implement me
	err := bs.br.DeleteBook(ctx, ID)
	return err
}

func CreateBookService(br persistence.IBookRepository, tx transaction.Transaction) IBookService {
	return &BookServiceImpl{
		br: br,
		tx: tx,
	}
}
