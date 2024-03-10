package product

import (
	"time"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/product/adapter/postgres"
	"wb-data-service-golang/wb-data-service/internal/module/product/core"
	"wb-data-service-golang/wb-data-service/internal/module/product/repository"
	"wb-data-service-golang/wb-data-service/internal/module/product/repository/shared"
	"wb-data-service-golang/wb-data-service/internal/module/product/usecase"
)

type Dependency struct {
	Logger   domain.Logger
	Database domain.DatabaseManager
	Timeout  time.Duration
	WbWorker domain.WbWorker
}

func NewProductModule(dependency Dependency) core.ProductUseCase {
	productPostgresAdapter := postgres.NewPostgresManagerAdapter[shared.ProductModel](dependency.Database)

	productRepository := repository.NewProductRepository(
		productPostgresAdapter,
	)

	return usecase.NewProductUseCase(
		dependency.Logger,
		productRepository,
		dependency.Timeout,
		dependency.WbWorker,
	)
}
