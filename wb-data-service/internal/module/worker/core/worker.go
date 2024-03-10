package core

import (
	"context"

	product "wb-data-service-golang/wb-data-service/internal/module/product/core"
)

type (
	WorkerUseCase interface {
		GetProduct(context.Context, product.Product) (bool, error)
		UpdateProduct(context.Context, int) (bool, error)
		UpdatePriceHistory(context.Context, int) (bool, error)
	}
)
