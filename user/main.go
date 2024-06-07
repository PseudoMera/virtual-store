package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/PseudoMera/virtual-store/shared"
	"github.com/PseudoMera/virtual-store/user/api"
	"github.com/PseudoMera/virtual-store/user/service"
	"github.com/PseudoMera/virtual-store/user/store"
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
	userService := service.NewUserService(store, logger)
	userAPI := api.NewUserAPI(userService)

	router.Post(fmt.Sprintf("%s/user", apiPath), userAPI.CreateUser)
	router.Get(fmt.Sprintf("%s/user", apiPath), userAPI.GetUser)
	router.Post(fmt.Sprintf("%s/user/profile", apiPath), userAPI.CreateUserProfile)
	router.Get(fmt.Sprintf("%s/user/profile", apiPath), userAPI.GetUserProfile)
	router.Put(fmt.Sprintf("%s/user/profile", apiPath), userAPI.UpdateUserProfile)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.httpServerPort), router); err != nil {
		panic(err)
	}
}
