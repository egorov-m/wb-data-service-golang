package core

import "context"

type (
	ProductRepository interface {
		MergeByNmId(context.Context, WbProductDetail) error
	}
)
