package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
	"wb-data-service-golang/wb-data-service/config"
	"wb-data-service-golang/wb-data-service/internal/infrastructure/cache"
	"wb-data-service-golang/wb-data-service/internal/infrastructure/crypto"
	"wb-data-service-golang/wb-data-service/internal/infrastructure/database"
	"wb-data-service-golang/wb-data-service/internal/infrastructure/jwt"
	"wb-data-service-golang/wb-data-service/internal/module/authorization"
	"wb-data-service-golang/wb-data-service/pkg/pgxpool"

	http2 "net/http"
	"wb-data-service-golang/wb-data-service/internal/infrastructure/delivery/http"
	"wb-data-service-golang/wb-data-service/internal/infrastructure/logger"
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

	// init data sources
	cache := cache.NewCache(time.Hour, time.Hour)
	pgxpool := pgxpool.NewDatabase(ctx, config.DatabaseDsn)
	defer pgxpool.Close()

	database := database.NewDatabaseManager(pgxpool, nil)

	// crypto
	cryptoManager := crypto.NewCryptoManager(config.PasswordSalt)

	// jwt
	jwtManager := jwt.NewTokenManager(config.TokenSalt)

	// init modules
	authModule := authorization.NewAuthorizationModule(authorization.Dependency{
		Logger:     logger,
		Cache:      cache,
		Database:   database,
		CryptoHash: cryptoManager,
		Token:      jwtManager,
		Timeout:    defaultUseCaseTimeout,
	})

	router := http.InitRoutes(authModule)

	if err := http2.ListenAndServe(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port), router); err != nil {
		logger.Error(err, nil)
	}
}
