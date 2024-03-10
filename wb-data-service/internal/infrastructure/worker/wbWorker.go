package worker

import (
	"context"
	"github.com/hibiken/asynq"
	"wb-data-service-golang/wb-data-service/internal/domain"
)

type _WbWorker struct {
	Client *asynq.Client
}

func NewWbWorker(client *asynq.Client) domain.WbWorker {
	return &_WbWorker{
		Client: client,
	}
}

func (worker *_WbWorker) ProcessTask(ctx context.Context, queueLevel string, task *asynq.Task) (*asynq.TaskInfo, error) {
	res, err := worker.Client.EnqueueContext(ctx, task, asynq.Queue(queueLevel))

	return res, err
}

func (worker *_WbWorker) ProcessTaskByName(ctx context.Context, queueLevel string, taskName string, payload []byte) (*asynq.TaskInfo, error) {
	task := asynq.NewTask(taskName, payload)
	res, err := worker.Client.EnqueueContext(ctx, task, asynq.Queue(queueLevel))

	return res, err
}
