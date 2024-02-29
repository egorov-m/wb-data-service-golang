package pgxpool

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"wb-data-service-golang/wb-data-worker/pkg/slog"
)

// Aliases
type (
	CommandTag   pgconn.CommandTag
	Rows         = pgx.Rows
	Row          = pgx.Row
	Batch        = pgx.Batch
	BatchResults = pgx.BatchResults
)

func NewDatabase(ctx context.Context, databaseDsn string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(databaseDsn)
	if err != nil {
		slog.Fatal(err.Error())
	}

	config.ConnConfig.RuntimeParams = map[string]string{
		"client_encoding": "LATIN1",
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		slog.Fatal(err.Error())
	}

	return pool
}
