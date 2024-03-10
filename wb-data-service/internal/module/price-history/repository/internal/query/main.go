package query

import (
	"wb-data-service-golang/wb-data-service/internal/module/price-history/repository/internal/model"
	"wb-data-service-golang/wb-data-service/pkg/goqu"
)

func GetInsert(model model.PriceHistory) (string, []any, error) {
	record := make(goqu.Record)

	record["id"] = model.Id
	record["nm_id"] = model.NmId
	record["dt"] = model.Dt
	record["price"] = model.Price

	return goqu.Dialect.
		Insert(model.TableName()).
		Rows(record).
		Returning(model).
		Prepared(true).
		ToSQL()
}

func GetSelectByNmId(model model.PriceHistory) (string, []any, error) {
	return goqu.Dialect.
		Select(model).
		From(model.TableName()).
		Where(
			goqu.C("nm_id").Eq(model.NmId),
		).Prepared(true).ToSQL()
}
