package repository

import (
	"context"
	"database/sql"
	"time"
	"winartodev/book-store-be/entity"
)

type BookRepository interface {
	// seller
	GetBooks(ctx context.Context) ([]entity.Book, error)
	GetBook(ctx context.Context, id int64) (entity.Book, error)
	CreateBook(ctx context.Context, book *entity.Book) error
	UpdateBook(ctx context.Context, id int64, book *entity.Book) error
	DeleteBook(ctx context.Context, id int64) error
}

type mysqlBook struct {
	DB *sql.DB
}

func NewMysqlBook(db *sql.DB) BookRepository {
	return &mysqlBook{DB: db}
}

func (mb *mysqlBook) GetBooks(ctx context.Context) ([]entity.Book, error) {
	var books []entity.Book

	rows, err := mb.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book entity.Book

		err := rows.Scan(&book.ID, &book.PublisherID, &book.CategoryID, &book.Title, &book.Author, &book.Publication, &book.Stock, &book.Price, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (mb *mysqlBook) GetBook(ctx context.Context, id int64) (entity.Book, error) {
	var book entity.Book

	err := mb.DB.QueryRow("SELECT * FROM books WHERE id=$1", id).Scan(&book.ID, &book.PublisherID, &book.CategoryID, &book.Title, &book.Author, &book.Publication, &book.Stock, &book.Price, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Book{}, nil
		}
		return entity.Book{}, err
	}

	return book, nil
}

func (mb *mysqlBook) CreateBook(ctx context.Context, book *entity.Book) error {
	stmt, err := mb.DB.Prepare("INSERT INTO books (publisher_id, category_id, title, author, year_of_publication, stock, price, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)")

	if err != nil {
		return err
	}

	startTime := time.Now()
	book.CreatedAt = startTime
	book.UpdatedAt = startTime

	res, err := stmt.Exec(&book.PublisherID, &book.CategoryID, &book.Title, &book.Author, &book.Publication, &book.Stock, &book.Price, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return err
	}

	if row, _ := res.RowsAffected(); row == 1 {
		return nil
	}

	return nil
}

func (mb *mysqlBook) UpdateBook(ctx context.Context, id int64, book *entity.Book) error {
	stmt, err := mb.DB.Prepare("UPDATE books SET publisher_id=$1, category_id=$2, title=$3, author=$4, year_of_publication=$5, stock=$6, price=$7, updated_at=$8 WHERE id=$9")
	if err != nil {
		return err
	}

	startTime := time.Now()
	book.UpdatedAt = startTime

	_, err = stmt.Exec(&book.PublisherID, &book.CategoryID, &book.Title, &book.Author, &book.Publication, &book.Stock, &book.Price, &book.UpdatedAt, id)
	if err != nil {
		return err
	}

	return nil
}

func (mb *mysqlBook) DeleteBook(ctx context.Context, id int64) error {
	stmt, err := mb.DB.Prepare("DELETE FROM books WHERE id=$1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
