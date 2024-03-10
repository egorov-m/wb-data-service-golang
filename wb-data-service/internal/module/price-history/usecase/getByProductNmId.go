package repository

import (
	"context"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/price-history/core"
)

func (useCase *_PriceHistoryUseCase) GetByProductNmId(ctx context.Context, nmId int) ([]core.PriceHistory, error) {
	ctx, cancel := context.WithTimeout(ctx, useCase.defaultContextTimeout)
	defer cancel()

	res, err := useCase.PriceHistoryRepository.GetByProductNmId(ctx, core.PriceHistory{NmId: nmId})
	if err != nil {
		useCase.Logger.Error(err, domain.LoggerArgs{
			"product_nm_id": nmId,
		})

		return []core.PriceHistory{}, domain.ErrorInternalServer
	}

	return res, nil
}
