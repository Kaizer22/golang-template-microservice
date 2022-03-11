package entity

// PromoResult example
type PromoResult struct {
	Winner Participant `json:"winner"`
	Prize  Prize       `json:"prize"`
}
