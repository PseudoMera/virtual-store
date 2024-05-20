package service

import (
	"context"
	"errors"
	"log/slog"

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
	db     *store.Store
	logger *slog.Logger
}

func NewUserService(db *store.Store, logger *slog.Logger) *UserService {
	return &UserService{
		db:     db,
		logger: logger,
	}
}

func (u *UserService) CreateUser(ctx context.Context, email, password string) (int, error) {
	if email == "" {
		u.logger.Info("error at CreateUser", slog.String("error", errEmptyAddress.Error()))
		return 0, errEmptyEmail
	}
	if password == "" {
		u.logger.Info("error at CreateUser", slog.String("error", errEmptyPassword.Error()))
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
		u.logger.Info("error at GetUser", slog.String("error", errEmptyAddress.Error()))
		return nil, errEmptyEmail
	}

	return u.db.RetrieveUser(ctx, email)
}

func (u *UserService) CreateUserProfile(ctx context.Context, userID int, name string, photo string, country string, address string, phone string) (int, error) {
	if userID == 0 {
		u.logger.Info("error at CreateUserProfile", slog.String("error", errEmptyUserID.Error()))
		return 0, errEmptyUserID
	}
	if name == "" {
		u.logger.Info("error at CreateUserProfile", slog.String("error", errEmptyName.Error()))
		return 0, errEmptyName
	}
	if photo == "" {
		u.logger.Info("error at CreateUserProfile", slog.String("error", errEmptyPhoto.Error()))
		return 0, errEmptyPhoto
	}
	if country == "" {
		u.logger.Info("error at CreateUserProfile", slog.String("error", errEmptyCountry.Error()))
		return 0, errEmptyCountry
	}
	if address == "" {
		u.logger.Info("error at CreateUserProfile", slog.String("error", errEmptyAddress.Error()))
		return 0, errEmptyAddress
	}
	if phone == "" {
		u.logger.Info("error at CreateUserProfile", slog.String("error", errEmptyPhone.Error()))
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
		u.logger.Info("error at RetrieveUserProfile", slog.String("error", errEmptyUserID.Error()))
		return nil, errEmptyUserID
	}

	return u.db.RetrieveUserProfile(ctx, userID)
}

func (u *UserService) UpdateUserProfile(ctx context.Context, userID int, name string, photo string, country string, address string, phone string) error {
	if userID == 0 {
		u.logger.Info("error at UpdateUserProfile", slog.String("error", errEmptyUserID.Error()))
		return errEmptyUserID
	}
	if name == "" {
		u.logger.Info("error at UpdateUserProfile", slog.String("error", errEmptyName.Error()))
		return errEmptyName
	}
	if photo == "" {
		u.logger.Info("error at UpdateUserProfile", slog.String("error", errEmptyPhoto.Error()))
		return errEmptyPhoto
	}
	if country == "" {
		u.logger.Info("error at UpdateUserProfile", slog.String("error", errEmptyCountry.Error()))
		return errEmptyCountry
	}
	if address == "" {
		u.logger.Info("error at UpdateUserProfile", slog.String("error", errEmptyAddress.Error()))
		return errEmptyAddress
	}
	if phone == "" {
		u.logger.Info("error at UpdateUserProfile", slog.String("error", errEmptyPhone.Error()))
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
