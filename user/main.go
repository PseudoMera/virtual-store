package main

import (
	"context"
	"net/http"
	"time"

	"github.com/PseudoMera/virtual-store/shared"
	"github.com/PseudoMera/virtual-store/user/api"
	"github.com/PseudoMera/virtual-store/user/service"
	"github.com/PseudoMera/virtual-store/user/store"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	database, err := shared.NewDatabase(ctx)
	if err != nil {
		panic(err)
	}
	defer cancel()

	logger := shared.NewLogger()
	router := api.NewRouter()
	store := store.NewStore(database.DB())
	userService := service.NewUserService(store, logger)
	userAPI := api.NewUserAPI(userService)

	router.Post("/api/v1/user", userAPI.CreateUser)

	if err := http.ListenAndServe(":3000", router); err != nil {
		panic(err)
	}
}
