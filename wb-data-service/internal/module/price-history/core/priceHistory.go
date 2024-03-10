package core

import "context"

type (
	PriceHistoryUseCase interface {
		Load(context.Context, PriceHistory) (PriceHistoryTask, error)
		GetByProductNmId(context.Context, int) ([]PriceHistory, error)
	}

	PriceHistoryRepository interface {
		GetByProductNmId(context.Context, PriceHistory) ([]PriceHistory, error)
	}
)
