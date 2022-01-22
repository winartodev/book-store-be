package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"winartodev/book-store-be/entity"
	"winartodev/book-store-be/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {
	rightQuery := "SELECT (.+)"
	wrongQuery := "SELECT (.+)"

	testCases := []struct {
		name    string
		rows    []entity.Book
		query   string
		isError bool
		err     error
	}{
		{
			name:    "success",
			rows:    []entity.Book{{ID: 1, PublisherID: 1, CategoryID: 1, Title: "Book Title", Author: "Book Author", Publication: 2021, Stock: 4, Price: 100000, CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			query:   rightQuery,
			isError: false,
			err:     nil,
		},
		{
			name:    "wrong query",
			rows:    []entity.Book{},
			query:   wrongQuery,
			isError: true,
			err:     errors.New("Dummy Error"),
		},
		{
			name:    "failed",
			rows:    []entity.Book{{ID: 1, PublisherID: 1, CategoryID: 1, Title: "Book Title", Author: "Book Author", Publication: 2021, Stock: 4, Price: 100000, CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			query:   rightQuery,
			isError: false,
			err:     nil,
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
				rows := sqlmock.NewRows([]string{"id", "publisher_id", "category_id", "title", "author", "year_of_publication", "stock", "price", "created_at", "updated_at"})
				for _, row := range test.rows {
					rows.AddRow(row.ID, row.PublisherID, row.CategoryID, row.Title, row.Author, row.Publication, row.Stock, row.Price, row.CreatedAt, row.UpdatedAt)
				}
				mock.ExpectQuery(test.query).WillReturnRows(rows)
			} else {
				mock.ExpectQuery(test.query).WillReturnError(test.err)
			}

			mysqlBook := repository.NewMysqlBook(db)
			ret, err := mysqlBook.GetBooks(context.Background())

			assert.Equal(t, test.isError, err != nil)

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

func TestGetBook(t *testing.T) {
	rightQuery := "SELECT (.+) "
	wrongQuery := "SELECT (.+)"

	testCases := []struct {
		name    string
		id      int64
		row     entity.Book
		query   string
		isError bool
		err     error
	}{
		{
			name:    "success",
			id:      1,
			row:     entity.Book{ID: 1, PublisherID: 1, CategoryID: 1, Title: "Book Title", Author: "Book Author", Publication: 2021, Stock: 4, Price: 100000, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			query:   rightQuery,
			isError: false,
			err:     nil,
		},
		{
			name:    "success but data is empty",
			id:      1,
			row:     entity.Book{},
			query:   rightQuery,
			isError: false,
			err:     sql.ErrNoRows,
		},
		{
			name:    "failed",
			id:      1,
			row:     entity.Book{},
			query:   rightQuery,
			isError: true,
			err:     errors.New("Dummy Error"),
		},
		{
			name:    "wrong query",
			id:      1,
			row:     entity.Book{},
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
				row := sqlmock.NewRows([]string{"id", "publisher_id", "category_id", "title", "author", "year_of_publication", "stock", "price", "created_at", "updated_at"}).
					AddRow(test.row.ID, test.row.PublisherID, test.row.CategoryID, test.row.Title, test.row.Author, test.row.Publication, test.row.Stock, test.row.Price, test.row.CreatedAt, test.row.UpdatedAt)

				mock.ExpectQuery(test.query).WithArgs(test.id).WillReturnRows(row)
			} else {
				mock.ExpectQuery(test.query).WithArgs(test.id).WillReturnError(test.err)
			}

			mysqlBook := repository.NewMysqlBook(db)
			ret, err := mysqlBook.GetBook(context.Background(), test.id)

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

func TestCreateBook(t *testing.T) {
	rightQuery := "INSERT INTO (.+)"
	wrongQuery := "INSERT INTO"

	testCases := []struct {
		name    string
		book    entity.Book
		query   string
		isError bool
		err     error
	}{
		{
			name:    "success",
			book:    entity.Book{PublisherID: 1, CategoryID: 1, Title: "Book Title", Author: "Book Author", Publication: 2021, Stock: 4, Price: 10000, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			query:   rightQuery,
			isError: false,
			err:     nil,
		},
		{
			name:    "failed",
			book:    entity.Book{},
			query:   rightQuery,
			isError: true,
			err:     errors.New("Dummy Error"),
		},
		{
			name:    "wrong query",
			book:    entity.Book{},
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
				mock.ExpectPrepare(test.query).
					ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				mock.ExpectPrepare(test.query).
					ExpectExec().WillReturnError(test.err)
			}

			mysqlBook := repository.NewMysqlBook(db)
			err = mysqlBook.CreateBook(context.Background(), &test.book)
			assert.Equal(t, test.isError, err != nil)
		})
	}
}

func TestUpdateBook(t *testing.T) {
	rightQuery := "UPDATE books (.+)"
	wrongQuery := "UPDATEING books"

	testCases := []struct {
		name    string
		id      int64
		book    entity.Book
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
			name:    "failed",
			id:      1,
			book:    entity.Book{},
			query:   rightQuery,
			isError: true,
			err:     errors.New("Dummy Error"),
		},
		{
			name:    "wrong query",
			id:      1,
			book:    entity.Book{},
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
				mock.ExpectPrepare(test.query).
					ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				mock.ExpectPrepare(test.query).
					ExpectExec().WillReturnError(test.err)
			}

			mysqlBook := repository.NewMysqlBook(db)
			err = mysqlBook.UpdateBook(context.Background(), test.id, &test.book)
			fmt.Println(err)
			assert.Equal(t, test.isError, err != nil)
		})
	}
}

func TestDeleteBook(t *testing.T) {
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
			name:    "failed",
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

			mysqlBook := repository.NewMysqlBook(db)
			err = mysqlBook.DeleteBook(context.Background(), test.id)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}
