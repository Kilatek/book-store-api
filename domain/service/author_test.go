package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"bookstore.com/domain/entity"
	"bookstore.com/port/payload"
	"bookstore.com/repository"
	"bookstore.com/test"
	"go.uber.org/mock/gomock"
)

func Test_authorService_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name       string
		authorRepo func() repository.AuthorRepository
		id         string
		want       *payload.AuthorResponse
		wantErr    bool
	}{
		{
			name: "find author succesfully",
			id:   test.AuthorId1,
			want: &payload.AuthorResponse{
				Id:          test.AuthorId1,
				FirstName:   test.AuthorFirstName1,
				LastName:    test.AuthorLastName1,
				BirthDate:   test.AuthorBirthDate1,
				Nationality: test.AuthorNationality1,
				CreatedAt:   test.CreatedAtStr,
				UpdatedAt:   test.UpdatedAtStr,
			},
			wantErr: false,
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
		},
		{
			name:    "id empty",
			id:      "",
			want:    nil,
			wantErr: true,
			authorRepo: func() repository.AuthorRepository {
				authorRepo := repository.NewMockAuthorRepository(ctrl)
				return authorRepo
			},
		},
		{
			name:    "find author failed",
			id:      test.AuthorId1,
			want:    nil,
			wantErr: true,
			authorRepo: func() repository.AuthorRepository {
				authorRepo := repository.NewMockAuthorRepository(ctrl)
				authorRepo.EXPECT().Find(gomock.Any(), test.AuthorId1).Return(nil, errors.New("error occur"))

				return authorRepo
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewAuthorService(tt.authorRepo(), nil)
			got, err := s.Find(context.TODO(), tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("authorService.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authorService.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authorService_Store(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name       string
		AuthorRepo func() repository.AuthorRepository
		req        *payload.AuthorRequest
		wantErr    bool
	}{
		{
			name: "store author successfully",
			req: &payload.AuthorRequest{
				FirstName:   test.AuthorFirstName1,
				LastName:    test.AuthorLastName1,
				BirthDate:   test.AuthorBirthDate1,
				Nationality: test.AuthorNationality1,
			},
			wantErr: false,
			AuthorRepo: func() repository.AuthorRepository {
				authorRepo := repository.NewMockAuthorRepository(ctrl)
				authorRepo.EXPECT().Store(gomock.Any(), &entity.Author{
					FirstName:   test.AuthorFirstName1,
					LastName:    test.AuthorLastName1,
					BirthDate:   test.AuthorBirthDate1,
					Nationality: test.AuthorNationality1,
				}).Return(nil)

				return authorRepo
			},
		},
		{
			name: "firstname empty",
			req: &payload.AuthorRequest{
				LastName:    test.AuthorLastName1,
				BirthDate:   test.AuthorBirthDate1,
				Nationality: test.AuthorNationality1,
			},
			wantErr: true,
			AuthorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
		},
		{
			name: "lastname empty",
			req: &payload.AuthorRequest{
				FirstName:   test.AuthorFirstName1,
				BirthDate:   test.AuthorBirthDate1,
				Nationality: test.AuthorNationality1,
			},
			wantErr: true,
			AuthorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
		},
		{
			name: "birthdate empty",
			req: &payload.AuthorRequest{
				FirstName:   test.AuthorFirstName1,
				LastName:    test.AuthorLastName1,
				Nationality: test.AuthorNationality1,
			},
			wantErr: true,
			AuthorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
		},
		{
			name: "nationality empty",
			req: &payload.AuthorRequest{
				FirstName: test.AuthorFirstName1,
				LastName:  test.AuthorLastName1,
				BirthDate: test.AuthorBirthDate1,
			},
			wantErr: true,
			AuthorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
		},
		{
			name: "store author failed",
			req: &payload.AuthorRequest{
				FirstName:   test.AuthorFirstName1,
				LastName:    test.AuthorLastName1,
				BirthDate:   test.AuthorBirthDate1,
				Nationality: test.AuthorNationality1,
			},
			wantErr: true,
			AuthorRepo: func() repository.AuthorRepository {
				authorRepo := repository.NewMockAuthorRepository(ctrl)
				authorRepo.EXPECT().Store(gomock.Any(), &entity.Author{
					FirstName:   test.AuthorFirstName1,
					LastName:    test.AuthorLastName1,
					BirthDate:   test.AuthorBirthDate1,
					Nationality: test.AuthorNationality1,
				}).Return(errors.New("error occur"))

				return authorRepo
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewAuthorService(tt.AuthorRepo(), nil)
			if err := s.Store(context.TODO(), tt.req); (err != nil) != tt.wantErr {
				t.Errorf("authorService.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_authorService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name       string
		AuthorRepo func() repository.AuthorRepository
		id         string
		req        *payload.AuthorRequest
		wantErr    bool
	}{
		{
			name: "update author successfully",
			id:   test.AuthorId1,
			req: &payload.AuthorRequest{
				FirstName:   test.AuthorFirstName1,
				LastName:    test.AuthorLastName1,
				BirthDate:   test.AuthorBirthDate1,
				Nationality: test.AuthorNationality1,
			},
			wantErr: false,
			AuthorRepo: func() repository.AuthorRepository {
				authorRepo := repository.NewMockAuthorRepository(ctrl)
				authorRepo.EXPECT().Find(gomock.Any(), test.AuthorId1).Return(&entity.Author{}, nil)
				authorRepo.EXPECT().Update(gomock.Any(), &entity.Author{
					Id:          test.AuthorId1,
					FirstName:   test.AuthorFirstName1,
					LastName:    test.AuthorLastName1,
					BirthDate:   test.AuthorBirthDate1,
					Nationality: test.AuthorNationality1,
				}).Return(nil)

				return authorRepo
			},
		},
		{
			name: "firstname empty",
			id:   test.AuthorId1,
			req: &payload.AuthorRequest{
				LastName:    test.AuthorLastName1,
				BirthDate:   test.AuthorBirthDate1,
				Nationality: test.AuthorNationality1,
			},
			wantErr: true,
			AuthorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
		},
		{
			name: "lastname empty",
			id:   test.AuthorId1,
			req: &payload.AuthorRequest{
				FirstName:   test.AuthorFirstName1,
				BirthDate:   test.AuthorBirthDate1,
				Nationality: test.AuthorNationality1,
			},
			wantErr: true,
			AuthorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
		},
		{
			name: "birthdate empty",
			id:   test.AuthorId1,
			req: &payload.AuthorRequest{
				FirstName:   test.AuthorFirstName1,
				LastName:    test.AuthorLastName1,
				Nationality: test.AuthorNationality1,
			},
			wantErr: true,
			AuthorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
		},
		{
			name: "nationality empty",
			id:   test.AuthorId1,
			req: &payload.AuthorRequest{
				FirstName: test.AuthorFirstName1,
				LastName:  test.AuthorLastName1,
				BirthDate: test.AuthorBirthDate1,
			},
			wantErr: true,
			AuthorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
		},
		{
			name: "update author failed",
			id:   test.AuthorId1,
			req: &payload.AuthorRequest{
				FirstName:   test.AuthorFirstName1,
				LastName:    test.AuthorLastName1,
				BirthDate:   test.AuthorBirthDate1,
				Nationality: test.AuthorNationality1,
			},
			wantErr: true,
			AuthorRepo: func() repository.AuthorRepository {
				authorRepo := repository.NewMockAuthorRepository(ctrl)
				authorRepo.EXPECT().Find(gomock.Any(), test.AuthorId1).Return(&entity.Author{}, nil)
				authorRepo.EXPECT().Update(gomock.Any(), &entity.Author{
					Id:          test.AuthorId1,
					FirstName:   test.AuthorFirstName1,
					LastName:    test.AuthorLastName1,
					BirthDate:   test.AuthorBirthDate1,
					Nationality: test.AuthorNationality1,
				}).Return(errors.New("error occur"))

				return authorRepo
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewAuthorService(tt.AuthorRepo(), nil)
			if err := s.Update(context.TODO(), tt.id, tt.req); (err != nil) != tt.wantErr {
				t.Errorf("authorService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_authorService_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name       string
		AuthorRepo func() repository.AuthorRepository
		want       []*payload.AuthorResponse
		wantErr    bool
	}{
		{
			name:    "find all authors successfully",
			wantErr: false,
			want: []*payload.AuthorResponse{
				{
					Id:          test.AuthorId1,
					FirstName:   test.AuthorFirstName1,
					LastName:    test.AuthorLastName1,
					BirthDate:   test.AuthorBirthDate1,
					Nationality: test.AuthorNationality1,
					CreatedAt:   test.CreatedAtStr,
					UpdatedAt:   test.UpdatedAtStr,
				},
			},
			AuthorRepo: func() repository.AuthorRepository {
				authorRepo := repository.NewMockAuthorRepository(ctrl)
				authorRepo.EXPECT().FindAll(gomock.Any()).Return([]*entity.Author{
					{
						Id:          test.AuthorId1,
						FirstName:   test.AuthorFirstName1,
						LastName:    test.AuthorLastName1,
						BirthDate:   test.AuthorBirthDate1,
						Nationality: test.AuthorNationality1,
						CreatedAt:   test.CreatedAt,
						UpdatedAt:   test.CreatedAt,
					},
				}, nil)

				return authorRepo
			},
		},
		{
			name:    "find all authors failed",
			wantErr: true,
			want:    nil,
			AuthorRepo: func() repository.AuthorRepository {
				authorRepo := repository.NewMockAuthorRepository(ctrl)
				authorRepo.EXPECT().FindAll(gomock.Any()).Return(nil, errors.New("error occured"))

				return authorRepo
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewAuthorService(tt.AuthorRepo(), nil)
			got, err := s.FindAll(context.TODO())
			if (err != nil) != tt.wantErr {
				t.Errorf("authorService.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authorService.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authorService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name       string
		AuthorRepo func() repository.AuthorRepository
		id         string
		wantErr    bool
	}{
		{
			name: "delete author successfully",
			id:   test.AuthorId1,
			AuthorRepo: func() repository.AuthorRepository {
				authorRepo := repository.NewMockAuthorRepository(ctrl)
				authorRepo.EXPECT().Find(gomock.Any(), test.AuthorId1).Return(&entity.Author{}, nil)
				authorRepo.EXPECT().Delete(gomock.Any(), test.AuthorId1).Return(nil)

				return authorRepo
			},
		},
		{
			name:    "id empty",
			id:      "",
			wantErr: true,
			AuthorRepo: func() repository.AuthorRepository {
				return repository.NewMockAuthorRepository(ctrl)
			},
		},
		{
			name:    "delete author failed",
			id:      test.AuthorId1,
			wantErr: true,
			AuthorRepo: func() repository.AuthorRepository {
				authorRepo := repository.NewMockAuthorRepository(ctrl)
				authorRepo.EXPECT().Find(gomock.Any(), test.AuthorId1).Return(&entity.Author{}, nil)
				authorRepo.EXPECT().Delete(gomock.Any(), test.AuthorId1).Return(errors.New("error occur"))

				return authorRepo
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewAuthorService(tt.AuthorRepo(), nil)
			if err := s.Delete(context.TODO(), tt.id); (err != nil) != tt.wantErr {
				t.Errorf("authorService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
