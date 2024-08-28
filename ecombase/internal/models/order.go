package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerName string     `json:"customer_name"`
	TotalAmount  float64    `json:"total_amount"`
	Status       string     `json:"status"`
	LineItems    []LineItem `json:"line_items" gorm:"foreignKey:OrderID"`
}
