package repository

import (
	"wb-data-service-golang/wb-data-service/internal/module/user/adapter/cache"
	"wb-data-service-golang/wb-data-service/internal/module/user/adapter/postgres"
	"wb-data-service-golang/wb-data-service/internal/module/user/core"
	"wb-data-service-golang/wb-data-service/internal/module/user/repository/shared"
)

type _UserRepository struct {
	CacheAdapter    cache.CacheAdapter
	PostgresAdapter postgres.PostgresManagerAdapter[shared.UserModel]
}

func NewUserRepository(
	cacheAdapter cache.CacheAdapter,
	postgresAdapter postgres.PostgresManagerAdapter[shared.UserModel],
) core.UserRepository {
	return &_UserRepository{
		CacheAdapter:    cacheAdapter,
		PostgresAdapter: postgresAdapter,
	}
}
