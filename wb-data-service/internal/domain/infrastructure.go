package domain

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"time"
)

type (
	LoggerArgs map[string]any

	Logger interface {
		Debug(msg string, args LoggerArgs)
		Error(err error, args LoggerArgs)
		Warn(msg string, args LoggerArgs)
		Info(msg string, args LoggerArgs)
	}
)

type (
	Cache interface {
		Set(string, any, time.Duration)
		Get(string) (any, error)
		Del(string)
	}
)

type (
	CryptoManager interface {
		Encrypt(string) (string, error)
	}
)

type (
	TokenClaims struct {
		UserId     int
		Expiration time.Time
	}

	TokenManager interface {
		GenerateAccess(int) (string, error)
		GenerateRefresh(int) (string, error)
		Parse(string) (TokenClaims, error)
	}
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

type (
	WbWorker interface {
		ProcessTask(context.Context, string, *asynq.Task) (*asynq.TaskInfo, error)
		ProcessTaskByName(context.Context, string, string, []byte) (*asynq.TaskInfo, error)
	}
)
