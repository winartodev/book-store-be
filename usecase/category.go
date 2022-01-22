package usecase

import (
	"context"
	"winartodev/book-store-be/entity"
	"winartodev/book-store-be/repository"
)

type CategoryUsecase interface {
	GetCategories(ctx context.Context) ([]entity.Category, error)
	GetCategory(ctx context.Context, id int64) (entity.Category, error)
	CreateCategory(ctx context.Context, category *entity.Category) error
	UpdateCategory(ctx context.Context, id int64, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int64) error
}

type CategoryRepository struct {
	CategoryRepo repository.CategoryRepository
}

func NewCategoryUsecase(repo *CategoryRepository) CategoryUsecase {
	return &CategoryRepository{
		CategoryRepo: repo.CategoryRepo,
	}
}

func (r *CategoryRepository) GetCategories(ctx context.Context) ([]entity.Category, error) {
	res, err := r.CategoryRepo.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *CategoryRepository) GetCategory(ctx context.Context, id int64) (entity.Category, error) {
	res, err := r.CategoryRepo.GetCategory(ctx, id)
	if err != nil {
		return entity.Category{}, err
	}

	return res, nil
}

func (r *CategoryRepository) CreateCategory(ctx context.Context, category *entity.Category) error {
	err := r.CategoryRepo.CreateCategory(ctx, category)
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, id int64, category *entity.Category) error {
	err := r.CategoryRepo.UpdateCategory(ctx, id, category)
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, id int64) error {
	err := r.CategoryRepo.DeleteCategory(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
