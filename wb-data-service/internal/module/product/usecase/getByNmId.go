package usecase

import (
	"context"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/product/core"
)

func (useCase *_ProductUseCase) GetByNmId(ctx context.Context, entity core.Product) (core.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, useCase.defaultContextTimeout)
	defer cancel()

	entity, err := useCase.ProductRepository.GetByNmId(ctx, entity)
	if err != nil {
		useCase.Logger.Error(err, domain.LoggerArgs{
			"product_nm_id": entity.NmId,
		})

		return core.Product{}, domain.ErrorInternalServer
	}

	if entity.IsEmpty() {
		useCase.Logger.Info("product empty", domain.LoggerArgs{
			"product_nm_id": entity.NmId,
		})
		return core.Product{}, domain.ErrorNotFound
	}

	return entity, nil
}
