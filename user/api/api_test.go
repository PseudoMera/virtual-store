package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PseudoMera/virtual-store/shared"
	"github.com/PseudoMera/virtual-store/user/service"
	"github.com/PseudoMera/virtual-store/user/store"
	"github.com/go-chi/chi/v5"
)

const (
	projectRootPath = "../../"

	testEmail    = "test@test.test"
	testPassword = "testPassword!!!"
	testName     = "tester"
	testPhoto    = "testPhoto"
	testCountry  = "Testing Lands"
	testAddress  = "testAddress"
	testPhone    = "test-test-test"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	db, _, err := shared.SetupPostgresClient(ctx, projectRootPath)
	if err != nil {
		t.Fatal(err)
	}

	s := store.NewStore(db.DB())
	serv := service.NewUserService(s, slog.Default())
	api := NewUserAPI(serv)
	router := chi.NewRouter()
	router.Post("/api/v1/user", api.CreateUser)

	ts := httptest.NewServer(router)
	defer ts.Close()

	createUser := CreateUserRequest{
		Email:    testEmail,
		Password: testPassword,
	}
	createUserBytes, err := json.Marshal(createUser)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", ts.URL+"/api/v1/user", bytes.NewBuffer(createUserBytes))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusCreated {
		t.Fatalf("wanted %d, got %d", http.StatusCreated, status)
	}
}

func TestGetUser(t *testing.T) {
	ctx := context.Background()
	db, _, err := shared.SetupPostgresClient(ctx, projectRootPath)
	if err != nil {
		t.Fatal(err)
	}

	s := store.NewStore(db.DB())

	_, err = s.StoreUser(ctx, store.User{
		Email:    testEmail,
		Password: testPassword,
	})
	if err != nil {
		t.Fatal(err)
	}

	serv := service.NewUserService(s, slog.Default())
	api := NewUserAPI(serv)
	router := chi.NewRouter()
	router.Get("/api/v1/user", api.GetUser)

	ts := httptest.NewServer(router)
	defer ts.Close()

	getUserReq := GetUserRequest{
		Email: testEmail,
	}
	getUserReqBytes, err := json.Marshal(getUserReq)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", ts.URL+"/api/v1/user", bytes.NewBuffer(getUserReqBytes))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		t.Fatalf("wanted %d, got %d", http.StatusOK, status)
	}
}

func TestCreateUserProfile(t *testing.T) {
	ctx := context.Background()
	db, _, err := shared.SetupPostgresClient(ctx, projectRootPath)
	if err != nil {
		t.Fatal(err)
	}

	s := store.NewStore(db.DB())
	userID, err := s.StoreUser(ctx, store.User{
		Email:    testEmail,
		Password: testPassword,
	})
	if err != nil {
		t.Fatal(err)
	}

	serv := service.NewUserService(s, slog.Default())
	api := NewUserAPI(serv)
	router := chi.NewRouter()
	router.Post("/api/v1/user/profile", api.CreateUserProfile)

	ts := httptest.NewServer(router)
	defer ts.Close()

	createUserProfile := CreateUserProfileRequest{
		UserID:  userID,
		Name:    testName,
		Photo:   testPhoto,
		Country: testCountry,
		Address: testAddress,
		Phone:   testPhone,
	}
	createUserProfileBytes, err := json.Marshal(createUserProfile)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", ts.URL+"/api/v1/user/profile", bytes.NewBuffer(createUserProfileBytes))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusCreated {
		t.Fatalf("wanted %d, got %d", http.StatusCreated, status)
	}
}

func TestGetUserProfiel(t *testing.T) {
	ctx := context.Background()
	db, _, err := shared.SetupPostgresClient(ctx, projectRootPath)
	if err != nil {
		t.Fatal(err)
	}

	s := store.NewStore(db.DB())
	userID, err := s.StoreUser(ctx, store.User{
		Email:    testEmail,
		Password: testPassword,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = s.StoreUserProfile(ctx, store.Profile{
		UserID:  userID,
		Name:    testName,
		Photo:   testPhoto,
		Country: testCountry,
		Address: testAddress,
		Phone:   testPhone,
	})
	if err != nil {
		t.Fatal(err)
	}

	serv := service.NewUserService(s, slog.Default())
	api := NewUserAPI(serv)
	router := chi.NewRouter()
	router.Get("/api/v1/user/profile", api.GetUserProfile)

	ts := httptest.NewServer(router)
	defer ts.Close()

	getUserProfileReq := GetUserProfileRequest{
		UserID: userID,
	}
	getUserProfileReqBytes, err := json.Marshal(getUserProfileReq)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", ts.URL+"/api/v1/user/profile", bytes.NewBuffer(getUserProfileReqBytes))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		t.Fatalf("wanted %d, got %d", http.StatusOK, status)
	}
}

func TestUpdateUserProfile(t *testing.T) {
	ctx := context.Background()
	db, _, err := shared.SetupPostgresClient(ctx, projectRootPath)
	if err != nil {
		t.Fatal(err)
	}

	s := store.NewStore(db.DB())
	userID, err := s.StoreUser(ctx, store.User{
		Email:    testEmail,
		Password: testPassword,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = s.StoreUserProfile(ctx, store.Profile{
		UserID:  userID,
		Name:    testName,
		Photo:   testPhoto,
		Country: testCountry,
		Address: testAddress,
		Phone:   testPhone,
	})
	if err != nil {
		t.Fatal(err)
	}

	serv := service.NewUserService(s, slog.Default())
	api := NewUserAPI(serv)
	router := chi.NewRouter()
	router.Put("/api/v1/user/profile", api.UpdateUserProfile)

	ts := httptest.NewServer(router)
	defer ts.Close()

	updateUserProfileReq := UpdateUserProfileRequest{
		UserID:  userID,
		Name:    testName + "UPDATED",
		Photo:   testPhoto,
		Country: testCountry,
		Address: testAddress,
		Phone:   testPhone,
	}
	updateUserProfileReqBytes, err := json.Marshal(updateUserProfileReq)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", ts.URL+"/api/v1/user/profile", bytes.NewBuffer(updateUserProfileReqBytes))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusNoContent {
		t.Fatalf("wanted %d, got %d", http.StatusNoContent, status)
	}
}
