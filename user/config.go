package main

import (
	"errors"
	"os"
)

const (
	connectionString = "CONNECTION_STRING"
	httpServerPort   = "HTTP_SERVER_PORT"
)

var (
	errEmptyConnectionString = errors.New("env variable 'CONNECTION_STRING' cannot be empty")
	errEmptyHTTPServerPort   = errors.New("env variable 'HTTP_SERVER_PORT' cannot be empty")
)

type config struct {
	connectionString string
	httpServerPort   string
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

	return config{
		connectionString: cstr,
		httpServerPort:   httpPort,
	}
}
