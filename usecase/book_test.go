package usecase_test

import (
	"context"
	"errors"
	"testing"

	"winartodev/book-store-be/entity"
	"winartodev/book-store-be/mocks"
	"winartodev/book-store-be/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockBookProvider struct {
	BookRepo *mocks.BookRepository
}

func bookProvider() mockBookProvider {
	return mockBookProvider{
		BookRepo: new(mocks.BookRepository),
	}
}

func newBookUseCaseMock(repo *usecase.BookRepository) usecase.BookUsecase {
	return usecase.NewBookUsecase(repo)
}

func TestGetBooks(t *testing.T) {
	testCases := []struct {
		name    string
		books   []entity.Book
		expRes  []entity.Book
		isError bool
		wantErr error
	}{
		{
			name:    "success",
			books:   []entity.Book{{ID: 1, PublisherID: 1, CategoryID: 1, Title: "Book Title", Author: "Book Author", Publication: 2021, Stock: 4}},
			expRes:  []entity.Book{{ID: 1, PublisherID: 1, CategoryID: 1, Title: "Book Title", Author: "Book Author", Publication: 2021, Stock: 4}},
			isError: false,
			wantErr: nil,
		},
		{
			name:    "failed",
			books:   []entity.Book{{ID: 1, PublisherID: 1, CategoryID: 1, Title: "Book Title", Author: "Book Author", Publication: 2021, Stock: 4}},
			expRes:  nil,
			isError: true,
			wantErr: errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			prov := bookProvider()
			prov.BookRepo.On("GetBooks", mock.Anything).Return(test.books, test.wantErr)

			bookUsecase := newBookUseCaseMock(&usecase.BookRepository{prov.BookRepo})

			ctx := context.Background()
			res, err := bookUsecase.GetBooks(ctx)

			assert.Equal(t, test.isError, err != nil)
			if !test.isError {
				assert.NotNil(t, res)
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
				assert.Nil(t, res)
			}
		})
	}
}

func TestGetBook(t *testing.T) {
	testCases := []struct {
		name    string
		ID      int64
		book    entity.Book
		expBook entity.Book
		isError bool
		wantErr error
	}{
		{
			name:    "success",
			ID:      1,
			book:    entity.Book{ID: 1, PublisherID: 1, CategoryID: 1, Title: "Book Title", Author: "Book Author", Publication: 2021, Stock: 4},
			expBook: entity.Book{ID: 1, PublisherID: 1, CategoryID: 1, Title: "Book Title", Author: "Book Author", Publication: 2021, Stock: 4},
			isError: false,
			wantErr: nil,
		},
		{
			name:    "failed",
			ID:      1,
			book:    entity.Book{},
			expBook: entity.Book{},
			isError: true,
			wantErr: errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			prov := bookProvider()
			prov.BookRepo.On("GetBook", mock.Anything, mock.AnythingOfType("int64")).Return(test.book, test.wantErr)

			bookUsecase := newBookUseCaseMock(&usecase.BookRepository{prov.BookRepo})
			ctx := context.Background()
			res, err := bookUsecase.GetBook(ctx, test.ID)

			assert.Equal(t, test.isError, err != nil)
			if !test.isError {
				assert.NotNil(t, res)
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestCreateBook(t *testing.T) {
	testCases := []struct {
		name    string
		book    entity.Book
		isError bool
		wantErr error
	}{
		{
			name:    "success",
			book:    entity.Book{ID: 1, PublisherID: 1, CategoryID: 1, Title: "Book Title", Author: "Book Author", Publication: 2021, Stock: 4},
			isError: false,
			wantErr: nil,
		},
		{
			name:    "failed",
			book:    entity.Book{},
			isError: true,
			wantErr: errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			prov := bookProvider()
			prov.BookRepo.On("CreateBook", mock.Anything, mock.Anything).Return(test.wantErr)

			bookUsecase := newBookUseCaseMock(&usecase.BookRepository{prov.BookRepo})
			ctx := context.Background()
			err := bookUsecase.CreateBook(ctx, &test.book)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}

func TestUpdateBook(t *testing.T) {
	testCases := []struct {
		name    string
		ID      int64
		book    entity.Book
		isError bool
		wantErr error
	}{
		{
			name:    "success",
			ID:      1,
			book:    entity.Book{ID: 1, PublisherID: 1, CategoryID: 1, Title: "Book Title", Author: "Book Author", Publication: 2021, Stock: 4},
			isError: false,
			wantErr: nil,
		},
		{
			name:    "failed",
			ID:      1,
			book:    entity.Book{},
			isError: true,
			wantErr: errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			prov := bookProvider()
			prov.BookRepo.On("UpdateBook", mock.Anything, mock.AnythingOfType("int64"), mock.Anything).Return(test.wantErr)

			bookUsecase := newBookUseCaseMock(&usecase.BookRepository{prov.BookRepo})
			ctx := context.Background()
			err := bookUsecase.UpdateBook(ctx, test.ID, &test.book)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}

func TestDeleteBook(t *testing.T) {
	testCases := []struct {
		name    string
		ID      int64
		book    entity.Book
		isError bool
		wantErr error
	}{
		{
			name:    "success",
			ID:      1,
			book:    entity.Book{ID: 1, PublisherID: 1, CategoryID: 1, Title: "Book Title", Author: "Book Author", Publication: 2021, Stock: 4},
			isError: false,
			wantErr: nil,
		},
		{
			name:    "failed",
			ID:      1,
			book:    entity.Book{},
			isError: true,
			wantErr: errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			prov := bookProvider()
			prov.BookRepo.On("DeleteBook", mock.Anything, mock.AnythingOfType("int64")).Return(test.wantErr)

			bookUsecase := newBookUseCaseMock(&usecase.BookRepository{prov.BookRepo})
			ctx := context.Background()
			err := bookUsecase.DeleteBook(ctx, test.ID)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}
