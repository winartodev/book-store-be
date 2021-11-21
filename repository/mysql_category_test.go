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
		query    string
		isError  bool
		err      error
	}{
		{
			name:     "success",
			category: []bookstorebe.Category{{ID: 1, Name: "Classics"}, {ID: 2, Name: "Detective and Mystery"}},
			expRows:  []bookstorebe.Category{{ID: 1, Name: "Classics"}, {ID: 2, Name: "Detective and Mystery"}},
			query:    "SELECT (.+) FROM category",
			isError:  false,
			err:      nil,
		},
		{
			name:     "failed",
			category: []bookstorebe.Category{},
			expRows:  nil,
			query:    "SELECT (.+) FROM category",
			isError:  true,
			err:      errors.New("Dummy Error"),
		},
		{
			name:     "wrong query",
			category: []bookstorebe.Category{},
			expRows:  nil,
			query:    "SELECT * FROM category",
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

				mock.ExpectQuery(test.query).WillReturnRows(rows)
			} else {
				mock.ExpectQuery(test.query).WillReturnError(test.err)
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
		query    string
		isError  bool
		err      error
	}{
		{
			name:     "success",
			id:       1,
			category: bookstorebe.Category{ID: 1, Name: "Classics"},
			expRows:  bookstorebe.Category{ID: 1, Name: "Classics"},
			query:    "SELECT (.+) FROM category WHERE id=\\?",
			isError:  false,
			err:      nil,
		},
		{
			name:     "failed",
			id:       1,
			category: bookstorebe.Category{},
			expRows:  bookstorebe.Category{},
			query:    "SELECT (.+) FROM category WHERE id=\\?",
			isError:  true,
			err:      errors.New("Dummy Error"),
		},
		{
			name:     "row not found",
			id:       1,
			category: bookstorebe.Category{},
			expRows:  bookstorebe.Category{},
			query:    "SELECT (.+) FROM category WHERE id=\\?",
			isError:  true,
			err:      errors.New("sql: no rows in result set"),
		},
		{
			name:     "wrong query",
			id:       1,
			category: bookstorebe.Category{},
			expRows:  bookstorebe.Category{},
			query:    "SELECT * FROM category WHERE id=?",
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
				mock.ExpectQuery(test.query).WillReturnRows(rows)
			} else {
				mock.ExpectQuery(test.query).WillReturnError(test.err)
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
		query    string
		isError  bool
		err      error
	}{
		{
			name:     "success",
			category: bookstorebe.Category{ID: 1, Name: "Classics"},
			query:    "INSERT INTO category VALUES\\(NULL, \\?\\)",
			isError:  false,
			err:      nil,
		},
		{
			name:     "failed",
			category: bookstorebe.Category{},
			query:    "INSERT INTO category VALUES\\(NULL, \\?\\)",
			isError:  true,
			err:      errors.New("Dummy Error"),
		},
		{
			name:     "wrong query",
			category: bookstorebe.Category{},
			query:    "INSERT INTO category VALUES(NULL, ?)",
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
				mock.ExpectPrepare(test.query).
					ExpectExec().WithArgs(&test.category.Name).
					WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				mock.ExpectPrepare(test.query).
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
		query    string
		isError  bool
		err      error
	}{
		{
			name:     "success",
			id:       1,
			category: bookstorebe.Category{ID: 1, Name: "Classics"},
			query:    "UPDATE category SET name=\\? WHERE id=\\?",
			isError:  false,
			err:      nil,
		},
		{
			name:     "failed",
			id:       1,
			category: bookstorebe.Category{},
			query:    "UPDATE category SET name=\\? WHERE id=\\?",
			isError:  true,
			err:      errors.New("Dummy Error"),
		},
		{
			name:     "wrong query",
			id:       1,
			category: bookstorebe.Category{},
			query:    "UPDATE category SET name=? WHERE id=?",
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
				mock.ExpectPrepare(test.query).
					ExpectExec().WithArgs(&test.category.Name, test.id).
					WillReturnResult(sqlmock.NewResult(0, 1))
			} else {
				mock.ExpectPrepare(test.query).
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
		query   string
		isError bool
		err     error
	}{
		{
			name:    "success",
			id:      1,
			query:   "DELETE category WHERE id=\\?",
			isError: false,
			err:     nil,
		},
		{
			name:    "falied",
			id:      1,
			query:   "DELETE category WHERE id=\\?",
			isError: true,
			err:     errors.New("Dummy Error"),
		},
		{
			name:    "wrong query",
			id:      1,
			query:   "DELETE category WHERE ids=?",
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
				mock.ExpectPrepare(test.query).ExpectExec().WithArgs(test.id).WillReturnResult(sqlmock.NewResult(0, 0))
			} else {
				mock.ExpectPrepare(test.query).ExpectExec().WithArgs(test.id).WillReturnError(test.err)
			}

			mysqlCategory := repository.NewMysqlCategory(db)
			err = mysqlCategory.DeleteCategory(context.Background(), test.id)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}
