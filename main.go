package main

import (
	"context"
	"time"

	"github.com/PseudoMera/virtual-store/db"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := db.NewDatabase(ctx)
	if err != nil {
		panic(err)
	}
	defer cancel()
}
