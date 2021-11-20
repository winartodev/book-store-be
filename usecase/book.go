package usecase

import (
	"context"
	bookstorebe "winartodev/book-store-be"
)

type BookUsecase struct {
	BookRepo bookstorebe.BookRepository
}

func NewBookUsecase(uc *BookUsecase) bookstorebe.BookUsecase {
	return &BookUsecase{BookRepo: uc.BookRepo}
}

func (uc *BookUsecase) GetBooks(ctx context.Context) ([]bookstorebe.Book, error) {
	res, err := uc.BookRepo.GetBooks(ctx)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *BookUsecase) GetBook(ctx context.Context, id int64) (bookstorebe.Book, error) {
	res, err := uc.BookRepo.GetBook(ctx, id)
	if err != nil {
		return bookstorebe.Book{}, err
	}

	return res, nil
}

func (uc *BookUsecase) CreateBook(ctx context.Context, book *bookstorebe.Book) error {
	err := uc.BookRepo.CreateBook(ctx, book)
	if err != nil {
		return err
	}

	return nil
}

func (uc *BookUsecase) UpdateBook(ctx context.Context, id int64, book *bookstorebe.Book) error {
	err := uc.BookRepo.UpdateBook(ctx, id, book)
	if err != nil {
		return err
	}

	return nil
}

func (uc *BookUsecase) DeleteBook(ctx context.Context, id int64) error {
	err := uc.BookRepo.DeleteBook(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
