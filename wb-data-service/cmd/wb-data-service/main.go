package main

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"os/signal"
	"time"
	"wb-data-service-golang/wb-data-service/config"
	_ "wb-data-service-golang/wb-data-service/docs"
	"wb-data-service-golang/wb-data-service/internal/infrastructure/cache"
	"wb-data-service-golang/wb-data-service/internal/infrastructure/crypto"
	"wb-data-service-golang/wb-data-service/internal/infrastructure/database"
	"wb-data-service-golang/wb-data-service/internal/infrastructure/jwt"
	"wb-data-service-golang/wb-data-service/internal/module/authorization"
	"wb-data-service-golang/wb-data-service/internal/module/price-history"
	"wb-data-service-golang/wb-data-service/internal/module/product"
	"wb-data-service-golang/wb-data-service/pkg/pgxpool"

	"wb-data-service-golang/wb-data-service/internal/infrastructure/delivery/http"
	"wb-data-service-golang/wb-data-service/internal/infrastructure/logger"
)

const (
	defaultUseCaseTimeout = 5 * time.Second
)

// @title        Wildberries product service API.
// @version      0.0.1
// @license.name MIT license
// @license.url  https://github.com/egorov-m/wb-data-service-golang/blob/main/LICENSE
// @BasePath /api/v1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	config := config.New(".")
	slogger := logger.NewLogger(config.Env)
	logger := logger.New(slogger)

	// init data sources
	cache := cache.NewCache(time.Hour, time.Hour)
	pgxpool := pgxpool.NewDatabase(ctx, config.DatabaseDsn)
	defer pgxpool.Close()

	database := database.NewDatabaseManager(pgxpool, nil)

	// crypto
	cryptoManager := crypto.NewCryptoManager(config.PasswordSalt)

	// jwt
	jwtManager := jwt.NewTokenManager(config.TokenSalt)

	workerClient := asynq.NewClient(asynq.RedisClientOpt{
		Addr: fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
	})

	defer func(workerClient *asynq.Client) {
		err := workerClient.Close()
		if err != nil {
			logger.Error(err, nil)
		}
	}(workerClient)

	// init modules
	authModule := authorization.NewAuthorizationModule(authorization.Dependency{
		Logger:     logger,
		Cache:      cache,
		Database:   database,
		CryptoHash: cryptoManager,
		Token:      jwtManager,
		Timeout:    defaultUseCaseTimeout,
	})

	productModule := product.NewProductModule(product.Dependency{
		Logger:   logger,
		Database: database,
		Timeout:  defaultUseCaseTimeout,
		WbWorker: nil,
	})

	priceHistoryModule := price_history.NewPriceHistoryModule(price_history.Dependency{
		Logger:   logger,
		Database: database,
		Timeout:  defaultUseCaseTimeout,
		WbWorker: nil,
	})

	router := http.InitRoutes(authModule, productModule, priceHistoryModule)

	if config.Env != "prod" {
		//docs.SwaggerInfo.BasePath = "/docs"
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	if err := router.Run(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)); err != nil {
		logger.Error(err, nil)
	}
}
