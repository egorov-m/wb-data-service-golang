package repository

import (
	"wb-data-service-golang/wb-data-service/internal/module/product/adapter/postgres"
	"wb-data-service-golang/wb-data-service/internal/module/product/core"
	"wb-data-service-golang/wb-data-service/internal/module/product/repository/shared"
)

type _ProductRepository struct {
	PostgresAdapter postgres.PostgresManagerAdapter[shared.ProductModel]
}

func NewProductRepository(
	postgresAdapter postgres.PostgresManagerAdapter[shared.ProductModel],
) core.ProductRepository {
	return &_ProductRepository{
		PostgresAdapter: postgresAdapter,
	}
}
