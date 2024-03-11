package repository

import (
	"context"
	"github.com/pkg/errors"
	"wb-data-service-golang/wb-data-worker/internal/module/product/core"
	"wb-data-service-golang/wb-data-worker/internal/module/product/repository/internal/model"
	"wb-data-service-golang/wb-data-worker/internal/module/product/repository/internal/query"
)

func (repository *_ProductRepository) MergeProductByNmId(ctx context.Context, entity core.WbProductDetail) error {
	insertModel, err := model.NewProduct(entity)
	if err != nil {
		return err
	}

	sql, args, err := query.GetSelectByNmId(insertModel)
	if err != nil {
		return errors.Wrap(err, "generate select product by id sql-query error")
	}
	transaction, err := repository.PostgresAdapter.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "begin merge product transaction error")
	}
	connection := transaction.GetConnect()

	_, err = connection.Query(ctx, sql, args...)
	if err != nil {
		sql, args, err := query.GetInsert(insertModel)
		if err != nil {
			return errors.Wrap(err, "insert product sql-query error")
		}
		_, err = connection.Query(ctx, sql, args...)
		if err != nil {
			if err := transaction.Rollback(ctx); err != nil {
				return errors.Wrap(err, "rollback insert product transaction after execute query error")
			}
			return errors.Wrap(err, "insert product in database error")
		}
	} else {
		sql, args, err := query.GetUpdate(insertModel)
		if err != nil {
			return errors.Wrap(err, "update product sql-query error")
		}
		err = connection.Exec(ctx, sql, args...)
		if err != nil {
			if err := transaction.Rollback(ctx); err != nil {
				return errors.Wrap(err, "rollback update product transaction after execute query error")
			}
			return errors.Wrap(err, "update product in database error")
		}
	}

	if err := transaction.Commit(ctx); err != nil {
		return errors.Wrap(err, "commit merge product transaction error")
	}

	return nil
}
