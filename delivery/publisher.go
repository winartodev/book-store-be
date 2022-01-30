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

type PublsiherHandler struct {
	uc       usecase.PublisherUsecase
	username string
	pasword  string
}

func NewPublisherHandler(usecase usecase.PublisherUsecase, username string, password string) PublsiherHandler {
	return PublsiherHandler{
		uc:       usecase,
		username: username,
		pasword:  password,
	}
}

func (h *PublsiherHandler) Register(r *httprouter.Router) error {
	if r == nil {
		return errors.New("router cannot be empty")
	}

	r.GET("/bookstore/publisher", handler.Decorate(h.GetPublishers, middleware.MiddlewareBasicAuth(h.username, h.pasword)))
	r.GET("/bookstore/publisher/:id", handler.Decorate(h.GetPublisher, middleware.MiddlewareBasicAuth(h.username, h.pasword)))
	r.POST("/bookstore/publisher", handler.Decorate(h.CreatePublisher, middleware.MiddlewareBasicAuth(h.username, h.pasword)))
	r.PUT("/bookstore/publisher/:id", handler.Decorate(h.UpdatePublisher, middleware.MiddlewareBasicAuth(h.username, h.pasword)))
	r.DELETE("/bookstore/publisher/:id", handler.Decorate(h.DeletePublisher, middleware.MiddlewareBasicAuth(h.username, h.pasword)))

	return nil
}

func (h *PublsiherHandler) GetPublishers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	ctx := r.Context()
	data, err := h.uc.GetPublishers(ctx)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return err
	}

	if len(data) == 0 {
		response.SuccessResponse(w, http.StatusOK, "Publisher is empty")
		return nil
	}

	response.SuccessResponse(w, http.StatusOK, data)
	return nil
}

func (h *PublsiherHandler) GetPublisher(w http.ResponseWriter, r *http.Request, param httprouter.Params) error {
	id, _ := strconv.ParseInt(param.ByName("id"), 10, 64)

	ctx := r.Context()
	data, err := h.uc.GetPublisher(ctx, id)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return err
	}

	if data.ID == 0 {
		response.FailedResponse(w, http.StatusNotFound, fmt.Sprintf("Publisher ID %d Was Not Found", id))
		return nil
	}

	response.SuccessResponse(w, http.StatusOK, data)
	return nil
}

func (h *PublsiherHandler) CreatePublisher(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	var publisher entity.Publisher
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&publisher); err != nil {
		response.FailedResponse(w, 1, err.Error())
		return err
	}

	ctx := r.Context()
	err := h.uc.CreatePublisher(ctx, &publisher)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return nil
	}

	response.SuccessResponse(w, http.StatusCreated, "Created")
	return nil
}

func (h *PublsiherHandler) UpdatePublisher(w http.ResponseWriter, r *http.Request, param httprouter.Params) error {
	id, _ := strconv.ParseInt(param.ByName("id"), 10, 64)

	var publisher entity.Publisher
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&publisher); err != nil {
		response.FailedResponse(w, 1, err.Error())
		return err
	}

	ctx := r.Context()
	err := h.uc.UpdatePublisher(ctx, id, &publisher)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return nil
	}

	response.SuccessResponse(w, http.StatusOK, "Publisher Has Been Updated")
	return nil
}

func (h *PublsiherHandler) DeletePublisher(w http.ResponseWriter, r *http.Request, param httprouter.Params) error {
	id, _ := strconv.ParseInt(param.ByName("id"), 10, 64)

	ctx := r.Context()
	err := h.uc.DeletePublisher(ctx, id)
	if err != nil {
		response.FailedResponse(w, http.StatusForbidden, err.Error())
		return err
	}

	response.SuccessResponse(w, http.StatusOK, "Publisher Has Been Deleted")
	return nil
}
