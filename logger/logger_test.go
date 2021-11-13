package logger_test

import (
	"errors"
	"testing"
	"winartodev/book-store-be/logger"
)

func TestInit(t *testing.T) {
	logger.Init()
}

func TestInfo(t *testing.T) {
	testCases := []struct {
		name    string
		message string
		fileds  logger.Fields
	}{
		{
			name:    "Logger Info Test",
			message: "Ïnfo",
			fileds:  logger.Fields{"key1": "val1"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			logger.Info(test.message, test.fileds)
		})
	}
}

func TestError(t *testing.T) {
	testCases := []struct {
		name   string
		err    error
		fields logger.Fields
	}{
		{
			name:   "Logger Error Test",
			err:    errors.New("error logs"),
			fields: logger.Fields{"key1": "val1"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			logger.Error(test.err, test.fields)
		})
	}
}
