package service

import (
	"insight/internal/database"
	"insight/internal/models"
	"insight/pkg/consts"
	"insight/pkg/utils"
)

type SupplierService struct {
	db database.Suppliers
}

func NewSupplierService(db database.Suppliers) *SupplierService {
	return &SupplierService{db: db}
}

func (s *SupplierService) AddNewSupplier(params *models.Supplier) (*models.Supplier, error) {
	if params.CompanyLogo != "" {
		filename := utils.FilePathGen("")
		path := consts.GlobalFilePath
		err := utils.SaveImageFromBase64(params.CompanyLogo, path+filename)
		if err != nil {
			return nil, err
		}
		params.CompanyLogo = filename
		return s.db.AddNewSupplier(params)
	}
	return s.db.AddNewSupplier(params)
}

func (s *SupplierService) UpdateSupplierParams(params *models.Supplier) error {
	return s.db.UpdateSupplierParams(params)
}

func (s *SupplierService) GetAllSuppliers(page, limit int, search string) ([]*models.Supplier, int, error) {
	offset := (page * limit) - limit
	return s.db.GetAllSuppliers(limit, offset, search)
}
func (s *SupplierService) GetSupplier(supplierId int) (*models.Supplier, error) {
	shop, err := s.db.GetSupplier(supplierId)
	if err != nil {
		return nil, err
	}
	if shop.CompanyLogo != "" {
		file, err := utils.ConvertImageToBase64(consts.GlobalFilePath, shop.CompanyLogo)
		if err != nil {
			return nil, err
		}
		shop.CompanyLogo = file
		return shop, nil
	}
	return s.db.GetSupplier(supplierId)
}
func (s *SupplierService) DeleteSupplier(shopId int) error {
	return s.db.DeleteSupplier(shopId)
}
