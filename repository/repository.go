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
