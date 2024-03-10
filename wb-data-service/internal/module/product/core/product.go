package core

import "context"

type (
	ProductUseCase interface {
		Load(context.Context, Product) (ProductTask, error)
		GetByNmId(context.Context, Product) (Product, error)
		GetAll(context.Context) ([]Product, error)
		GetCount(context.Context) (int, error)
		GetQuantity(context.Context) (int, error)
		DeleteByNmId(context.Context, Product) error
	}

	ProductRepository interface {
		GetByNmId(context.Context, Product) (Product, error)
		GetAll(context.Context, Product) ([]Product, error)
		GetCount(context.Context, Product) (int, error)
		GetQuantity(context.Context, Product) (int, error)
		DeleteByNmId(context.Context, Product) error
	}
)
