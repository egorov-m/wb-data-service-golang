package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"net/http"
	"wb-data-service-golang/wb-data-worker/internal/module/tasks/core"
	"wb-data-service-golang/wb-data-worker/internal/module/tasks/utils"
)

func (useCase *_WbTasksUseCase) LoadPriceHistory(ctx context.Context, task *asynq.Task) error {
	payload := task.Payload()
	var data core.NmIdPayload
	err := json.Unmarshal(payload, &data)
	if err != nil {
		return err
	}

	nmIdFloat := data["nm_id"].(float64)
	url := utils.GetUrlPriceHistory(int(nmIdFloat))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := useCase.HttpClient.SendRequest(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println(resp.StatusCode)

	return nil
}
