package usecase

import (
	"context"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/product/core"
)

func (useCase *_ProductUseCase) GetQuantity(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, useCase.defaultContextTimeout)
	defer cancel()

	res, err := useCase.ProductRepository.GetQuantity(ctx, core.Product{})
	if err != nil {
		useCase.Logger.Error(err, nil)
		return 0, domain.ErrorInternalServer
	}

	return res, nil
}
