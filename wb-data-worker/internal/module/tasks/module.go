package tasks

import (
	"net/http"
	"time"
	"wb-data-service-golang/wb-data-worker/internal/domain"
	"wb-data-service-golang/wb-data-worker/internal/infrastructure/httpClient"
	priceHistoryPostgres "wb-data-service-golang/wb-data-worker/internal/module/price-history/adapter/postgres"
	priceHistory "wb-data-service-golang/wb-data-worker/internal/module/price-history/repository"
	"wb-data-service-golang/wb-data-worker/internal/module/price-history/repository/shared"
	productPostgres "wb-data-service-golang/wb-data-worker/internal/module/product/adapter/postgres"
	product "wb-data-service-golang/wb-data-worker/internal/module/product/repository"
	productShared "wb-data-service-golang/wb-data-worker/internal/module/product/repository/shared"
	"wb-data-service-golang/wb-data-worker/internal/module/tasks/core"
	"wb-data-service-golang/wb-data-worker/internal/module/tasks/usecase"
)

type Dependency struct {
	Logger   domain.Logger
	Database domain.DatabaseManager
	Timeout  time.Duration
}

func NewTasksModule(dependency Dependency) core.WbTasksUseCase {
	productPostgresAdapter := productPostgres.NewPostgresManagerAdapter[productShared.ProductModel](dependency.Database)
	productRepository := product.NewProductRepository(
		productPostgresAdapter,
	)
	priceHistoryPostgresAdapter := priceHistoryPostgres.NewPostgresManagerAdapter[shared.PriceHistoryModel](dependency.Database)
	priceHistoryRepository := priceHistory.NewPriceHistoryRepository(
		priceHistoryPostgresAdapter,
	)

	baseHttpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	client := httpClient.NewHttpClient(baseHttpClient, map[int]bool{200: true})

	return usecase.NewWbTasksUseCase(
		dependency.Logger,
		client,
		productRepository,
		priceHistoryRepository,
		dependency.Timeout,
	)
}
