package handler

import (
	"net/http"
	"winartodev/book-store-be/response"

	"github.com/julienschmidt/httprouter"
)

type Registration interface {
	Register(r *httprouter.Router) error
}

func NotFound(w http.ResponseWriter, _ *http.Request) {
	response.Write(w, "endpoint not found", http.StatusNotFound)
}

func NewHandler(register ...Registration) http.Handler {
	router := httprouter.New()
	router.HandleMethodNotAllowed = false

	for _, r := range register {
		r.Register(router)
	}

	router.NotFound = http.HandlerFunc(NotFound)
	return router
}
