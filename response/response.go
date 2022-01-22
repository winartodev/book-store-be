package response

import (
	"encoding/json"
	"net/http"
)

type successBody struct {
	Status   string      `json:"status"`
	HttpCode int         `json:"status_code"`
	Data     interface{} `json:"data"`
}

type failedBody struct {
	Status   string `json:"status"`
	HttpCode int    `json:"status_code"`
	Message  string `json:"message"`
}

func Write(w http.ResponseWriter, result interface{}, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(result)
}

func statusOK(status int, data interface{}) successBody {
	return successBody{
		Status:   http.StatusText(status),
		HttpCode: status,
		Data:     data,
	}
}

func statusFailed(status int, message string) failedBody {
	return failedBody{
		Status:   http.StatusText(status),
		HttpCode: status,
		Message:  message,
	}
}

func SuccessResponse(w http.ResponseWriter, status int, data interface{}) {
	Write(w, statusOK(status, data), status)
}

func FailedResponse(w http.ResponseWriter, status int, message string) {
	Write(w, statusFailed(status, message), status)
}
