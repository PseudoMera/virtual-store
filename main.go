package main

import (
	"context"
	"time"

	"github.com/PseudoMera/virtual-store/db"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_ = db.NewDatabase(ctx)
	defer cancel()
}
