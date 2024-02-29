package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"

	"wb-data-service-golang/wb-data-worker/internal/module/product/task/asynq/payload"
)

func HandlerGetProductTask(c context.Context, t *asynq.Task) error {
	var p payload.GetProductPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	// TODO

	return nil
}

func HandlerUpdateProductTask(c context.Context, t *asynq.Task) error {
	var p payload.UpdateProductPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	// TODO

	return nil
}

func HandlerUpdatePriceHistoryTask(c context.Context, t *asynq.Task) error {
	var p payload.UpdatePriceHistoryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	// TODO

	return nil
}
