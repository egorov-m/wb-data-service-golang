package repository

import (
	"context"
	"github.com/pkg/errors"
	"wb-data-service-golang/wb-data-service/internal/module/user/core"
	"wb-data-service-golang/wb-data-service/internal/module/user/repository/internal/model"
	"wb-data-service-golang/wb-data-service/internal/module/user/repository/internal/query"
)

func (repository *_UserRepository) Update(ctx context.Context, entity core.User) error {
	repository.CacheAdapter.Del(entity)

	model := model.NewUser(entity)
	sql, args, err := query.GetUpdate(model)
	if err != nil {
		return errors.Wrap(err, "generate update user sql-query error")
	}

	connection := repository.PostgresAdapter.GetConnect()
	if err := connection.Exec(ctx, sql, args...); err != nil {
		return errors.Wrap(err, "execute update user query error")
	}

	return nil
}
