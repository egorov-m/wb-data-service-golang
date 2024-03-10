package model

import (
	"database/sql"
	"wb-data-service-golang/wb-data-service/internal/module/product/core"
)

type Product struct {
	NmId        int            `db:"nm_id"`
	Name        string         `db:"name"`
	Brand       string         `db:"brand"`
	BrandId     int            `db:"brand_id"`
	SiteBrandId int            `db:"site_brand_id"`
	SupplierId  int            `db:"supplier_id"`
	Sale        int            `db:"sale"`
	Price       int            `db:"price"`
	SalePrice   int            `db:"sale_price"`
	Rating      float32        `db:"rating"`
	Feedbacks   int            `db:"feedbacks"`
	Colors      sql.NullString `db:"colors"`
	Quantity    int            `db:"quantity"`
	CreatedAt   sql.NullTime   `db:"created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
}

func NewProduct(entity core.Product) Product {
	return Product{
		NmId:        entity.NmId,
		Name:        entity.Name,
		Brand:       entity.Brand,
		BrandId:     entity.BrandId,
		SiteBrandId: entity.SiteBrandId,
		SupplierId:  entity.SupplierId,
		Sale:        entity.Sale,
		Price:       entity.Price,
		SalePrice:   entity.SalePrice,
		Rating:      entity.Rating,
		Feedbacks:   entity.Feedbacks,
		Colors: sql.NullString{
			String: entity.Colors,
			Valid:  entity.Colors != "",
		},
		Quantity:  entity.Quantity,
		CreatedAt: sql.NullTime{Time: entity.CreatedAt, Valid: true},
		UpdatedAt: sql.NullTime{Time: entity.UpdatedAt, Valid: true},
	}
}

func (model Product) TableName() string {
	return "product"
}

func (model Product) ToEntity() core.Product {
	return core.Product{
		NmId:        model.NmId,
		Name:        model.Name,
		Brand:       model.Brand,
		BrandId:     model.BrandId,
		SiteBrandId: model.SiteBrandId,
		SupplierId:  model.SupplierId,
		Sale:        model.Sale,
		Price:       model.Price,
		SalePrice:   model.SalePrice,
		Rating:      model.Rating,
		Feedbacks:   model.Feedbacks,
		Colors:      model.Colors.String,
		Quantity:    model.Quantity,
		CreatedAt:   model.CreatedAt.Time,
		UpdatedAt:   model.UpdatedAt.Time,
	}
}
