package delivery_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"winartodev/book-store-be/delivery"
	"winartodev/book-store-be/entity"
	"winartodev/book-store-be/fixture"
	"winartodev/book-store-be/handler"
	"winartodev/book-store-be/logger"
	"winartodev/book-store-be/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func init() {
	logger.Init()
}

func newPublisherHandler() (http.Handler, *mocks.PublisherUsecase) {
	os.Setenv("BOOKSTORE_USERNAME", fixture.DummyUsername)
	os.Setenv("BOOKSTORE_PASSWORD", fixture.DummyPassword)

	username := fixture.DummyUsername
	password := fixture.DummyPassword

	uc := new(mocks.PublisherUsecase)
	publisher := delivery.NewPublisherHandler(uc, username, password)
	h := handler.NewHandler(&publisher)
	return h, uc
}

func TestGetPublishers(t *testing.T) {
	testCases := []struct {
		name      string
		publisher []entity.Publisher
		wantError bool
		getErr    error
	}{
		{
			name: "success",
			publisher: []entity.Publisher{
				{
					ID:          1,
					Name:        "asdf",
					Address:     "asdf",
					PhoneNumber: "123",
				},
			},
		},
		{
			name:      "success publisher data is empty",
			publisher: []entity.Publisher{},
			wantError: false,
		},
		{
			name:      "failed to get publisher data",
			publisher: []entity.Publisher{},
			wantError: true,
			getErr:    errors.New("failed to get publisher data"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			handler, publisher := newPublisherHandler()
			publisher.On("GetPublishers", mock.Anything).Return(test.publisher, test.getErr)

			recoder := httptest.NewRecorder()
			request := fixture.HTTPBasicAuth(http.MethodGet, "/bookstore/publisher", fixture.DummyUsername, fixture.DummyPassword, nil)
			handler.ServeHTTP(recoder, request)

			assert.Equal(t, test.wantError, recoder.Code != http.StatusOK)
		})
	}
}

func TestGetPublisher(t *testing.T) {
	testCases := []struct {
		name      string
		id        int64
		publisher entity.Publisher
		wantError bool
		getErr    error
	}{
		{
			name: "success",
			id:   1,
			publisher: entity.Publisher{
				ID:          1,
				Name:        "asdf",
				Address:     "asdf",
				PhoneNumber: "123",
			},
			wantError: false,
			getErr:    nil,
		},
		{
			name:      "failed publisher data not found",
			id:        1,
			publisher: entity.Publisher{},
			wantError: true,
		},
		{
			name:      "failed to get publisher data",
			id:        1,
			publisher: entity.Publisher{},
			wantError: true,
			getErr:    errors.New("failed to get publisher data"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			handler, publisher := newPublisherHandler()
			publisher.On("GetPublisher", mock.Anything, mock.Anything).Return(test.publisher, test.getErr)

			recoder := httptest.NewRecorder()
			request := fixture.HTTPBasicAuth(http.MethodGet, fmt.Sprintf("/bookstore/publisher/%d", test.id), fixture.DummyUsername, fixture.DummyPassword, nil)
			handler.ServeHTTP(recoder, request)

			assert.Equal(t, test.wantError, recoder.Code != http.StatusOK)
		})
	}
}

func TestCreatePublisher(t *testing.T) {
	testCases := []struct {
		name      string
		publisher entity.Publisher
		wantError bool
		createErr error
	}{
		{
			name: "success",
			publisher: entity.Publisher{
				ID:          1,
				Name:        "asdf",
				Address:     "asdf",
				PhoneNumber: "123",
			},
		},
		{
			name:      "failed to create publisher data",
			publisher: entity.Publisher{},
			wantError: true,
			createErr: errors.New("failed to create publisher data"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			handler, publisher := newPublisherHandler()
			publisher.On("CreatePublisher", mock.Anything, mock.Anything).Return(test.createErr)

			body, _ := json.Marshal(test.publisher)
			recoder := httptest.NewRecorder()
			request := fixture.HTTPBasicAuth(http.MethodPost, "/bookstore/publisher", fixture.DummyUsername, fixture.DummyPassword, body)
			handler.ServeHTTP(recoder, request)
			fmt.Print(recoder.Body)
			assert.Equal(t, test.wantError, recoder.Code != http.StatusCreated)
		})
	}
}

func TestUpdatePublisher(t *testing.T) {
	testCases := []struct {
		name      string
		id        int64
		publisher entity.Publisher
		wantError bool
		updateErr error
	}{
		{
			name: "success",
			id:   1,
			publisher: entity.Publisher{
				ID:          1,
				Name:        "asdf",
				Address:     "asdf",
				PhoneNumber: "123",
			},
			wantError: false,
			updateErr: nil,
		},
		{
			name:      "failed to update publisher data",
			id:        1,
			publisher: entity.Publisher{},
			wantError: true,
			updateErr: errors.New("failed to update publisher data"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			handler, publisher := newPublisherHandler()
			publisher.On("UpdatePublisher", mock.Anything, mock.Anything, mock.Anything).Return(test.updateErr)

			body, _ := json.Marshal(test.publisher)
			recoder := httptest.NewRecorder()
			request := fixture.HTTPBasicAuth(http.MethodPut, fmt.Sprintf("/bookstore/publisher/%d", test.id), fixture.DummyUsername, fixture.DummyPassword, body)

			handler.ServeHTTP(recoder, request)

			assert.Equal(t, test.wantError, recoder.Code != http.StatusOK)
		})
	}
}

func TestDeletePublisher(t *testing.T) {
	testCases := []struct {
		name      string
		id        int64
		wantError bool
		deleteErr error
	}{
		{
			name:      "success",
			id:        1,
			wantError: false,
			deleteErr: nil,
		},
		{
			name:      "failed to delete publisher data",
			id:        1,
			wantError: true,
			deleteErr: errors.New("failed to update publisher data"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			handler, publisher := newPublisherHandler()
			publisher.On("DeletePublisher", mock.Anything, mock.Anything).Return(test.deleteErr)

			recoder := httptest.NewRecorder()
			request := fixture.HTTPBasicAuth(http.MethodDelete, fmt.Sprintf("/bookstore/publisher/%d", test.id), fixture.DummyUsername, fixture.DummyPassword, nil)
			handler.ServeHTTP(recoder, request)

			assert.Equal(t, test.wantError, recoder.Code != http.StatusOK)
		})
	}
}
