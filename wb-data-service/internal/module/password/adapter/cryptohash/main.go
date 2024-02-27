package cryptohash

import "wb-data-service-golang/wb-data-service/internal/domain"

type CryptoHashAdapter interface {
	GeneratePasswordHash(password string) (string, error)
}

type _CryptoHashAdapter struct {
	Client domain.CryptoManager
}

func NewCryptoHashAdapter(client domain.CryptoManager) CryptoHashAdapter {
	return &_CryptoHashAdapter{Client: client}
}

func (adapter *_CryptoHashAdapter) GeneratePasswordHash(password string) (string, error) {
	return adapter.Client.Encrypt(password)
}
