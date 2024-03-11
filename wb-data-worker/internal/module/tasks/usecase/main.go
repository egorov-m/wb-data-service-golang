package usecase

import (
	"time"
	"wb-data-service-golang/wb-data-worker/internal/domain"
	product "wb-data-service-golang/wb-data-worker/internal/module/product/core"
	"wb-data-service-golang/wb-data-worker/internal/module/tasks/core"
)

type _WbTasksUseCase struct {
	Logger            domain.Logger
	HttpClient        domain.HttpClient
	ProductRepository product.ProductRepository

	defaultContextTimeout time.Duration
}

func NewWbTasksUseCase(
	logger domain.Logger,
	httpClient domain.HttpClient,
	productRepository product.ProductRepository,
	contextTimeout time.Duration,
) core.WbTasksUseCase {
	return &_WbTasksUseCase{
		Logger:                logger,
		HttpClient:            httpClient,
		ProductRepository:     productRepository,
		defaultContextTimeout: contextTimeout,
	}
}
