package repository

import (
	"context"
	"database/sql"
	bookstorebe "winartodev/book-store-be"
)

type mysqlPublisher struct {
	DB *sql.DB
}

func NewMysqlPublisher(db *sql.DB) bookstorebe.PublisherRepository {
	return &mysqlPublisher{DB: db}
}

func (mp *mysqlPublisher) GetPublishers(ctx context.Context) ([]bookstorebe.Publisher, error) {
	var publishers = []bookstorebe.Publisher{}

	rows, err := mp.DB.Query("SELECT * FROM publisher")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var publisher bookstorebe.Publisher

		err := rows.Scan(&publisher.ID, &publisher.Name, &publisher.Address, &publisher.PhoneNumber)
		if err != nil {
			return nil, err
		}

		publishers = append(publishers, publisher)
	}

	return publishers, nil
}

func (mp *mysqlPublisher) GetPublisher(ctx context.Context, id int64) (bookstorebe.Publisher, error) {
	var publisher bookstorebe.Publisher

	err := mp.DB.QueryRow("SELECT * FROM publisher WHERE id=?", id).Scan(&publisher.ID, &publisher.Name, &publisher.Address, &publisher.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return bookstorebe.Publisher{}, nil
		}
		return bookstorebe.Publisher{}, err
	}

	return publisher, nil
}

func (mp *mysqlPublisher) CreatePublisher(ctx context.Context, publisher *bookstorebe.Publisher) error {
	stmt, err := mp.DB.Prepare("INSERT INTO publisher VALUES(NULL, ?, ?, ?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&publisher.Name, publisher.Address, &publisher.PhoneNumber)
	if err != nil {
		return err
	}

	if row, _ := res.RowsAffected(); row == 1 {
		return nil
	}

	return nil
}

func (mp *mysqlPublisher) UpdatePublisher(ctx context.Context, id int64, publisher *bookstorebe.Publisher) error {
	stmt, err := mp.DB.Prepare("UPDATE publisher SET name=?, address=?, phone_number=? WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&publisher.Name, &publisher.Address, &publisher.PhoneNumber, id)
	if err != nil {
		return err
	}

	return nil
}

func (mp *mysqlPublisher) DeletePublisher(ctx context.Context, id int64) error {
	stmt, err := mp.DB.Prepare("DELETE FROM publisher WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
