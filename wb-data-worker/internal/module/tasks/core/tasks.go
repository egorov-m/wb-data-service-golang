package core

import (
	"context"
	"github.com/hibiken/asynq"
)

type (
	WbTasksUseCase interface {
		LoadProduct(context.Context, *asynq.Task) error
		LoadPriceHistory(context.Context, *asynq.Task) error
	}
)
