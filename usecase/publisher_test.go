package usecase_test

import (
	"context"
	"errors"
	"testing"
	bookstorebe "winartodev/book-store-be"
	"winartodev/book-store-be/mocks"
	"winartodev/book-store-be/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockPublisherProvider struct {
	publisherRepo *mocks.PublisherRepository
}

func publihserProvider() mockPublisherProvider {
	return mockPublisherProvider{
		publisherRepo: new(mocks.PublisherRepository),
	}
}

func newPublisherUsecase(uc *usecase.PublisherUsecase) bookstorebe.PublisherUsecase {
	return usecase.NewPublihserUsecase(uc)
}

func TestGetPublishers(t *testing.T) {
	testCases := []struct {
		name      string
		publisher []bookstorebe.Publisher
		expRes    []bookstorebe.Publisher
		isError   bool
		wantErr   error
	}{
		{
			name:      "success",
			publisher: []bookstorebe.Publisher{{ID: 1, Name: "Publisher Name", Address: "Publisher Address", PhoneNumber: "123456789"}},
			expRes:    []bookstorebe.Publisher{{ID: 1, Name: "Publisher Name", Address: "Publisher Address", PhoneNumber: "123456789"}},
			isError:   false,
			wantErr:   nil,
		},
		{
			name:      "failed",
			publisher: []bookstorebe.Publisher{},
			expRes:    nil,
			isError:   true,
			wantErr:   errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			prov := publihserProvider()
			prov.publisherRepo.On("GetPublishers", mock.Anything).Return(test.publisher, test.wantErr)

			publisherUsecase := newPublisherUsecase(&usecase.PublisherUsecase{prov.publisherRepo})
			ctx := context.Background()
			res, err := publisherUsecase.GetPublishers(ctx)

			assert.Equal(t, test.isError, err != nil)
			if !test.isError {
				assert.NotNil(t, res)
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
				assert.Nil(t, res)
			}
		})
	}
}

func TestGetPublisher(t *testing.T) {
	testCases := []struct {
		name      string
		ID        int64
		publsiher bookstorebe.Publisher
		expRes    bookstorebe.Publisher
		isError   bool
		wantErr   error
	}{
		{
			name:      "success",
			ID:        1,
			publsiher: bookstorebe.Publisher{ID: 1, Name: "Publisher Name", Address: "Publisher Address", PhoneNumber: "123456789"},
			expRes:    bookstorebe.Publisher{ID: 1, Name: "Publisher Name", Address: "Publisher Address", PhoneNumber: "123456789"},
			isError:   false,
			wantErr:   nil,
		},
		{
			name:      "failed",
			ID:        1,
			publsiher: bookstorebe.Publisher{},
			expRes:    bookstorebe.Publisher{},
			isError:   true,
			wantErr:   errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			prov := publihserProvider()
			prov.publisherRepo.On("GetPublisher", mock.Anything, mock.AnythingOfType("int64")).Return(test.publsiher, test.wantErr)

			publisherUsecase := newPublisherUsecase(&usecase.PublisherUsecase{prov.publisherRepo})
			ctx := context.Background()
			res, err := publisherUsecase.GetPublisher(ctx, test.ID)

			assert.Equal(t, test.isError, err != nil)
			if !test.isError {
				assert.NotNil(t, res)
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestCreatePublisher(t *testing.T) {
	testCases := []struct {
		name      string
		publisher bookstorebe.Publisher
		isError   bool
		wantErr   error
	}{
		{
			name:      "success",
			publisher: bookstorebe.Publisher{ID: 1, Name: "Publisher Name", Address: "Publisher Address", PhoneNumber: "123456789"},
			isError:   false,
			wantErr:   nil,
		},
		{
			name:      "failed",
			publisher: bookstorebe.Publisher{},
			isError:   true,
			wantErr:   errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			prov := publihserProvider()
			prov.publisherRepo.On("CreatePublisher", mock.Anything, mock.Anything).Return(test.wantErr)

			publisherUsecase := newPublisherUsecase(&usecase.PublisherUsecase{prov.publisherRepo})
			ctx := context.Background()
			err := publisherUsecase.CreatePublisher(ctx, &test.publisher)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}

func TestUpdatePublisher(t *testing.T) {
	testCases := []struct {
		name      string
		ID        int64
		publisher bookstorebe.Publisher
		isError   bool
		wantErr   error
	}{
		{
			name:      "success",
			ID:        1,
			publisher: bookstorebe.Publisher{ID: 1, Name: "Publisher Name", Address: "Publisher Address", PhoneNumber: "123456789"},
			isError:   false,
			wantErr:   nil,
		},
		{
			name:      "failed",
			ID:        1,
			publisher: bookstorebe.Publisher{},
			isError:   true,
			wantErr:   errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			prov := publihserProvider()
			prov.publisherRepo.On("UpdatePublisher", mock.Anything, mock.AnythingOfType("int64"), mock.Anything).Return(test.wantErr)

			publisherUsecase := newPublisherUsecase(&usecase.PublisherUsecase{prov.publisherRepo})
			ctx := context.Background()
			err := publisherUsecase.UpdatePublisher(ctx, test.ID, &test.publisher)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}

func TestDeletePublisher(t *testing.T) {
	testCases := []struct {
		name      string
		ID        int64
		isError   bool
		wantError error
	}{
		{
			name:      "success",
			ID:        1,
			isError:   false,
			wantError: nil,
		},
		{
			name:      "failed",
			ID:        1,
			isError:   true,
			wantError: errors.New("Dummy Error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			prov := publihserProvider()
			prov.publisherRepo.On("DeletePublisher", mock.Anything, mock.AnythingOfType("int64")).Return(test.wantError)

			publisherUsecase := newPublisherUsecase(&usecase.PublisherUsecase{prov.publisherRepo})
			ctx := context.Background()
			err := publisherUsecase.DeletePublisher(ctx, test.ID)

			assert.Equal(t, test.isError, err != nil)
		})
	}
}
