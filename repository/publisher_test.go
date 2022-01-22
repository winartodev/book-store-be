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

func TestGetPublisher(t *testing.T) {
	rightQuery := "SELECT (.+)"
	wrongQuery := "SELECT (.+)"

	testCases := []struct {
		name    string
		rows    []entity.Publisher
		query   string
		isError bool
		err     error
	}{
		{
			name:    "success",
			rows:    []entity.Publisher{{ID: 1, Name: "Publisher Name", Address: "Publisher Address", PhoneNumber: "123456789", CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			query:   rightQuery,
			isError: false,
			err:     nil,
		},
		{
			name:    "failed",
			rows:    []entity.Publisher{},
			query:   rightQuery,
			isError: true,
			err:     errors.New("Dummy Error"),
		},
		{
			name:    "wrong query",
			rows:    []entity.Publisher{},
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
				rows := sqlmock.NewRows([]string{"id", "name", "address", "phone_number", "created_at", "updated_at"})
				for _, row := range test.rows {
					rows.AddRow(row.ID, row.Name, row.Address, row.PhoneNumber, row.CreatedAt, row.UpdatedAt)
				}
				mock.ExpectQuery(test.query).WillReturnRows(rows)
			} else {
				mock.ExpectQuery(test.query).WillReturnError(test.err)
			}

			mysqlPublisher := repository.NewMysqlPublisher(db)
			ret, err := mysqlPublisher.GetPublishers(context.Background())

			assert.Equal(t, err != nil, test.isError)
			if !test.isError {
				assert.NotNil(t, ret)
				assert.Nil(t, err)
			} else {
				assert.Nil(t, ret)
				assert.NotNil(t, err)
			}
		})
	}
}

func TestGetPublishers(t *testing.T) {
	rightQuery := "SELECT (.+)"
	wrongQuery := "SELECT (.+)"

	testCases := []struct {
		name    string
		id      int64
		row     entity.Publisher
		query   string
		isError bool
		err     error
	}{
		{
			name:    "success",
			id:      1,
			row:     entity.Publisher{ID: 1, Name: "Publisher Name", Address: "Publisher Address", PhoneNumber: "123456789", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			query:   rightQuery,
			isError: false,
			err:     nil,
		},
		{
			name:    "failed",
			id:      3,
			row:     entity.Publisher{},
			query:   rightQuery,
			isError: true,
			err:     errors.New("Dummy Error"),
		},
		{
			name:    "wrong query",
			id:      3,
			row:     entity.Publisher{},
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
				rows := sqlmock.NewRows([]string{"id", "name", "address", "phone_number", "created_at", "updated_at"}).AddRow(&test.row.ID, &test.row.Name, &test.row.Address, &test.row.PhoneNumber, test.row.CreatedAt, test.row.UpdatedAt)
				mock.ExpectQuery(test.query).WithArgs(test.id).WillReturnRows(rows)
			} else {
				mock.ExpectQuery(test.query).WithArgs(test.id).WillReturnError(test.err)
			}

			mysqlPublisher := repository.NewMysqlPublisher(db)
			ret, err := mysqlPublisher.GetPublisher(context.Background(), test.id)

			assert.Equal(t, err != nil, test.isError)
			if !test.isError {
				assert.NotNil(t, ret)
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestCreatePublisher(t *testing.T) {
	rightQuery := "INSERT (.+)"
	wrongQuery := "INSERT (.+)"

	testCases := []struct {
		name      string
		publisher entity.Publisher
		query     string
		isError   bool
		err       error
	}{
		{
			name:      "success",
			publisher: entity.Publisher{ID: 1, Name: "Publisher Name", Address: "Publisher Address", PhoneNumber: "123456789", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			query:     rightQuery,
			isError:   false,
			err:       nil,
		},
		{
			name:      "failed",
			publisher: entity.Publisher{},
			query:     rightQuery,
			isError:   true,
			err:       errors.New("Dummy Error"),
		},
		{
			name:      "wrong query",
			publisher: entity.Publisher{},
			query:     wrongQuery,
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
				mock.ExpectPrepare(test.query).
					ExpectExec().WithArgs(&test.publisher.Name, &test.publisher.Address, &test.publisher.PhoneNumber, &test.publisher.CreatedAt, &test.publisher.UpdatedAt).
					WillReturnResult(sqlmock.NewResult(0, 1))
			} else {
				mock.ExpectPrepare(test.query).
					ExpectExec().WithArgs(&test.publisher.Name, &test.publisher.Address, &test.publisher.PhoneNumber, &test.publisher.CreatedAt, &test.publisher.UpdatedAt).
					WillReturnError(test.err)
			}

			mysqlPublisher := repository.NewMysqlPublisher(db)
			err = mysqlPublisher.CreatePublisher(context.Background(), &test.publisher)
			assert.Equal(t, test.isError, err != nil)
		})
	}
}

func TestUpdatePublisher(t *testing.T) {
	rightQuery := "UPDATE (.+)"
	wrongQuery := "UPDATE (.+)"

	testCases := []struct {
		name      string
		id        int64
		publisher entity.Publisher
		query     string
		isError   bool
		err       error
	}{
		{
			name:      "success",
			id:        1,
			publisher: entity.Publisher{ID: 1, Name: "Publisher Name", Address: "Publisher Address", PhoneNumber: "123456789", UpdatedAt: time.Now()},
			query:     rightQuery,
			isError:   false,
			err:       nil,
		},
		{
			name:      "failed",
			id:        1,
			publisher: entity.Publisher{ID: 1, Name: "Publisher Name", Address: "Publisher Address", PhoneNumber: "123456789"},
			query:     rightQuery,
			isError:   true,
			err:       errors.New("Dummy Error"),
		},
		{
			name:      "wrong query",
			id:        1,
			publisher: entity.Publisher{ID: 1, Name: "Publisher Name", Address: "Publisher Address", PhoneNumber: "123456789"},
			query:     wrongQuery,
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
				mock.ExpectPrepare(test.query).
					ExpectExec().WithArgs(&test.publisher.Name, &test.publisher.Address, &test.publisher.PhoneNumber, &test.publisher.UpdatedAt, test.id).
					WillReturnResult(sqlmock.NewResult(0, 0))
			} else {
				mock.ExpectPrepare(test.query).
					ExpectExec().WithArgs(&test.publisher.Name, &test.publisher.Address, &test.publisher.PhoneNumber, &test.publisher.UpdatedAt, test.id).
					WillReturnError(test.err)
			}

			mysqlPublisher := repository.NewMysqlPublisher(db)
			err = mysqlPublisher.UpdatePublisher(context.Background(), test.id, &test.publisher)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}

func TestDeletePublishetr(t *testing.T) {
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

			mysqlPublisher := repository.NewMysqlPublisher(db)
			err = mysqlPublisher.DeletePublisher(context.Background(), test.id)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}
