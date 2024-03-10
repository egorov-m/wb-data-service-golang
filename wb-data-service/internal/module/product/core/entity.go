package core

import "time"

type Product struct {
	NmId        int       `json:"nm_id" example:"139760729"`
	Name        string    `json:"name" example:"iPhone 14 Pro Max 1TB"`
	Brand       string    `json:"brand" example:"Apple"`
	BrandId     int       `json:"brand_id" example:"6049"`
	SiteBrandId int       `json:"site_brand_id" example:"16049"`
	SupplierId  int       `json:"supplier_id" example:"887491"`
	Sale        int       `json:"sale" example:"6"`
	Price       int       `json:"price" example:"20199000"`
	SalePrice   int       `json:"sale_price" example:"18987000"`
	Rating      float32   `json:"rating" example:"5.0"`
	Feedbacks   int       `json:"feedbacks" example:"31"`
	Colors      string    `json:"colors" example:"фиолетовый"`
	Quantity    int       `json:"quantity" example:"0"`
	CreatedAt   time.Time `json:"created_at" example:"2024-02-11 18:57:11.811169+00"`
	UpdatedAt   time.Time `json:"updated_at" example:"2024-02-11 18:57:11.811169+00"`
}

type ProductTask struct {
	Id   string `json:"task_id"`
	Type string `json:"type"`
}

func (entity Product) IsEmpty() bool {
	return entity == Product{}
}
