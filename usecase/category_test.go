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

type mockCategoryProvider struct {
	categoryRepo *mocks.CategoryRepository
}

func categoryProvider() mockCategoryProvider {
	return mockCategoryProvider{
		categoryRepo: new(mocks.CategoryRepository),
	}
}

func newCategoryUsecase(repo *usecase.CategoryRepository) usecase.CategoryUsecase {
	return usecase.NewCategoryUsecase(repo)
}

func TestGetCategories(t *testing.T) {
	testCases := []struct {
		name     string
		category []entity.Category
		expRes   []entity.Category
		isError  bool
		wantErr  error
	}{
		{
			name:     "success",
			category: []entity.Category{{ID: 1, Name: "Classics"}, {ID: 2, Name: "Detective and Mystery"}},
			expRes:   []entity.Category{{ID: 1, Name: "Classics"}, {ID: 2, Name: "Detective and Mystery"}},
			isError:  false,
			wantErr:  nil,
		},
		{
			name:     "failed",
			category: []entity.Category{},
			expRes:   []entity.Category{},
			isError:  true,
			wantErr:  errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			prov := categoryProvider()
			prov.categoryRepo.On("GetCategories", mock.Anything).Return(test.category, test.wantErr)

			categoryUsecase := newCategoryUsecase(&usecase.CategoryRepository{prov.categoryRepo})
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
		category entity.Category
		expRes   entity.Category
		isError  bool
		wantErr  error
	}{
		{
			name:     "success",
			ID:       1,
			category: entity.Category{ID: 1, Name: "Classics"},
			expRes:   entity.Category{ID: 1, Name: "Classics"},
			isError:  false,
			wantErr:  nil,
		},
		{
			name:     "failed",
			ID:       1,
			category: entity.Category{ID: 2, Name: "Detective and Mystery"},
			expRes:   entity.Category{ID: 2, Name: "Detective and Mystery"},
			isError:  true,
			wantErr:  errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			prov := categoryProvider()
			prov.categoryRepo.On("GetCategory", mock.Anything, mock.AnythingOfType("int64")).Return(test.category, test.wantErr)

			categoryUsecase := newCategoryUsecase(&usecase.CategoryRepository{prov.categoryRepo})
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
		category entity.Category
		isError  bool
		wantErr  error
	}{
		{
			name:     "success",
			category: entity.Category{ID: 1, Name: "Classics"},
			isError:  false,
			wantErr:  nil,
		},
		{
			name:     "failed",
			category: entity.Category{},
			isError:  true,
			wantErr:  errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			prov := categoryProvider()
			prov.categoryRepo.On("CreateCategory", mock.Anything, mock.Anything).Return(test.wantErr)

			categoryUsecase := newCategoryUsecase(&usecase.CategoryRepository{prov.categoryRepo})
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
		category entity.Category
		isError  bool
		wantErr  error
	}{
		{
			name:     "success",
			ID:       1,
			category: entity.Category{ID: 1, Name: "Classics"},
			isError:  false,
			wantErr:  nil,
		},
		{
			name:     "failed",
			ID:       1,
			category: entity.Category{ID: 2, Name: "Detective and Mystery"},
			isError:  true,
			wantErr:  errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			prov := categoryProvider()
			prov.categoryRepo.On("UpdateCategory", mock.Anything, mock.AnythingOfType("int64"), mock.Anything).Return(test.wantErr)

			categoryUsecase := newCategoryUsecase(&usecase.CategoryRepository{prov.categoryRepo})
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

			categoryUsecase := newCategoryUsecase(&usecase.CategoryRepository{prov.categoryRepo})
			ctx := context.Background()
			err := categoryUsecase.DeleteCategory(ctx, test.ID)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}
