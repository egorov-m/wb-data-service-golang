package tasks

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"net/http"
)

const (
// CardDetailBaseUrl = "https://card.wb.ru/cards/v1/detail?" +
//
//	"appType=1&" +
//	"curr=rub&" +
//	"dest=-1257786&" +
//	"spp=30&" +
//	"nm="
)

func HandleProductLoadTask(c context.Context, t *asynq.Task) error {
	payload := t.Payload()
	var data NmIdPayload
	err := json.Unmarshal(payload, &data)
	if err != nil {
		return err
	}

	url := CardDetailBaseUrl + data["nm_id"].(string)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	return nil
}

func HandlePriceHistoryLoadTask(c context.Context, t *asynq.Task) error {
	nmId, err := t.Payload.GetInt("nm_id")

	return nil
}
