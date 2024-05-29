package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/PseudoMera/virtual-store/order/api"
	"github.com/PseudoMera/virtual-store/order/service"
	"github.com/PseudoMera/virtual-store/order/store"
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
	orderService := service.NewOrderService(store, logger)
	orderAPI := api.NewOrderAPI(orderService)

	router.Post("/api/v1/order", orderAPI.CreateOrder)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.httpServerPort), router); err != nil {
		panic(err)
	}
}
