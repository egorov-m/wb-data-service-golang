package repository

import (
	"context"
	"github.com/pkg/errors"
	"wb-data-service-golang/wb-data-service/internal/module/user/core"
	"wb-data-service-golang/wb-data-service/internal/module/user/repository/internal/model"
	"wb-data-service-golang/wb-data-service/internal/module/user/repository/internal/query"
)

func (repository *_UserRepository) Delete(ctx context.Context, entity core.User) error {
	connection := repository.PostgresAdapter.GetConnect()

	repository.CacheAdapter.Del(entity)

	deleteModel := model.NewUser(entity)
	sql, args, err := query.GetDelete(deleteModel)
	if err != nil {
		return errors.Wrap(err, "generate delete user sql-query error")
	}

	if err := connection.Exec(ctx, sql, args...); err != nil {
		return errors.Wrap(err, "delete user from database error")
	}

	return nil
}
