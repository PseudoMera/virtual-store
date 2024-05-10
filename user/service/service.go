package service

import (
	"context"
	"errors"

	"github.com/PseudoMera/virtual-store/user/store"
)

var (
	errEmptyEmail    = errors.New("email field cannot be empty")
	errEmptyPassword = errors.New("password field cannot be empty")
)

type UserService struct {
	db *store.Store
}

func NewUserService(db *store.Store) *UserService {
	return &UserService{
		db: db,
	}
}

func (u *UserService) CreateUser(ctx context.Context, email, password string) error {
	if email == "" {
		return errEmptyEmail
	}
	if password == "" {
		return errEmptyPassword
	}

	user := store.User{
		Email:    email,
		Password: password,
	}
	return u.db.StoreUser(ctx, user)
}

func (u *UserService) GetUser(ctx context.Context, email string) (*store.User, error) {
	if email == "" {
		return nil, errEmptyEmail
	}

	return u.db.RetrieveUser(ctx, email)
}
