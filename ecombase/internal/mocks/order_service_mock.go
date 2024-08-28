// internal/mocks/order_service_mock.go
package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/hauchu1196/ecombase/internal/models"
)

type MockOrderService struct {
	mock.Mock
}

func (m *MockOrderService) CreateOrder(order *models.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderService) GetOrder(id uint) (*models.Order, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Order), args.Error(1)
}

func (m *MockOrderService) ListOrders(page, pageSize int) ([]models.Order, int, error) {
    args := m.Called(page, pageSize)
    return args.Get(0).([]models.Order), args.Int(1), args.Error(2)
}

func (m *MockOrderService) UpdateOrder(order *models.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderService) DeleteOrder(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}