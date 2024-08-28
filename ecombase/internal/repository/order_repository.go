// internal/repository/order_repository.go
package repository

import (
	"github.com/hauchu1196/ecombase/internal/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func (r *OrderRepository) GetDB() *gorm.DB {
	return r.db
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) GetByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("LineItems").First(&order, id).Error
	return &order, err
}

func (r *OrderRepository) List(page, pageSize int) ([]models.Order, int, error) {
	var orders []models.Order
	var total int64

	offset := (page - 1) * pageSize

	err := r.db.Model(&models.Order{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("LineItems").
		Select("orders.id").
		Offset(offset).
		Limit(pageSize).
        Find(&orders).
        Error
	if err != nil {
		return nil, 0, err
	}

	return orders, int(total), nil
}

func (r *OrderRepository) Update(order *models.Order) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(order).Error
}

func (r *OrderRepository) Delete(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("order_id = ?", id).Delete(&models.LineItem{}).Error; err != nil {
			return err
		}
		return tx.Delete(&models.Order{}, id).Error
	})
}

func (r *OrderRepository) UpdateWithTx(tx *gorm.DB, order *models.Order) error {
	return tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(order).Error
}

// 2 update trong 1 repository