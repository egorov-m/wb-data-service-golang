package tasks

import (
	"encoding/json"
	"github.com/hibiken/asynq"
)

const (
	TypeLoadProduct = "product_load_task"

	TypeLoadPriceHistory = "price_history_load_task"
)

type (
	NmIdPayload = map[string]interface{}
)

func NewProductLoadTask(nmId int) (*asynq.Task, error) {
	payload := NmIdPayload{
		"nm_id": nmId,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeLoadProduct, jsonData), nil
}

func NewPriceHistoryLoadTask(nmId int) (*asynq.Task, error) {
	payload := NmIdPayload{
		"nm_id": nmId,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeLoadPriceHistory, jsonData), nil
}
