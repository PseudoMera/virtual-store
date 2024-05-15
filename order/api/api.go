package api

import (
	"encoding/json"
	"net/http"

	"github.com/PseudoMera/virtual-store/order/service"
	"github.com/PseudoMera/virtual-store/order/store"
	"github.com/PseudoMera/virtual-store/shared"
)

type OrderAPI struct {
	service *service.OrderService
}

func NewOrderAPI(service *service.OrderService) *OrderAPI {
	return &OrderAPI{
		service: service,
	}
}

type CreateOrderRequest struct {
	UserID     int     `json:"user_id"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
}

type CreateOrderResponse struct {
	ID int `json:"id"`
}

func (o *OrderAPI) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	id, err := o.service.CreateOrder(r.Context(), req.UserID, req.TotalPrice, store.OrderStatus(req.Status))
	if err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	shared.WriteResponse(http.StatusCreated, CreateOrderResponse{
		ID: id,
	}, w)
}

type GetOrderRequest struct {
	ID int `json:"id"`
}

func (o *OrderAPI) GetOrder(w http.ResponseWriter, r *http.Request) {
	var req GetOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	order, err := o.service.GetOrder(r.Context(), req.ID)
	if err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	shared.WriteResponse(http.StatusOK, order, w)
}

type GetOrdersByUserRequest struct {
	UserID int `json:"user_id"`
}

func (o *OrderAPI) GetOrdersByUser(w http.ResponseWriter, r *http.Request) {
	var req GetOrdersByUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	orders, err := o.service.GetOrdersByUser(r.Context(), req.UserID)
	if err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	shared.WriteResponse(http.StatusOK, orders, w)
}

type UpdateOrderRequest struct {
	ID         int     `json:"id"`
	Status     string  `json:"status"`
	TotalPrice float64 `json:"total_price"`
}

func (o *OrderAPI) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var req UpdateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err := o.service.UpdateOrder(r.Context(), req.ID, store.OrderStatus(req.Status), req.TotalPrice); err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type UpdateOrderStatusRequest struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func (o *OrderAPI) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	var req UpdateOrderStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err := o.service.UpdateOrderStatus(r.Context(), req.ID, store.OrderStatus(req.Status)); err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
