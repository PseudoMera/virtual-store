package main

import (
	"errors"
	"os"
)

const (
	connectionString = "CONNECTION_STRING"
	httpServerPort   = "HTTP_SERVER_PORT"
	grpcServerPort   = "GRPC_SERVER_PORT"
)

var (
	errEmptyConnectionString = errors.New("env variable 'CONNECTION_STRING' cannot be empty")
	errEmptyHTTPServerPort   = errors.New("env variable 'HTTP_SERVER_PORT' cannot be empty")
	errEmptyGRPCServerPort   = errors.New("env variable 'GRPC_SERVER_PORT' cannot be empty")
)

type config struct {
	connectionString string
	httpServerPort   string
	grpcServerPort   string
}

func getConfig() config {
	cstr := os.Getenv(connectionString)
	if cstr == "" {
		panic(errEmptyConnectionString)
	}

	httpPort := os.Getenv(httpServerPort)
	if httpPort == "" {
		panic(errEmptyHTTPServerPort)
	}

	grpcPort := os.Getenv(grpcServerPort)
	if grpcPort == "" {
		panic(errEmptyGRPCServerPort)
	}

	return config{
		connectionString: cstr,
		httpServerPort:   httpPort,
		grpcServerPort:   grpcPort,
	}
}
