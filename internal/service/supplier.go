package service

import (
	"insight/internal/database"
	"insight/internal/models"
)

type SupplierService struct {
	db database.Suppliers
}

func NewSupplierService(db database.Suppliers) *SupplierService {
	return &SupplierService{db: db}
}

func (s *SupplierService) AddNewSupplier(params *models.Supplier) error {
	return s.db.AddNewSupplier(params)
}

func (s *SupplierService) UpdateSupplierParams(params *models.Supplier) error {
	return s.db.UpdateSupplierParams(params)
}

func (s *SupplierService) GetAllSuppliers(page, limit int, search string) ([]*models.Supplier, error) {
	offset := (page * limit) - limit
	return s.db.GetAllSuppliers(limit, offset, search)
}
func (s *SupplierService) GetSupplier(shopId int) (*models.Supplier, error) {
	return s.db.GetSupplier(shopId)
}
func (s *SupplierService) DeleteSupplier(shopId int) error {
	return s.db.DeleteSupplier(shopId)
}
