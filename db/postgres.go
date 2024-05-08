package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	DB *pgxpool.Pool
}

func NewDatabase(ctx context.Context) *PostgresDB {
	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &PostgresDB{
		DB: pool,
	}
}
