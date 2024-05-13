package main

import (
	"context"
	"time"

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

	_ = store.NewStore(database.DB())
}
