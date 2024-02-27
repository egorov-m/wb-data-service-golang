package repository

import (
	"context"
	"github.com/pkg/errors"
	"wb-data-service-golang/wb-data-service/internal/module/password/adapter/cryptohash"
	"wb-data-service-golang/wb-data-service/internal/module/password/core"
)

type _PasswordRepository struct {
	CryptoHashAdapter cryptohash.CryptoHashAdapter
}

func NewPasswordRepository(
	cryptoHashAdapter cryptohash.CryptoHashAdapter,
) core.PasswordRepository {
	return &_PasswordRepository{
		CryptoHashAdapter: cryptoHashAdapter,
	}
}

func (repository *_PasswordRepository) GeneratePassword(ctx context.Context, password string) (string, error) {
	password, err := repository.CryptoHashAdapter.GeneratePasswordHash(password)
	if err != nil {
		return "", errors.Wrap(err, "generate password hash error")
	}

	return password, nil
}
