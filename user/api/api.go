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

type CreateUserResponse struct {
	ID int `json:"id"`
}

func (u *UserAPI) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	id, err := u.service.CreateUser(r.Context(), req.Email, req.Password)
	if err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	shared.WriteResponse(http.StatusCreated, CreateUserResponse{
		ID: id,
	}, w)
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

type CreateUserProfileRequest struct {
	UserID  int    `json:"user_id"`
	Name    string `json:"name"`
	Photo   string `json:"photo"`
	Country string `json:"country"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type CreateUserProfileResponse struct {
	ID int `json:"id"`
}

func (u *UserAPI) CreateUserProfile(w http.ResponseWriter, r *http.Request) {
	var req CreateUserProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := u.service.CreateUserProfile(r.Context(), req.UserID, req.Name, req.Photo, req.Country, req.Address, req.Phone)
	if err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	shared.WriteResponse(http.StatusCreated, CreateUserProfileResponse{
		ID: id,
	}, w)
}

type GetUserProfileRequest struct {
	UserID int `json:"user_id"`
}

func (u *UserAPI) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	var req GetUserProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	profile, err := u.service.RetrieveUserProfile(r.Context(), req.UserID)
	if err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	shared.WriteResponse(http.StatusOK, profile, w)
}

type UpdateUserProfileRequest struct {
	UserID  int    `json:"user_id"`
	Name    string `json:"name"`
	Photo   string `json:"photo"`
	Country string `json:"country"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func (u *UserAPI) UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	var req UpdateUserProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := u.service.UpdateUserProfile(r.Context(), req.UserID, req.Name, req.Photo, req.Country, req.Address, req.Phone); err != nil {
		shared.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
