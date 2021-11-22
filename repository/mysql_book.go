package repository

import (
	"context"
	"database/sql"
	bookstorebe "winartodev/book-store-be"
)

type mysqlBook struct {
	DB *sql.DB
}

func NewMysqlBook(db *sql.DB) bookstorebe.BookRepository {
	return &mysqlBook{DB: db}
}

func (mb *mysqlBook) GetBooks(ctx context.Context) ([]bookstorebe.Book, error) {
	var books []bookstorebe.Book

	rows, err := mb.DB.Query("SELECT * FROM book")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book bookstorebe.Book

		err := rows.Scan(&book.ID, &book.PublisherID, &book.CategoryID, &book.Title, &book.Author, &book.Publication, &book.Stock)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (mb *mysqlBook) GetBook(ctx context.Context, id int64) (bookstorebe.Book, error) {
	var book bookstorebe.Book

	err := mb.DB.QueryRow("SELECT * FROM book WHERE id=?", id).Scan(&book.ID, &book.PublisherID, &book.CategoryID, &book.Title, &book.Author, &book.Publication, &book.Stock)
	if err != nil {
		if err == sql.ErrNoRows {
			return bookstorebe.Book{}, nil
		}
		return bookstorebe.Book{}, err
	}

	return book, nil
}

func (mb *mysqlBook) CreateBook(ctx context.Context, book *bookstorebe.Book) error {
	stmt, err := mb.DB.Prepare("INSERT INTO book VALUES(NULL, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&book.PublisherID, &book.CategoryID, &book.Title, &book.Author, &book.Publication, &book.Stock)
	if err != nil {
		return err
	}

	if row, _ := res.RowsAffected(); row == 1 {
		return nil
	}

	return nil
}

func (mb *mysqlBook) UpdateBook(ctx context.Context, id int64, book *bookstorebe.Book) error {
	stmt, err := mb.DB.Prepare("UPDATE book SET publisher_id=?, category_id=?, title=?, author=?, year_of_publication=?, stock=? WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&book.PublisherID, &book.CategoryID, &book.Title, &book.Author, &book.Publication, &book.Stock, id)
	if err != nil {
		return err
	}

	return nil
}

func (mb *mysqlBook) DeleteBook(ctx context.Context, id int64) error {
	stmt, err := mb.DB.Prepare("DELETE FROM book WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
