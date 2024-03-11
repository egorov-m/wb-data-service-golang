package core

import "context"

type (
	ProductRepository interface {
		MergeProductByNmId(context.Context, WbProductDetail) error
	}
)
