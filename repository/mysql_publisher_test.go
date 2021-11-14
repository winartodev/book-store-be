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

func TestGetPublisher(t *testing.T) {
	testCases := []struct {
		name    string
		rows    []bookstorebe.Publisher
		expRows []bookstorebe.Publisher
		isError bool
		err     error
	}{
		{
			name:    "success",
			rows:    []bookstorebe.Publisher{{ID: 1, Name: "Publisher Name", Address: "Publisher Address", PhoneNumber: "123456789"}},
			expRows: []bookstorebe.Publisher{{ID: 1, Name: "Publisher Name", Address: "Publisher Address", PhoneNumber: "123456789"}},
			isError: false,
			err:     nil,
		},
		{
			name:    "failed",
			rows:    []bookstorebe.Publisher{},
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
				rows := sqlmock.NewRows([]string{"id", "name", "address", "phone_number"})
				for _, row := range test.rows {
					rows.AddRow(row.ID, row.Name, row.Address, row.PhoneNumber)
				}
				mock.ExpectQuery("SELECT (.+) FROM publisher").WillReturnRows(rows)
			} else {
				mock.ExpectQuery("SELECT (.+) FROM publisher").WillReturnError(test.err)
			}

			mysqlPublisher := repository.NewMysqlPublisher(db)
			ret, err := mysqlPublisher.GetPublishers(context.Background())

			assert.Equal(t, err != nil, test.isError)
			if !test.isError {
				assert.NotNil(t, ret)
				assert.Equal(t, test.expRows, ret)
				assert.Nil(t, err)
			} else {
				assert.Nil(t, ret)
				assert.NotNil(t, err)
			}
		})
	}
}

func TestGetPublishers(t *testing.T) {
	testCases := []struct {
		name    string
		id      int64
		row     bookstorebe.Publisher
		expRows bookstorebe.Publisher
		isError bool
		err     error
	}{
		{
			name:    "success",
			id:      1,
			row:     bookstorebe.Publisher{ID: 1, Name: "Publisher Name", Address: "Publisher Address", PhoneNumber: "123456789"},
			expRows: bookstorebe.Publisher{ID: 1, Name: "Publisher Name", Address: "Publisher Address", PhoneNumber: "123456789"},
			isError: false,
			err:     nil,
		},
		{
			name:    "failed",
			id:      3,
			row:     bookstorebe.Publisher{},
			expRows: bookstorebe.Publisher{},
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
				rows := sqlmock.NewRows([]string{"id", "name", "address", "phone_number"}).AddRow(&test.row.ID, &test.row.Name, &test.row.Address, &test.row.PhoneNumber)
				mock.ExpectQuery("SELECT (.+) FROM publisher WHERE id=\\?").WillReturnRows(rows)
			} else {
				mock.ExpectQuery("SELECT (.+) FROM publisher WHERE id=\\?").WillReturnError(test.err)
			}

			mysqlPublisher := repository.NewMysqlPublisher(db)
			ret, err := mysqlPublisher.GetPublisher(context.Background(), test.id)

			assert.Equal(t, err != nil, test.isError)
			if !test.isError {
				assert.Equal(t, test.expRows, ret)
				assert.NotNil(t, ret)
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestCreatePublisher(t *testing.T) {
	testCases := []struct {
		name      string
		publisher bookstorebe.Publisher
		isError   bool
		err       error
	}{
		{
			name:      "success",
			publisher: bookstorebe.Publisher{ID: 1, Name: "Publisher Name", Address: "Publisher Address", PhoneNumber: "123456789"},
			isError:   false,
			err:       nil,
		},
		{
			name:      "failed",
			publisher: bookstorebe.Publisher{},
			isError:   true,
			err:       errors.New("Dummy Error"),
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
				mock.ExpectPrepare("INSERT INTO publisher VALUES\\(NULL, \\?, \\?, \\?\\)").
					ExpectExec().WithArgs(&test.publisher.Name, &test.publisher.Address, &test.publisher.PhoneNumber).
					WillReturnResult(sqlmock.NewResult(0, 1))
			} else {
				mock.ExpectPrepare("INSERT INTO publisher VALUES\\(NULL, \\?, \\?, \\?\\)").
					ExpectExec().WithArgs(&test.publisher.Name, &test.publisher.Address, &test.publisher.PhoneNumber).
					WillReturnError(test.err)
			}

			mysqlPublisher := repository.NewMysqlPublisher(db)
			err = mysqlPublisher.CreatePublisher(context.Background(), &test.publisher)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}

func TestUpdatePublisher(t *testing.T) {
	testCases := []struct {
		name      string
		id        int64
		publisher bookstorebe.Publisher
		isError   bool
		err       error
	}{
		{
			name:      "success",
			id:        1,
			publisher: bookstorebe.Publisher{ID: 1, Name: "Publisher Name", Address: "Publisher Address", PhoneNumber: "123456789"},
			isError:   false,
			err:       nil,
		},
		{
			name:      "failed",
			id:        1,
			publisher: bookstorebe.Publisher{ID: 1, Name: "Publisher Name", Address: "Publisher Address", PhoneNumber: "123456789"},
			isError:   true,
			err:       errors.New("Dummy Error"),
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
				mock.ExpectPrepare("UPDATE publisher SET name=\\?, address=\\?, phone_number=\\? WHERE id=\\?").
					ExpectExec().WithArgs(&test.publisher.Name, &test.publisher.Address, &test.publisher.PhoneNumber, test.id).
					WillReturnResult(sqlmock.NewResult(0, 0))
			} else {
				mock.ExpectPrepare("UPDATE publisher SET name=\\?, address=\\?, phone_number=\\? WHERE id=\\?").
					ExpectExec().WithArgs(&test.publisher.Name, &test.publisher.Address, &test.publisher.PhoneNumber, test.id).
					WillReturnError(test.err)
			}

			mysqlPublisher := repository.NewMysqlPublisher(db)
			err = mysqlPublisher.UpdatePublisher(context.Background(), test.id, &test.publisher)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}

func TestDeletePublishetr(t *testing.T) {
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
				mock.ExpectPrepare("DELETE publisher WHERE id=\\?").ExpectExec().WithArgs(test.id).WillReturnResult(sqlmock.NewResult(0, 0))
			} else {
				mock.ExpectPrepare("DELETE publisher WHERE id=\\?").ExpectExec().WithArgs(test.id).WillReturnError(test.err)
			}

			mysqlPublisher := repository.NewMysqlPublisher(db)
			err = mysqlPublisher.DeletePublisher(context.Background(), test.id)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}
