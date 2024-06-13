package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/PseudoMera/virtual-store/order/api"
	"github.com/PseudoMera/virtual-store/order/grpc"
	"github.com/PseudoMera/virtual-store/order/service"
	"github.com/PseudoMera/virtual-store/order/store"
	"github.com/PseudoMera/virtual-store/shared"
	egrpc "google.golang.org/grpc"
)

const (
	apiPath = "/api/v1"
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

	router.Post(fmt.Sprintf("%s/order", apiPath), orderAPI.CreateOrder)
	router.Get(fmt.Sprintf("%s/order", apiPath), orderAPI.GetOrder)
	router.Get(fmt.Sprintf("%s/user-order", apiPath), orderAPI.GetOrdersByUser)
	router.Put(fmt.Sprintf("%s/order", apiPath), orderAPI.UpdateOrder)
	router.Put(fmt.Sprintf("%s/order/status", apiPath), orderAPI.UpdateOrderStatus)

	lis, err := net.Listen("tcp", config.grpcServerPort)
	if err != nil {
		panic(err)
	}

	var opts []egrpc.ServerOption
	grpcServer := egrpc.NewServer(opts...)
	serviceServer := grpc.NewOrderServer(store)
	grpc.RegisterOrderServiceServer(grpcServer, serviceServer)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.httpServerPort), router); err != nil {
		panic(err)
	}
}
