package repository

import (
	"context"
	"github.com/pkg/errors"
	"wb-data-service-golang/wb-data-service/internal/module/user/core"
	"wb-data-service-golang/wb-data-service/internal/module/user/repository/internal/model"
	"wb-data-service-golang/wb-data-service/internal/module/user/repository/internal/query"
)

func (repository *_UserRepository) Create(ctx context.Context, entity core.User) (core.User, error) {
	insertModel := model.NewUser(entity)

	sql, args, err := query.GetInsert(insertModel)
	if err != nil {
		return core.User{}, errors.Wrap(err, "generate create user sql-query error")
	}

	transaction, err := repository.PostgresAdapter.Begin(ctx)
	if err != nil {
		return core.User{}, errors.Wrap(err, "begin create user transaction error")
	}
	connection := transaction.GetConnect()

	model, err := connection.Query(ctx, sql, args...)
	if err != nil {
		if err := transaction.Rollback(ctx); err != nil {
			return core.User{}, errors.Wrap(err, "rollback create user transaction after execute query error")
		}
		return core.User{}, errors.Wrap(err, "create user in database error")
	}

	if err := transaction.Commit(ctx); err != nil {
		return core.User{}, errors.Wrap(err, "commit create user transaction error")
	}

	entity = model.ToEntity()

	repository.CacheAdapter.Set(entity)

	return entity, nil
}
