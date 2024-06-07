package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PseudoMera/virtual-store/order/service"
	"github.com/PseudoMera/virtual-store/order/store"
	"github.com/PseudoMera/virtual-store/shared"
	userStore "github.com/PseudoMera/virtual-store/user/store"
	"github.com/go-chi/chi/v5"
)

const (
	projectRootPath = "../../"

	testTotalPrice = 120.40
	testStatus     = "pending"
	testEmail      = "test@test.test"
	testPassword   = "testPassword!!!"
)

func TestCreateOrder(t *testing.T) {
	ctx := context.Background()
	db, _, err := shared.SetupPostgresClient(ctx, projectRootPath)
	if err != nil {
		t.Fatal(err)
	}

	s := store.NewStore(db.DB())
	serv := service.NewOrderService(s, slog.Default())
	api := NewOrderAPI(serv)
	router := chi.NewRouter()
	router.Post("/api/v1/order", api.CreateOrder)

	ts := httptest.NewServer(router)
	defer ts.Close()

	uStore := userStore.NewStore(db.DB())
	userID, err := uStore.StoreUser(ctx, userStore.User{
		Email:    testEmail,
		Password: testPassword,
	})
	if err != nil {
		t.Fatal(err)
	}

	createOrder := CreateOrderRequest{
		UserID:     userID,
		TotalPrice: testTotalPrice,
		Status:     testStatus,
	}
	createOrderBytes, err := json.Marshal(createOrder)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", ts.URL+"/api/v1/order", bytes.NewBuffer(createOrderBytes))
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

func TestGetOrder(t *testing.T) {
	ctx := context.Background()
	db, _, err := shared.SetupPostgresClient(ctx, projectRootPath)
	if err != nil {
		t.Fatal(err)
	}

	s := store.NewStore(db.DB())
	serv := service.NewOrderService(s, slog.Default())
	api := NewOrderAPI(serv)
	router := chi.NewRouter()
	router.Get("/api/v1/order", api.GetOrder)

	ts := httptest.NewServer(router)
	defer ts.Close()

	uStore := userStore.NewStore(db.DB())
	userID, err := uStore.StoreUser(ctx, userStore.User{
		Email:    testEmail,
		Password: testPassword,
	})
	if err != nil {
		t.Fatal(err)
	}

	orderID, err := s.StoreOrder(ctx, store.Order{
		UserID:     userID,
		TotalPrice: testTotalPrice,
		Status:     testStatus,
	})
	if err != nil {
		t.Fatal(err)
	}

	getOrder := GetOrderRequest{
		ID: orderID,
	}
	getOrderBytes, err := json.Marshal(getOrder)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", ts.URL+"/api/v1/order", bytes.NewBuffer(getOrderBytes))
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

func TestGetOrdersByUser(t *testing.T) {
	ctx := context.Background()
	db, _, err := shared.SetupPostgresClient(ctx, projectRootPath)
	if err != nil {
		t.Fatal(err)
	}

	s := store.NewStore(db.DB())
	serv := service.NewOrderService(s, slog.Default())
	api := NewOrderAPI(serv)
	router := chi.NewRouter()
	router.Get("/api/v1/order", api.GetOrdersByUser)

	ts := httptest.NewServer(router)
	defer ts.Close()

	uStore := userStore.NewStore(db.DB())
	userID, err := uStore.StoreUser(ctx, userStore.User{
		Email:    testEmail,
		Password: testPassword,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = s.StoreOrder(ctx, store.Order{
		UserID:     userID,
		TotalPrice: testTotalPrice,
		Status:     testStatus,
	})
	if err != nil {
		t.Fatal(err)
	}

	getOrder := GetOrdersByUserRequest{
		UserID: userID,
	}
	getOrderBytes, err := json.Marshal(getOrder)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", ts.URL+"/api/v1/order", bytes.NewBuffer(getOrderBytes))
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

func TestUpdateOrder(t *testing.T) {
	ctx := context.Background()
	db, _, err := shared.SetupPostgresClient(ctx, projectRootPath)
	if err != nil {
		t.Fatal(err)
	}

	s := store.NewStore(db.DB())
	serv := service.NewOrderService(s, slog.Default())
	api := NewOrderAPI(serv)
	router := chi.NewRouter()
	router.Put("/api/v1/order", api.UpdateOrder)

	ts := httptest.NewServer(router)
	defer ts.Close()

	uStore := userStore.NewStore(db.DB())
	userID, err := uStore.StoreUser(ctx, userStore.User{
		Email:    testEmail,
		Password: testPassword,
	})
	if err != nil {
		t.Fatal(err)
	}

	orderID, err := s.StoreOrder(ctx, store.Order{
		UserID:     userID,
		TotalPrice: testTotalPrice,
		Status:     testStatus,
	})
	if err != nil {
		t.Fatal(err)
	}

	updateOrder := UpdateOrderRequest{
		ID:         orderID,
		Status:     "completed",
		TotalPrice: 125555,
	}
	updateOrderBytes, err := json.Marshal(updateOrder)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", ts.URL+"/api/v1/order", bytes.NewBuffer(updateOrderBytes))
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

func TestUpdateOrderStatus(t *testing.T) {
	ctx := context.Background()
	db, _, err := shared.SetupPostgresClient(ctx, projectRootPath)
	if err != nil {
		t.Fatal(err)
	}

	s := store.NewStore(db.DB())
	serv := service.NewOrderService(s, slog.Default())
	api := NewOrderAPI(serv)
	router := chi.NewRouter()
	router.Put("/api/v1/order", api.UpdateOrderStatus)

	ts := httptest.NewServer(router)
	defer ts.Close()

	uStore := userStore.NewStore(db.DB())
	userID, err := uStore.StoreUser(ctx, userStore.User{
		Email:    testEmail,
		Password: testPassword,
	})
	if err != nil {
		t.Fatal(err)
	}

	orderID, err := s.StoreOrder(ctx, store.Order{
		UserID:     userID,
		TotalPrice: testTotalPrice,
		Status:     testStatus,
	})
	if err != nil {
		t.Fatal(err)
	}

	updateOrder := UpdateOrderStatusRequest{
		ID:     orderID,
		Status: "completed",
	}
	updateOrderBytes, err := json.Marshal(updateOrder)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", ts.URL+"/api/v1/order", bytes.NewBuffer(updateOrderBytes))
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
