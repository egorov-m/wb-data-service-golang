package payload

import (
	"encoding/json"
	"github.com/hibiken/asynq"
)

const (
	TypeGetProduct         = "product:get"
	TypeUpdateProduct      = "product:update"
	TypeUpdatePriceHistory = "price-history:update"
)

type GetProductPayload struct {
	NmId         int
	IsOverridden bool
}

type UpdateProductPayload struct {
	NmId int
}

type UpdatePriceHistoryPayload struct {
	NmId         int
	IsOverridden bool
}

func NewGetProductTask(nmId int, isOverridden bool) (*asynq.Task, error) {
	payload, err := json.Marshal(GetProductPayload{NmId: nmId, IsOverridden: isOverridden})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeGetProduct, payload), nil
}

func NewUpdateProductTask(nmId int) (*asynq.Task, error) {
	payload, err := json.Marshal(UpdateProductPayload{NmId: nmId})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeUpdateProduct, payload), nil
}

func NewUpdatePriceHistoryTask(nmId int) (*asynq.Task, error) {
	payload, err := json.Marshal(UpdatePriceHistoryPayload{NmId: nmId})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeUpdatePriceHistory, payload), nil
}
