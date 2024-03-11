package query

import (
	"wb-data-service-golang/wb-data-worker/internal/module/price-history/repository/internal/model"
	"wb-data-service-golang/wb-data-worker/pkg/goqu"
)

func GetInsert(model model.PriceHistory) (string, []any, error) {
	record := make(goqu.Record)

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

func GetSelectByNmIdAndDt(model model.PriceHistory) (string, []any, error) {
	return goqu.Dialect.
		Select(model).
		From(model.TableName()).
		Where(
			goqu.C("nm_id").Eq(model.NmId),
			goqu.C("dt").Eq(model.Dt),
		).Prepared(true).
		ToSQL()
}

func GetUpdate(model model.PriceHistory) (string, []any, error) {
	record := make(goqu.Record)
	record["nm_id"] = model.NmId
	record["dt"] = model.Dt
	record["price"] = model.Price

	return goqu.Dialect.
		Update(model.TableName()).
		Set(record).
		Where(
			goqu.C("nm_id").Eq(model.NmId),
			goqu.C("dt").Eq(model.Dt),
		).Prepared(true).
		ToSQL()
}
