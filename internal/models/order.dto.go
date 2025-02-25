package models

import "time"

type Order struct {
	OrderID int `json:"order_id"`
	UserID int `json:"user_id"`
	PID int `json:"p_id"`
	OrderedQuantity int `json:"ordered_quantity"`
	TotalPrice int `json:"total_price"`
	OrderDate time.Time `json:"order_date"`
}

type CreateOrder struct {
	UserID int `json:"user_id"`
	PID int `json:"p_id"`
	OrderedQuantity int `json:"ordered_quantity"`
}
