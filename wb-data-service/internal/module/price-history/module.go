package price_history

import (
	"time"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/price-history/adapter/postgres"
	"wb-data-service-golang/wb-data-service/internal/module/price-history/core"
	"wb-data-service-golang/wb-data-service/internal/module/price-history/repository"
	"wb-data-service-golang/wb-data-service/internal/module/price-history/repository/shared"
	usecase "wb-data-service-golang/wb-data-service/internal/module/price-history/usecase"
)

type Dependency struct {
	Logger   domain.Logger
	Database domain.DatabaseManager
	Timeout  time.Duration
	WbWorker domain.WbWorker
}

func NewPriceHistoryModule(dependency Dependency) core.PriceHistoryUseCase {
	priceHistoryPostgresAdapter := postgres.NewPostgresManagerAdapter[shared.PriceHistoryModel](dependency.Database)

	priceHistoryRepository := repository.NewPriceHistoryRepository(
		priceHistoryPostgresAdapter,
	)

	return usecase.NewPriceHistoryUseCase(
		dependency.Logger,
		priceHistoryRepository,
		dependency.Timeout,
	)
}
