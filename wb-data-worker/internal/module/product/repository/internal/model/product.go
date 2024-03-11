package model

import (
	"database/sql"
	"errors"
	"time"
	"wb-data-service-golang/wb-data-worker/internal/module/product/core"
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

func NewProduct(entity core.WbProductDetail) (Product, error) {

	data := entity["data"].(map[string]interface{})
	products := data["products"].([]interface{})
	if len(products) != 1 {
		return Product{}, errors.New("WbProductDetail contains an incorrect number of products, there should be one")
	}
	product := products[0].(map[string]interface{})
	color := ""
	colors := product["colors"].([]interface{})
	if len(colors) > 1 {
		return Product{}, errors.New("product contains an incorrect number of colors, there should be one or zero")
	}
	if len(colors) == 1 {
		color = colors[0].(map[string]interface{})["name"].(string)
	}
	quantity := 0
	sizes := product["sizes"].([]interface{})
	for _, value := range sizes {
		size := value.(map[string]interface{})
		stocks := size["stocks"].([]interface{})
		for _, stock := range stocks {
			quantity += int(stock.(map[string]interface{})["dty"].(float64))
		}
	}

	return Product{
		NmId:        int(product["id"].(float64)),
		Name:        product["name"].(string),
		Brand:       product["brand"].(string),
		BrandId:     int(product["brandId"].(float64)),
		SiteBrandId: int(product["siteBrandId"].(float64)),
		SupplierId:  int(product["supplierId"].(float64)),
		Sale:        int(product["sale"].(float64)),
		Price:       int(product["priceU"].(float64)),
		SalePrice:   int(product["salePriceU"].(float64)),
		Rating:      float32(product["rating"].(float64)),
		Feedbacks:   int(product["feedbacks"].(float64)),
		Colors: sql.NullString{
			String: color,
			Valid:  color != "",
		},
		Quantity:  quantity,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}, nil
}

func (model Product) TableName() string {
	return "product"
}
