package service

import (
	"context"

	"bookstore.com/domain/entity"
	portError "bookstore.com/port/error"
	"bookstore.com/port/payload"
	"bookstore.com/repository"
	"bookstore.com/tools/mapper"
)

type authorService struct {
	authorRepo       repository.AuthorRepository
	notificationRepo repository.NotificationRepository
}

func NewAuthorService(authRepo repository.AuthorRepository, notificationRepo repository.NotificationRepository) AuthorService {
	return &authorService{authorRepo: authRepo, notificationRepo: notificationRepo}
}

func (s *authorService) Find(ctx context.Context, id string) (*payload.AuthorResponse, error) {
	if id == "" {
		return nil, portError.NewBadRequestError("Id is empty.", nil)
	}

	author, err := s.authorRepo.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	res := &payload.AuthorResponse{}
	if err := mapper.MapStructsWithJSONTags(author, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *authorService) Store(ctx context.Context, req *payload.AuthorRequest) error {
	if err := req.Validate(); err != nil {
		return portError.NewBadRequestError(err.Error(), nil)
	}

	author := &entity.Author{}
	if err := mapper.MapStructsWithJSONTags(req, author); err != nil {
		return err
	}
	if s.notificationRepo != nil {
		s.notificationRepo.AddAction(ctx, "addAuthor")
	}
	return s.authorRepo.Store(ctx, author)

}
func (s *authorService) Update(ctx context.Context, id string, req *payload.AuthorRequest) error {
	if id == "" {
		return portError.NewBadRequestError("id is empty", nil)
	}

	if err := req.Validate(); err != nil {
		return portError.NewBadRequestError(err.Error(), nil)
	}

	author, err := s.authorRepo.Find(ctx, id)
	if err != nil {
		return err
	}

	if err := mapper.MapStructsWithJSONTags(req, author); err != nil {
		return err
	}

	author.Id = id

	if s.notificationRepo != nil {
		s.notificationRepo.AddAction(ctx, "updateAuthor")
	}
	return s.authorRepo.Update(ctx, author)
}

func (s *authorService) FindAll(ctx context.Context) ([]*payload.AuthorResponse, error) {
	authors, err := s.authorRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	list := []*payload.AuthorResponse{}
	for _, author := range authors {
		authorRes := &payload.AuthorResponse{}
		if err := mapper.MapStructsWithJSONTags(author, authorRes); err != nil {
			return nil, err
		}
		list = append(list, authorRes)
	}

	return list, nil
}

func (s *authorService) Delete(ctx context.Context, id string) error {
	_, err := s.Find(ctx, id)
	if err != nil {
		return err
	}

	if s.notificationRepo != nil {
		s.notificationRepo.AddAction(ctx, "deleteAuthor")
	}
	return s.authorRepo.Delete(ctx, id)
}
