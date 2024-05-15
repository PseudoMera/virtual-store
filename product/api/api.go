package api

import (
	"encoding/json"
	"net/http"

	"github.com/PseudoMera/virtual-store/product/store"
	"github.com/PseudoMera/virtual-store/shared"
)

type ProductAPI struct {
	db *store.Store
}

func NewProductAPI(db *store.Store) *ProductAPI {
	return &ProductAPI{
		db: db,
	}
}

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

func (p *ProductAPI) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}
}
