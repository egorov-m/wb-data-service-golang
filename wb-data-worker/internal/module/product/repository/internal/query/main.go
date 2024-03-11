package query

import (
	"wb-data-service-golang/wb-data-worker/internal/module/product/repository/internal/model"
	"wb-data-service-golang/wb-data-worker/pkg/goqu"
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
		//OnConflict(goqu.DoUpdate("name", goqu.C("name").Set(goqu.I("exclude.name")))).
		//OnConflict(goqu.DoUpdate("brand", goqu.C("brand").Set(goqu.I("exclude.brand")))).
		//OnConflict(goqu.DoUpdate("brand_id", goqu.C("brand_id").Set(goqu.I("exclude.brand_id")))).
		//OnConflict(goqu.DoUpdate("site_brand_id", goqu.C("site_brand_id").Set(goqu.I("exclude.site_brand_id")))).
		//OnConflict(goqu.DoUpdate("supplier_id", goqu.C("supplier_id").Set(goqu.I("exclude.supplier_id")))).
		//OnConflict(goqu.DoUpdate("sale", goqu.C("sale").Set(goqu.I("exclude.sale")))).
		//OnConflict(goqu.DoUpdate("price", goqu.C("price").Set(goqu.I("exclude.price")))).
		//OnConflict(goqu.DoUpdate("sale_price", goqu.C("sale_price").Set(goqu.I("exclude.sale_price")))).
		//OnConflict(goqu.DoUpdate("rating", goqu.C("rating").Set(goqu.I("exclude.rating")))).
		//OnConflict(goqu.DoUpdate("feedbacks", goqu.C("feedbacks").Set(goqu.I("exclude.feedbacks")))).
		//OnConflict(goqu.DoUpdate("colors", goqu.C("colors").Set(goqu.I("exclude.colors")))).
		//OnConflict(goqu.DoUpdate("quantity", goqu.C("quantity").Set(goqu.I("exclude.quantity")))).
		//OnConflict(goqu.DoUpdate("create_at", goqu.C("create_at").Set(goqu.I("exclude.create_at")))).
		//OnConflict(goqu.DoUpdate("update_at", goqu.C("update_at").Set(goqu.I("exclude.update_at")))).
		Returning(model).
		Prepared(true).
		ToSQL()
}

func GetUpdate(model model.Product) (string, []any, error) {
	record := make(goqu.Record)
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
	if model.UpdatedAt.Valid {
		record["updated_at"] = model.UpdatedAt.Time
	}

	return goqu.Dialect.
		Update(model.TableName()).
		Set(record).
		Where(goqu.C("nm_id").Eq(model.NmId)).
		Prepared(true).
		ToSQL()
}

func GetSelectByNmId(model model.Product) (string, []any, error) {
	return goqu.Dialect.
		Select(model).
		From(model.TableName()).
		Where(
			goqu.C("nm_id").Eq(model.NmId),
		).Prepared(true).ToSQL()
}
