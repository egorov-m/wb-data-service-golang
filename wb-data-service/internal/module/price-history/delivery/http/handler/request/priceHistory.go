package request

import "wb-data-service-golang/wb-data-service/internal/module/price-history/core"

type LoadPriceHistoryInBody struct {
	NmId int `json:"nm_id" validate:"required,number"`
}

func (model LoadPriceHistoryInBody) ToEntity() core.PriceHistory {
	return core.PriceHistory{
		NmId: model.NmId,
	}
}
