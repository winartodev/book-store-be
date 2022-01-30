package delivery_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"winartodev/book-store-be/delivery"
	"winartodev/book-store-be/entity"
	"winartodev/book-store-be/fixture"
	"winartodev/book-store-be/handler"
	"winartodev/book-store-be/logger"
	"winartodev/book-store-be/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func init() {
	logger.Init()
}

func newBookHandler() (http.Handler, *mocks.BookUsecase) {
	os.Setenv("BOOKSTORE_USERNAME", fixture.DummyUsername)
	os.Setenv("BOOKSTORE_PASSWORD", fixture.DummyPassword)

	username := fixture.DummyUsername
	password := fixture.DummyPassword

	uc := new(mocks.BookUsecase)
	book := delivery.NewBookHandler(uc, username, password)
	h := handler.NewHandler(&book)

	return h, uc
}

func TestGetBooks(t *testing.T) {
	testCases := []struct {
		name     string
		books    []entity.Book
		wantErr  bool
		getError error
	}{
		{
			name: "success",
			books: []entity.Book{
				{
					ID:          1,
					PublisherID: 1,
					CategoryID:  1,
					Title:       "Clean Architecture: A Craftsman's Guide to Software Structure and Design",
					Author:      "Robert C. Martin",
					Publication: 2020,
					Stock:       10,
				},
			},
		},
		{
			name:  "success with no data",
			books: []entity.Book{},
		},
		{
			name:     "failed to get books",
			books:    []entity.Book{},
			wantErr:  true,
			getError: errors.New("failed to get books"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			handler, book := newBookHandler()
			book.On("GetBooks", mock.Anything).Return(test.books, test.getError)

			recoder := httptest.NewRecorder()
			request := fixture.HTTPBasicAuth(http.MethodGet, "/bookstore/book", fixture.DummyUsername, fixture.DummyPassword, nil)

			handler.ServeHTTP(recoder, request)

			assert.Equal(t, test.wantErr, recoder.Code != http.StatusOK)
		})
	}
}

func TestGetBook(t *testing.T) {
	testCases := []struct {
		name    string
		id      int64
		book    entity.Book
		wantErr bool
		getErr  error
	}{
		{
			name: "success",
			id:   1,
			book: entity.Book{
				ID:          1,
				PublisherID: 1,
				CategoryID:  1,
				Title:       "Clean Architecture: A Craftsman's Guide to Software Structure and Design",
				Author:      "Robert C. Martin",
				Publication: 2020,
				Stock:       10,
			},
		},
		{
			name:    "success with book not found",
			id:      1,
			book:    entity.Book{},
			wantErr: true,
		},
		{
			name:    "failed to get book data",
			id:      1,
			book:    entity.Book{},
			wantErr: true,
			getErr:  errors.New("failed to get book data"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			handler, book := newBookHandler()
			book.On("GetBook", mock.Anything, mock.Anything).Return(test.book, test.getErr)

			recoder := httptest.NewRecorder()
			request := fixture.HTTPBasicAuth(http.MethodGet, fmt.Sprintf("/bookstore/book/%d", test.id), fixture.DummyUsername, fixture.DummyPassword, nil)

			handler.ServeHTTP(recoder, request)

			assert.Equal(t, test.wantErr, recoder.Code != http.StatusOK)
		})
	}
}

func TestCreateBook(t *testing.T) {
	testCases := []struct {
		name      string
		book      entity.Book
		wantErr   bool
		createErr error
	}{
		{
			name: "success",
			book: entity.Book{
				ID:          1,
				PublisherID: 1,
				CategoryID:  1,
				Title:       "Clean Architecture: A Craftsman's Guide to Software Structure and Design",
				Author:      "Robert C. Martin",
				Publication: 2020,
				Stock:       10,
			},
		},
		{
			name: "failed create book",
			book: entity.Book{
				ID:          1,
				PublisherID: 1,
				CategoryID:  1,
				Title:       "Clean Architecture: A Craftsman's Guide to Software Structure and Design",
				Author:      "Robert C. Martin",
				Publication: 2020,
				Stock:       10,
			},
			wantErr:   true,
			createErr: errors.New("failed create book"),
		},
	}

	for _, test := range testCases {
		handler, book := newBookHandler()
		book.On("CreateBook", mock.Anything, mock.Anything).Return(test.createErr)

		body, _ := json.Marshal(test.book)
		recoder := httptest.NewRecorder()
		request := fixture.HTTPBasicAuth(http.MethodPost, "/bookstore/book", fixture.DummyUsername, fixture.DummyPassword, body)

		handler.ServeHTTP(recoder, request)

		assert.Equal(t, test.wantErr, recoder.Code != http.StatusCreated)
	}
}

func TestUpdateBook(t *testing.T) {
	testCases := []struct {
		name      string
		id        int64
		book      entity.Book
		wantErr   bool
		updateErr error
	}{
		{
			name: "success",
			id:   1,
			book: entity.Book{
				ID:          1,
				PublisherID: 1,
				CategoryID:  1,
				Title:       "Clean Architecture: A Craftsman's Guide to Software Structure and Design",
				Author:      "Robert C. Martin",
				Publication: 2020,
				Stock:       10,
			},
		},
		{
			name: "failed to create book",
			id:   1,
			book: entity.Book{
				ID:          1,
				PublisherID: 1,
				CategoryID:  1,
				Title:       "Clean Architecture: A Craftsman's Guide to Software Structure and Design",
				Author:      "Robert C. Martin",
				Publication: 2020,
				Stock:       10,
			},
			wantErr:   true,
			updateErr: errors.New("failed to create book"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			handler, book := newBookHandler()
			book.On("UpdateBook", mock.Anything, mock.Anything, mock.Anything).Return(test.updateErr)

			body, _ := json.Marshal(test.book)
			recoder := httptest.NewRecorder()
			request := fixture.HTTPBasicAuth(http.MethodPut, fmt.Sprintf("/bookstore/book/%d", test.id), fixture.DummyUsername, fixture.DummyPassword, body)

			handler.ServeHTTP(recoder, request)

			assert.Equal(t, test.wantErr, recoder.Code != http.StatusOK)
		})
	}
}

func TestDeleteBook(t *testing.T) {
	testCases := []struct {
		name      string
		id        int64
		wantError bool
		deleteErr error
	}{
		{
			name: "success",
			id:   1,
		},
		{
			name:      "failed delete data",
			id:        1,
			wantError: true,
			deleteErr: errors.New("failed delete data"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			handler, book := newBookHandler()
			book.On("DeleteBook", mock.Anything, mock.Anything).Return(test.deleteErr)

			recoder := httptest.NewRecorder()
			request := fixture.HTTPBasicAuth(http.MethodDelete, fmt.Sprintf("/bookstore/book/%d", test.id), fixture.DummyUsername, fixture.DummyPassword, nil)

			handler.ServeHTTP(recoder, request)

			assert.Equal(t, test.wantError, recoder.Code != http.StatusOK)
		})
	}
}
