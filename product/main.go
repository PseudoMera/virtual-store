package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/PseudoMera/virtual-store/product/api"
	"github.com/PseudoMera/virtual-store/product/grpc"
	"github.com/PseudoMera/virtual-store/product/service"
	"github.com/PseudoMera/virtual-store/product/store"
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
	productService := service.NewProductService(store, logger)
	productAPI := api.NewProductAPI(productService)

	router.Post(fmt.Sprintf("%s/product", apiPath), productAPI.CreateProduct)
	router.Get(fmt.Sprintf("%s/product", apiPath), productAPI.GetProduct)
	router.Get(fmt.Sprintf("%s/products", apiPath), productAPI.GetProducts)
	router.Put(fmt.Sprintf("%s/product", apiPath), productAPI.UpdateProduct)
	router.Put(fmt.Sprintf("%s/product/stock", apiPath), productAPI.UpdateProductStock)

	lis, err := net.Listen("tcp", config.grpcServerPort)
	if err != nil {
		panic(err)
	}

	var opts []egrpc.ServerOption
	grpcServer := egrpc.NewServer(opts...)
	serviceServer := grpc.NewProductServer(store)
	grpc.RegisterProductServiceServer(grpcServer, serviceServer)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.httpServerPort), router); err != nil {
		panic(err)
	}
}
