package repository

import (
	"context"
	"github.com/pkg/errors"
	"wb-data-service-golang/wb-data-service/internal/module/price-history/core"
	"wb-data-service-golang/wb-data-service/internal/module/price-history/repository/internal/model"
	"wb-data-service-golang/wb-data-service/internal/module/price-history/repository/internal/query"
)

func (repository *_PriceHistoryRepository) GetByProductNmId(ctx context.Context, entity core.PriceHistory) ([]core.PriceHistory, error) {
	selectModel := model.NewPriceHistory(entity)
	sql, args, err := query.GetSelectByNmId(selectModel)
	if err != nil {
		return []core.PriceHistory{}, errors.Wrap(err, "generate select price history by nm_id sql query error")
	}

	connection := repository.PostgresAdapter.GetConnect()
	res, err := connection.QueryMany(ctx, sql, args...)
	if err != nil {
		return []core.PriceHistory{}, errors.Wrap(err, "execute select price history by nm_id query error")
	}

	t := make([]core.PriceHistory, len(res))
	for i, entity := range res {
		t[i] = entity.ToEntity()
	}
	return t, nil
}
