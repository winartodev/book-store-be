package repository_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"winartodev/book-store-be/entity"
	"winartodev/book-store-be/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetCategories(t *testing.T) {
	rightQuery := "SELECT (.+) "
	wrongQuery := "SELECT (.+)"

	testCases := []struct {
		name     string
		category []entity.Category
		query    string
		isError  bool
		err      error
	}{
		{
			name:     "success",
			category: []entity.Category{{ID: 1, Name: "Classics", CreatedAt: time.Now(), UpdatedAt: time.Now()}, {ID: 2, Name: "Detective and Mystery", CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			query:    rightQuery,
			isError:  false,
			err:      nil,
		},
		{
			name:     "failed",
			category: []entity.Category{},
			query:    rightQuery,
			isError:  true,
			err:      errors.New("Dummy Error"),
		},
		{
			name:     "wrong query",
			category: []entity.Category{},
			query:    wrongQuery,
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
				rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"})
				for _, row := range test.category {
					rows.AddRow(&row.ID, &row.Name, &row.CreatedAt, &row.UpdatedAt)
				}

				mock.ExpectQuery(test.query).WillReturnRows(rows)
			} else {
				mock.ExpectQuery(test.query).WillReturnError(test.err)
			}

			mysqlCategory := repository.NewMysqlCategory(db)
			ret, err := mysqlCategory.GetCategories(context.Background())

			assert.Equal(t, err != nil, test.isError)
			if !test.isError {
				assert.NotNil(t, ret)
				assert.NoError(t, err)
			} else {
				assert.Nil(t, ret)
				assert.Error(t, err)
			}
		})
	}
}

func TestGetCategory(t *testing.T) {
	rightQuery := "SELECT (.+)"
	wrongQuery := "SELECT (.+)"

	testCases := []struct {
		name     string
		id       int64
		category entity.Category
		query    string
		isError  bool
		err      error
	}{
		{
			name:     "success",
			id:       1,
			category: entity.Category{},
			query:    rightQuery,
			isError:  false,
			err:      nil,
		},
		{
			name:     "failed",
			id:       1,
			category: entity.Category{},
			query:    rightQuery,
			isError:  true,
			err:      errors.New("Dummy Error"),
		},
		{
			name:     "row not found",
			id:       1,
			category: entity.Category{},
			query:    rightQuery,
			isError:  true,
			err:      errors.New("sql: no rows in result set"),
		},
		{
			name:     "wrong query",
			id:       1,
			category: entity.Category{},
			query:    wrongQuery,
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
				rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).AddRow(&test.category.ID, &test.category.Name, &test.category.CreatedAt, &test.category.UpdatedAt)
				mock.ExpectQuery(test.query).WillReturnRows(rows)
			} else {
				mock.ExpectQuery(test.query).WillReturnError(test.err)
			}

			mysqlCategory := repository.NewMysqlCategory(db)
			ret, err := mysqlCategory.GetCategory(context.Background(), test.id)

			assert.Equal(t, test.isError, err != nil)
			if !test.isError {
				assert.NotNil(t, ret)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestCreateCategory(t *testing.T) {
	rightQuery := "INSERT (.+)"
	wrongQuery := "INSERT (.+))"

	testCases := []struct {
		name     string
		category entity.Category
		query    string
		isError  bool
		err      error
	}{
		{
			name:     "success",
			category: entity.Category{ID: 1, Name: "Classics", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			query:    rightQuery,
			isError:  false,
			err:      nil,
		},
		{
			name:     "failed",
			category: entity.Category{},
			query:    rightQuery,
			isError:  true,
			err:      errors.New("Dummy Error"),
		},
		{
			name:     "wrong query",
			category: entity.Category{},
			query:    wrongQuery,
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
					ExpectExec().WithArgs(&test.category.Name, &test.category.CreatedAt, &test.category.UpdatedAt).
					WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				mock.ExpectPrepare(test.query).
					ExpectExec().WithArgs(&test.category.Name, &test.category.CreatedAt, &test.category.UpdatedAt).
					WillReturnError(test.err)
			}

			mysqlCategory := repository.NewMysqlCategory(db)
			err = mysqlCategory.CreateCategory(context.Background(), &test.category)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}

func TestUpdateCategory(t *testing.T) {
	rightQuery := "UPDATE (.+)"
	wrongQuery := "UPDATE (.+)"

	testCases := []struct {
		name     string
		id       int64
		category entity.Category
		query    string
		isError  bool
		err      error
	}{
		{
			name:     "success",
			id:       1,
			category: entity.Category{ID: 1, Name: "Classics", UpdatedAt: time.Now()},
			query:    rightQuery,
			isError:  false,
			err:      nil,
		},
		{
			name:     "failed",
			id:       1,
			category: entity.Category{},
			query:    rightQuery,
			isError:  true,
			err:      errors.New("Dummy Error"),
		},
		{
			name:     "wrong query",
			id:       1,
			category: entity.Category{},
			query:    wrongQuery,
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
					ExpectExec().WithArgs(&test.category.Name, &test.category.UpdatedAt, test.id).
					WillReturnResult(sqlmock.NewResult(0, 1))
			} else {
				mock.ExpectPrepare(test.query).
					ExpectExec().WithArgs(&test.category.Name, &test.category.UpdatedAt, test.id).
					WillReturnError(test.err)
			}

			mysqlCategory := repository.NewMysqlCategory(db)
			err = mysqlCategory.UpdateCategory(context.Background(), test.id, &test.category)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}

func TestDeleteCategory(t *testing.T) {
	rightQuery := "DELETE (.+)"
	wrongQuery := "DELETE (.+)"

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
			query:   rightQuery,
			isError: false,
			err:     nil,
		},
		{
			name:    "falied",
			id:      1,
			query:   rightQuery,
			isError: true,
			err:     errors.New("Dummy Error"),
		},
		{
			name:    "wrong query",
			id:      1,
			query:   wrongQuery,
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
