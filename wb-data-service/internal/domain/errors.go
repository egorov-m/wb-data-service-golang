package domain

import "errors"

var (
	ErrorInternalServer  = errors.New("internal server error")
	ErrorNotFound        = errors.New("not found error")
	ErrorInvalidType     = errors.New("invalid type error")
	ErrorInvalidToken    = errors.New("invalid token")
	ErrorExpirationToken = errors.New("token expiration")
)
