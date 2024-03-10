package repository

import (
	"context"
	"wb-data-service-golang/wb-data-service/internal/module/price-history/core"
)

func (useCase *_PriceHistoryUseCase) Load(ctx context.Context, entity core.PriceHistory) (bool, error) {

	// TODO Put the task in the queue for the Worker

	return false, nil
}
