package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

const DSN = "postgres://user:password@localhost:5432/testdb?pool_max_conns=50"

func GetDBPool() *pgxpool.Pool {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, DSN)
	if err != nil {
		panic(err)
	}
	return pool
}
