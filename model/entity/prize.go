package entity

// Prize example
type Prize struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

type PrizeInfo struct {
	Description string `json:"description"`
}
