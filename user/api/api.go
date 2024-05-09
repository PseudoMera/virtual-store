package api

import (
	"encoding/json"
	"net/http"

	"github.com/PseudoMera/virtual-store/user/service"
	"github.com/PseudoMera/virtual-store/user/store"
)

type UserAPI struct {
	db      *store.Store
	service *service.UserService
}

func NewUserAPI(db *store.Store, service *service.UserService) *UserAPI {
	return &UserAPI{
		db:      db,
		service: service,
	}
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserAPI) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := u.service.CreateUser(r.Context(), req.Email, req.Password); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
