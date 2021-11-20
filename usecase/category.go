package usecase

import (
	"context"
	bookstorebe "winartodev/book-store-be"
)

type CategoryUsecase struct {
	CategoryRepo bookstorebe.CategoryRepository
}

func NewCategoryUsecase(uc *CategoryUsecase) bookstorebe.CategoryUsecase {
	return &CategoryUsecase{
		CategoryRepo: uc.CategoryRepo,
	}
}

func (uc *CategoryUsecase) GetCategories(ctx context.Context) ([]bookstorebe.Category, error) {
	res, err := uc.CategoryRepo.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *CategoryUsecase) GetCategory(ctx context.Context, id int64) (bookstorebe.Category, error) {
	res, err := uc.CategoryRepo.GetCategory(ctx, id)
	if err != nil {
		return bookstorebe.Category{}, err
	}

	return res, nil
}

func (uc *CategoryUsecase) CreateCategory(ctx context.Context, category *bookstorebe.Category) error {
	err := uc.CategoryRepo.CreateCategory(ctx, category)
	if err != nil {
		return err
	}

	return nil
}

func (uc *CategoryUsecase) UpdateCategory(ctx context.Context, id int64, category *bookstorebe.Category) error {
	err := uc.CategoryRepo.UpdateCategory(ctx, id, category)
	if err != nil {
		return err
	}

	return nil
}

func (uc *CategoryUsecase) DeleteCategory(ctx context.Context, id int64) error {
	err := uc.CategoryRepo.DeleteCategory(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
