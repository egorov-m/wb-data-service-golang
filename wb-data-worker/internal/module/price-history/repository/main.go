package repository

import (
	"wb-data-service-golang/wb-data-worker/internal/module/price-history/adapter/postgres"
	"wb-data-service-golang/wb-data-worker/internal/module/price-history/core"
	"wb-data-service-golang/wb-data-worker/internal/module/price-history/repository/shared"
)

type _PriceHistoryRepository struct {
	PostgresAdapter postgres.PostgresManagerAdapter[shared.PriceHistoryModel]
}

func NewPriceHistoryRepository(
	postgresAdapter postgres.PostgresManagerAdapter[shared.PriceHistoryModel],
) core.PriceHistoryRepository {
	return &_PriceHistoryRepository{
		PostgresAdapter: postgresAdapter,
	}
}
