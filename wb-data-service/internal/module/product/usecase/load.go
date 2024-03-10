package usecase

import (
	"context"
	"wb-data-service-golang/wb-data-service/internal/module/product/core"
)

func (useCase *_ProductUseCase) Load(ctx context.Context, entity core.Product) (bool, error) {

	// TODO Put the task in the queue for the Worker

	return false, nil
}
