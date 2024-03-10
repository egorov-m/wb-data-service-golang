package core

import "time"

type PriceHistory struct {
	Id    int       `json:"id" example:"1"`
	NmId  int       `json:"nm_id" example:"139760729"`
	Dt    time.Time `json:"dt" example:"2024-02-11 18:57:11.811169+00"`
	Price int       `json:"price" example:"20199000"`
}

type PriceHistoryTask struct {
	Id   string `json:"task_id"`
	Type string `json:"type"`
}
