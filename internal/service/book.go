package service

import "context"

type Book struct{}

func NewBook() *Book {
	return &Book{}
}

func (b *Book) Setup(ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}
