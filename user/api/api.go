package api

import (
	"encoding/json"
	"net/http"

	"github.com/PseudoMera/virtual-store/shared"
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
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err := u.service.CreateUser(r.Context(), req.Email, req.Password); err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type GetUserRequest struct {
	Email string `json:"email"`
}

func (u *UserAPI) GetUser(w http.ResponseWriter, r *http.Request) {
	var req GetUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := u.service.GetUser(r.Context(), req.Email)
	if err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	shared.WriteResponse(http.StatusOK, user, w)
}
