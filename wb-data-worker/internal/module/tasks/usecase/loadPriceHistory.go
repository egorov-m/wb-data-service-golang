package usecase

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"io"
	"net/http"
	"wb-data-service-golang/wb-data-worker/internal/domain"
	"wb-data-service-golang/wb-data-worker/internal/module/price-history/core"
	tasksCore "wb-data-service-golang/wb-data-worker/internal/module/tasks/core"
	"wb-data-service-golang/wb-data-worker/internal/module/tasks/utils"
)

func (useCase *_WbTasksUseCase) LoadPriceHistory(ctx context.Context, task *asynq.Task) error {
	payload := task.Payload()
	var data tasksCore.NmIdPayload
	err := json.Unmarshal(payload, &data)
	if err != nil {
		useCase.Logger.Error(err, domain.LoggerArgs{
			"nm_id": data["nm_id"],
		})
		return err
	}

	nmIdFloat := data["nm_id"].(float64)
	url := utils.GetUrlPriceHistory(int(nmIdFloat))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		useCase.Logger.Error(err, domain.LoggerArgs{
			"nm_id":   data["nm_id"],
			"request": req,
		})
		return err
	}

	resp, err := useCase.HttpClient.SendRequest(ctx, req)
	if err != nil {
		useCase.Logger.Error(err, domain.LoggerArgs{
			"nm_id":    data["nm_id"],
			"response": resp,
		})
		return err
	}
	byteData, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		useCase.Logger.Error(err, domain.LoggerArgs{
			"nm_id": data["nm_id"],
		})
		return err
	}
	var priceHistoryData core.WbPricesHistory
	err = json.Unmarshal(byteData, &priceHistoryData)
	if err != nil {
		useCase.Logger.Error(err, domain.LoggerArgs{
			"nm_id": data["nm_id"],
		})
		return err
	}

	err = useCase.PriceHistoryRepository.MergeByProductNmId(ctx, priceHistoryData, int(nmIdFloat))
	if err != nil {
		useCase.Logger.Error(err, domain.LoggerArgs{
			"nm_id": data["nm_id"],
		})
		return err
	}

	useCase.Logger.Info("Success load price history", domain.LoggerArgs{
		"nm_id": data["nm_id"],
	})

	return nil
}
