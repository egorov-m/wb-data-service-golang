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

	url := GetUrlPriceHistory(data["nm_id"].(int))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := wbTasks.HttpSession.SendRequest(ctx, req)
	fmt.Println(resp.StatusCode)

	return nil
}
