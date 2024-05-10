package service

import (
	"context"
	"errors"

	"github.com/PseudoMera/virtual-store/user/store"
)

var (
	errEmptyEmail    = errors.New("email field cannot be empty")
	errEmptyPassword = errors.New("password field cannot be empty")
	errEmptyUserID   = errors.New("user_id field cannot be empty")
	errEmptyName     = errors.New("name field cannot be empty")
	errEmptyPhoto    = errors.New("photo field cannot be empty")
	errEmptyCountry  = errors.New("country field cannot be empty")
	errEmptyAddress  = errors.New("address field cannot be empty")
	errEmptyPhone    = errors.New("phone field cannot be empty")
)

type UserService struct {
	db *store.Store
}

func NewUserService(db *store.Store) *UserService {
	return &UserService{
		db: db,
	}
}

func (u *UserService) CreateUser(ctx context.Context, email, password string) (int, error) {
	if email == "" {
		return 0, errEmptyEmail
	}
	if password == "" {
		return 0, errEmptyPassword
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

func (u *UserService) CreateUserProfile(ctx context.Context, userID int, name string, photo string, country string, address string, phone string) (int, error) {
	if userID == 0 {
		return 0, errEmptyUserID
	}
	if name == "" {
		return 0, errEmptyName
	}
	if photo == "" {
		return 0, errEmptyPhoto
	}
	if country == "" {
		return 0, errEmptyCountry
	}
	if address == "" {
		return 0, errEmptyAddress
	}
	if phone == "" {
		return 0, errEmptyPhone
	}

	return u.db.StoreUserProfile(ctx, store.Profile{
		UserID:  userID,
		Name:    name,
		Photo:   photo,
		Country: country,
		Address: address,
		Phone:   phone,
	})
}

func (u *UserService) RetrieveUserProfile(ctx context.Context, userID int) (*store.Profile, error) {
	if userID == 0 {
		return nil, errEmptyUserID
	}

	return u.db.RetrieveUserProfile(ctx, userID)
}

func (u *UserService) UpdateUserProfile(ctx context.Context, userID int, name string, photo string, country string, address string, phone string) error {
	if userID == 0 {
		return errEmptyUserID
	}
	if name == "" {
		return errEmptyName
	}
	if photo == "" {
		return errEmptyPhoto
	}
	if country == "" {
		return errEmptyCountry
	}
	if address == "" {
		return errEmptyAddress
	}
	if phone == "" {
		return errEmptyPhone
	}

	return u.db.UpdateUserProfile(ctx, store.Profile{
		UserID:  userID,
		Name:    name,
		Photo:   photo,
		Country: country,
		Address: address,
		Phone:   phone,
	})
}
