package postgres

import (
	"context"
	"wb-data-service-golang/wb-data-worker/internal/domain"
	"wb-data-service-golang/wb-data-worker/internal/infrastructure/database"
	"wb-data-service-golang/wb-data-worker/internal/module/product/repository/shared"
)

type PostgresAdapter[T shared.ProductModel] interface {
	Exec(ctx context.Context, sql string, args ...any) error
	Query(ctx context.Context, sql string, args ...any) (T, error)
}

type _PostgresAdapter[T shared.ProductModel] struct {
	Database domain.Database
}

func NewPostgresAdapter[T shared.ProductModel](database domain.Database) PostgresAdapter[T] {
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
