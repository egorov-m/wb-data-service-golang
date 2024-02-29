package domain

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type (
	DatabaseManager interface {
		Begin(context.Context) (DatabaseManager, error)
		Commit(context.Context) error
		Rollback(context.Context) error

		GetConnect() Database
	}

	Database interface {
		Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
		Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
		QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
		SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	}
)
