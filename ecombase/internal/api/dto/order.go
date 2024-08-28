// internal/api/dto/order.go
package dto

import (
	"time"

	"github.com/hauchu1196/ecombase/internal/models"
)

type OrderResponse struct {
	ID           uint               `json:"id"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	CustomerName string             `json:"customer_name"`
	TotalAmount  float64            `json:"total_amount"`
	Status       string             `json:"status"`
	LineItems    []LineItemResponse `json:"line_items"`
}

type LineItemResponse struct {
	ID          uint    `json:"id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	TotalPrice  float64 `json:"total_price"`
}

type OrderListResponse struct {
	Orders []OrderSummary `json:"orders"`
	Total  int            `json:"total"`
}

type OrderSummary struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	CustomerName string    `json:"customer_name"`
	TotalAmount  float64   `json:"total_amount"`
	Status       string    `json:"status"`
	ItemCount    int       `json:"item_count"`
}

func ToOrderListResponse(orders []models.Order) OrderListResponse {
	summaries := make([]OrderSummary, len(orders))
	for i, order := range orders {
		summaries[i] = OrderSummary{
			ID:           order.ID,
			CreatedAt:    order.CreatedAt,
			CustomerName: order.CustomerName,
			TotalAmount:  order.TotalAmount,
			Status:       order.Status,
			ItemCount:    len(order.LineItems),
		}
	}
	return OrderListResponse{
		Orders: summaries,
		Total:  len(summaries),
	}
}

func ToOrderResponse(order *models.Order) OrderResponse {
	lineItems := make([]LineItemResponse, len(order.LineItems))
	for i, item := range order.LineItems {
		lineItems[i] = LineItemResponse{
			ID:          item.ID,
			ProductName: item.ProductName,
			Quantity:    item.Quantity,
			UnitPrice:   item.UnitPrice,
			TotalPrice:  item.TotalPrice,
		}
	}

	return OrderResponse{
		ID:           order.ID,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
		CustomerName: order.CustomerName,
		TotalAmount:  order.TotalAmount,
		Status:       order.Status,
		LineItems:    lineItems,
	}
}
