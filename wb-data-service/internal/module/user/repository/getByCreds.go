package repository

import (
	"context"
	"github.com/pkg/errors"
	"wb-data-service-golang/wb-data-service/internal/module/user/core"
	"wb-data-service-golang/wb-data-service/internal/module/user/repository/internal/model"
	"wb-data-service-golang/wb-data-service/internal/module/user/repository/internal/query"
)

func (repository *_UserRepository) GetByCreds(ctx context.Context, entity core.User) (core.User, error) {
	connection := repository.PostgresAdapter.GetConnect()

	selectModel := model.NewUser(entity)

	sql, args, err := query.GetSelectByCreds(selectModel)
	if err != nil {
		return core.User{}, errors.Wrap(err, "generate select user by creds sql query error")
	}

	model, err := connection.Query(ctx, sql, args...)
	if err != nil {
		return core.User{}, errors.Wrap(err, "execute select user by creds sql query error")
	}

	return model.ToEntity(), nil
}
