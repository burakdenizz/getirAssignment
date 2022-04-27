package models

type GetRequest struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Gte       string `json:"gte"`
	Lte       string `json:"lte"`
}
