package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"net/http"
)

func (wbTasks *_WbTasks) LoadPriceHistory(ctx context.Context, task *asynq.Task) error {
	payload := task.Payload()
	var data NmIdPayload
	err := json.Unmarshal(payload, &data)
	if err != nil {
		return err
	}

	nmIdFloat := data["nm_id"].(float64)
	url := GetUrlPriceHistory(int(nmIdFloat))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := wbTasks.HttpSession.SendRequest(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println(resp.StatusCode)

	return nil
}
