package usecase

import (
	"context"
	"encoding/json"
	"wb-data-service-golang/wb-data-service/internal/module/product/core"
)

func (useCase *_ProductUseCase) Load(ctx context.Context, entity core.Product) (core.ProductTask, error) {

	payload := map[string]interface{}{
		"nm_id": entity.NmId,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return core.ProductTask{}, err
	}

	res, err := useCase.Worker.ProcessTaskByName(ctx, "low", "product_load_task", jsonData)

	return core.ProductTask{Id: res.ID, Type: res.Type}, nil
}
