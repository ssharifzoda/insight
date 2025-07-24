package database

import (
	"gorm.io/gorm"
	"insight/internal/models"
	"time"
)

type SupplierDb struct {
	conn *gorm.DB
}

func NewSupplierDb(conn *gorm.DB) *SupplierDb {
	return &SupplierDb{conn: conn}
}

func (s *SupplierDb) AddNewSupplier(params *models.Supplier) (*models.Supplier, error) {
	return params, s.conn.Table("suppliers").Create(&params).Error
}

func (s *SupplierDb) UpdateSupplierParams(params *models.Supplier) error {
	return s.conn.Table("suppliers").Where("id").Updates(*params).Error
}

func (s *SupplierDb) GetAllSuppliers(limit, offset int, search string) ([]*models.Supplier, error) {
	var shops []*models.Supplier
	tx := s.conn.Table("suppliers").Where("status", 1)
	if search != "" {
		tx = tx.Where("name LIKE ?", "%"+search+"%").Find(&shops)
		return shops, tx.Error
	}
	tx = tx.Limit(limit).Offset(offset).Order("order_number DESC").Find(&shops)
	return shops, tx.Error
}
func (s *SupplierDb) GetSupplier(supplierId int) (*models.Supplier, error) {
	var shop *models.Supplier
	err := s.conn.Table("suppliers").Where("active = 1 and id = ?", supplierId).First(&shop).Error
	return shop, err
}

func (s *SupplierDb) DeleteSupplier(supplierId int) error {
	return s.conn.Table("suppliers").Where("id", supplierId).UpdateColumns(map[string]interface{}{
		"active":     0,
		"deleted_at": time.Now(),
	}).Error
}
