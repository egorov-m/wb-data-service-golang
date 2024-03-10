package postgres

import (
	"context"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/infrastructure/database"
	"wb-data-service-golang/wb-data-service/internal/module/price-history/repository/shared"
)

type PostgresAdapter[T shared.PriceHistoryModel] interface {
	Exec(ctx context.Context, sql string, args ...any) error
	Query(ctx context.Context, sql string, args ...any) (T, error)
	QueryMany(ctx context.Context, sql string, args ...any) ([]*T, error)
}

type _PostgresAdapter[T shared.PriceHistoryModel] struct {
	Database domain.Database
}

func NewPostgresAdapter[T shared.PriceHistoryModel](database domain.Database) PostgresAdapter[T] {
	return &_PostgresAdapter[T]{Database: database}
}

func (adapter *_PostgresAdapter[T]) Exec(ctx context.Context, sql string, args ...any) error {
	_, err := adapter.Database.Exec(ctx, sql, args...)
	return err
}

func (adapter *_PostgresAdapter[T]) Query(ctx context.Context, sql string, args ...any) (T, error) {
	rows, err := adapter.Database.Query(ctx, sql, args...)
	if err != nil {
		return T{}, err
	}

	result, err := database.CollectOneRow[T](rows)
	if err != nil {
		return T{}, err
	}

	return *result, nil
}

func (adapter *_PostgresAdapter[T]) QueryMany(ctx context.Context, sql string, args ...any) ([]*T, error) {
	rows, err := adapter.Database.Query(ctx, sql, args...)
	if err != nil {
		return []*T{}, err
	}

	result, err := database.CollectRows[T](rows)
	if err != nil {
		return []*T{}, err
	}

	return result, nil
}
