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

type CategoryHandler struct {
	uc usecase.CategoryUsecase
}

func NewCategoryHandler(usecase usecase.CategoryUsecase) CategoryHandler {
	return CategoryHandler{
		uc: usecase,
	}
}

func (h *CategoryHandler) Register(r *httprouter.Router) error {
	if r == nil {
		return errors.New("router cannot be empty")
	}

	r.GET("/bookstore/category", h.GetCategories)
	r.GET("/bookstore/category/:id", h.GetCategory)
	r.POST("/bookstore/category", h.CreateCategory)
	r.PUT("/bookstore/category/:id", h.UpdateCategory)
	r.DELETE("/bookstore/category/:id", h.DeleteCategory)

	return nil
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()
	data, err := h.uc.GetCategories(ctx)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return
	}

	if len(data) == 0 {
		response.SuccessResponse(w, http.StatusOK, "Category is empty")
		return
	}

	response.SuccessResponse(w, http.StatusOK, data)
}

func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id, _ := strconv.ParseInt(param.ByName("id"), 10, 64)

	ctx := r.Context()
	data, err := h.uc.GetCategory(ctx, id)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return
	}

	if data.ID == 0 {
		response.FailedResponse(w, http.StatusNotFound, fmt.Sprintf("Category ID %d Was Not Found", id))
		return
	}

	response.SuccessResponse(w, http.StatusOK, data)
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	var category entity.Category
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&category); err != nil {
		response.FailedResponse(w, 1, err.Error())
		return
	}

	ctx := r.Context()
	err := h.uc.CreateCategory(ctx, &category)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusCreated, "Created")
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id, _ := strconv.ParseInt(param.ByName("id"), 10, 64)

	var category entity.Category
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		response.FailedResponse(w, 1, err.Error())
		return
	}

	ctx := r.Context()
	err := h.uc.UpdateCategory(ctx, id, &category)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, "Category Has Been Updated")
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id, _ := strconv.ParseInt(param.ByName("id"), 10, 64)

	ctx := r.Context()
	err := h.uc.DeleteCategory(ctx, id)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, "Category Has Been Deleted")
}
