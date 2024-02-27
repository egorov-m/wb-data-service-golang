package repository

import (
	"context"
	"github.com/pkg/errors"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/user/core"
	"wb-data-service-golang/wb-data-service/internal/module/user/repository/internal/model"
	"wb-data-service-golang/wb-data-service/internal/module/user/repository/internal/query"
)

func (repository *_UserRepository) GetById(ctx context.Context, entity core.User) (core.User, error) {
	if entity, err := repository.CacheAdapter.Get(entity); err != nil {
		if !errors.Is(err, domain.ErrorNotFound) {
			return core.User{}, errors.Wrap(err, "cache user error")
		}
	} else {
		return entity, nil
	}

	selectModel := model.NewUser(entity)
	sql, args, err := query.GetSelectById(selectModel)
	if err != nil {
		return core.User{}, errors.Wrap(err, "generate select user sql query error")
	}

	connection := repository.PostgresAdapter.GetConnect()

	model, err := connection.Query(ctx, sql, args...)
	if err != nil {
		return core.User{}, errors.Wrap(err, "execute select user query error")
	}

	return model.ToEntity(), nil
}
