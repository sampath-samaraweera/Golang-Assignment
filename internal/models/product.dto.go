package models

type Product struct {
	PId   int    `json:"p_id"`
	Name string `json:"name"`
	PType string `json:"p_type"`
	Price int `json:"price"`
	Quantity int `json:"quantity"`
}
type UpdateProduct struct {
	Name string `json:"name"`
	PType string `json:"p_type"`
	Price int `json:"price"`
	Quantity int `json:"quantity"`
}