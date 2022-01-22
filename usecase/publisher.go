package usecase

import (
	"context"
	"winartodev/book-store-be/entity"
	"winartodev/book-store-be/repository"
)

type PublisherUsecase interface {
	GetPublishers(ctx context.Context) ([]entity.Publisher, error)
	GetPublisher(ctx context.Context, id int64) (entity.Publisher, error)
	CreatePublisher(ctx context.Context, publisher *entity.Publisher) error
	UpdatePublisher(ctx context.Context, id int64, publisher *entity.Publisher) error
	DeletePublisher(ctx context.Context, id int64) error
}

type PublisherRepository struct {
	PublisherRepo repository.PublisherRepository
}

func NewPublihserUsecase(repo *PublisherRepository) PublisherUsecase {
	return &PublisherRepository{
		PublisherRepo: repo.PublisherRepo,
	}
}

func (uc *PublisherRepository) GetPublishers(ctx context.Context) ([]entity.Publisher, error) {
	res, err := uc.PublisherRepo.GetPublishers(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *PublisherRepository) GetPublisher(ctx context.Context, id int64) (entity.Publisher, error) {
	res, err := uc.PublisherRepo.GetPublisher(ctx, id)
	if err != nil {
		return entity.Publisher{}, err
	}

	return res, nil
}

func (uc *PublisherRepository) CreatePublisher(ctx context.Context, publisher *entity.Publisher) error {
	err := uc.PublisherRepo.CreatePublisher(ctx, publisher)
	if err != nil {
		return err
	}

	return nil
}

func (uc *PublisherRepository) UpdatePublisher(ctx context.Context, id int64, publisher *entity.Publisher) error {
	err := uc.PublisherRepo.UpdatePublisher(ctx, id, publisher)
	if err != nil {
		return err
	}

	return nil
}

func (uc *PublisherRepository) DeletePublisher(ctx context.Context, id int64) error {
	err := uc.PublisherRepo.DeletePublisher(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
