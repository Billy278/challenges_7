package repository

import (
	model "challenges_7/module/model/book"
	"context"
)

type BookRepository interface {
	CreateBook(ctx context.Context, bookIn model.Book) (book model.Book, err error)
	UpdateBook(ctx context.Context, bookIn model.Book) (book model.Book, err error)
	FindByIdBook(ctx context.Context, idIn uint64) (book model.Book, err error)
	FindAllBook(ctx context.Context) (books []model.Book, err error)
	DeleteBook(ctx context.Context, bookId uint64) (book model.Book, err error)
}
