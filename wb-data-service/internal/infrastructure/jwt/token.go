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
	UserId     int    `json:"user_id"`
	Expiration int64  `json:"expiration"`
	Type       string `json:"type"`
}

type _TokenManager struct {
	TokenSalt []byte
}

func NewTokenManager(tokenSalt string) domain.TokenManager {
	return &_TokenManager{
		TokenSalt: byteconv.Bytes(tokenSalt),
	}
}

func (tokenManager *_TokenManager) generateToken(userId int, expiration time.Duration, tokenType string) (string, error) {
	claims := _TokenClaims{
		UserId:     userId,
		Expiration: time.Now().Add(expiration).Unix(),
		Type:       tokenType,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(tokenManager.TokenSalt)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (tokenManager *_TokenManager) GenerateAccess(userId int) (string, error) {
	return tokenManager.generateToken(userId, expirationAccessTime, "access")
}

func (tokenManager *_TokenManager) GenerateRefresh(userId int) (string, error) {
	return tokenManager.generateToken(userId, expirationRefreshTime, "refresh")
}

func (tokenManager *_TokenManager) Parse(token string) (domain.TokenClaims, error) {
	signedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
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

	claims, ok := signedToken.Claims.(jwt.MapClaims)
	if !ok {
		return domain.TokenClaims{}, domain.ErrorInvalidToken
	}

	return domain.TokenClaims{
		UserId:     int(claims["user_id"].(float64)),
		Expiration: time.Unix(int64(claims["expiration"].(float64)), 0),
		Type:       claims["type"].(string),
	}, nil
}
