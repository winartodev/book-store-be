// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "winartodev/book-store-be/entity"

	mock "github.com/stretchr/testify/mock"
)

// PublisherRepository is an autogenerated mock type for the PublisherRepository type
type PublisherRepository struct {
	mock.Mock
}

// CreatePublisher provides a mock function with given fields: ctx, publisher
func (_m *PublisherRepository) CreatePublisher(ctx context.Context, publisher *entity.Publisher) error {
	ret := _m.Called(ctx, publisher)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Publisher) error); ok {
		r0 = rf(ctx, publisher)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePublisher provides a mock function with given fields: ctx, id
func (_m *PublisherRepository) DeletePublisher(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPublisher provides a mock function with given fields: ctx, id
func (_m *PublisherRepository) GetPublisher(ctx context.Context, id int64) (entity.Publisher, error) {
	ret := _m.Called(ctx, id)

	var r0 entity.Publisher
	if rf, ok := ret.Get(0).(func(context.Context, int64) entity.Publisher); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entity.Publisher)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPublishers provides a mock function with given fields: ctx
func (_m *PublisherRepository) GetPublishers(ctx context.Context) ([]entity.Publisher, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Publisher
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Publisher); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Publisher)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePublisher provides a mock function with given fields: ctx, id, publisher
func (_m *PublisherRepository) UpdatePublisher(ctx context.Context, id int64, publisher *entity.Publisher) error {
	ret := _m.Called(ctx, id, publisher)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *entity.Publisher) error); ok {
		r0 = rf(ctx, id, publisher)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
