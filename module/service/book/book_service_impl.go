package service

import (
	model "challenges_7/module/model/book"
	repository "challenges_7/module/repository/book"
	"context"
	"log"
)

type BookServiceImpl struct {
	RepoBook repository.BookRepository
}

func NewBookServiceImpl(rpbook repository.BookRepository) BookService {
	return &BookServiceImpl{
		RepoBook: rpbook,
	}
}

func (booksrv_imp *BookServiceImpl) CreateBookSrv(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	book, err = booksrv_imp.RepoBook.CreateBook(ctx, bookIn)
	if err != nil {
		log.Printf("[ERROR] error Insert Book :%v\n", err)
	}
	return book, err

}
func (booksrv_imp *BookServiceImpl) UpdateBookSrv(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	book, err = booksrv_imp.RepoBook.UpdateBook(ctx, bookIn)
	if err != nil {
		log.Printf("[ERROR] error Update Book :%v\n", err)
	}
	return book, err
}
func (booksrv_imp *BookServiceImpl) FindByIdBookSrv(ctx context.Context, idIn uint64) (book model.Book, err error) {
	book, err = booksrv_imp.RepoBook.FindByIdBook(ctx, idIn)
	if err != nil {
		log.Printf("[ERROR] error findbook Book :%v\n", err)
		return book, err
	}
	return book, err
}
func (booksrv_imp *BookServiceImpl) FindAllBookSrv(ctx context.Context) (books []model.Book, err error) {
	books, err = booksrv_imp.RepoBook.FindAllBook(ctx)
	if err != nil {
		log.Printf("[ERROR] error findbook Book :%v\n", err)
		return books, err
	}
	return books, err
}
func (booksrv_imp *BookServiceImpl) DeleteBookSrv(ctx context.Context, bookId uint64) (book model.Book, err error) {
	book, err = booksrv_imp.RepoBook.DeleteBook(ctx, bookId)
	if err != nil {
		log.Printf("[ERROR] error deletebook Book :%v\n", err)
		return book, err
	}
	return book, err
}
