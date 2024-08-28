package models

import (
	"gorm.io/gorm"
)

type LineItem struct {
	gorm.Model
	OrderID     uint    `json:"order_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	TotalPrice  float64 `json:"total_price"`
}
