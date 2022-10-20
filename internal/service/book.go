package service

import (
	"context"

	v1 "golayout/api/helloworld/v1"
	"golayout/internal/dao"
)

type Book struct {
	bookDao *dao.Book
}

func NewBook(bookDao *dao.Book) *Book {
	return &Book{
		bookDao: bookDao,
	}
}

func (b *Book) SayHello(ctx context.Context, request *v1.HelloRequest) (*v1.HelloReply, error) {
	// TODO implement me
	panic("implement me")
}
