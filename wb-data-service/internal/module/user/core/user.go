package core

import "context"

type (
	UserUseCase interface {
		Save(context.Context, User) (User, error)
		GetByCreds(context.Context, User) (User, error)
	}

	UserRepository interface {
		Create(context.Context, User) (User, error)
		GetById(context.Context, User) (User, error)
		GetByCreds(context.Context, User) (User, error)
		Update(context.Context, User) error
		UpdatePassword(context.Context, User) error
		Delete(context.Context, User) error
	}
)
