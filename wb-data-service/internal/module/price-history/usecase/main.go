package repository

import (
	"time"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/price-history/core"
)

type _PriceHistoryUseCase struct {
	Logger                 domain.Logger
	PriceHistoryRepository core.PriceHistoryRepository

	defaultContextTimeout time.Duration
}

func NewPriceHistoryUseCase(
	logger domain.Logger,
	priceHistoryRepository core.PriceHistoryRepository,
	contextTimeout time.Duration,
) core.PriceHistoryUseCase {
	return &_PriceHistoryUseCase{
		Logger:                 logger,
		PriceHistoryRepository: priceHistoryRepository,
		defaultContextTimeout:  contextTimeout,
	}
}
