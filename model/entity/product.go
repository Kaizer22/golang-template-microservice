package entity

// Product example
type Product struct {
	Id          int64  `json:"id" example:"1" format:"int64"`
	Name        string `json:"name" example:"Pepsi"`
	Description string `json:"description" example:"Carbonated sweet drink"`
	CategoryId  int64  `json:"category_id" example:"3"`
}

// ProductData example
type ProductData struct {
	Name        string `json:"name" example:"Pepsi"`
	Description string `json:"description" example:"Carbonated sweet drink"`
	CategoryId  int64  `json:"category_id" example:"3"`
}
