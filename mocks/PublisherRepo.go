package mocks

import (
	"context"
	"fmt"
	bookstorebe "winartodev/book-store-be"

	"github.com/stretchr/testify/mock"
)

type PublisherRepo struct {
	mock.Mock
}

func (m *PublisherRepo) GetPublishers(ctx context.Context) ([]bookstorebe.Publisher, error) {
	fmt.Println("1")
	ret := m.Called(ctx)
	return ret.Get(0).([]bookstorebe.Publisher), ret.Error(1)
}

func (m *PublisherRepo) GetPublisher(ctx context.Context, id int64) (bookstorebe.Publisher, error) {
	ret := m.Called(ctx, id)
	return ret.Get(0).(bookstorebe.Publisher), ret.Error(1)
}

func (m *PublisherRepo) CreatePublisher(ctx context.Context, Publisher *bookstorebe.Publisher) error {
	ret := m.Called(ctx, Publisher)
	return ret.Error(0)
}

func (m *PublisherRepo) UpdatePublisher(ctx context.Context, id int64, Publisher *bookstorebe.Publisher) error {
	ret := m.Called(ctx, id, Publisher)
	return ret.Error(0)
}

func (m *PublisherRepo) DeletePublisher(ctx context.Context, id int64) error {
	ret := m.Called(ctx, id)
	return ret.Error(0)
}
