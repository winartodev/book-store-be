// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	bookstorebe "winartodev/book-store-be"

	mock "github.com/stretchr/testify/mock"
)

// BookUsecase is an autogenerated mock type for the BookUsecase type
type BookUsecase struct {
	mock.Mock
}

// CreateBook provides a mock function with given fields: ctx, book
func (_m *BookUsecase) CreateBook(ctx context.Context, book *bookstorebe.Book) error {
	ret := _m.Called(ctx, book)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *bookstorebe.Book) error); ok {
		r0 = rf(ctx, book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBook provides a mock function with given fields: ctx, id
func (_m *BookUsecase) DeleteBook(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBook provides a mock function with given fields: ctx, id
func (_m *BookUsecase) GetBook(ctx context.Context, id int64) (bookstorebe.Book, error) {
	ret := _m.Called(ctx, id)

	var r0 bookstorebe.Book
	if rf, ok := ret.Get(0).(func(context.Context, int64) bookstorebe.Book); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bookstorebe.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBooks provides a mock function with given fields: ctx
func (_m *BookUsecase) GetBooks(ctx context.Context) ([]bookstorebe.Book, error) {
	ret := _m.Called(ctx)

	var r0 []bookstorebe.Book
	if rf, ok := ret.Get(0).(func(context.Context) []bookstorebe.Book); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bookstorebe.Book)
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

// UpdateBook provides a mock function with given fields: ctx, id, book
func (_m *BookUsecase) UpdateBook(ctx context.Context, id int64, book *bookstorebe.Book) error {
	ret := _m.Called(ctx, id, book)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *bookstorebe.Book) error); ok {
		r0 = rf(ctx, id, book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
