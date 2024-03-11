package utils

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

func getBasketID(productID int) int {
	shortID := productID / 100000
	var basket int

	switch {
	case 0 <= shortID && shortID <= 143:
		basket = 1
	case 144 <= shortID && shortID <= 287:
		basket = 2
	case 288 <= shortID && shortID <= 431:
		basket = 3
	case 432 <= shortID && shortID <= 719:
		basket = 4
	case 720 <= shortID && shortID <= 1007:
		basket = 5
	case 1008 <= shortID && shortID <= 1061:
		basket = 6
	case 1062 <= shortID && shortID <= 1115:
		basket = 7
	case 1116 <= shortID && shortID <= 1169:
		basket = 8
	case 1170 <= shortID && shortID <= 1313:
		basket = 9
	case 1314 <= shortID && shortID <= 1601:
		basket = 10
	case 1602 <= shortID && shortID <= 1655:
		basket = 11
	case 1656 <= shortID && shortID <= 1919:
		basket = 12
	default:
		basket = 13
	}

	return basket
}
