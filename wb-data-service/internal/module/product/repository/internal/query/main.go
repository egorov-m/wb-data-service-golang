package query

import (
	"wb-data-service-golang/wb-data-service/internal/module/product/repository/internal/model"
	"wb-data-service-golang/wb-data-service/pkg/goqu"
)

func GetInsert(model model.Product) (string, []any, error) {
	record := make(goqu.Record)

	record["nm_id"] = model.NmId
	record["name"] = model.Name
	record["brand"] = model.Brand
	record["brand_id"] = model.BrandId
	record["site_brand_id"] = model.SiteBrandId
	record["supplier_id"] = model.SupplierId
	record["sale"] = model.Sale
	record["price"] = model.Price
	record["sale_price"] = model.SalePrice
	record["rating"] = model.Rating
	record["feedbacks"] = model.Feedbacks

	if model.Colors.Valid {
		record["colors"] = model.Colors.String
	}

	record["quantity"] = model.Quantity

	if model.CreatedAt.Valid {
		record["created_at"] = model.CreatedAt.Time
	}

	if model.UpdatedAt.Valid {
		record["updated_at"] = model.UpdatedAt.Time
	}

	return goqu.Dialect.
		Insert(model.TableName()).
		Rows(record).
		Returning(model).
		Prepared(true).
		ToSQL()
}

func GetSelectById(model model.Product) (string, []any, error) {
	return goqu.Dialect.
		Select(model).
		From(model.TableName()).
		Where(
			goqu.C("nm_id").Eq(model.NmId),
		).Prepared(true).ToSQL()
}

func GetCount(model model.Product) (string, []any, error) {
	return goqu.Dialect.
		Select(goqu.COUNT("*")).
		From(model.TableName()).
		Prepared(true).
		ToSQL()
}

func GetQuantity(model model.Product) (string, []any, error) {
	return goqu.Dialect.
		Select(goqu.COALESCE(goqu.SUM("quantity"), 0)).
		From(model.TableName()).
		Prepared(true).
		ToSQL()
}

func GetAll(model model.Product) (string, []any, error) {
	return goqu.Dialect.
		Select(model).
		From(model.TableName()).
		Prepared(true).ToSQL()
}

func GetDelete(model model.Product) (string, []any, error) {
	return goqu.Dialect.
		Delete(model.TableName()).
		Where(
			goqu.C("nm_id").Eq(model.NmId),
		).Prepared(true).ToSQL()
}
