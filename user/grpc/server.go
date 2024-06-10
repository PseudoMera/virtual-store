package grpc

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
	email := req.Email
	if email == "" {
		return nil, errEmptyEmail
	}

	user, err := us.db.RetrieveUser(ctx, email)
	if err != nil {
		return nil, errRetrievingUser
	}

	cID := int64(user.ID)
	cUser := &User{
		Email:    user.Email,
		Id:       cID,
		Password: user.Password,
	}
	return cUser, nil
}

func (us *UserServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	email, password := req.Email, req.Password
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
		Id:       cID,
		Email:    email,
		Password: password,
	}, err
}

func (us *UserServer) CreateUserProfile(ctx context.Context, req *CreateUserProfileRequest) (*CreateUserProfileResponse, error) {
	userID, name, photo, country, address, phone := req.Id, req.Name, req.Photo, req.Country, req.Address, req.Phone
	if userID == 0 {
		return nil, errEmptyUserID
	}
	if name == "" {
		return nil, errEmptyName
	}
	if photo == "" {
		return nil, errEmptyPhoto
	}
	if country == "" {
		return nil, errEmptyCountry
	}
	if address == "" {
		return nil, errEmptyAddress
	}
	if phone == "" {
		return nil, errEmptyPhone
	}

	profileID, err := us.db.StoreUserProfile(ctx, store.Profile{
		UserID:  int(userID),
		Name:    name,
		Photo:   photo,
		Country: country,
		Address: address,
		Phone:   phone,
	})
	return &CreateUserProfileResponse{Id: int64(profileID)}, err
}

func (us *UserServer) GetUserProfile(ctx context.Context, req *GetUserProfileRequest) (*Profile, error) {
	if req.Id == 0 {
		return nil, errEmptyUserID
	}

	profile, err := us.db.RetrieveUserProfile(ctx, int(req.Id))

	return &Profile{
		Id:      int64(profile.ID),
		UserID:  int64(profile.UserID),
		Name:    profile.Name,
		Photo:   profile.Photo,
		Country: profile.Country,
		Address: profile.Address,
		Phone:   profile.Phone,
	}, err
}

func (us *UserServer) UpdateUserProfile(ctx context.Context, req *UpdateUserProfileRequest) (*SuccessResponse, error) {
	userID, name, photo, country, address, phone := req.UserID, req.Name, req.Photo, req.Country, req.Address, req.Phone
	if userID == 0 {
		return nil, errEmptyUserID
	}
	if name == "" {
		return nil, errEmptyName
	}
	if photo == "" {
		return nil, errEmptyPhoto
	}
	if country == "" {
		return nil, errEmptyCountry
	}
	if address == "" {
		return nil, errEmptyAddress
	}
	if phone == "" {
		return nil, errEmptyPhone
	}

	if err := us.db.UpdateUserProfile(ctx, store.Profile{
		UserID:  int(userID),
		Name:    name,
		Photo:   photo,
		Country: country,
		Address: address,
		Phone:   phone,
	}); err != nil {
		return nil, err
	}

	return &SuccessResponse{
		Msg: "Success!",
	}, nil
}
