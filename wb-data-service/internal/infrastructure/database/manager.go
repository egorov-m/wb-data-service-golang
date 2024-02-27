package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"wb-data-service-golang/wb-data-service/internal/domain"
)

type _DatabaseManager struct {
	Conn *pgxpool.Pool
	Tx   pgx.Tx
}

func NewDatabaseManager(conn *pgxpool.Pool, tx pgx.Tx) domain.DatabaseManager {
	return &_DatabaseManager{
		Conn: conn,
		Tx:   tx,
	}
}

func (database *_DatabaseManager) Begin(ctx context.Context) (domain.DatabaseManager, error) {
	transaction, err := database.Conn.Begin(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "begin database transaction error")
	}

	return NewDatabaseManager(
		database.Conn,
		transaction,
	), nil
}

func (database *_DatabaseManager) ClearTransaction() {
	database.Tx = nil
}

func (database *_DatabaseManager) Commit(ctx context.Context) error {
	defer database.ClearTransaction()
	if err := database.Tx.Commit(ctx); err != nil {
		return errors.Wrap(err, "commit database transaction error")
	}

	return nil
}

func (database *_DatabaseManager) Rollback(ctx context.Context) error {
	defer database.ClearTransaction()
	if err := database.Tx.Rollback(ctx); err != nil {
		return errors.Wrap(err, "rollback database transaction error")
	}

	return nil
}

func (database *_DatabaseManager) GetConnect() domain.Database {
	if database.Tx != nil {
		return database.Tx
	}

	return database.Conn
}
