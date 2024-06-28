package store

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type Store struct {
	db *pgxpool.Pool
}

// NewStore takes a pgx pool as parameter and returns a Store struct which can be used to interact
// with a postgres database
func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		db: db,
	}
}

type User struct {
	ID        int
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Profile struct {
	ID        int
	UserID    int
	Name      string
	Photo     string
	Country   string
	Address   string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// StoreUser creates a new user and stores in the databse.
// Returns the created user ID if everything went well,
// an error if there was an issue.
func (s *Store) StoreUser(ctx context.Context, user User) (int, error) {
	hPassword, err := hashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	var id int
	err = s.db.QueryRow(ctx, "INSERT INTO vstore_user(email, password) VALUES($1, $2) RETURNING id", user.Email, hPassword).Scan(&id)
	return id, err
}

// RetrieveUser retrieves the user with the given email.
// Returns a User or an error depending on the result of the query.
func (s *Store) RetrieveUser(ctx context.Context, email string) (*User, error) {
	user := new(User)
	err := s.db.QueryRow(ctx, "SELECT id, email, created_at FROM vstore_user WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.CreatedAt)
	return user, err
}

// StoreUserProfile creates a new profile for the specified user.
// The profile must contain the user id to match the profile.
func (s *Store) StoreUserProfile(ctx context.Context, profile Profile) (int, error) {
	var id int
	err := s.db.QueryRow(ctx, "INSERT INTO user_profile(user_id, name, photo, country, address, phone) VALUES($1, $2, $3, $4, $5, $6) RETURNING id", profile.UserID, profile.Name, profile.Photo, profile.Country, profile.Address, profile.Phone).Scan(&id)
	return id, err
}

// RetrieveUserProfile retrieves the user profile with the given user ID.
func (s *Store) RetrieveUserProfile(ctx context.Context, userID int) (*Profile, error) {
	profile := new(Profile)
	err := s.db.QueryRow(ctx, "SELECT id, user_id, name, photo, country, address, phone, created_at FROM user_profile WHERE user_id = $1", userID).Scan(&profile.ID, &profile.UserID, &profile.Name, &profile.Photo, &profile.Country, &profile.Address, &profile.Phone, &profile.CreatedAt)
	return profile, err
}

// UpdateUserProfile updates the user profile data.
func (s *Store) UpdateUserProfile(ctx context.Context, profile Profile) error {
	_, err := s.db.Exec(ctx, "UPDATE user_profile SET name = $2, photo = $3, country = $4, address = $5, phone = $6 WHERE user_id = $1", profile.UserID, profile.Name, profile.Photo, profile.Country, profile.Address, profile.Phone)
	return err
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPassword is a utility function that compares a hashedpassword and a plain string password to see if they match.
// Returns true if they both match, false otherwise.
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
