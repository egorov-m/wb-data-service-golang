package usecase

import (
	"context"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/product/core"
)

func (useCase *_ProductUseCase) GetAll(ctx context.Context) ([]core.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, useCase.defaultContextTimeout)
	defer cancel()

	res, err := useCase.ProductRepository.GetAll(ctx, core.Product{})
	if err != nil {
		useCase.Logger.Error(err, nil)
		return []core.Product{}, domain.ErrorInternalServer
	}

	return res, nil
}
