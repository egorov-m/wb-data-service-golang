package repository

import (
	"context"
	"github.com/pkg/errors"
	"wb-data-service-golang/wb-data-service/internal/module/product/core"
	"wb-data-service-golang/wb-data-service/internal/module/product/repository/internal/model"
	"wb-data-service-golang/wb-data-service/internal/module/product/repository/internal/query"
)

func (repository *_ProductRepository) DeleteByNmId(ctx context.Context, entity core.Product) error {
	deleteModel := model.NewProduct(entity)
	sql, args, err := query.GetDelete(deleteModel)
	if err != nil {
		return errors.Wrap(err, "generate delete product sql query error")
	}

	connection := repository.PostgresAdapter.GetConnect()
	if err := connection.Exec(ctx, sql, args...); err != nil {
		return errors.Wrap(err, "delete product from database error")
	}

	return nil
}
