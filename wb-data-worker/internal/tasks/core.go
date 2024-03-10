package tasks

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
)

type (
	WbTasks interface {
		LoadProduct(context.Context, *asynq.Task) error
		LoadPriceHistory(context.Context, *asynq.Task) error
	}
)

const (
	CardDetailBaseUrl = "https://card.wb.ru/cards/v1/detail?" +
		"appType=1&" +
		"curr=rub&" +
		"dest=-1257786&" +
		"spp=30&" +
		"nm="
)

func GetUrlPriceHistory(nmID int) string {
	url := fmt.Sprintf("https://basket-%02d.wbbasket.ru/vol%d/part%d/%d/info/price-history.json",
		getBasketID(nmID), nmID/100000, nmID/1000, nmID)

	return url
}

func getBasketID(nmID int) int {
	return nmID % 1000 / 100
}
