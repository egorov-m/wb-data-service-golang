package token

import "wb-data-service-golang/wb-data-service/internal/domain"

type TokenAdapter interface {
	GenerateAccess(userId int) (string, error)
	GenerateRefresh(userId int) (string, error)
	Parse(token string) (domain.TokenClaims, error)
}

type _TokenAdapter struct {
	Manager domain.TokenManager
}

func NewTokenAdapter(manager domain.TokenManager) TokenAdapter {
	return &_TokenAdapter{Manager: manager}
}

func (adapter *_TokenAdapter) GenerateAccess(userId int) (string, error) {
	return adapter.Manager.GenerateAccess(userId)
}

func (adapter *_TokenAdapter) GenerateRefresh(userId int) (string, error) {
	return adapter.Manager.GenerateRefresh(userId)
}

func (adapter *_TokenAdapter) Parse(token string) (domain.TokenClaims, error) {
	return adapter.Manager.Parse(token)
}
