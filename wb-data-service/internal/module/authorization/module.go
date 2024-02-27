package authorization

import (
	"time"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/authorization/core"
	"wb-data-service-golang/wb-data-service/internal/module/authorization/usecase"
	passwordCryptoAdapter "wb-data-service-golang/wb-data-service/internal/module/password/adapter/cryptohash"
	passwordRepository "wb-data-service-golang/wb-data-service/internal/module/password/repository"
	tokenTokenAdapter "wb-data-service-golang/wb-data-service/internal/module/token/adapter/token"
	tokenRepository "wb-data-service-golang/wb-data-service/internal/module/token/repository"
	userCacheAdapter "wb-data-service-golang/wb-data-service/internal/module/user/adapter/cache"
	userPostgresAdapter "wb-data-service-golang/wb-data-service/internal/module/user/adapter/postgres"
	userRepository "wb-data-service-golang/wb-data-service/internal/module/user/repository"
	userShared "wb-data-service-golang/wb-data-service/internal/module/user/repository/shared"
)

type Dependency struct {
	Logger     domain.Logger
	Cache      domain.Cache
	Database   domain.DatabaseManager
	CryptoHash domain.CryptoManager
	Token      domain.TokenManager
	Timeout    time.Duration
}

func NewAuthorizationModule(dependency Dependency) core.AuthorizationUseCase {
	usrCacheAdapter := userCacheAdapter.NewCacheAdapter(dependency.Cache)
	usrPostgresAdapter := userPostgresAdapter.NewPostgresManagerAdapter[userShared.UserModel](dependency.Database)
	usrRepository := userRepository.NewUserRepository(
		usrCacheAdapter,
		usrPostgresAdapter,
	)

	passwordCryptoAdapter := passwordCryptoAdapter.NewCryptoHashAdapter(dependency.CryptoHash)
	passwordRepository := passwordRepository.NewPasswordRepository(passwordCryptoAdapter)

	tokenTokenAdapter := tokenTokenAdapter.NewTokenAdapter(dependency.Token)
	tokenRepository := tokenRepository.NewTokenRepository(tokenTokenAdapter)

	return usecase.NewAuthorizationUseCase(
		dependency.Logger,
		usrRepository,
		passwordRepository,
		tokenRepository,
		dependency.Timeout,
	)
}
