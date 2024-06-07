package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PseudoMera/virtual-store/product/service"
	"github.com/PseudoMera/virtual-store/product/store"
	"github.com/PseudoMera/virtual-store/shared"
	"github.com/go-chi/chi/v5"
)

const (
	projectRootPath = "../../"

	testName  = "testName"
	testPrice = 12.05
	teststock = 120
)

func TestCreateProduct(t *testing.T) {
	ctx := context.Background()
	db, _, err := shared.SetupPostgresClient(ctx, projectRootPath)
	if err != nil {
		t.Fatal(err)
	}

	s := store.NewStore(db.DB())
	serv := service.NewProductService(s, slog.Default())
	api := NewProductAPI(serv)
	router := chi.NewRouter()
	router.Post("/api/v1/product", api.CreateProduct)

	ts := httptest.NewServer(router)
	defer ts.Close()

	createProduct := CreateProductRequest{
		Name:  testName,
		Price: testPrice,
		Stock: teststock,
	}
	createProductBytes, err := json.Marshal(createProduct)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", ts.URL+"/api/v1/product", bytes.NewBuffer(createProductBytes))
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

func TestGetProduct(t *testing.T) {
	ctx := context.Background()
	db, _, err := shared.SetupPostgresClient(ctx, projectRootPath)
	if err != nil {
		t.Fatal(err)
	}

	s := store.NewStore(db.DB())
	productID, err := s.StoreProduct(ctx, store.Product{
		Name:  testName,
		Price: testPrice,
		Stock: teststock,
	})
	if err != nil {
		t.Fatal(err)
	}

	serv := service.NewProductService(s, slog.Default())
	api := NewProductAPI(serv)
	router := chi.NewRouter()
	router.Get("/api/v1/product", api.GetProduct)

	ts := httptest.NewServer(router)
	defer ts.Close()

	getProduct := GetProductRequest{
		ID: productID,
	}
	getProductBytes, err := json.Marshal(getProduct)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", ts.URL+"/api/v1/product", bytes.NewBuffer(getProductBytes))
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

func TestGetProducts(t *testing.T) {
	ctx := context.Background()
	db, _, err := shared.SetupPostgresClient(ctx, projectRootPath)
	if err != nil {
		t.Fatal(err)
	}

	s := store.NewStore(db.DB())
	_, err = s.StoreProduct(ctx, store.Product{
		Name:  testName,
		Price: testPrice,
		Stock: teststock,
	})
	if err != nil {
		t.Fatal(err)
	}

	serv := service.NewProductService(s, slog.Default())
	api := NewProductAPI(serv)
	router := chi.NewRouter()
	router.Get("/api/v1/product", api.GetProducts)

	ts := httptest.NewServer(router)
	defer ts.Close()

	getProducts := GetProductsRequest{
		Name: testName,
	}
	getProductsBytes, err := json.Marshal(getProducts)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", ts.URL+"/api/v1/product", bytes.NewBuffer(getProductsBytes))
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

func TestUpdateProduct(t *testing.T) {
	ctx := context.Background()
	db, _, err := shared.SetupPostgresClient(ctx, projectRootPath)
	if err != nil {
		t.Fatal(err)
	}

	s := store.NewStore(db.DB())
	productID, err := s.StoreProduct(ctx, store.Product{
		Name:  testName,
		Price: testPrice,
		Stock: teststock,
	})
	if err != nil {
		t.Fatal(err)
	}

	serv := service.NewProductService(s, slog.Default())
	api := NewProductAPI(serv)
	router := chi.NewRouter()
	router.Put("/api/v1/product", api.UpdateProduct)

	ts := httptest.NewServer(router)
	defer ts.Close()

	updateProduct := UpdateProductRequest{
		ID:    productID,
		Name:  testName + "UPDATED",
		Price: testPrice,
		Stock: teststock,
	}
	updateProductBytes, err := json.Marshal(updateProduct)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", ts.URL+"/api/v1/product", bytes.NewBuffer(updateProductBytes))
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

func TestUpdateProductStock(t *testing.T) {
	ctx := context.Background()
	db, _, err := shared.SetupPostgresClient(ctx, projectRootPath)
	if err != nil {
		t.Fatal(err)
	}

	s := store.NewStore(db.DB())
	productID, err := s.StoreProduct(ctx, store.Product{
		Name:  testName,
		Price: testPrice,
		Stock: teststock,
	})
	if err != nil {
		t.Fatal(err)
	}

	serv := service.NewProductService(s, slog.Default())
	api := NewProductAPI(serv)
	router := chi.NewRouter()
	router.Put("/api/v1/product", api.UpdateProductStock)

	ts := httptest.NewServer(router)
	defer ts.Close()

	updateProductStock := UpdateProductStockRequest{
		ID:    productID,
		Stock: teststock * 2,
	}
	updateProductStockBytes, err := json.Marshal(updateProductStock)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", ts.URL+"/api/v1/product", bytes.NewBuffer(updateProductStockBytes))
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
