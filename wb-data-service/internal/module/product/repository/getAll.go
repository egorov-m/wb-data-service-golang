package repository

import (
	"context"
	"github.com/pkg/errors"
	"wb-data-service-golang/wb-data-service/internal/module/product/core"
	"wb-data-service-golang/wb-data-service/internal/module/product/repository/internal/model"
	"wb-data-service-golang/wb-data-service/internal/module/product/repository/internal/query"
)

func (repository *_ProductRepository) GetAll(ctx context.Context, entity core.Product) ([]core.Product, error) {

	selectModel := model.NewProduct(entity)
	sql, args, err := query.GetAll(selectModel)
	if err != nil {
		return []core.Product{}, errors.Wrap(err, "generate select product sql query error")
	}

	connection := repository.PostgresAdapter.GetConnect()

	model, err := connection.QueryMany(ctx, sql, args...)
	if err != nil {
		return []core.Product{}, errors.Wrap(err, "execute select product query error")
	}

	t := make([]core.Product, len(model))
	for i, entity := range model {
		t[i] = entity.ToEntity()
	}
	return t, nil
}
