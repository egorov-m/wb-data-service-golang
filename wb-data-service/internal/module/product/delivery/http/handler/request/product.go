package request

import "wb-data-service-golang/wb-data-service/internal/module/product/core"

type LoadProductInBody struct {
	NmId int `json:"nm_id" validate:"required,number"`
}

func (model LoadProductInBody) ToEntity() core.Product {
	return core.Product{
		NmId: model.NmId,
	}
}
