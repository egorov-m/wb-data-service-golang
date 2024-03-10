package model

import (
	"time"
	"wb-data-service-golang/wb-data-service/internal/module/price-history/core"
)

type PriceHistory struct {
	Id    int       `db:"id"`
	NmId  int       `db:"nm_id"`
	Dt    time.Time `db:"dt"`
	Price int       `db:"price"`
}

func NewPriceHistory(entity core.PriceHistory) PriceHistory {
	return PriceHistory{
		Id:    entity.Id,
		NmId:  entity.NmId,
		Dt:    entity.Dt,
		Price: entity.Price,
	}
}

func (model PriceHistory) TableName() string {
	return "price_history"
}

func (model PriceHistory) ToEntity() core.PriceHistory {
	return core.PriceHistory{
		Id:    model.Id,
		NmId:  model.NmId,
		Dt:    model.Dt,
		Price: model.Price,
	}
}
