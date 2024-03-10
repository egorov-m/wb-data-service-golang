package repository

import (
	"context"
	"github.com/pkg/errors"
	"wb-data-service-golang/wb-data-service/internal/module/product/core"
	"wb-data-service-golang/wb-data-service/internal/module/product/repository/internal/model"
	"wb-data-service-golang/wb-data-service/internal/module/product/repository/internal/query"
)

func (repository *_ProductRepository) GetQuantity(ctx context.Context, entity core.Product) (int, error) {

	selectModel := model.NewProduct(entity)
	sql, args, err := query.GetQuantity(selectModel)
	if err != nil {
		return 0, errors.Wrap(err, "generate select quantity product sql query error")
	}

	connection := repository.PostgresAdapter.GetConnect()

	model, err := connection.QueryInt(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "execute select quantity product query error")
	}

	return model, nil
}
