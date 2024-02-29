package main

import (
	"fmt"
	"github.com/hibiken/asynq"
	"log"
	"wb-data-service-golang/wb-data-worker/config"
	"wb-data-service-golang/wb-data-worker/internal/module/product/task/asynq/handler"
	"wb-data-service-golang/wb-data-worker/internal/module/product/task/asynq/payload"
)

func main() {
	//ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	//defer cancel()

	config := config.New(".")

	redisConnection := asynq.RedisClientOpt{
		Addr: fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
	}

	//pgxpool := pgxpool.NewDatabase(ctx, config.DatabaseDsn)
	//defer pgxpool.Close()
	//database := database.NewDatabaseManager(pgxpool, nil)

	worker := asynq.NewServer(redisConnection, asynq.Config{
		// Specify how many concurrent workers to use.
		Concurrency: config.Concurrency,
		// Specify multiple queues with different priority.
		Queues: map[string]int{
			"critical": 6, // processed 60% of the time
			"default":  3, // processed 30% of the time
			"low":      1, // processed 10% of the time
		},
	})

	mux := asynq.NewServeMux()

	mux.HandleFunc(payload.TypeGetProduct, handler.HandlerGetProductTask)
	mux.HandleFunc(payload.TypeUpdateProduct, handler.HandlerUpdateProductTask)
	mux.HandleFunc(payload.TypeUpdatePriceHistory, handler.HandlerUpdatePriceHistoryTask)

	// Run worker server.
	if err := worker.Run(mux); err != nil {
		log.Fatal(err)
	}
}
