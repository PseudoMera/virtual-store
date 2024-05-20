package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/PseudoMera/virtual-store/product/api"
	"github.com/PseudoMera/virtual-store/product/service"
	"github.com/PseudoMera/virtual-store/product/store"
	"github.com/PseudoMera/virtual-store/shared"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	database, err := shared.NewPostgresDatabase(ctx, os.Getenv("CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}
	defer cancel()

	logger := shared.NewLogger()
	router := api.NewRouter()
	store := store.NewStore(database.DB())
	productService := service.NewProductService(store, logger)
	productAPI := api.NewProductAPI(productService)

	router.Post("/api/v1/product", productAPI.CreateProduct)

	if err := http.ListenAndServe(":3000", router); err != nil {
		panic(err)
	}
}
