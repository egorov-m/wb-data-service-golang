package core

import (
	"context"
	token "wb-data-service-golang/wb-data-service/internal/module/token/core"
	user "wb-data-service-golang/wb-data-service/internal/module/user/core"
)

type (
	AuthorizationUseCase interface {
		SignIn(context.Context, user.User) (token.Token, error)
		SignUp(context.Context, user.User) (token.Token, error)
		CheckToken(context.Context, string) (int, error)
	}
)
