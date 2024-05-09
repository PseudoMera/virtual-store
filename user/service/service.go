package service

import (
	"context"

	"github.com/PseudoMera/virtual-store/user/store"
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
		return nil
	}
	if password == "" {
		return nil
	}

	user := store.User{
		Email:    email,
		Password: password,
	}
	return u.db.StoreUser(ctx, user)
}
