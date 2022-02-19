package model

type MongoRequest struct {
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	MinCount  int `json:"minCount"`
	MaxCount  int `json:"maxCount"`
}