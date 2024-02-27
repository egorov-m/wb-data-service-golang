package usecase

import (
	"context"
	"github.com/pkg/errors"
	"time"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/token/core"
	"wb-data-service-golang/wb-data-service/internal/module/token/usecase/internal/utils"
)

type _TokenUseCase struct {
	Logger          domain.Logger
	TokenRepository core.TokenRepository

	defaultTimeout time.Duration
}

func NewTokenUseCase(
	logger domain.Logger,
	tokenRepository core.TokenRepository,
	defaultTimeout time.Duration,
) core.TokenUseCase {
	return &_TokenUseCase{
		Logger:          logger,
		TokenRepository: tokenRepository,
		defaultTimeout:  defaultTimeout,
	}
}

func (useCase *_TokenUseCase) RefreshToken(ctx context.Context, entity core.Token) (core.Token, error) {
	ctx, cancel := context.WithTimeout(ctx, useCase.defaultTimeout)
	defer cancel()

	decodeToken, err := useCase.TokenRepository.DecodeToken(ctx, entity.RefreshToken)
	if err != nil {
		useCase.Logger.Debug(err.Error(), nil)
		return core.Token{}, domain.ErrorInternalServer
	}

	// check expiration time
	if utils.CheckRefreshExpirationTime(decodeToken.Expiration) {
		refreshToken, err := useCase.TokenRepository.GenerateRefreshToken(ctx, decodeToken.UserId)
		if err != nil {
			useCase.Logger.Error(errors.Wrap(err, "generate access token error"), domain.LoggerArgs{
				"user_id": decodeToken.UserId,
			})
			return core.Token{}, domain.ErrorInternalServer
		}

		entity.RefreshToken = refreshToken
	}

	accessToken, err := useCase.TokenRepository.GenerateAccessToken(ctx, decodeToken.UserId)
	if err != nil {
		useCase.Logger.Error(errors.Wrap(err, "generate access token error"), domain.LoggerArgs{
			"user_id": decodeToken.UserId,
		})
		return core.Token{}, domain.ErrorInternalServer
	}

	entity.AccessToken = accessToken

	return entity, nil
}
