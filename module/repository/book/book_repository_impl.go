package repository

import (
	"challenges_7/config"
	model "challenges_7/module/model/book"
	"context"
	"errors"
	"fmt"
)

type BookRepositoryImpl struct {
	DataStore config.DataStore
}

func NewBookRepositoryImpl(dt config.DataStore) BookRepository {
	return &BookRepositoryImpl{
		DataStore: dt,
	}
}

func (book_impl *BookRepositoryImpl) FindByIdBook(ctx context.Context, idIn uint64) (book model.Book, err error) {
	//find by id
	book, ok := book_impl.DataStore.DataBook[idIn]
	if !ok {
		return book, errors.New("NOT FOUND")
	}
	return book, err
}

func (book_impl *BookRepositoryImpl) CreateBook(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	//intsert
	bookIn.Id = uint64(len(book_impl.DataStore.DataBook) + 1)
	book_impl.DataStore.DataBook[bookIn.Id] = bookIn
	fmt.Println(book_impl.DataStore)
	return bookIn, err

}
func (book_impl *BookRepositoryImpl) UpdateBook(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	_, err = book_impl.FindByIdBook(ctx, bookIn.Id)
	if err != nil {
		return book, err
	}
	//update
	book_impl.DataStore.DataBook[bookIn.Id] = bookIn
	return bookIn, err

}

func (book_impl *BookRepositoryImpl) FindAllBook(ctx context.Context) (books []model.Book, err error) {
	//find all
	for _, v := range book_impl.DataStore.DataBook {
		books = append(books, v)
	}
	return books, err
}

func (book_impl *BookRepositoryImpl) DeleteBook(ctx context.Context, bookId uint64) (book model.Book, err error) {
	book, err = book_impl.FindByIdBook(ctx, bookId)
	if err != nil {
		return book, err
	}
	//delete
	delete(book_impl.DataStore.DataBook, bookId)
	return book, err
}
