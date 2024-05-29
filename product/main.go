package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/PseudoMera/virtual-store/product/api"
	"github.com/PseudoMera/virtual-store/product/service"
	"github.com/PseudoMera/virtual-store/product/store"
	"github.com/PseudoMera/virtual-store/shared"
)

func main() {
	config := getConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	database, err := shared.NewPostgresDatabase(ctx, config.connectionString)
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

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.httpServerPort), router); err != nil {
		panic(err)
	}
}
