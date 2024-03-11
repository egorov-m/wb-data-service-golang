package tasks

import (
	"net/http"
	"time"
	"wb-data-service-golang/wb-data-worker/internal/domain"
	"wb-data-service-golang/wb-data-worker/internal/infrastructure/httpClient"
	"wb-data-service-golang/wb-data-worker/internal/module/product/adapter/postgres"
	"wb-data-service-golang/wb-data-worker/internal/module/product/repository"
	"wb-data-service-golang/wb-data-worker/internal/module/product/repository/shared"
	"wb-data-service-golang/wb-data-worker/internal/module/tasks/core"
	"wb-data-service-golang/wb-data-worker/internal/module/tasks/usecase"
)

type Dependency struct {
	Logger   domain.Logger
	Database domain.DatabaseManager
	Timeout  time.Duration
}

func NewTasksModule(dependency Dependency) core.WbTasksUseCase {
	productPostgresAdapter := postgres.NewPostgresManagerAdapter[shared.ProductModel](dependency.Database)
	productRepository := repository.NewProductRepository(
		productPostgresAdapter,
	)

	baseHttpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	client := httpClient.NewHttpClient(baseHttpClient, map[int]bool{200: true})

	return usecase.NewWbTasksUseCase(
		dependency.Logger,
		client,
		productRepository,
		dependency.Timeout,
	)
}
