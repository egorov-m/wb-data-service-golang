package main

import (
	"fmt"
	"github.com/hibiken/asynq"
	"log"
	"net/http"
	"time"
	"wb-data-service-golang/wb-data-worker/config"
	"wb-data-service-golang/wb-data-worker/internal/infrastructure/httpSession"
	"wb-data-service-golang/wb-data-worker/internal/infrastructure/logger"
	tasks2 "wb-data-service-golang/wb-data-worker/internal/tasks"
)

func main() {
	//ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	//defer cancel()

	config := config.New(".")

	slogger := logger.NewLogger(config.Env)
	logger := logger.New(slogger)

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
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	session := httpSession.NewHttpSession(httpClient, map[int]bool{200: true})
	wbTasks := tasks2.NewWbTasks(logger, session)

	mux := asynq.NewServeMux()

	mux.HandleFunc(tasks2.TypeLoadProduct, wbTasks.LoadProduct)
	mux.HandleFunc(tasks2.TypeLoadPriceHistory, wbTasks.LoadPriceHistory)

	// Run worker server.
	if err := worker.Run(mux); err != nil {
		log.Fatal(err)
	}
}
