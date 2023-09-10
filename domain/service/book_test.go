package service

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"bookstore.com/domain/entity"
	"bookstore.com/port/payload"
	"bookstore.com/repository"
	"bookstore.com/test"
	"go.uber.org/mock/gomock"
)

func Test_bookService_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name       string
		bookRepo   func() repository.BookRepository
		authorRepo func() repository.AuthorRepository
		id         string
		want       *payload.BookResponse
		wantErr    bool
	}{
		{
			name: "find book successfully",
			bookRepo: func() repository.BookRepository {
				bookRepo := repository.NewMockBookRepository(ctrl)
				bookRepo.EXPECT().Find(gomock.Any(), test.BookId1).Return(&entity.Book{
					Id:       test.BookId1,
					AuthorId: test.AuthorId1,
					Author: &entity.Author{
						Id:          test.AuthorId1,
						FirstName:   test.AuthorFirstName1,
						LastName:    test.AuthorLastName1,
						BirthDate:   test.AuthorBirthDate1,
						Nationality: test.AuthorNationality1,
						CreatedAt:   test.CreatedAt,
						UpdatedAt:   test.UpdatedAt,
					},
					Name:            test.BookName1,
					Description:     test.BookDescription1,
					PublicationDate: test.PublicationDate1,
					Price:           test.Price1,
					CreatedAt:       test.CreatedAt,
					UpdatedAt:       test.UpdatedAt,
				}, nil)

				return bookRepo
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			id: test.BookId1,
			want: &payload.BookResponse{
				Id: test.BookId1,
				Author: &payload.AuthorResponse{
					Id:          test.AuthorId1,
					FirstName:   test.AuthorFirstName1,
					LastName:    test.AuthorLastName1,
					BirthDate:   test.AuthorBirthDate1,
					Nationality: test.AuthorNationality1,
					CreatedAt:   test.CreatedAtStr,
					UpdatedAt:   test.UpdatedAtStr,
				},
				Name:            test.BookName1,
				Description:     test.BookDescription1,
				PublicationDate: test.PublicationDate1,
				Price:           test.Price1,
				CreatedAt:       test.CreatedAtStr,
				UpdatedAt:       test.UpdatedAtStr,
			},
			wantErr: false,
		},
		{
			name: "find book failed",
			bookRepo: func() repository.BookRepository {
				bookRepo := repository.NewMockBookRepository(ctrl)
				bookRepo.EXPECT().Find(gomock.Any(), test.BookId1).Return(nil, errors.New("error occur"))

				return bookRepo
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			id:      test.BookId1,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewBookService(tt.bookRepo(), tt.authorRepo(), nil)
			got, err := s.Find(context.TODO(), tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("bookService.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			rawGotData, _ := json.Marshal(got)
			rawWantData, _ := json.Marshal(tt.want)
			if !reflect.DeepEqual(rawGotData, rawWantData) {
				t.Errorf("bookService.Find() = %v, want %v", string(rawGotData), string(rawWantData))
			}
		})
	}
}

func Test_bookService_Store(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name       string
		bookRepo   func() repository.BookRepository
		authorRepo func() repository.AuthorRepository
		req        *payload.BookRequest
		wantErr    bool
	}{
		{
			name: "store book successfully",
			bookRepo: func() repository.BookRepository {
				bookRepo := repository.NewMockBookRepository(ctrl)
				bookRepo.EXPECT().Store(gomock.Any(), &entity.Book{
					AuthorId:        test.AuthorId1,
					Name:            test.BookName1,
					Description:     test.BookDescription1,
					PublicationDate: test.PublicationDate1,
					Price:           test.Price1,
				}).Return(nil)

				return bookRepo
			},
			authorRepo: func() repository.AuthorRepository {
				authorRepo := repository.NewMockAuthorRepository(ctrl)
				authorRepo.EXPECT().Find(gomock.Any(), test.AuthorId1).Return(&entity.Author{
					Id:          test.AuthorId1,
					FirstName:   test.AuthorFirstName1,
					LastName:    test.AuthorLastName1,
					BirthDate:   test.AuthorBirthDate1,
					Nationality: test.AuthorNationality1,
					CreatedAt:   test.CreatedAt,
					UpdatedAt:   test.UpdatedAt,
				}, nil)

				return authorRepo
			},
			req: &payload.BookRequest{
				AuthorId:        test.AuthorId1,
				Name:            test.BookName1,
				Description:     test.BookDescription1,
				PublicationDate: test.PublicationDate1,
				Price:           test.Price1,
			},
		},
		{
			name: "author not found",
			bookRepo: func() repository.BookRepository {
				return repository.NewMockBookRepository(ctrl)
			},
			authorRepo: func() repository.AuthorRepository {
				authorRepo := repository.NewMockAuthorRepository(ctrl)
				authorRepo.EXPECT().Find(gomock.Any(), test.AuthorId1).Return(nil, errors.New("author not found"))

				return authorRepo
			},
			req: &payload.BookRequest{
				AuthorId:        test.AuthorId1,
				Name:            test.BookName1,
				Description:     test.BookDescription1,
				PublicationDate: test.PublicationDate1,
				Price:           test.Price1,
			},
			wantErr: true,
		},
		{
			name: "author id empty",
			bookRepo: func() repository.BookRepository {
				return repository.NewMockBookRepository(ctrl)
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			req: &payload.BookRequest{
				Name:            test.BookName1,
				Description:     test.BookDescription1,
				PublicationDate: test.PublicationDate1,
				Price:           test.Price1,
			},
			wantErr: true,
		},
		{
			name: "name empty",
			bookRepo: func() repository.BookRepository {
				return repository.NewMockBookRepository(ctrl)
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			req: &payload.BookRequest{
				AuthorId:        test.AuthorId1,
				Description:     test.BookDescription1,
				PublicationDate: test.PublicationDate1,
				Price:           test.Price1,
			},
			wantErr: true,
		},
		{
			name: "description empty",
			bookRepo: func() repository.BookRepository {
				return repository.NewMockBookRepository(ctrl)
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			req: &payload.BookRequest{
				AuthorId:        test.AuthorId1,
				Name:            test.BookName1,
				PublicationDate: test.PublicationDate1,
				Price:           test.Price1,
			},
			wantErr: true,
		},
		{
			name: "publication date empty",
			bookRepo: func() repository.BookRepository {
				return repository.NewMockBookRepository(ctrl)
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			req: &payload.BookRequest{
				AuthorId:    test.AuthorId1,
				Name:        test.BookName1,
				Description: test.BookDescription1,
				Price:       test.Price1,
			},
			wantErr: true,
		},
		{
			name: "price equal to 0",
			bookRepo: func() repository.BookRepository {
				return repository.NewMockBookRepository(ctrl)
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			req: &payload.BookRequest{
				AuthorId:        test.AuthorId1,
				Name:            test.BookName1,
				Description:     test.BookDescription1,
				PublicationDate: test.PublicationDate1,
			},
			wantErr: true,
		},
		{
			name: "invalid price",
			bookRepo: func() repository.BookRepository {
				return repository.NewMockBookRepository(ctrl)
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			req: &payload.BookRequest{
				AuthorId:        test.AuthorId1,
				Name:            test.BookName1,
				Description:     test.BookDescription1,
				PublicationDate: test.PublicationDate1,
				Price:           test.InvalidPrice1,
			},
			wantErr: true,
		},
		{
			name: "store book failed",
			bookRepo: func() repository.BookRepository {
				bookRepo := repository.NewMockBookRepository(ctrl)
				bookRepo.EXPECT().Store(gomock.Any(), &entity.Book{
					AuthorId:        test.AuthorId1,
					Name:            test.BookName1,
					Description:     test.BookDescription1,
					PublicationDate: test.PublicationDate1,
					Price:           test.Price1,
				}).Return(errors.New("error occur"))

				return bookRepo
			},
			authorRepo: func() repository.AuthorRepository {
				authorRepo := repository.NewMockAuthorRepository(ctrl)
				authorRepo.EXPECT().Find(gomock.Any(), test.AuthorId1).Return(&entity.Author{
					Id:          test.AuthorId1,
					FirstName:   test.AuthorFirstName1,
					LastName:    test.AuthorLastName1,
					BirthDate:   test.AuthorBirthDate1,
					Nationality: test.AuthorNationality1,
					CreatedAt:   test.CreatedAt,
					UpdatedAt:   test.UpdatedAt,
				}, nil)

				return authorRepo
			},
			req: &payload.BookRequest{
				AuthorId:        test.AuthorId1,
				Name:            test.BookName1,
				Description:     test.BookDescription1,
				PublicationDate: test.PublicationDate1,
				Price:           test.Price1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewBookService(tt.bookRepo(), tt.authorRepo(), nil)
			if err := s.Store(context.TODO(), tt.req); (err != nil) != tt.wantErr {
				t.Errorf("bookService.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_bookService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name       string
		bookRepo   func() repository.BookRepository
		authorRepo func() repository.AuthorRepository
		id         string
		req        *payload.BookRequest
		wantErr    bool
	}{
		{
			name: "update book successfully",
			bookRepo: func() repository.BookRepository {
				bookRepo := repository.NewMockBookRepository(ctrl)
				bookRepo.EXPECT().Find(gomock.Any(), test.BookId1).Return(&entity.Book{}, nil)
				bookRepo.EXPECT().Update(gomock.Any(), &entity.Book{
					Id:              test.BookId1,
					AuthorId:        test.AuthorId1,
					Name:            test.BookName1,
					Description:     test.BookDescription1,
					PublicationDate: test.PublicationDate1,
					Price:           test.Price1,
				}).Return(nil)

				return bookRepo
			},
			authorRepo: func() repository.AuthorRepository {
				authorRepo := repository.NewMockAuthorRepository(ctrl)
				authorRepo.EXPECT().Find(gomock.Any(), test.AuthorId1).Return(&entity.Author{
					Id:          test.AuthorId1,
					FirstName:   test.AuthorFirstName1,
					LastName:    test.AuthorLastName1,
					BirthDate:   test.AuthorBirthDate1,
					Nationality: test.AuthorNationality1,
					CreatedAt:   test.CreatedAt,
					UpdatedAt:   test.UpdatedAt,
				}, nil)

				return authorRepo
			},
			id: test.BookId1,
			req: &payload.BookRequest{
				AuthorId:        test.AuthorId1,
				Name:            test.BookName1,
				Description:     test.BookDescription1,
				PublicationDate: test.PublicationDate1,
				Price:           test.Price1,
			},
		},
		{
			name: "id empty",
			bookRepo: func() repository.BookRepository {
				return repository.NewMockBookRepository(ctrl)
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			id: "",
			req: &payload.BookRequest{
				AuthorId:        test.AuthorId1,
				Name:            test.BookName1,
				Description:     test.BookDescription1,
				PublicationDate: test.PublicationDate1,
				Price:           test.Price1,
			},
			wantErr: true,
		},
		{
			name: "author id empty",
			bookRepo: func() repository.BookRepository {
				return repository.NewMockBookRepository(ctrl)
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			id: test.BookId1,
			req: &payload.BookRequest{
				Name:            test.BookName1,
				Description:     test.BookDescription1,
				PublicationDate: test.PublicationDate1,
				Price:           test.Price1,
			},
			wantErr: true,
		},
		{
			name: "name empty",
			bookRepo: func() repository.BookRepository {
				return repository.NewMockBookRepository(ctrl)
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			id: test.BookId1,
			req: &payload.BookRequest{
				AuthorId:        test.AuthorId1,
				Description:     test.BookDescription1,
				PublicationDate: test.PublicationDate1,
				Price:           test.Price1,
			},
			wantErr: true,
		},
		{
			name: "description id empty",
			bookRepo: func() repository.BookRepository {
				return repository.NewMockBookRepository(ctrl)
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			id: test.BookId1,
			req: &payload.BookRequest{
				AuthorId:        test.AuthorId1,
				Name:            test.BookName1,
				PublicationDate: test.PublicationDate1,
				Price:           test.Price1,
			},
			wantErr: true,
		},
		{
			name: "publication date empty",
			bookRepo: func() repository.BookRepository {
				return repository.NewMockBookRepository(ctrl)
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			id: test.BookId1,
			req: &payload.BookRequest{
				AuthorId:    test.AuthorId1,
				Name:        test.BookName1,
				Description: test.BookDescription1,
				Price:       test.Price1,
			},
			wantErr: true,
		},
		{
			name: "price equal to 0",
			bookRepo: func() repository.BookRepository {
				return repository.NewMockBookRepository(ctrl)
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			id: test.BookId1,
			req: &payload.BookRequest{
				AuthorId:        test.AuthorId1,
				Name:            test.BookName1,
				Description:     test.BookDescription1,
				PublicationDate: test.PublicationDate1,
			},
			wantErr: true,
		},
		{
			name: "invalid price",
			bookRepo: func() repository.BookRepository {
				return repository.NewMockBookRepository(ctrl)
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			id: test.BookId1,
			req: &payload.BookRequest{
				AuthorId:        test.AuthorId1,
				Name:            test.BookName1,
				Description:     test.BookDescription1,
				PublicationDate: test.PublicationDate1,
				Price:           test.InvalidPrice1,
			},
			wantErr: true,
		},
		{
			name: "find book failed",
			bookRepo: func() repository.BookRepository {
				bookRepo := repository.NewMockBookRepository(ctrl)
				bookRepo.EXPECT().Find(gomock.Any(), test.BookId1).Return(nil, errors.New("error occur"))

				return bookRepo
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			id: test.BookId1,
			req: &payload.BookRequest{
				AuthorId:        test.AuthorId1,
				Name:            test.BookName1,
				Description:     test.BookDescription1,
				PublicationDate: test.PublicationDate1,
				Price:           test.Price1,
			},
			wantErr: true,
		},
		{
			name: "find author failed",
			bookRepo: func() repository.BookRepository {
				bookRepo := repository.NewMockBookRepository(ctrl)
				bookRepo.EXPECT().Find(gomock.Any(), test.BookId1).Return(&entity.Book{}, nil)

				return bookRepo
			},
			authorRepo: func() repository.AuthorRepository {
				authorRepo := repository.NewMockAuthorRepository(ctrl)
				authorRepo.EXPECT().Find(gomock.Any(), test.AuthorId1).Return(nil, errors.New("error occur"))

				return authorRepo
			},
			id: test.BookId1,
			req: &payload.BookRequest{
				AuthorId:        test.AuthorId1,
				Name:            test.BookName1,
				Description:     test.BookDescription1,
				PublicationDate: test.PublicationDate1,
				Price:           test.Price1,
			},
			wantErr: true,
		},
		{
			name: "update book failed",
			bookRepo: func() repository.BookRepository {
				bookRepo := repository.NewMockBookRepository(ctrl)
				bookRepo.EXPECT().Find(gomock.Any(), test.BookId1).Return(&entity.Book{}, nil)
				bookRepo.EXPECT().Update(gomock.Any(), &entity.Book{
					Id:              test.BookId1,
					AuthorId:        test.AuthorId1,
					Name:            test.BookName1,
					Description:     test.BookDescription1,
					PublicationDate: test.PublicationDate1,
					Price:           test.Price1,
				}).Return(errors.New("error occur"))

				return bookRepo
			},
			authorRepo: func() repository.AuthorRepository {
				authorRepo := repository.NewMockAuthorRepository(ctrl)
				authorRepo.EXPECT().Find(gomock.Any(), test.AuthorId1).Return(&entity.Author{
					Id:          test.AuthorId1,
					FirstName:   test.AuthorFirstName1,
					LastName:    test.AuthorLastName1,
					BirthDate:   test.AuthorBirthDate1,
					Nationality: test.AuthorNationality1,
					CreatedAt:   test.CreatedAt,
					UpdatedAt:   test.UpdatedAt,
				}, nil)

				return authorRepo
			},
			id: test.BookId1,
			req: &payload.BookRequest{
				AuthorId:        test.AuthorId1,
				Name:            test.BookName1,
				Description:     test.BookDescription1,
				PublicationDate: test.PublicationDate1,
				Price:           test.Price1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewBookService(tt.bookRepo(), tt.authorRepo(), nil)
			if err := s.Update(context.TODO(), tt.id, tt.req); (err != nil) != tt.wantErr {
				t.Errorf("bookService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_bookService_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name       string
		bookRepo   func() repository.BookRepository
		authorRepo func() repository.AuthorRepository
		want       []*payload.BookResponse
		wantErr    bool
	}{
		{
			name: "find all books successfully",
			bookRepo: func() repository.BookRepository {
				bookRepo := repository.NewMockBookRepository(ctrl)
				bookRepo.EXPECT().FindAll(gomock.Any()).Return([]*entity.Book{
					{
						Id:       test.BookId1,
						AuthorId: test.AuthorId1,
						Author: &entity.Author{
							Id:          test.AuthorId1,
							FirstName:   test.AuthorFirstName1,
							LastName:    test.AuthorLastName1,
							BirthDate:   test.AuthorBirthDate1,
							Nationality: test.AuthorNationality1,
							CreatedAt:   test.CreatedAt,
							UpdatedAt:   test.UpdatedAt,
						},
						Name:            test.BookName1,
						Description:     test.BookDescription1,
						PublicationDate: test.PublicationDate1,
						Price:           test.Price1,
						CreatedAt:       test.CreatedAt,
						UpdatedAt:       test.UpdatedAt,
					},
				}, nil)

				return bookRepo
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			want: []*payload.BookResponse{
				{
					Id: test.BookId1,
					Author: &payload.AuthorResponse{
						Id:          test.AuthorId1,
						FirstName:   test.AuthorFirstName1,
						LastName:    test.AuthorLastName1,
						BirthDate:   test.AuthorBirthDate1,
						Nationality: test.AuthorNationality1,
						CreatedAt:   test.CreatedAtStr,
						UpdatedAt:   test.UpdatedAtStr,
					},
					Name:            test.BookName1,
					Description:     test.BookDescription1,
					PublicationDate: test.PublicationDate1,
					Price:           test.Price1,
					CreatedAt:       test.CreatedAtStr,
					UpdatedAt:       test.UpdatedAtStr,
				},
			},
		},
		{
			name: "find all books failed",
			bookRepo: func() repository.BookRepository {
				bookRepo := repository.NewMockBookRepository(ctrl)
				bookRepo.EXPECT().FindAll(gomock.Any()).Return(nil, errors.New("error occur"))

				return bookRepo
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewBookService(tt.bookRepo(), tt.authorRepo(), nil)
			got, err := s.FindAll(context.TODO())
			if (err != nil) != tt.wantErr {
				t.Errorf("bookService.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bookService.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name       string
		bookRepo   func() repository.BookRepository
		authorRepo func() repository.AuthorRepository
		id         string
		wantErr    bool
	}{
		{
			name: "delete book successfully",
			bookRepo: func() repository.BookRepository {
				bookRepo := repository.NewMockBookRepository(ctrl)
				bookRepo.EXPECT().Find(gomock.Any(), test.BookId1).Return(&entity.Book{}, nil)
				bookRepo.EXPECT().Delete(gomock.Any(), test.BookId1).Return(nil)

				return bookRepo
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			id: test.BookId1,
		},
		{
			name: "find book failed",
			bookRepo: func() repository.BookRepository {
				bookRepo := repository.NewMockBookRepository(ctrl)
				bookRepo.EXPECT().Find(gomock.Any(), test.BookId1).Return(&entity.Book{}, errors.New("error occur"))
				return bookRepo
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			id:      test.BookId1,
			wantErr: true,
		},
		{
			name: "delete book failed",
			bookRepo: func() repository.BookRepository {
				bookRepo := repository.NewMockBookRepository(ctrl)
				bookRepo.EXPECT().Find(gomock.Any(), test.BookId1).Return(&entity.Book{}, nil)
				bookRepo.EXPECT().Delete(gomock.Any(), test.BookId1).Return(errors.New("error occur"))

				return bookRepo
			},
			authorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
			id:      test.BookId1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewBookService(tt.bookRepo(), tt.authorRepo(), nil)
			if err := s.Delete(context.TODO(), tt.id); (err != nil) != tt.wantErr {
				t.Errorf("bookService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
