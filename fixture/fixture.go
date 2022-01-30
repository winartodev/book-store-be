package fixture

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

const (
	DummyUsername = "username"
	DummyPassword = "password"
)

// HTTPBasicAuth will return http.request with basic authentication
func HTTPBasicAuth(method, path, username, password string, body []byte) *http.Request {
	request := httptest.NewRequest(method, "http://localhost"+path, bytes.NewBuffer(body))
	request.SetBasicAuth(username, password)
	request.Header.Set("Content-Type", "application/json")
	return request
}
