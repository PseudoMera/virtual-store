package grpc

import (
	"context"
	"errors"

	"github.com/PseudoMera/virtual-store/user/store"
)

var (
	errEmptyEmail    = errors.New("email field cannot be empty")
	errEmptyPassword = errors.New("password field cannot be empty")

	errRetrievingUser = errors.New("something went wrong while retrieving user")
)

type UserServer struct {
	db *store.Store
	UnimplementedUserServiceServer
}

func NewUserServer(db *store.Store) *UserServer {
	return &UserServer{
		db: db,
	}
}

func (us *UserServer) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
	email := *req.Email
	if email == "" {
		return nil, errEmptyEmail
	}

	user, err := us.db.RetrieveUser(ctx, email)
	if err != nil {
		return nil, errRetrievingUser
	}

	cUser := &User{
		Email:    &user.Email,
		Id:       &user.ID,
		Password: &user.Password,
	}
	return cUser, nil
}

func (us *UserServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	email, password := *req.Email, *req.Password

	if email == "" {
		return nil, errEmptyEmail
	}
	if password == "" {
		return nil, errEmptyPassword
	}

	id, err := us.db.StoreUser(ctx, store.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	cID := int64(id)

	return &User{
		Id:       &cID,
		Email:    &email,
		Password: &password,
	}, err
}
