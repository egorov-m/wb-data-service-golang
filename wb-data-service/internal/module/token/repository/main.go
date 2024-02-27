package repository

import (
	"context"
	"github.com/pkg/errors"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/token/adapter/token"
	"wb-data-service-golang/wb-data-service/internal/module/token/core"
)

type _TokenRepository struct {
	TokenAdapter token.TokenAdapter
}

func NewTokenRepository(tokenAdapter token.TokenAdapter) core.TokenRepository {
	return &_TokenRepository{TokenAdapter: tokenAdapter}
}

func (repository *_TokenRepository) DecodeToken(_ context.Context, token string) (domain.TokenClaims, error) {
	claims, err := repository.TokenAdapter.Parse(token)
	if err != nil {
		return domain.TokenClaims{}, errors.Wrap(err, "parse token error")
	}

	return claims, nil
}

func (repository *_TokenRepository) GenerateAccessToken(_ context.Context, userId int) (string, error) {
	accessToken, err := repository.TokenAdapter.GenerateAccess(userId)
	if err != nil {
		return "", errors.Wrap(err, "generate access token error")
	}

	return accessToken, nil
}

func (repository *_TokenRepository) GenerateRefreshToken(_ context.Context, userId int) (string, error) {
	refreshToken, err := repository.TokenAdapter.GenerateRefresh(userId)
	if err != nil {
		return "", errors.Wrap(err, "generate refresh token error")
	}

	return refreshToken, nil
}
