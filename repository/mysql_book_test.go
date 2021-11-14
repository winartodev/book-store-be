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

func TestGetBooks(t *testing.T) {
	testCases := []struct {
		name    string
		rows    []bookstorebe.Book
		expRows []bookstorebe.Book
		isError bool
		err     error
	}{
		{
			name:    "success",
			rows:    []bookstorebe.Book{{ID: 1, PublisherID: 1, CategoryID: 1, Title: "Book Title", Author: "Book Author", Publication: 2021, Stock: 4}},
			expRows: []bookstorebe.Book{{ID: 1, PublisherID: 1, CategoryID: 1, Title: "Book Title", Author: "Book Author", Publication: 2021, Stock: 4}},
			isError: false,
			err:     nil,
		},
		{
			name:    "failed",
			rows:    []bookstorebe.Book{},
			expRows: nil,
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
				rows := sqlmock.NewRows([]string{"id", "publisher_id", "category_id", "title", "author", "year_of_publication", "stock"})
				for _, row := range test.rows {
					rows.AddRow(row.ID, row.PublisherID, row.CategoryID, row.Title, row.Author, row.Publication, row.Stock)
				}
				mock.ExpectQuery("SELECT (.+) FROM book").WillReturnRows(rows)
			} else {
				mock.ExpectQuery("SELECT (.+) FROM book").WillReturnError(test.err)
			}

			mysqlBook := repository.NewMysqlBook(db)
			ret, err := mysqlBook.GetBooks(context.Background())

			assert.Equal(t, test.isError, err != nil)

			if !test.isError {
				assert.NotNil(t, ret)
				assert.Equal(t, test.expRows, ret)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestGetBook(t *testing.T) {
	testCases := []struct {
		name    string
		id      int64
		row     bookstorebe.Book
		expRow  bookstorebe.Book
		isError bool
		err     error
	}{
		{
			name:    "success",
			id:      1,
			row:     bookstorebe.Book{ID: 1, PublisherID: 1, CategoryID: 1, Title: "Book Title", Author: "Book Author", Publication: 2021, Stock: 4},
			expRow:  bookstorebe.Book{ID: 1, PublisherID: 1, CategoryID: 1, Title: "Book Title", Author: "Book Author", Publication: 2021, Stock: 4},
			isError: false,
			err:     nil,
		},
		{
			name:    "failed",
			id:      1,
			row:     bookstorebe.Book{},
			expRow:  bookstorebe.Book{},
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
				row := sqlmock.NewRows([]string{"id", "publisher_id", "category_id", "title", "author", "year_of_publication", "stock"}).
					AddRow(test.row.ID, test.row.PublisherID, test.row.CategoryID, test.row.Title, test.row.Author, test.row.Publication, test.row.Stock)

				mock.ExpectQuery("SELECT (.+) FROM book WHERE id=\\?").WillReturnRows(row)
			} else {
				mock.ExpectQuery("SELECT (.+) FROM book WHERE id=\\?").WillReturnError(test.err)
			}

			mysqlBook := repository.NewMysqlBook(db)
			ret, err := mysqlBook.GetBook(context.Background(), test.id)

			assert.Equal(t, test.isError, err != nil)

			if !test.isError {
				assert.NotNil(t, ret)
				assert.Equal(t, test.expRow, ret)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestCreateBook(t *testing.T) {
	testCases := []struct {
		name    string
		book    bookstorebe.Book
		isError bool
		err     error
	}{
		{
			name:    "success",
			book:    bookstorebe.Book{ID: 1, PublisherID: 1, CategoryID: 1, Title: "Book Title", Author: "Book Author", Publication: 2021, Stock: 4},
			isError: false,
			err:     nil,
		},
		{
			name:    "failed",
			book:    bookstorebe.Book{},
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
				mock.ExpectPrepare("INSERT INTO book VALUES\\(NULL, \\?, \\?, \\?, \\?, \\?. \\?\\)").
					ExpectExec().WithArgs(&test.book.PublisherID, &test.book.CategoryID, &test.book.Title, &test.book.Author, &test.book.Publication, &test.book.Stock).
					WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				mock.ExpectPrepare("INSERT INTO book VALUES\\(NULL, \\?, \\?, \\?, \\?, \\?. \\?\\)").
					ExpectExec().WithArgs(&test.book.PublisherID, &test.book.CategoryID, &test.book.Title, &test.book.Author, &test.book.Publication, &test.book.Stock).
					WillReturnError(test.err)
			}

			mysqlBook := repository.NewMysqlBook(db)
			err = mysqlBook.CreateBook(context.Background(), &test.book)
			fmt.Println(err)
			assert.Equal(t, test.isError, err != nil)
		})
	}
}

func TestUpdateBook(t *testing.T) {
	testCases := []struct {
		name    string
		id      int64
		book    bookstorebe.Book
		isError bool
		err     error
	}{
		{
			name:    "success",
			id:      1,
			book:    bookstorebe.Book{ID: 1, PublisherID: 1, CategoryID: 1, Title: "Book Title", Author: "Book Author", Publication: 2021, Stock: 4},
			isError: false,
			err:     nil,
		},
		{
			name:    "failed",
			id:      1,
			book:    bookstorebe.Book{},
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
				mock.ExpectPrepare("UPDATE book SET publisher_id=\\?, category_id=\\?, title=\\?, author=\\?, year_of_publication=\\?, stock=\\? WHERE id=\\?").
					ExpectExec().WithArgs(test.book.PublisherID, test.book.CategoryID, test.book.Title, test.book.Author, test.book.Publication, test.book.Stock, test.id).
					WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				mock.ExpectPrepare("UPDATE book SET publisher_id=\\?, category_id=\\?, title=\\?, author=\\?, year_of_publication=\\?, stock=\\? WHERE id=\\?").
					ExpectExec().WithArgs(test.book.PublisherID, test.book.CategoryID, test.book.Title, test.book.Author, test.book.Publication, test.book.Stock, test.id).
					WillReturnError(test.err)
			}

			mysqlBook := repository.NewMysqlBook(db)
			err = mysqlBook.UpdateBook(context.Background(), test.id, &test.book)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}

func TestDeleteBook(t *testing.T) {
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
			name:    "failed",
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
				mock.ExpectPrepare("DELETE book WHERE id=\\?").ExpectExec().WithArgs(test.id).WillReturnResult(sqlmock.NewResult(0, 0))
			} else {
				mock.ExpectPrepare("DELETE book WHERE id=\\?").ExpectExec().WithArgs(test.id).WillReturnError(test.err)
			}

			mysqlBook := repository.NewMysqlBook(db)
			err = mysqlBook.DeleteBook(context.Background(), test.id)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}
