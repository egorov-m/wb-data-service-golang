package user

import (
	"time"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/user/adapter/cache"
	"wb-data-service-golang/wb-data-service/internal/module/user/adapter/postgres"
	"wb-data-service-golang/wb-data-service/internal/module/user/core"
	"wb-data-service-golang/wb-data-service/internal/module/user/repository"
	"wb-data-service-golang/wb-data-service/internal/module/user/repository/shared"
	"wb-data-service-golang/wb-data-service/internal/module/user/usecase"
)

type Dependency struct {
	Logger   domain.Logger
	Cache    domain.Cache
	Database domain.DatabaseManager
	Timeout  time.Duration
}

func NewUserModule(dependency Dependency) core.UserUseCase {
	userCacheAdapter := cache.NewCacheAdapter(dependency.Cache)
	userPostgresAdapter := postgres.NewPostgresManagerAdapter[shared.UserModel](dependency.Database)

	userRepository := repository.NewUserRepository(
		userCacheAdapter,
		userPostgresAdapter,
	)

	return usecase.NewUserUseCase(
		dependency.Logger,
		userRepository,
		dependency.Timeout,
	)
}
