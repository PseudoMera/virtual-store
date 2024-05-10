package shared

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	db *pgxpool.Pool
}

func NewDatabase(ctx context.Context) (*PostgresDB, error) {
	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return &PostgresDB{
		db: pool,
	}, nil
}

func (pg *PostgresDB) DB() *pgxpool.Pool {
	return pg.db
}
