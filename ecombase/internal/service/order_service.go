// internal/service/order_service.go
package service

import (
	"github.com/hauchu1196/ecombase/internal/models"
	"github.com/hauchu1196/ecombase/internal/repository"
	"gorm.io/gorm"
)

type OrderServiceInterface interface {
	CreateOrder(order *models.Order) error
	GetOrder(id uint) (*models.Order, error)
	ListOrders(page, pageSize int) ([]models.Order, int, error)
	UpdateOrder(order *models.Order) error
	DeleteOrder(id uint) error
}

type OrderService struct {
	repo     *repository.OrderRepository
	lineRepo *repository.LineItemRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	// Calculate total amount
	var totalAmount float64
	for i, item := range order.LineItems {
		item.TotalPrice = float64(item.Quantity) * item.UnitPrice
		totalAmount += item.TotalPrice
		order.LineItems[i] = item
	}
	order.TotalAmount = totalAmount

	return s.repo.Create(order)
}

func (s *OrderService) GetOrder(id uint) (*models.Order, error) {
	return s.repo.GetByID(id)
}

func (s *OrderService) ListOrders(page, pageSize int) ([]models.Order, int, error) {
	return s.repo.List(page, pageSize)
}
func (s *OrderService) UpdateOrder(order *models.Order) error {
	// Recalculate total amount
	var totalAmount float64
	for i, item := range order.LineItems {
		item.TotalPrice = float64(item.Quantity) * item.UnitPrice
		totalAmount += item.TotalPrice
		order.LineItems[i] = item
	}
	order.TotalAmount = totalAmount

	tx := s.repo.GetDB().Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := s.repo.UpdateWithTx(tx, order); err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range order.LineItems {
		if err := s.lineRepo.UpdateWithTx(tx, &item); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (s *OrderService) DeleteOrder(id uint) error {
	return s.repo.Delete(id)
}
