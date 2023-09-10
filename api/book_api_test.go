package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"bookstore.com/domain/service"
	"bookstore.com/port/payload"
	"bookstore.com/test"
	"go.uber.org/mock/gomock"
)

func Test_bookHandler_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	books := []*payload.BookResponse{
		{
			Id: test.BookId1,
			Author: &payload.AuthorResponse{
				Id:          test.BookId1,
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
	}
	expectedBooksJson, _ := json.Marshal(books)

	tests := []struct {
		name           string
		bookService    func() service.BookService
		args           args
		expected       string
		expectedStatus int
	}{
		{
			name: "success to retrieve books",
			bookService: func() service.BookService {
				bookService := service.NewMockBookService(ctrl)
				bookService.EXPECT().FindAll(gomock.Any()).Return(books, nil)

				return bookService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/api/v1/books", nil),
			},
			expected:       string(expectedBooksJson),
			expectedStatus: http.StatusOK,
		},
		{
			name: "failed to retrieve books",
			bookService: func() service.BookService {
				bookService := service.NewMockBookService(ctrl)
				bookService.EXPECT().FindAll(gomock.Any()).Return(nil, errors.New("error occur"))

				return bookService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/api/v1/books", nil),
			},
			expected:       string(`{"message":"Some thing wrong with the server"}`),
			expectedStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewBookHandler(tt.bookService())
			h.GetAll(tt.args.w, tt.args.r)

			if tt.args.w.Body.String() != tt.expected {
				t.Errorf("Expected json response %s, got %s", tt.expected, tt.args.w.Body.String())
			}

			if tt.args.w.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, tt.args.w.Code)
			}
		})
	}

}

func Test_bookHandler_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	book := &payload.BookResponse{
		Id: test.BookId1,
		Author: &payload.AuthorResponse{
			Id:          test.BookId1,
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
	}
	expectedBookJson, _ := json.Marshal(book)

	tests := []struct {
		name           string
		bookService    func() service.BookService
		args           args
		expected       string
		expectedStatus int
	}{
		{
			name: "success to retrieve book",
			bookService: func() service.BookService {
				bookService := service.NewMockBookService(ctrl)
				bookService.EXPECT().Find(gomock.Any(), gomock.Any()).Return(book, nil)

				return bookService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/api/v1/books/1", nil),
			},
			expected:       string(expectedBookJson),
			expectedStatus: http.StatusOK,
		},
		{
			name: "failed to retrieve book",
			bookService: func() service.BookService {
				bookService := service.NewMockBookService(ctrl)
				bookService.EXPECT().Find(gomock.Any(), gomock.Any()).Return(nil, errors.New("error occur"))

				return bookService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/api/v1/books/1", nil),
			},
			expected:       string(`{"message":"Some thing wrong with the server"}`),
			expectedStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewBookHandler(tt.bookService())
			h.Get(tt.args.w, tt.args.r)

			if tt.args.w.Body.String() != tt.expected {
				t.Errorf("Expected json response %s, got %s", tt.expected, tt.args.w.Body.String())
			}

			if tt.args.w.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, tt.args.w.Code)
			}
		})
	}

}

func Test_bookHandler_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	book := &payload.BookRequest{
		AuthorId:        test.AuthorId1,
		Name:            test.BookName1,
		Description:     test.BookDescription1,
		PublicationDate: test.PublicationDate1,
		Price:           test.Price1,
	}
	bodyData, _ := json.Marshal(book)

	tests := []struct {
		name           string
		bookService    func() service.BookService
		args           args
		expected       string
		expectedStatus int
	}{
		{
			name: "success to create book",
			bookService: func() service.BookService {
				bookService := service.NewMockBookService(ctrl)
				bookService.EXPECT().Store(gomock.Any(), gomock.Any()).Return(nil)

				return bookService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/books", bytes.NewReader(bodyData)),
			},
			expected:       string(bodyData),
			expectedStatus: http.StatusOK,
		},
		{
			name: "failed to to create book",
			bookService: func() service.BookService {
				bookService := service.NewMockBookService(ctrl)
				bookService.EXPECT().Store(gomock.Any(), gomock.Any()).Return(errors.New("error occur"))

				return bookService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/api/v1/books", bytes.NewReader(bodyData)),
			},
			expected:       string(`{"message":"Some thing wrong with the server"}`),
			expectedStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewBookHandler(tt.bookService())
			h.Post(tt.args.w, tt.args.r)

			if tt.args.w.Body.String() != tt.expected {
				t.Errorf("Expected json response %s, got %s", tt.expected, tt.args.w.Body.String())
			}

			if tt.args.w.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, tt.args.w.Code)
			}
		})
	}

}

func Test_bookHandler_Put(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	book := &payload.BookRequest{
		AuthorId:        test.AuthorId1,
		Name:            test.BookName1,
		Description:     test.BookDescription1,
		PublicationDate: test.PublicationDate1,
		Price:           test.Price1,
	}
	bodyData, _ := json.Marshal(book)

	tests := []struct {
		name           string
		bookService    func() service.BookService
		args           args
		expected       string
		expectedStatus int
	}{
		{
			name: "success to update book",
			bookService: func() service.BookService {
				bookService := service.NewMockBookService(ctrl)
				bookService.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

				return bookService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("PUT", "/books", bytes.NewReader(bodyData)),
			},
			expected:       string(bodyData),
			expectedStatus: http.StatusOK,
		},
		{
			name: "failed to to update book",
			bookService: func() service.BookService {
				bookService := service.NewMockBookService(ctrl)
				bookService.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error occur"))

				return bookService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("PUT", "/api/v1/books", bytes.NewReader(bodyData)),
			},
			expected:       string(`{"message":"Some thing wrong with the server"}`),
			expectedStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewBookHandler(tt.bookService())
			h.Put(tt.args.w, tt.args.r)

			if tt.args.w.Body.String() != tt.expected {
				t.Errorf("Expected json response %s, got %s", tt.expected, tt.args.w.Body.String())
			}

			if tt.args.w.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, tt.args.w.Code)
			}
		})
	}

}

func Test_bookHandler_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	tests := []struct {
		name           string
		bookService    func() service.BookService
		args           args
		expected       string
		expectedStatus int
	}{
		{
			name: "success to delete book",
			bookService: func() service.BookService {
				bookService := service.NewMockBookService(ctrl)
				bookService.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)

				return bookService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("DELETE", "/books", nil),
			},
			expected:       string(`{"message":"Deleted book successfully!"}`),
			expectedStatus: http.StatusOK,
		},
		{
			name: "failed to to delete book",
			bookService: func() service.BookService {
				bookService := service.NewMockBookService(ctrl)
				bookService.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(errors.New("error occur"))

				return bookService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("DELETE", "/api/v1/books/1", nil),
			},
			expected:       string(`{"message":"Some thing wrong with the server"}`),
			expectedStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewBookHandler(tt.bookService())
			h.Delete(tt.args.w, tt.args.r)

			if tt.args.w.Body.String() != tt.expected {
				t.Errorf("Expected json response %s, got %s", tt.expected, tt.args.w.Body.String())
			}

			if tt.args.w.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, tt.args.w.Code)
			}
		})
	}

}
