package usecase

import (
	"context"
	"github.com/pkg/errors"
	"wb-data-service-golang/wb-data-service/internal/domain"
	token "wb-data-service-golang/wb-data-service/internal/module/token/core"
	user "wb-data-service-golang/wb-data-service/internal/module/user/core"
)

func (useCase *_AuthorizationUseCase) SignUp(ctx context.Context, entity user.User) (token.Token, error) {
	ctx, cancel := context.WithTimeout(ctx, useCase.defaultTimeout)
	defer cancel()

	if err := useCase.GenerateAndSetPassword(ctx, &entity); err != nil {
		useCase.Logger.Error(errors.Wrap(err, "generate password error"), domain.LoggerArgs{
			"user_id": entity.Id,
		})
		return token.Token{}, domain.ErrorInternalServer
	}

	user, err := useCase.UserRepository.Create(ctx, entity)
	if err != nil {
		useCase.Logger.Error(errors.Wrap(err, "create user error"), nil)
		return token.Token{}, domain.ErrorInternalServer
	}

	tokenModel, err := useCase.GenerateUserTokens(ctx, user)
	if err != nil {
		useCase.Logger.Error(errors.Wrap(err, "generate tokens error"), domain.LoggerArgs{
			"user_id": user.Id,
		})
		return token.Token{}, domain.ErrorInternalServer
	}

	return tokenModel, nil
}
