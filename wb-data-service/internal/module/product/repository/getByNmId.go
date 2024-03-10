package repository

import (
	"context"
	"github.com/pkg/errors"
	"wb-data-service-golang/wb-data-service/internal/module/product/core"
	"wb-data-service-golang/wb-data-service/internal/module/product/repository/internal/model"
	"wb-data-service-golang/wb-data-service/internal/module/product/repository/internal/query"
)

func (repository *_ProductRepository) GetByNmId(ctx context.Context, entity core.Product) (core.Product, error) {

	selectModel := model.NewProduct(entity)
	sql, args, err := query.GetSelectById(selectModel)
	if err != nil {
		return core.Product{}, errors.Wrap(err, "generate select product sql query error")
	}

	connection := repository.PostgresAdapter.GetConnect()

	model, err := connection.Query(ctx, sql, args...)
	if err != nil {
		return core.Product{}, errors.Wrap(err, "execute select product query error")
	}

	return model.ToEntity(), nil
}
