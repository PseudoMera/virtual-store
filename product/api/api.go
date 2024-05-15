package api

import (
	"encoding/json"
	"net/http"

	"github.com/PseudoMera/virtual-store/product/service"
	"github.com/PseudoMera/virtual-store/product/store"
	"github.com/PseudoMera/virtual-store/shared"
)

type ProductAPI struct {
	db      *store.Store
	service *service.ProductService
}

func NewProductAPI(db *store.Store, service *service.ProductService) *ProductAPI {
	return &ProductAPI{
		db:      db,
		service: service,
	}
}

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type CreateProductResponse struct {
	ID int `json:"id"`
}

func (p *ProductAPI) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	id, err := p.service.CreateProduct(r.Context(), req.Name, req.Price, req.Stock)
	if err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	shared.WriteResponse(http.StatusCreated, CreateProductResponse{
		ID: id,
	}, w)
}

type GetProductRequest struct {
	ID int `json:"id"`
}

func (p *ProductAPI) GetProduct(w http.ResponseWriter, r *http.Request) {
	var req GetProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	product, err := p.service.GetProduct(r.Context(), req.ID)
	if err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	shared.WriteResponse(http.StatusOK, product, w)
}

type GetProducsRequest struct {
	Name string `json:"name"`
}

func (p *ProductAPI) GetProducts(w http.ResponseWriter, r *http.Request) {
	var req GetProducsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	products, err := p.service.GetProducts(r.Context(), req.Name)
	if err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	shared.WriteResponse(http.StatusOK, products, w)
}

type UpdateProductRequest struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

func (p *ProductAPI) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var req UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err := p.service.UpdateProduct(r.Context(), req.ID, req.Name, req.Price, req.Stock)
	if err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type UpdateProductStockRequest struct {
	ID    int `json:"id"`
	Stock int `json:"stock"`
}

func (p *ProductAPI) UpdateProductStock(w http.ResponseWriter, r *http.Request) {
	var req UpdateProductStockRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err := p.service.UpdateProductStock(r.Context(), req.ID, req.Stock)
	if err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
