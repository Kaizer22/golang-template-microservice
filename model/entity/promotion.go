package entity

// Promotion example
type Promotion struct {
	Id           int64          `json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Prizes       []*Prize       `json:"prizes"`
	Participants []*Participant `json:"participants"`
}

// PromoInfo example
type PromoInfo struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// PromoShort example
type PromoShort struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
