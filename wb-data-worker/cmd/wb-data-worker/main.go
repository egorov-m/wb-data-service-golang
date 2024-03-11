package main

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
	"os"
	"os/signal"
	"time"
	"wb-data-service-golang/wb-data-worker/config"
	"wb-data-service-golang/wb-data-worker/internal/infrastructure/database"
	"wb-data-service-golang/wb-data-worker/internal/infrastructure/logger"
	"wb-data-service-golang/wb-data-worker/internal/module/tasks"
	"wb-data-service-golang/wb-data-worker/internal/module/tasks/core"
	"wb-data-service-golang/wb-data-worker/pkg/pgxpool"
)

const (
	defaultUseCaseTimeout = 5 * time.Second
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	config := config.New(".")

	slogger := logger.NewLogger(config.Env)
	logger := logger.New(slogger)

	redisConnection := asynq.RedisClientOpt{
		Addr: fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
	}

	pgxpool := pgxpool.NewDatabase(ctx, config.DatabaseDsn)
	defer pgxpool.Close()
	database := database.NewDatabaseManager(pgxpool, nil)

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

	tasksModule := tasks.NewTasksModule(tasks.Dependency{
		Logger:   logger,
		Database: database,
		Timeout:  defaultUseCaseTimeout,
	})

	mux := asynq.NewServeMux()

	mux.HandleFunc(core.TypeLoadProduct, tasksModule.LoadProduct)
	mux.HandleFunc(core.TypeLoadPriceHistory, tasksModule.LoadPriceHistory)

	// Run worker server.
	if err := worker.Run(mux); err != nil {
		log.Fatal(err)
	}
}
