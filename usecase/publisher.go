package usecase

import (
	"context"
	bookstorebe "winartodev/book-store-be"
)

type PublisherUsecase struct {
	PublisherRepo bookstorebe.PublisherRepository
}

func NewPublihserUsecase(uc *PublisherUsecase) bookstorebe.PublisherUsecase {
	return &PublisherUsecase{
		PublisherRepo: uc.PublisherRepo,
	}
}

func (uc *PublisherUsecase) GetPublishers(ctx context.Context) ([]bookstorebe.Publisher, error) {
	res, err := uc.PublisherRepo.GetPublishers(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *PublisherUsecase) GetPublisher(ctx context.Context, id int64) (bookstorebe.Publisher, error) {
	res, err := uc.PublisherRepo.GetPublisher(ctx, id)
	if err != nil {
		return bookstorebe.Publisher{}, err
	}

	return res, nil
}

func (uc *PublisherUsecase) CreatePublisher(ctx context.Context, publisher *bookstorebe.Publisher) error {
	err := uc.PublisherRepo.CreatePublisher(ctx, publisher)
	if err != nil {
		return err
	}

	return nil
}

func (uc *PublisherUsecase) UpdatePublisher(ctx context.Context, id int64, publisher *bookstorebe.Publisher) error {
	err := uc.PublisherRepo.UpdatePublisher(ctx, id, publisher)
	if err != nil {
		return err
	}

	return nil
}

func (uc *PublisherUsecase) DeletePublisher(ctx context.Context, id int64) error {
	err := uc.PublisherRepo.DeletePublisher(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
