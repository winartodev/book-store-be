package repository

import (
	"context"
	"database/sql"
	"time"
	"winartodev/book-store-be/entity"
)

type PublisherRepository interface {
	// seller
	GetPublishers(ctx context.Context) ([]entity.Publisher, error)
	GetPublisher(ctx context.Context, id int64) (entity.Publisher, error)
	CreatePublisher(ctx context.Context, publisher *entity.Publisher) error
	UpdatePublisher(ctx context.Context, id int64, publisher *entity.Publisher) error
	DeletePublisher(ctx context.Context, id int64) error
}

type mysqlPublisher struct {
	DB *sql.DB
}

func NewMysqlPublisher(db *sql.DB) PublisherRepository {
	return &mysqlPublisher{DB: db}
}

func (mp *mysqlPublisher) GetPublishers(ctx context.Context) ([]entity.Publisher, error) {
	var publishers = []entity.Publisher{}

	rows, err := mp.DB.Query("SELECT * FROM publishers")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var publisher entity.Publisher

		err := rows.Scan(&publisher.ID, &publisher.Name, &publisher.Address, &publisher.PhoneNumber, &publisher.CreatedAt, &publisher.UpdatedAt)
		if err != nil {
			return nil, err
		}

		publishers = append(publishers, publisher)
	}

	return publishers, nil
}

func (mp *mysqlPublisher) GetPublisher(ctx context.Context, id int64) (entity.Publisher, error) {
	var publisher entity.Publisher

	err := mp.DB.QueryRow("SELECT * FROM publishers WHERE id=$1", id).Scan(&publisher.ID, &publisher.Name, &publisher.Address, &publisher.PhoneNumber, &publisher.CreatedAt, &publisher.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Publisher{}, nil
		}
		return entity.Publisher{}, err
	}

	return publisher, nil
}

func (mp *mysqlPublisher) CreatePublisher(ctx context.Context, publisher *entity.Publisher) error {
	stmt, err := mp.DB.Prepare("INSERT INTO publishers (name, address, phone_number, created_at, updated_at) VALUES($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}

	startTime := time.Now()
	publisher.CreatedAt = startTime
	publisher.UpdatedAt = startTime

	res, err := stmt.Exec(&publisher.Name, publisher.Address, &publisher.PhoneNumber, &publisher.CreatedAt, &publisher.UpdatedAt)
	if err != nil {
		return err
	}

	if row, _ := res.RowsAffected(); row == 1 {
		return nil
	}

	return nil
}

func (mp *mysqlPublisher) UpdatePublisher(ctx context.Context, id int64, publisher *entity.Publisher) error {
	stmt, err := mp.DB.Prepare("UPDATE publishers SET name=$1, address=$2, phone_number=$3, updated_at=$4 WHERE id=$5")
	if err != nil {
		return err
	}

	startTime := time.Now()
	publisher.UpdatedAt = startTime

	_, err = stmt.Exec(&publisher.Name, &publisher.Address, &publisher.PhoneNumber, &publisher.UpdatedAt, id)
	if err != nil {
		return err
	}

	return nil
}

func (mp *mysqlPublisher) DeletePublisher(ctx context.Context, id int64) error {
	stmt, err := mp.DB.Prepare("DELETE FROM publishers WHERE id=$1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
