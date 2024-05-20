package shared

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	db *pgxpool.Pool
}

func NewPostgresDatabase(ctx context.Context, connectionStr string) (*PostgresDB, error) {
	pool, err := pgxpool.New(ctx, connectionStr)
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
