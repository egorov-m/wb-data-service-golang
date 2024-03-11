package model

import (
	"time"
	"wb-data-service-golang/wb-data-worker/internal/module/price-history/core"
)

type PriceHistory struct {
	Id    int       `db:"id"`
	NmId  int       `db:"nm_id"`
	Dt    time.Time `db:"dt"`
	Price int       `db:"price"`
}

func NewPriceHistory(entity core.WbPricesHistory, nmId int) []PriceHistory {
	res := make([]PriceHistory, len(entity))

	for index, value := range entity {
		priceHistory := value.(map[string]interface{})
		price := priceHistory["price"].(map[string]interface{})
		res[index] = PriceHistory{
			NmId:  nmId,
			Dt:    time.Unix(int64(priceHistory["dt"].(float64)), 0),
			Price: int(price["RUB"].(float64)),
		}
	}

	return res
}

func (model PriceHistory) TableName() string {
	return "price_history"
}
