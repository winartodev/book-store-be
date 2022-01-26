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

type PublsiherHandler struct {
	uc usecase.PublisherUsecase
}

func NewPublisherHandler(usecase usecase.PublisherUsecase) PublsiherHandler {
	return PublsiherHandler{
		uc: usecase,
	}
}

func (h *PublsiherHandler) Register(r *httprouter.Router) error {
	if r == nil {
		return errors.New("router cannot be empty")
	}

	r.GET("/bookstore/publisher", h.GetPublishers)
	r.GET("/bookstore/publisher/:id", h.GetPublisher)
	r.POST("/bookstore/publisher", h.CreatePublisher)
	r.PUT("/bookstore/publisher/:id", h.UpdatePublisher)
	r.DELETE("/bookstore/publisher/:id", h.DeletePublisher)

	return nil
}

func (h *PublsiherHandler) GetPublishers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()
	data, err := h.uc.GetPublishers(ctx)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return
	}

	if len(data) == 0 {
		response.SuccessResponse(w, http.StatusOK, "Publisher is empty")
		return
	}

	response.SuccessResponse(w, http.StatusOK, data)
}

func (h *PublsiherHandler) GetPublisher(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id, _ := strconv.ParseInt(param.ByName("id"), 10, 64)

	ctx := r.Context()
	data, err := h.uc.GetPublisher(ctx, id)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return
	}

	if data.ID == 0 {
		response.FailedResponse(w, http.StatusNotFound, fmt.Sprintf("Publisher ID %d Was Not Found", id))
		return
	}

	response.SuccessResponse(w, http.StatusOK, data)
}

func (h *PublsiherHandler) CreatePublisher(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var publisher entity.Publisher
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&publisher); err != nil {
		response.FailedResponse(w, 1, err.Error())
		return
	}

	ctx := r.Context()
	err := h.uc.CreatePublisher(ctx, &publisher)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusCreated, "Created")
}

func (h *PublsiherHandler) UpdatePublisher(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id, _ := strconv.ParseInt(param.ByName("id"), 10, 64)

	var publisher entity.Publisher
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&publisher); err != nil {
		response.FailedResponse(w, 1, err.Error())
		return
	}

	ctx := r.Context()
	err := h.uc.UpdatePublisher(ctx, id, &publisher)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, "Publisher Has Been Updated")
}

func (h *PublsiherHandler) DeletePublisher(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id, _ := strconv.ParseInt(param.ByName("id"), 10, 64)

	ctx := r.Context()
	err := h.uc.DeletePublisher(ctx, id)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, "Publisher Has Been Deleted")
}
