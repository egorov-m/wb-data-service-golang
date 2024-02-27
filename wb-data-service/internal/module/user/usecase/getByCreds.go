package usecase

import (
	"context"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/user/core"
)

func (useCase *_UserUseCase) GetByCreds(ctx context.Context, entity core.User) (core.User, error) {
	ctx, cancel := context.WithTimeout(ctx, useCase.defaultContextTimeout)
	defer cancel()

	entity, err := useCase.UserRepository.GetByCreds(ctx, entity)
	if err != nil {
		useCase.Logger.Error(err, domain.LoggerArgs{
			"user_id":    entity.Id,
			"user_email": entity.Email,
		})
		return core.User{}, domain.ErrorInternalServer
	}

	if entity.IsEmpty() {
		useCase.Logger.Info("empty user", domain.LoggerArgs{
			"user_id": entity.Id,
		})
		return core.User{}, domain.ErrorNotFound
	}

	return entity, nil
}
