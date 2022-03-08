package entity

type Product struct {
	Id          int64  `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  int64  `json:"category_id"`
}

type AddProduct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  int64  `json:"category_id"`
}
