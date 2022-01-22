package repository

import (
	"context"
	"database/sql"
	"time"
	"winartodev/book-store-be/entity"
)

type CategoryRepository interface {
	// seller
	GetCategories(ctx context.Context) ([]entity.Category, error)
	GetCategory(ctx context.Context, id int64) (entity.Category, error)
	CreateCategory(ctx context.Context, category *entity.Category) error
	UpdateCategory(ctx context.Context, id int64, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int64) error
}

type mysqlCategory struct {
	DB *sql.DB
}

func NewMysqlCategory(db *sql.DB) CategoryRepository {
	return &mysqlCategory{DB: db}
}

func (mc *mysqlCategory) GetCategories(ctx context.Context) ([]entity.Category, error) {
	var categories []entity.Category

	rows, err := mc.DB.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category entity.Category

		err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (mc *mysqlCategory) GetCategory(ctx context.Context, id int64) (entity.Category, error) {
	var category entity.Category

	err := mc.DB.QueryRow("SELECT * FROM categories WHERE id=$1", id).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Category{}, nil
		}
		return entity.Category{}, err
	}

	return category, nil
}

func (mc *mysqlCategory) CreateCategory(ctx context.Context, category *entity.Category) error {
	stmt, err := mc.DB.Prepare("INSERT INTO categories (name, created_at, updated_at) VALUES($1, $2, $3)")
	if err != nil {
		return err
	}

	startTime := time.Now()
	category.CreatedAt = startTime
	category.UpdatedAt = startTime

	res, err := stmt.Exec(&category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return err
	}

	if row, _ := res.RowsAffected(); row == 1 {
		return nil
	}

	return nil
}

func (mc *mysqlCategory) UpdateCategory(ctx context.Context, id int64, category *entity.Category) error {
	stmt, err := mc.DB.Prepare("UPDATE categories SET name=$1, updated_at=$2 WHERE id=$3")
	if err != nil {
		return err
	}

	category.UpdatedAt = time.Now()
	_, err = stmt.Exec(&category.Name, &category.UpdatedAt, id)
	if err != nil {
		return err
	}

	return nil
}

func (mc *mysqlCategory) DeleteCategory(ctx context.Context, id int64) error {
	stmt, err := mc.DB.Prepare("DELETE FROM categories WHERE id=$1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
