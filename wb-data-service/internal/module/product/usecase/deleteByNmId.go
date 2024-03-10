package usecase

import (
	"context"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/product/core"
)

func (useCase *_ProductUseCase) DeleteByNmId(ctx context.Context, entity core.Product) error {
	ctx, cancel := context.WithTimeout(ctx, useCase.defaultContextTimeout)
	defer cancel()

	err := useCase.ProductRepository.DeleteByNmId(ctx, entity)
	if err != nil {
		useCase.Logger.Error(err, domain.LoggerArgs{
			"product_nm_id": entity.NmId,
		})
		return domain.ErrorInternalServer
	}

	return nil
}
