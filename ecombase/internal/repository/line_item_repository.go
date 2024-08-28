package repository

import (
	"github.com/hauchu1196/ecombase/internal/models"
	"gorm.io/gorm"
)

type LineItemRepository struct {
	db *gorm.DB
}

func (r *LineItemRepository) UpdateWithTx(tx *gorm.DB, lineItem *models.LineItem) error {
	return tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(lineItem).Error
}
