package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"winartodev/book-store-be/entity"
	"winartodev/book-store-be/handler"
	"winartodev/book-store-be/middleware"
	"winartodev/book-store-be/response"
	"winartodev/book-store-be/usecase"

	"github.com/julienschmidt/httprouter"
)

type BookHandler struct {
	uc       usecase.BookUsecase
	username string
	password string
}

func NewBookHandler(usecase usecase.BookUsecase, username string, password string) BookHandler {
	return BookHandler{
		uc:       usecase,
		username: username,
		password: password,
	}
}

func (h *BookHandler) Register(r *httprouter.Router) error {
	if r == nil {
		return errors.New("router cannot be empty")
	}

	r.GET("/bookstore/book", handler.Decorate(h.GetBooks, middleware.MiddlewareBasicAuth(h.username, h.password)))
	r.GET("/bookstore/book/:id", handler.Decorate(h.GetBook, middleware.MiddlewareBasicAuth(h.username, h.password)))
	r.POST("/bookstore/book", handler.Decorate(h.CreateBook, middleware.MiddlewareBasicAuth(h.username, h.password)))
	r.PUT("/bookstore/book/:id", handler.Decorate(h.UpdateBook, middleware.MiddlewareBasicAuth(h.username, h.password)))
	r.DELETE("/bookstore/book/:id", handler.Decorate(h.DeleteBook, middleware.MiddlewareBasicAuth(h.username, h.password)))

	return nil
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	ctx := r.Context()
	data, err := h.uc.GetBooks(ctx)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return err
	}

	if len(data) == 0 {
		response.SuccessResponse(w, http.StatusOK, "Book is empty")
		return nil
	}

	response.SuccessResponse(w, http.StatusOK, data)
	return nil
}

func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request, param httprouter.Params) error {
	id, _ := strconv.ParseInt(param.ByName("id"), 10, 64)

	ctx := r.Context()
	data, err := h.uc.GetBook(ctx, id)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return err
	}

	if data.ID == 0 {
		response.FailedResponse(w, http.StatusNotFound, fmt.Sprintf("Book ID %d Was Not Found", id))
		return nil
	}

	response.SuccessResponse(w, http.StatusOK, data)
	return nil
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	var book entity.Book
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&book); err != nil {
		response.FailedResponse(w, 1, err.Error())
		return err
	}

	ctx := r.Context()
	err := h.uc.CreateBook(ctx, &book)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return nil
	}

	response.SuccessResponse(w, http.StatusCreated, "Created")
	return nil
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request, param httprouter.Params) error {
	id, _ := strconv.ParseInt(param.ByName("id"), 10, 64)

	var book entity.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		response.FailedResponse(w, 1, err.Error())
		return err
	}

	ctx := r.Context()
	err := h.uc.UpdateBook(ctx, id, &book)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return nil
	}

	response.SuccessResponse(w, http.StatusOK, "Book Has Been Updated")
	return nil
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request, param httprouter.Params) error {
	id, _ := strconv.ParseInt(param.ByName("id"), 10, 64)

	ctx := r.Context()
	err := h.uc.DeleteBook(ctx, id)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return err
	}

	response.SuccessResponse(w, http.StatusOK, "Book Has Been Deleted")
	return nil
}
