package usecase_test

import (
	"context"
	"errors"
	"testing"
	bookstorebe "winartodev/book-store-be"
	"winartodev/book-store-be/mocks"
	"winartodev/book-store-be/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCategoryProvider struct {
	categoryRepo *mocks.CategoryRepo
}

func categoryProvider() mockCategoryProvider {
	return mockCategoryProvider{
		categoryRepo: new(mocks.CategoryRepo),
	}
}

func newCategoryUsecase(uc *usecase.CategoryUsecase) bookstorebe.CategoryUsecase {
	return usecase.NewCategoryUsecase(uc)
}

func TestGetCategories(t *testing.T) {
	testCases := []struct {
		name     string
		category []bookstorebe.Category
		expRes   []bookstorebe.Category
		isError  bool
		wantErr  error
	}{
		{
			name:     "success",
			category: []bookstorebe.Category{{ID: 1, Name: "Classics"}, {ID: 2, Name: "Detective and Mystery"}},
			expRes:   []bookstorebe.Category{{ID: 1, Name: "Classics"}, {ID: 2, Name: "Detective and Mystery"}},
			isError:  false,
			wantErr:  nil,
		},
		{
			name:     "failed",
			category: []bookstorebe.Category{},
			expRes:   []bookstorebe.Category{},
			isError:  true,
			wantErr:  errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			prov := categoryProvider()
			prov.categoryRepo.On("GetCategories", mock.Anything).Return(test.category, test.wantErr)

			categoryUsecase := newCategoryUsecase(&usecase.CategoryUsecase{prov.categoryRepo})
			ctx := context.Background()
			res, err := categoryUsecase.GetCategories(ctx)

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

func TestGetCategory(t *testing.T) {
	testCases := []struct {
		name     string
		ID       int64
		category bookstorebe.Category
		expRes   bookstorebe.Category
		isError  bool
		wantErr  error
	}{
		{
			name:     "success",
			ID:       1,
			category: bookstorebe.Category{ID: 1, Name: "Classics"},
			expRes:   bookstorebe.Category{ID: 1, Name: "Classics"},
			isError:  false,
			wantErr:  nil,
		},
		{
			name:     "failed",
			ID:       1,
			category: bookstorebe.Category{ID: 2, Name: "Detective and Mystery"},
			expRes:   bookstorebe.Category{ID: 2, Name: "Detective and Mystery"},
			isError:  true,
			wantErr:  errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			prov := categoryProvider()
			prov.categoryRepo.On("GetCategory", mock.Anything, mock.AnythingOfType("int64")).Return(test.category, test.wantErr)

			categoryUsecase := newCategoryUsecase(&usecase.CategoryUsecase{prov.categoryRepo})
			ctx := context.Background()
			res, err := categoryUsecase.GetCategory(ctx, test.ID)

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

func TestCreateCategory(t *testing.T) {
	testCases := []struct {
		name     string
		category bookstorebe.Category
		isError  bool
		wantErr  error
	}{
		{
			name:     "success",
			category: bookstorebe.Category{ID: 1, Name: "Classics"},
			isError:  false,
			wantErr:  nil,
		},
		{
			name:     "failed",
			category: bookstorebe.Category{},
			isError:  true,
			wantErr:  errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			prov := categoryProvider()
			prov.categoryRepo.On("CreateCategory", mock.Anything, mock.Anything).Return(test.wantErr)

			categoryUsecase := newCategoryUsecase(&usecase.CategoryUsecase{prov.categoryRepo})
			ctx := context.Background()
			err := categoryUsecase.CreateCategory(ctx, &test.category)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}

func TestUpateCategory(t *testing.T) {
	testCases := []struct {
		name     string
		ID       int64
		category bookstorebe.Category
		isError  bool
		wantErr  error
	}{
		{
			name:     "success",
			ID:       1,
			category: bookstorebe.Category{ID: 1, Name: "Classics"},
			isError:  false,
			wantErr:  nil,
		},
		{
			name:     "failed",
			ID:       1,
			category: bookstorebe.Category{ID: 2, Name: "Detective and Mystery"},
			isError:  true,
			wantErr:  errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			prov := categoryProvider()
			prov.categoryRepo.On("UpdateCategory", mock.Anything, mock.AnythingOfType("int64"), mock.Anything).Return(test.wantErr)

			categoryUsecase := newCategoryUsecase(&usecase.CategoryUsecase{prov.categoryRepo})
			ctx := context.Background()
			err := categoryUsecase.UpdateCategory(ctx, test.ID, &test.category)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}

func TestDeleteCategory(t *testing.T) {
	testCases := []struct {
		name    string
		ID      int64
		isError bool
		wantErr error
	}{
		{
			name:    "success",
			ID:      1,
			isError: false,
			wantErr: nil,
		},
		{
			name:    "failed",
			ID:      1,
			isError: true,
			wantErr: errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			prov := categoryProvider()
			prov.categoryRepo.On("DeleteCategory", mock.Anything, mock.AnythingOfType("int64")).Return(test.wantErr)

			categoryUsecase := newCategoryUsecase(&usecase.CategoryUsecase{prov.categoryRepo})
			ctx := context.Background()
			err := categoryUsecase.DeleteCategory(ctx, test.ID)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}
