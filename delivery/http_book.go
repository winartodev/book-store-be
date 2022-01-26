package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"winartodev/book-store-be/entity"
	"winartodev/book-store-be/response"
	"winartodev/book-store-be/usecase"

	"github.com/julienschmidt/httprouter"
)

type BookHandler struct {
	uc usecase.BookUsecase
}

func NewBookHandler(usecase usecase.BookUsecase) BookHandler {
	return BookHandler{
		uc: usecase,
	}
}

func (h *BookHandler) Register(r *httprouter.Router) error {
	if r == nil {
		return errors.New("router cannot be empty")
	}

	r.GET("/bookstore/book", h.GetBooks)
	r.GET("/bookstore/book/:id", h.GetBook)
	r.POST("/bookstore/book", h.CreateBook)
	r.PUT("/bookstore/book/:id", h.UpdateBook)
	r.DELETE("/bookstore/book/:id", h.DeleteBook)

	return nil
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()
	data, err := h.uc.GetBooks(ctx)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return
	}

	if len(data) == 0 {
		response.SuccessResponse(w, http.StatusOK, "Book is empty")
		return
	}

	response.SuccessResponse(w, http.StatusOK, data)
}

func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id, _ := strconv.ParseInt(param.ByName("id"), 10, 64)

	ctx := r.Context()
	data, err := h.uc.GetBook(ctx, id)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return
	}

	if data.ID == 0 {
		response.FailedResponse(w, http.StatusNotFound, fmt.Sprintf("Book ID %d Was Not Found", id))
		return
	}

	response.SuccessResponse(w, http.StatusOK, data)
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var book entity.Book
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&book); err != nil {
		response.FailedResponse(w, 1, err.Error())
		return
	}

	ctx := r.Context()
	err := h.uc.CreateBook(ctx, &book)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusCreated, "Created")
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id, _ := strconv.ParseInt(param.ByName("id"), 10, 64)

	var book entity.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		response.FailedResponse(w, 1, err.Error())
		return
	}

	ctx := r.Context()
	err := h.uc.UpdateBook(ctx, id, &book)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, "Book Has Been Updated")
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id, _ := strconv.ParseInt(param.ByName("id"), 10, 64)

	ctx := r.Context()
	err := h.uc.DeleteBook(ctx, id)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, "Book Has Been Deleted")
}
