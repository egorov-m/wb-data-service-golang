package core

import "context"

type (
	PriceHistoryRepository interface {
		MergeByProductNmId(context.Context, WbPricesHistory, int) error
	}
)
