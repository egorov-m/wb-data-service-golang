package usecase

import (
	"context"
	"github.com/pkg/errors"
	"time"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/authorization/core"
	"wb-data-service-golang/wb-data-service/internal/module/authorization/usecase/internal/utils"
	password "wb-data-service-golang/wb-data-service/internal/module/password/core"
	token "wb-data-service-golang/wb-data-service/internal/module/token/core"
	user "wb-data-service-golang/wb-data-service/internal/module/user/core"
)

type _AuthorizationUseCase struct {
	Logger             domain.Logger
	UserRepository     user.UserRepository
	PasswordRepository password.PasswordRepository
	TokenRepository    token.TokenRepository

	defaultTimeout time.Duration
}

func NewAuthorizationUseCase(
	logger domain.Logger,
	userRepository user.UserRepository,
	passwordRepository password.PasswordRepository,
	tokenRepository token.TokenRepository,
	timeout time.Duration,
) core.AuthorizationUseCase {
	return &_AuthorizationUseCase{
		Logger:             logger,
		UserRepository:     userRepository,
		PasswordRepository: passwordRepository,
		TokenRepository:    tokenRepository,
		defaultTimeout:     timeout,
	}
}

func (useCase *_AuthorizationUseCase) GenerateUserTokens(ctx context.Context, entity user.User) (token.Token, error) {
	accessToken, err := useCase.TokenRepository.GenerateAccessToken(ctx, entity.Id)
	if err != nil {
		return token.Token{}, errors.Wrap(err, "generate access token error")
	}

	//refreshToken, err := useCase.TokenRepository.GenerateRefreshToken(ctx, entity.Id)
	//if err != nil {
	//	return token.Token{}, errors.Wrap(err, "generate refresh token error")
	//}

	return utils.NewToken(accessToken), nil
}

func (useCase *_AuthorizationUseCase) GenerateAndSetPassword(ctx context.Context, entity *user.User) error {
	pass, err := useCase.PasswordRepository.GeneratePassword(ctx, entity.Password)
	if err != nil {
		return errors.Wrap(err, "generate password error")
	}
	entity.Password = pass

	return nil
}

func (useCase *_AuthorizationUseCase) CheckToken(ctx context.Context, bearerToken string, tokenType string) (int, error) {
	claims, err := useCase.TokenRepository.DecodeToken(ctx, bearerToken)
	if err != nil {
		return claims.UserId, err
	}
	if claims.Type != tokenType {
		return claims.UserId, domain.ErrorInvalidToken
	}
	if time.Now().Before(claims.Expiration) {
		return claims.UserId, nil
	} else {
		return claims.UserId, domain.ErrorExpirationToken
	}
}
