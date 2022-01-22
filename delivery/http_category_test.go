package delivery_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"winartodev/book-store-be/delivery"
	"winartodev/book-store-be/entity"
	"winartodev/book-store-be/handler"
	"winartodev/book-store-be/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func newCategoryHandler() (http.Handler, *mocks.CategoryUsecase) {
	uc := new(mocks.CategoryUsecase)
	category := delivery.NewCategoryHandler(uc)
	h := handler.NewHandler(&category)
	return h, uc
}

func TestGetCategories(t *testing.T) {
	testCases := []struct {
		name      string
		endpoint  string
		category  []entity.Category
		wantError bool
		getError  error
	}{
		{
			name:      "success",
			endpoint:  "/bookstore/category",
			category:  []entity.Category{{ID: 1, Name: "Classics"}, {ID: 2, Name: "Detective and Mystery"}},
			wantError: false,
			getError:  nil,
		},
		{
			name:      "success with no data",
			endpoint:  "/bookstore/category",
			category:  []entity.Category{},
			wantError: false,
			getError:  nil,
		},
		{
			name:      "failed get category data",
			endpoint:  "/bookstore/category",
			category:  []entity.Category{},
			wantError: true,
			getError:  errors.New("failed get category data"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			handler, category := newCategoryHandler()
			category.On("GetCategories", mock.Anything).Return(test.category, test.getError)

			recoder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, test.endpoint, nil)
			handler.ServeHTTP(recoder, request)

			assert.Equal(t, test.wantError, recoder.Code != http.StatusOK)
		})
	}
}

func TestGetCategory(t *testing.T) {
	testCases := []struct {
		name     string
		id       int64
		category entity.Category
		wantErr  bool
		getError error
	}{
		{
			name:     "success",
			id:       1,
			category: entity.Category{ID: 1, Name: "Classics"},
			wantErr:  false,
			getError: nil,
		},
		{
			name:     "failed with category not found",
			id:       1,
			category: entity.Category{},
			wantErr:  true,
			getError: nil,
		},
		{
			name:     "failed to get category",
			id:       1,
			category: entity.Category{ID: 1, Name: "Classics"},
			wantErr:  true,
			getError: errors.New("failed to get category"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			handler, category := newCategoryHandler()
			category.On("GetCategory", mock.Anything, mock.Anything).Return(test.category, test.getError)

			recoder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/bookstore/category/%v", test.id), nil)
			handler.ServeHTTP(recoder, request)

			assert.Equal(t, test.wantErr, recoder.Code != http.StatusOK)
		})
	}
}

func TestCreateCategory(t *testing.T) {
	testCases := []struct {
		name      string
		endpoint  string
		category  entity.Category
		wantError bool
		createErr error
	}{
		{
			name:      "success",
			endpoint:  "/bookstore/category",
			category:  entity.Category{ID: 1, Name: "Classics"},
			wantError: false,
			createErr: nil,
		},
		{
			name:      "failed created category",
			endpoint:  "/bookstore/category",
			category:  entity.Category{ID: 1, Name: "Classics"},
			wantError: true,
			createErr: errors.New("failed created category"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			handler, category := newCategoryHandler()
			category.On("CreateCategory", mock.Anything, mock.Anything).Return(test.createErr)

			body, _ := json.Marshal(test.category)
			recoder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost, test.endpoint, bytes.NewBuffer(body))

			handler.ServeHTTP(recoder, request)
			assert.Equal(t, test.wantError, recoder.Code != http.StatusCreated)
		})
	}
}

func TestUpdateCategory(t *testing.T) {
	testcases := []struct {
		name      string
		id        int64
		category  entity.Category
		wantError bool
		updateErr error
	}{
		{
			name:      "success",
			id:        1,
			category:  entity.Category{ID: 1, Name: "Classics"},
			wantError: false,
			updateErr: nil,
		},
		{
			name:      "failed update category",
			id:        1,
			category:  entity.Category{ID: 1, Name: "Classics"},
			wantError: true,
			updateErr: errors.New("fail update category"),
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			handler, category := newCategoryHandler()
			category.On("UpdateCategory", mock.Anything, mock.Anything, mock.Anything).Return(test.updateErr)

			body, _ := json.Marshal(test.category)

			recoder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/bookstore/category/%d", test.id), bytes.NewBuffer(body))
			handler.ServeHTTP(recoder, request)

			assert.Equal(t, test.wantError, recoder.Code != http.StatusOK)
		})
	}
}

func TestDeleteCategory(t *testing.T) {
	testcases := []struct {
		name      string
		id        int64
		wantErr   bool
		deleteErr error
	}{
		{
			name:      "success",
			id:        1,
			wantErr:   false,
			deleteErr: nil,
		},
		{
			name:      "failed to delete category",
			id:        1,
			wantErr:   true,
			deleteErr: errors.New("failed to delete category"),
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			handler, category := newCategoryHandler()
			category.On("DeleteCategory", mock.Anything, mock.Anything).Return(test.deleteErr)

			recoder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/bookstore/category/%d", test.id), nil)
			handler.ServeHTTP(recoder, request)

			assert.Equal(t, test.wantErr, recoder.Code != http.StatusOK)
		})
	}
}
