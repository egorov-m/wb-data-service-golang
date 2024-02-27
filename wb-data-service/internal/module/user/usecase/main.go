package usecase

import (
	"time"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/user/core"
)

type _UserUseCase struct {
	Logger         domain.Logger
	UserRepository core.UserRepository

	defaultContextTimeout time.Duration
}

func NewUserUseCase(
	logger domain.Logger,
	userRepository core.UserRepository,
	contextTimeout time.Duration,
) core.UserUseCase {
	return &_UserUseCase{
		Logger:                logger,
		UserRepository:        userRepository,
		defaultContextTimeout: contextTimeout,
	}
}
