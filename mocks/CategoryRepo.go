package mocks

import (
	"context"
	bookstorebe "winartodev/book-store-be"

	"github.com/stretchr/testify/mock"
)

type CategoryRepo struct {
	mock.Mock
}

func (m *CategoryRepo) GetCategories(ctx context.Context) ([]bookstorebe.Category, error) {
	ret := m.Called(ctx)
	return ret.Get(0).([]bookstorebe.Category), ret.Error(1)
}

func (m *CategoryRepo) GetCategory(ctx context.Context, id int64) (bookstorebe.Category, error) {
	ret := m.Called(ctx, id)
	return ret.Get(0).(bookstorebe.Category), ret.Error(1)
}

func (m *CategoryRepo) CreateCategory(ctx context.Context, category *bookstorebe.Category) error {
	ret := m.Called(ctx, category)
	return ret.Error(0)
}

func (m *CategoryRepo) UpdateCategory(ctx context.Context, id int64, category *bookstorebe.Category) error {
	ret := m.Called(ctx, id, category)
	return ret.Error(0)
}

func (m *CategoryRepo) DeleteCategory(ctx context.Context, id int64) error {
	ret := m.Called(ctx, id)
	return ret.Error(0)
}
