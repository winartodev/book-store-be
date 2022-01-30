package handler

import (
	"fmt"
	"net/http"
	"winartodev/book-store-be/middleware"
	"winartodev/book-store-be/response"

	"github.com/julienschmidt/httprouter"
)

type Registration interface {
	Register(r *httprouter.Router) error
}

func NotFound(w http.ResponseWriter, _ *http.Request) {
	response.Write(w, "endpoint not found", http.StatusNotFound)
}

func Decorate(handle middleware.HandleWithError, ds ...middleware.Decorator) httprouter.Handle {
	return middleware.HTTP(middleware.ApplyDecorators(handle, ds...))
}

func NewHandler(register ...Registration) http.Handler {
	router := httprouter.New()
	router.HandleMethodNotAllowed = false

	router.GET("/healthz", healthz)

	for _, r := range register {
		r.Register(router)
	}

	router.NotFound = http.HandlerFunc(NotFound)
	return router
}

func healthz(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "OK")
}
