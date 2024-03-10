package utils

import token "wb-data-service-golang/wb-data-service/internal/module/token/core"

func NewToken(access string) token.Token {
	return token.Token{
		AccessToken: access,
		//RefreshToken: refresh,
	}
}
