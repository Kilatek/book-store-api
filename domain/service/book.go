package service

import (
	"context"

	"bookstore.com/domain/entity"
	portError "bookstore.com/port/error"
	"bookstore.com/port/payload"
	"bookstore.com/repository"
	"bookstore.com/tools/mapper"
)

type bookService struct {
	bookRepo         repository.BookRepository
	authorRepo       repository.AuthorRepository
	notificationRepo repository.NotificationRepository
}

func NewBookService(
	bookRepo repository.BookRepository,
	authorRepo repository.AuthorRepository,
	notificationRepo repository.NotificationRepository,
) BookService {
	return &bookService{bookRepo: bookRepo, authorRepo: authorRepo, notificationRepo: notificationRepo}
}

func (s *bookService) Find(ctx context.Context, id string) (*payload.BookResponse, error) {
	if id == "" {
		return nil, portError.NewBadRequestError("id is empty", nil)
	}

	book, err := s.bookRepo.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	res := &payload.BookResponse{}
	if err := mapper.MapStructsWithJSONTags(book, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *bookService) Store(ctx context.Context, req *payload.BookRequest) error {
	if err := req.Validate(); err != nil {
		return portError.NewBadRequestError(err.Error(), nil)
	}

	_, err := s.authorRepo.Find(ctx, req.AuthorId)
	if err != nil {
		return err
	}

	book := &entity.Book{}
	if err := mapper.MapStructsWithJSONTags(req, book); err != nil {
		return err
	}

	newBook, err := s.bookRepo.Store(ctx, book)
	if s.notificationRepo != nil && newBook != nil {
		s.notificationRepo.Store(ctx, newBook)
	}

	return err
}
func (s *bookService) Update(ctx context.Context, id string, req *payload.BookRequest) error {
	if id == "" {
		return portError.NewBadRequestError("id is empty", nil)
	}

	if err := req.Validate(); err != nil {
		return portError.NewBadRequestError(err.Error(), nil)
	}

	book, err := s.bookRepo.Find(ctx, id)
	if err != nil {
		return err
	}

	_, err = s.authorRepo.Find(ctx, req.AuthorId)
	if err != nil {
		return err
	}

	if err := mapper.MapStructsWithJSONTags(req, book); err != nil {
		return err
	}

	book.Id = id

	return s.bookRepo.Update(ctx, book)
}

func (s *bookService) FindAll(ctx context.Context) ([]*payload.BookResponse, error) {
	books, err := s.bookRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	list := []*payload.BookResponse{}
	for _, book := range books {
		bookRes := &payload.BookResponse{}
		if err := mapper.MapStructsWithJSONTags(book, bookRes); err != nil {
			return nil, err
		}
		list = append(list, bookRes)
	}

	return list, nil
}

func (s *bookService) Delete(ctx context.Context, id string) error {
	_, err := s.Find(ctx, id)
	if err != nil {
		return err
	}

	return s.bookRepo.Delete(ctx, id)
}
