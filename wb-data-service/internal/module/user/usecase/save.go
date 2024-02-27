package usecase

import (
	"context"
	"github.com/pkg/errors"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/user/core"
)

func (useCase *_UserUseCase) Save(ctx context.Context, entity core.User) (core.User, error) {
	ctx, cancel := context.WithTimeout(ctx, useCase.defaultContextTimeout)
	defer cancel()

	if entity.Id != 0 {
		if err := useCase.UserRepository.Update(ctx, entity); err != nil {
			useCase.Logger.Error(errors.Wrap(err, "update user error"), domain.LoggerArgs{
				"user_id": entity.Id,
			})
			return core.User{}, domain.ErrorInternalServer
		}

		entity, err := useCase.UserRepository.GetById(ctx, entity)
		if err != nil {
			useCase.Logger.Error(errors.Wrap(err, "get user by id error"), domain.LoggerArgs{
				"user_id": entity.Id,
			})
			return core.User{}, domain.ErrorInternalServer
		}

		return entity, nil
	}

	entity, err := useCase.UserRepository.Create(ctx, entity)
	if err != nil {
		useCase.Logger.Error(errors.Wrap(err, "create user error"), nil)
		return core.User{}, domain.ErrorInternalServer
	}

	return entity, nil
}
