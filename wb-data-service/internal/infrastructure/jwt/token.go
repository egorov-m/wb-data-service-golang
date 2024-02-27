package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/pkg/byteconv"
)

const (
	expirationAccessTime  = time.Hour * 24
	expirationRefreshTime = expirationAccessTime * 30
)

type _TokenClaims struct {
	jwt.Claims
	UserId     int   `json:"user_id"`
	Expiration int64 `json:"expiration"`
}

type _TokenManager struct {
	TokenSalt []byte
}

func NewTokenManager(tokenSalt string) domain.TokenManager {
	return &_TokenManager{
		TokenSalt: byteconv.Bytes(tokenSalt),
	}
}

func (tokenManager *_TokenManager) generateToken(userId int, expiration time.Duration) (string, error) {
	claims := _TokenClaims{
		UserId:     userId,
		Expiration: time.Now().Add(expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(tokenManager.TokenSalt)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (tokenManager *_TokenManager) GenerateAccess(userId int) (string, error) {
	return tokenManager.generateToken(userId, expirationAccessTime)
}

func (tokenManager *_TokenManager) GenerateRefresh(userId int) (string, error) {
	return tokenManager.generateToken(userId, expirationRefreshTime)
}

func (tokenManager *_TokenManager) Parse(token string) (domain.TokenClaims, error) {
	signedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrorInvalidType
		}

		return tokenManager.TokenSalt, nil
	})
	if err != nil {
		return domain.TokenClaims{}, err
	}

	if !signedToken.Valid {
		return domain.TokenClaims{}, domain.ErrorInvalidToken
	}

	claims, ok := signedToken.Claims.(_TokenClaims)
	if !ok {
		return domain.TokenClaims{}, domain.ErrorInvalidToken
	}

	return domain.TokenClaims{
		UserId:     claims.UserId,
		Expiration: time.Unix(claims.Expiration, 0),
	}, nil
}
