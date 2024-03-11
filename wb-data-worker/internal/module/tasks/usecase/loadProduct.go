package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"io"
	"net/http"
	"wb-data-service-golang/wb-data-worker/internal/domain"
	"wb-data-service-golang/wb-data-worker/internal/module/product/core"
	core2 "wb-data-service-golang/wb-data-worker/internal/module/tasks/core"
	"wb-data-service-golang/wb-data-worker/internal/module/tasks/utils"
)

func (useCase *_WbTasksUseCase) LoadProduct(ctx context.Context, task *asynq.Task) error {
	payload := task.Payload()
	var data core2.NmIdPayload
	err := json.Unmarshal(payload, &data)
	if err != nil {
		useCase.Logger.Error(err, domain.LoggerArgs{
			"nm_id": data["nm_id"],
		})
		return err
	}

	url := utils.CardDetailBaseUrl + fmt.Sprintf("%d", int(data["nm_id"].(float64)))
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
	var productDetailData core.WbProductDetail
	err = json.Unmarshal(byteData, &productDetailData)
	if err != nil {
		useCase.Logger.Error(err, domain.LoggerArgs{
			"nm_id": data["nm_id"],
		})
		return err
	}

	err = useCase.ProductRepository.MergeProductByNmId(ctx, productDetailData)
	if err != nil {
		useCase.Logger.Error(err, domain.LoggerArgs{
			"nm_id": data["nm_id"],
		})
		return err
	}

	useCase.Logger.Info("Success load product ", domain.LoggerArgs{
		"nm_id": data["nm_id"],
	})

	return nil
}
