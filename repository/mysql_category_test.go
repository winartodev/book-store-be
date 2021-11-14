package repository_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	bookstorebe "winartodev/book-store-be"
	"winartodev/book-store-be/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetCategories(t *testing.T) {
	testCases := []struct {
		name     string
		category []bookstorebe.Category
		expRows  []bookstorebe.Category
		isError  bool
		err      error
	}{
		{
			name:     "success",
			category: []bookstorebe.Category{{ID: 1, Name: "Classics"}, {ID: 2, Name: "Detective and Mystery"}},
			expRows:  []bookstorebe.Category{{ID: 1, Name: "Classics"}, {ID: 2, Name: "Detective and Mystery"}},
			isError:  false,
			err:      nil,
		},
		{
			name:     "failed",
			category: []bookstorebe.Category{},
			expRows:  nil,
			isError:  true,
			err:      errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				panic(fmt.Sprintf("Database Not Connect %s", err))
			}
			defer db.Close()

			if !test.isError {
				rows := sqlmock.NewRows([]string{"id", "name"})
				for _, row := range test.category {
					rows.AddRow(&row.ID, &row.Name)
				}

				mock.ExpectQuery("SELECT (.+) FROM category").WillReturnRows(rows)
			} else {
				mock.ExpectQuery("SELECT (.+) FROM category").WillReturnError(test.err)
			}

			mysqlCategory := repository.NewMysqlCategory(db)
			ret, err := mysqlCategory.GetCategories(context.Background())

			assert.Equal(t, err != nil, test.isError)
			if !test.isError {
				assert.Equal(t, test.expRows, ret)
			} else {
				assert.Equal(t, test.expRows, ret)
			}
		})
	}
}

func TestGetCategory(t *testing.T) {
	testCases := []struct {
		name     string
		id       int64
		category bookstorebe.Category
		expRows  bookstorebe.Category
		isError  bool
		err      error
	}{
		{
			name:     "success",
			id:       1,
			category: bookstorebe.Category{ID: 1, Name: "Classics"},
			expRows:  bookstorebe.Category{ID: 1, Name: "Classics"},
			isError:  false,
			err:      nil,
		},
		{
			name:     "failed",
			id:       1,
			category: bookstorebe.Category{},
			expRows:  bookstorebe.Category{},
			isError:  true,
			err:      errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				panic(fmt.Sprintf("Database Not Connect %s", err))
			}
			defer db.Close()

			if !test.isError {
				rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(&test.category.ID, &test.category.Name)
				mock.ExpectQuery("SELECT (.+) FROM category WHERE id=\\?").WillReturnRows(rows)
			} else {
				mock.ExpectQuery("SELECT (.+) FROM category WHERE id=\\?").WillReturnError(test.err)
			}

			mysqlCategory := repository.NewMysqlCategory(db)
			ret, err := mysqlCategory.GetCategory(context.Background(), test.id)

			assert.Equal(t, test.isError, err != nil)
			if !test.isError {
				assert.Equal(t, test.expRows, ret)
			} else {
				assert.Equal(t, test.expRows, ret)
			}
		})
	}
}

func TestCreateCategory(t *testing.T) {
	testCases := []struct {
		name     string
		category bookstorebe.Category
		isError  bool
		err      error
	}{
		{
			name:     "success",
			category: bookstorebe.Category{ID: 1, Name: "Classics"},
			isError:  false,
			err:      nil,
		},
		{
			name:     "failed",
			category: bookstorebe.Category{},
			isError:  true,
			err:      errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				panic(fmt.Sprintf("Database Not Connect %s", err))
			}
			defer db.Close()

			if !test.isError {
				mock.ExpectPrepare("INSERT INTO category VALUES\\(NULL, \\?\\)").
					ExpectExec().WithArgs(&test.category.Name).
					WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				mock.ExpectPrepare("INSERT INTO category VALUES\\(NULL, \\?\\)").
					ExpectExec().WithArgs(&test.category.Name).
					WillReturnError(test.err)
			}

			mysqlCategory := repository.NewMysqlCategory(db)
			err = mysqlCategory.CreateCategory(context.Background(), &test.category)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}

func TestUpdateCategory(t *testing.T) {
	testCases := []struct {
		name     string
		id       int64
		category bookstorebe.Category
		isError  bool
		err      error
	}{
		{
			name:     "success",
			id:       1,
			category: bookstorebe.Category{ID: 1, Name: "Classics"},
			isError:  false,
			err:      nil,
		},
		{
			name:     "failed",
			id:       1,
			category: bookstorebe.Category{},
			isError:  true,
			err:      errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				panic(fmt.Sprintf("Database Not Connect %s", err))
			}
			defer db.Close()

			if !test.isError {
				mock.ExpectPrepare("UPDATE category SET name=\\? WHERE id=\\?").
					ExpectExec().WithArgs(&test.category.Name, test.id).
					WillReturnResult(sqlmock.NewResult(0, 1))
			} else {
				mock.ExpectPrepare("UPDATE category SET name=\\? WHERE id=\\?").
					ExpectExec().WithArgs(&test.category.Name, test.id).
					WillReturnError(test.err)
			}

			mysqlCategory := repository.NewMysqlCategory(db)
			err = mysqlCategory.UpdateCategory(context.Background(), test.id, &test.category)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}

func TestDeleteCategory(t *testing.T) {
	testCases := []struct {
		name    string
		id      int64
		isError bool
		err     error
	}{
		{
			name:    "success",
			id:      1,
			isError: false,
			err:     nil,
		},
		{
			name:    "falied",
			id:      1,
			isError: true,
			err:     errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				panic(fmt.Sprintf("Database Not Connect %s", err))
			}
			defer db.Close()

			if !test.isError {
				mock.ExpectPrepare("DELETE category WHERE id=\\?").ExpectExec().WithArgs(test.id).WillReturnResult(sqlmock.NewResult(0, 0))
			} else {
				mock.ExpectPrepare("DELETE category WHERE id=\\?").ExpectExec().WithArgs(test.id).WillReturnError(test.err)
			}

			mysqlCategory := repository.NewMysqlCategory(db)
			err = mysqlCategory.DeleteCategory(context.Background(), test.id)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}
