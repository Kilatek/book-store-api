package service

import (
	"context"

	"bookstore.com/port/payload"
)

type AuthorService interface {
	Find(ctx context.Context, id string) (*payload.AuthorResponse, error)
	Store(ctx context.Context, author *payload.AuthorRequest) error
	Update(ctx context.Context, id string, author *payload.AuthorRequest) error
	FindAll(ctx context.Context) ([]*payload.AuthorResponse, error)
	Delete(ctx context.Context, id string) error
}

type BookService interface {
	Find(ctx context.Context, id string) (*payload.BookResponse, error)
	Store(ctx context.Context, author *payload.BookRequest) error
	Update(ctx context.Context, id string, author *payload.BookRequest) error
	FindAll(ctx context.Context) ([]*payload.BookResponse, error)
	Delete(ctx context.Context, id string) error
}
