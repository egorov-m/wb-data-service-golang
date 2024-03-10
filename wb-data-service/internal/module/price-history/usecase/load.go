package repository

import (
	"context"
	"encoding/json"
	"wb-data-service-golang/wb-data-service/internal/module/price-history/core"
)

func (useCase *_PriceHistoryUseCase) Load(ctx context.Context, entity core.PriceHistory) (core.PriceHistoryTask, error) {

	payload := map[string]interface{}{
		"nm_id": entity.NmId,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return core.PriceHistoryTask{}, err
	}

	res, err := useCase.Worker.ProcessTaskByName(ctx, "low", "price_history_load_task", jsonData)

	return core.PriceHistoryTask{Id: res.ID, Type: res.Type}, nil
}
