// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "winartodev/book-store-be/entity"

	mock "github.com/stretchr/testify/mock"
)

// CategoryRepository is an autogenerated mock type for the CategoryRepository type
type CategoryRepository struct {
	mock.Mock
}

// CreateCategory provides a mock function with given fields: ctx, category
func (_m *CategoryRepository) CreateCategory(ctx context.Context, category *entity.Category) error {
	ret := _m.Called(ctx, category)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Category) error); ok {
		r0 = rf(ctx, category)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCategory provides a mock function with given fields: ctx, id
func (_m *CategoryRepository) DeleteCategory(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCategories provides a mock function with given fields: ctx
func (_m *CategoryRepository) GetCategories(ctx context.Context) ([]entity.Category, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Category
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Category); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Category)
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

// GetCategory provides a mock function with given fields: ctx, id
func (_m *CategoryRepository) GetCategory(ctx context.Context, id int64) (entity.Category, error) {
	ret := _m.Called(ctx, id)

	var r0 entity.Category
	if rf, ok := ret.Get(0).(func(context.Context, int64) entity.Category); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entity.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCategory provides a mock function with given fields: ctx, id, category
func (_m *CategoryRepository) UpdateCategory(ctx context.Context, id int64, category *entity.Category) error {
	ret := _m.Called(ctx, id, category)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *entity.Category) error); ok {
		r0 = rf(ctx, id, category)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
