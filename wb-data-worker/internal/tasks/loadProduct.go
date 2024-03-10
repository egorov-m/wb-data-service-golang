package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"net/http"
)

func (wbTasks *_WbTasks) LoadProduct(ctx context.Context, task *asynq.Task) error {
	payload := task.Payload()
	var data NmIdPayload
	err := json.Unmarshal(payload, &data)
	if err != nil {
		return err
	}

	url := CardDetailBaseUrl + fmt.Sprintf("%d", data["nm_id"])
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
