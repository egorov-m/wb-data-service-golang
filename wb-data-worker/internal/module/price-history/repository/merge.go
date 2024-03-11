package repository

import (
	"context"
	"github.com/pkg/errors"
	"wb-data-service-golang/wb-data-worker/internal/module/price-history/core"
	"wb-data-service-golang/wb-data-worker/internal/module/price-history/repository/internal/model"
	"wb-data-service-golang/wb-data-worker/internal/module/price-history/repository/internal/query"
)

func (repository *_PriceHistoryRepository) MergeByProductNmId(ctx context.Context, entity core.WbPricesHistory, nmId int) error {
	insertModels := model.NewPriceHistory(entity, nmId)

	transaction, err := repository.PostgresAdapter.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "begin merge price history transaction error")
	}
	connection := transaction.GetConnect()

	for _, value := range insertModels {
		sql, args, err := query.GetSelectByNmIdAndDt(value)
		if err != nil {
			return errors.Wrap(err, "generate select price history by nm_id and dt sql-query error")
		}
		_, err = connection.Query(ctx, sql, args...)
		if err != nil {
			sql, args, err := query.GetInsert(value)
			if err != nil {
				return errors.Wrap(err, "insert price history sql-query error")
			}
			_, err = connection.Query(ctx, sql, args...)
			if err != nil {
				if err := transaction.Rollback(ctx); err != nil {
					return errors.Wrap(err, "rollback insert price history transaction after execute query error")
				}
				return errors.Wrap(err, "insert price history in database error")
			}
		} else {
			sql, args, err := query.GetUpdate(value)
			if err != nil {
				return errors.Wrap(err, "update price history sql-query error")
			}
			err = connection.Exec(ctx, sql, args...)
			if err != nil {
				if err := transaction.Rollback(ctx); err != nil {
					return errors.Wrap(err, "rollback update price history transaction after execute query error")
				}
				return errors.Wrap(err, "update price history in database error")
			}
		}
	}

	if err := transaction.Commit(ctx); err != nil {
		return errors.Wrap(err, "commit merge price history transaction error")
	}

	return nil
}
