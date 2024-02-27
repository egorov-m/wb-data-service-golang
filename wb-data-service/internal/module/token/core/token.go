package core

import (
	"context"
	"wb-data-service-golang/wb-data-service/internal/domain"
)

type (
	TokenUseCase interface {
		RefreshToken(context.Context, Token) (Token, error)
	}

	TokenRepository interface {
		GenerateAccessToken(context.Context, int) (string, error)
		GenerateRefreshToken(context.Context, int) (string, error)
		DecodeToken(context.Context, string) (domain.TokenClaims, error)
	}
)
