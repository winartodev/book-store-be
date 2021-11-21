package repository

import (
	"context"
	"database/sql"
	bookstorebe "winartodev/book-store-be"
)

type mysqlCategory struct {
	DB *sql.DB
}

func NewMysqlCategory(db *sql.DB) bookstorebe.CategoryRepository {
	return &mysqlCategory{DB: db}
}

func (mc *mysqlCategory) GetCategories(ctx context.Context) ([]bookstorebe.Category, error) {
	var categories []bookstorebe.Category

	rows, err := mc.DB.Query("SELECT * FROM category")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category bookstorebe.Category

		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (mc *mysqlCategory) GetCategory(ctx context.Context, id int64) (bookstorebe.Category, error) {
	var category bookstorebe.Category

	err := mc.DB.QueryRow("SELECT * FROM category WHERE id=?", id).Scan(&category.ID, &category.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return bookstorebe.Category{}, nil
		}
		return bookstorebe.Category{}, err
	}

	return category, nil
}

func (mc *mysqlCategory) CreateCategory(ctx context.Context, category *bookstorebe.Category) error {
	stmt, err := mc.DB.Prepare("INSERT INTO category VALUES(NULL, ?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&category.Name)
	if err != nil {
		return err
	}

	if row, _ := res.RowsAffected(); row == 1 {
		return nil
	}

	return nil
}

func (mc *mysqlCategory) UpdateCategory(ctx context.Context, id int64, category *bookstorebe.Category) error {
	stmt, err := mc.DB.Prepare("UPDATE category SET name=? WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&category.Name, id)
	if err != nil {
		return err
	}

	return nil
}

func (mc *mysqlCategory) DeleteCategory(ctx context.Context, id int64) error {
	stmt, err := mc.DB.Prepare("DELETE category WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
