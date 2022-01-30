package middleware

import (
	"errors"
	"net/http"
	"winartodev/book-store-be/response"

	"github.com/julienschmidt/httprouter"
)

// HandleWithError is httprouter.Handle that return error
type HandleWithError func(http.ResponseWriter, *http.Request, httprouter.Params) error

// Decorator decorate HandleWithError
type Decorator func(handle HandleWithError) HandleWithError

// ApplyDecorators will return HandleWithError
func ApplyDecorators(handle HandleWithError, ds ...Decorator) HandleWithError {
	for _, d := range ds {
		handle = d(handle)
	}

	return handle
}

// HTTP runs HandleWithError and converts it to httprouter.Handle
func HTTP(handle HandleWithError) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		handle(w, r, params)
	}
}

// MiddlewareBasicAuth will authenticate with BasicAuth
func MiddlewareBasicAuth(username, password string) Decorator {
	return func(handle HandleWithError) HandleWithError {
		return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
			u, p, ok := r.BasicAuth()

			if ok && u == username && p == password {
				return handle(w, r, params)
			}

			w.Header().Set("WWW-Authenticate", "Basic")
			response.FailedResponse(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return errors.New("unauthorized basic auth")
		}
	}
}
