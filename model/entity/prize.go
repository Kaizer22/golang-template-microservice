package entity

// Prize example
type Prize struct {
	Id          int64  `json:"id"`
	Description string `json:"description"`
}

type PrizeInfo struct {
	Description string `json:"description"`
}
