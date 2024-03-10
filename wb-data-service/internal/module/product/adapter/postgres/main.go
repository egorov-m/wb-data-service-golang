package postgres

import (
	"context"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/product/repository/shared"
)

type PostgresManagerAdapter[T shared.ProductModel] interface {
	Begin(ctx context.Context) (PostgresManagerAdapter[T], error)
	Commit(ctx context.Context) error
	GetConnect() PostgresAdapter[T]
	Rollback(ctx context.Context) error
}

type _PostgresManagerAdapter[T shared.ProductModel] struct {
	DatabaseManager domain.DatabaseManager
}

func NewPostgresManagerAdapter[T shared.ProductModel](databaseManager domain.DatabaseManager) PostgresManagerAdapter[T] {
	return &_PostgresManagerAdapter[T]{
		DatabaseManager: databaseManager,
	}
}

func (adapter *_PostgresManagerAdapter[T]) Begin(ctx context.Context) (PostgresManagerAdapter[T], error) {
	databaseManager, err := adapter.DatabaseManager.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return NewPostgresManagerAdapter[T](databaseManager), nil
}

func (adapter *_PostgresManagerAdapter[T]) Commit(ctx context.Context) error {
	return adapter.DatabaseManager.Commit(ctx)
}

func (adapter *_PostgresManagerAdapter[T]) GetConnect() PostgresAdapter[T] {
	connection := adapter.DatabaseManager.GetConnect()
	return NewPostgresAdapter[T](connection)
}

func (adapter *_PostgresManagerAdapter[T]) Rollback(ctx context.Context) error {
	return adapter.DatabaseManager.Rollback(ctx)
}
