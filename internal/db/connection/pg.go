package connection

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var PgPool *pgxpool.Pool

func OpenDB() (*pgxpool.Pool, error) {
	pgconn := os.Getenv("POSTGRESQL_URL")
	if pgconn == "" {
		return nil, fmt.Errorf("POSTGRESQL_URL must be provided")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, pgconn)
	if err != nil {
		return nil, fmt.Errorf("database error: %w ", err)
	}

	pool.Config().MaxConns = 10
	pool.Config().MaxConnIdleTime = 20 * time.Second

	PgPool = pool

	return PgPool, nil
}
