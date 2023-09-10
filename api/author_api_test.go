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

func Test_authorHandler_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	authors := []*payload.AuthorResponse{
		{
			Id:          test.AuthorId1,
			FirstName:   test.AuthorFirstName1,
			LastName:    test.AuthorLastName1,
			BirthDate:   test.AuthorBirthDate1,
			Nationality: test.AuthorNationality1,
			CreatedAt:   test.CreatedAtStr,
			UpdatedAt:   test.UpdatedAtStr,
		},
	}
	expectedAuthorsJson, _ := json.Marshal(authors)

	tests := []struct {
		name           string
		authorService  func() service.AuthorService
		args           args
		expected       string
		expectedStatus int
	}{
		{
			name: "success to retrieve authors",
			authorService: func() service.AuthorService {
				authorService := service.NewMockAuthorService(ctrl)
				authorService.EXPECT().FindAll(gomock.Any()).Return(authors, nil)

				return authorService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/api/v1/authors", nil),
			},
			expected:       string(expectedAuthorsJson),
			expectedStatus: http.StatusOK,
		},
		{
			name: "failed to retrieve authors",
			authorService: func() service.AuthorService {
				authorService := service.NewMockAuthorService(ctrl)
				authorService.EXPECT().FindAll(gomock.Any()).Return(nil, errors.New("error occur"))

				return authorService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/api/v1/authors", nil),
			},
			expected:       string(`{"message":"Some thing wrong with the server"}`),
			expectedStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewAuthorHandler(tt.authorService())
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

func Test_authorHandler_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	author := &payload.AuthorResponse{
		Id:          test.AuthorId1,
		FirstName:   test.AuthorFirstName1,
		LastName:    test.AuthorLastName1,
		BirthDate:   test.AuthorBirthDate1,
		Nationality: test.AuthorNationality1,
		CreatedAt:   test.CreatedAtStr,
		UpdatedAt:   test.UpdatedAtStr,
	}
	expectedAuthorsJson, _ := json.Marshal(author)

	tests := []struct {
		name           string
		authorService  func() service.AuthorService
		args           args
		expected       string
		expectedStatus int
	}{
		{
			name: "success to retrieve author",
			authorService: func() service.AuthorService {
				authorService := service.NewMockAuthorService(ctrl)
				authorService.EXPECT().Find(gomock.Any(), gomock.Any()).Return(author, nil)

				return authorService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/api/v1/authors/1", nil),
			},
			expected:       string(expectedAuthorsJson),
			expectedStatus: http.StatusOK,
		},
		{
			name: "failed to retrieve author",
			authorService: func() service.AuthorService {
				authorService := service.NewMockAuthorService(ctrl)
				authorService.EXPECT().Find(gomock.Any(), gomock.Any()).Return(nil, errors.New("error occur"))

				return authorService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/api/v1/authors/1", nil),
			},
			expected:       string(`{"message":"Some thing wrong with the server"}`),
			expectedStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewAuthorHandler(tt.authorService())
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

func Test_authorHandler_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	author := &payload.AuthorRequest{
		FirstName:   test.AuthorFirstName1,
		LastName:    test.AuthorLastName1,
		BirthDate:   test.AuthorBirthDate1,
		Nationality: test.AuthorNationality1,
	}
	bodyData, _ := json.Marshal(author)

	tests := []struct {
		name           string
		authorService  func() service.AuthorService
		args           args
		expected       string
		expectedStatus int
	}{
		{
			name: "success to create author",
			authorService: func() service.AuthorService {
				authorService := service.NewMockAuthorService(ctrl)
				authorService.EXPECT().Store(gomock.Any(), gomock.Any()).Return(nil)

				return authorService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/authors", bytes.NewReader(bodyData)),
			},
			expected:       string(bodyData),
			expectedStatus: http.StatusOK,
		},
		{
			name: "failed to to create author",
			authorService: func() service.AuthorService {
				authorService := service.NewMockAuthorService(ctrl)
				authorService.EXPECT().Store(gomock.Any(), gomock.Any()).Return(errors.New("error occur"))

				return authorService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/api/v1/authors", bytes.NewReader(bodyData)),
			},
			expected:       string(`{"message":"Some thing wrong with the server"}`),
			expectedStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewAuthorHandler(tt.authorService())
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

func Test_authorHandler_Put(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	author := &payload.AuthorRequest{
		FirstName:   test.AuthorFirstName1,
		LastName:    test.AuthorLastName1,
		BirthDate:   test.AuthorBirthDate1,
		Nationality: test.AuthorNationality1,
	}
	bodyData, _ := json.Marshal(author)

	tests := []struct {
		name           string
		authorService  func() service.AuthorService
		args           args
		expected       string
		expectedStatus int
	}{
		{
			name: "success to update author",
			authorService: func() service.AuthorService {
				authorService := service.NewMockAuthorService(ctrl)
				authorService.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

				return authorService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("PUT", "/authors", bytes.NewReader(bodyData)),
			},
			expected:       string(bodyData),
			expectedStatus: http.StatusOK,
		},
		{
			name: "failed to to update author",
			authorService: func() service.AuthorService {
				authorService := service.NewMockAuthorService(ctrl)
				authorService.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error occur"))

				return authorService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("PUT", "/api/v1/authors", bytes.NewReader(bodyData)),
			},
			expected:       string(`{"message":"Some thing wrong with the server"}`),
			expectedStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewAuthorHandler(tt.authorService())
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

func Test_authorHandler_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	tests := []struct {
		name           string
		authorService  func() service.AuthorService
		args           args
		expected       string
		expectedStatus int
	}{
		{
			name: "success to delete author",
			authorService: func() service.AuthorService {
				authorService := service.NewMockAuthorService(ctrl)
				authorService.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)

				return authorService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("DELETE", "/authors", nil),
			},
			expected:       string(`{"message":"Deleted author successfully!"}`),
			expectedStatus: http.StatusOK,
		},
		{
			name: "failed to to delete author",
			authorService: func() service.AuthorService {
				authorService := service.NewMockAuthorService(ctrl)
				authorService.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(errors.New("error occur"))

				return authorService
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("DELETE", "/api/v1/authors/1", nil),
			},
			expected:       string(`{"message":"Some thing wrong with the server"}`),
			expectedStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewAuthorHandler(tt.authorService())
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
