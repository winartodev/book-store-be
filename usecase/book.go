package usecase

import (
	"context"

	"winartodev/book-store-be/entity"
	"winartodev/book-store-be/repository"
)

type BookUsecase interface {
	GetBooks(ctx context.Context) ([]entity.Book, error)
	GetBook(ctx context.Context, id int64) (entity.Book, error)
	CreateBook(ctx context.Context, book *entity.Book) error
	UpdateBook(ctx context.Context, id int64, book *entity.Book) error
	DeleteBook(ctx context.Context, id int64) error
}

type BookRepository struct {
	BookRepo repository.BookRepository
}

func NewBookUsecase(repo *BookRepository) BookUsecase {
	return &BookRepository{BookRepo: repo.BookRepo}
}

func (repo *BookRepository) GetBooks(ctx context.Context) ([]entity.Book, error) {
	res, err := repo.BookRepo.GetBooks(ctx)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *BookRepository) GetBook(ctx context.Context, id int64) (entity.Book, error) {
	res, err := repo.BookRepo.GetBook(ctx, id)
	if err != nil {
		return entity.Book{}, err
	}

	return res, nil
}

func (repo *BookRepository) CreateBook(ctx context.Context, book *entity.Book) error {
	err := repo.BookRepo.CreateBook(ctx, book)
	if err != nil {
		return err
	}

	return nil
}

func (repo *BookRepository) UpdateBook(ctx context.Context, id int64, book *entity.Book) error {
	err := repo.BookRepo.UpdateBook(ctx, id, book)
	if err != nil {
		return err
	}

	return nil
}

func (repo *BookRepository) DeleteBook(ctx context.Context, id int64) error {
	err := repo.BookRepo.DeleteBook(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
