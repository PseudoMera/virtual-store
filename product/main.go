package main

import (
	"context"
	"net/http"
	"time"

	"github.com/PseudoMera/virtual-store/product/api"
	"github.com/PseudoMera/virtual-store/product/service"
	"github.com/PseudoMera/virtual-store/product/store"
	"github.com/PseudoMera/virtual-store/shared"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	database, err := shared.NewDatabase(ctx)
	if err != nil {
		panic(err)
	}
	defer cancel()

	router := api.NewRouter()
	store := store.NewStore(database.DB())
	productService := service.NewProductService(store)
	productAPI := api.NewProductAPI(productService)

	router.Post("/api/v1/product", productAPI.CreateProduct)

	if err := http.ListenAndServe(":3000", router); err != nil {
		panic(err)
	}
}
