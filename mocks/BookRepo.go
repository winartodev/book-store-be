package mocks

import (
	"context"
	bookstorebe "winartodev/book-store-be"

	"github.com/stretchr/testify/mock"
)

type BookRepo struct {
	mock.Mock
}

func (m *BookRepo) GetBooks(ctx context.Context) ([]bookstorebe.Book, error) {
	ret := m.Called(ctx)
	return ret.Get(0).([]bookstorebe.Book), ret.Error(1)
}

func (m *BookRepo) GetBook(ctx context.Context, id int64) (bookstorebe.Book, error) {
	ret := m.Called(ctx, id)
	return ret.Get(0).(bookstorebe.Book), ret.Error(1)
}

func (m *BookRepo) CreateBook(ctx context.Context, book *bookstorebe.Book) error {
	ret := m.Called(ctx, book)
	return ret.Error(0)
}

func (m *BookRepo) UpdateBook(ctx context.Context, id int64, book *bookstorebe.Book) error {
	ret := m.Called(ctx, id, book)
	return ret.Error(0)
}

func (m *BookRepo) DeleteBook(ctx context.Context, id int64) error {
	ret := m.Called(ctx, id)
	return ret.Error(0)
}
