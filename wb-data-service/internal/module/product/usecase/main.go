package usecase

import (
	"time"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/product/core"
)

type _ProductUseCase struct {
	Logger            domain.Logger
	ProductRepository core.ProductRepository

	defaultContextTimeout time.Duration
}

func NewProductUseCase(
	logger domain.Logger,
	productRepository core.ProductRepository,
	contextTimeout time.Duration,
) core.ProductUseCase {
	return &_ProductUseCase{
		Logger:                logger,
		ProductRepository:     productRepository,
		defaultContextTimeout: contextTimeout,
	}
}
