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

type CategoryHandler struct {
	uc       usecase.CategoryUsecase
	username string
	passwrod string
}

func NewCategoryHandler(usecase usecase.CategoryUsecase, username string, password string) CategoryHandler {
	return CategoryHandler{
		uc:       usecase,
		username: username,
		passwrod: password,
	}
}

func (h *CategoryHandler) Register(r *httprouter.Router) error {
	if r == nil {
		return errors.New("router cannot be empty")
	}

	r.GET("/bookstore/category", handler.Decorate(h.GetCategories, middleware.MiddlewareBasicAuth(h.username, h.passwrod)))
	r.GET("/bookstore/category/:id", handler.Decorate(h.GetCategory, middleware.MiddlewareBasicAuth(h.username, h.passwrod)))
	r.POST("/bookstore/category", handler.Decorate(h.CreateCategory, middleware.MiddlewareBasicAuth(h.username, h.passwrod)))
	r.PUT("/bookstore/category/:id", handler.Decorate(h.UpdateCategory, middleware.MiddlewareBasicAuth(h.username, h.passwrod)))
	r.DELETE("/bookstore/category/:id", handler.Decorate(h.DeleteCategory, middleware.MiddlewareBasicAuth(h.username, h.passwrod)))

	return nil
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	ctx := r.Context()
	data, err := h.uc.GetCategories(ctx)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return err
	}

	if len(data) == 0 {
		response.SuccessResponse(w, http.StatusOK, "Category is empty")
		return nil
	}

	response.SuccessResponse(w, http.StatusOK, data)
	return nil
}

func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request, param httprouter.Params) error {
	id, _ := strconv.ParseInt(param.ByName("id"), 10, 64)

	ctx := r.Context()
	data, err := h.uc.GetCategory(ctx, id)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return err
	}

	if data.ID == 0 {
		response.FailedResponse(w, http.StatusNotFound, fmt.Sprintf("Category ID %d Was Not Found", id))
		return nil
	}

	response.SuccessResponse(w, http.StatusOK, data)
	return nil
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request, param httprouter.Params) error {
	var category entity.Category
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&category); err != nil {
		response.FailedResponse(w, 1, err.Error())
		return err
	}

	ctx := r.Context()
	err := h.uc.CreateCategory(ctx, &category)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return err
	}

	response.SuccessResponse(w, http.StatusCreated, "Created")
	return nil
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request, param httprouter.Params) error {
	id, _ := strconv.ParseInt(param.ByName("id"), 10, 64)

	var category entity.Category
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		response.FailedResponse(w, 1, err.Error())
		return err
	}

	ctx := r.Context()
	err := h.uc.UpdateCategory(ctx, id, &category)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return err
	}

	response.SuccessResponse(w, http.StatusOK, "Category Has Been Updated")
	return nil
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request, param httprouter.Params) error {
	id, _ := strconv.ParseInt(param.ByName("id"), 10, 64)

	ctx := r.Context()
	err := h.uc.DeleteCategory(ctx, id)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return err
	}

	response.SuccessResponse(w, http.StatusOK, "Category Has Been Deleted")
	return nil
}
