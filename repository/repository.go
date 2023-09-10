package repository

import (
	"context"

	"bookstore.com/domain/entity"
)

type AuthorRepository interface {
	Find(ctx context.Context, id string) (*entity.Author, error)
	Store(ctx context.Context, author *entity.Author) error
	Update(ctx context.Context, author *entity.Author) error
	FindAll(ctx context.Context) ([]*entity.Author, error)
	Delete(ctx context.Context, id string) error
}

type BookRepository interface {
	Find(ctx context.Context, id string) (*entity.Book, error)
	Store(ctx context.Context, author *entity.Book) (*entity.Book, error)
	Update(ctx context.Context, author *entity.Book) error
	FindAll(ctx context.Context) ([]*entity.Book, error)
	Delete(ctx context.Context, id string) error
}

type UserRepository interface {
	Find(ctx context.Context, username string) (*entity.User, error)
	Store(ctx context.Context, user *entity.User) error
}

type NotificationRepository interface {
	Store(ctx context.Context, book *entity.Book)
}
