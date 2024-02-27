package usecase

import (
	"context"
	"github.com/pkg/errors"
	"time"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/password/core"
	"wb-data-service-golang/wb-data-service/internal/module/password/usecase/internal/utils"
	user "wb-data-service-golang/wb-data-service/internal/module/user/core"
)

type _PasswordUseCase struct {
	Logger             domain.Logger
	UserRepository     user.UserRepository
	PasswordRepository core.PasswordRepository
	defaultTimeout     time.Duration
}

func NewPasswordUseCase(
	logger domain.Logger,
	userRepository user.UserRepository,
	passwordRepository core.PasswordRepository,
	timeout time.Duration,
) core.PasswordUseCase {
	return &_PasswordUseCase{
		Logger:             logger,
		UserRepository:     userRepository,
		PasswordRepository: passwordRepository,
		defaultTimeout:     timeout,
	}
}

func (useCase *_PasswordUseCase) ChangePassword(ctx context.Context, entity core.Password) error {
	ctx, cancel := context.WithTimeout(ctx, useCase.defaultTimeout)
	defer cancel()

	password, err := useCase.PasswordRepository.GeneratePassword(ctx, entity.Password)
	if err != nil {
		useCase.Logger.Error(errors.Wrap(err, "generate password error"), domain.LoggerArgs{
			"user_id": entity.UserId,
		})
		return domain.ErrorInternalServer
	}
	entity.Password = password

	usr := utils.ConvertPasswordToUser(entity)
	if err := useCase.UserRepository.UpdatePassword(ctx, usr); err != nil {
		useCase.Logger.Error(errors.Wrap(err, "generate password error"), domain.LoggerArgs{
			"user_id": entity.UserId,
		})
		return domain.ErrorInternalServer
	}

	return nil
}
