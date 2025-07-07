package service

import (
	"insight/internal/database"
	"insight/internal/models"
)

type SettingService struct {
	db database.Setting
}

func NewSettingService(db database.Setting) *SettingService {
	return &SettingService{db: db}
}

func (s *SettingService) AddBrand(params *models.Brand) error {
	return s.db.AddBrand(params)
}

func (s *SettingService) GetAllBrands(page, limit int) ([]*models.Brand, error) {
	offset := (limit * page) - limit
	return s.db.GetAllBrands(limit, offset)
}

func (s *SettingService) EditBrand(brand *models.Brand) error {
	return s.db.EditBrand(brand)
}

func (s *SettingService) DeleteBrand(brandId int) error {
	return s.db.DeleteBrand(brandId)
}

func (s *SettingService) AddNewCategory(category *models.Category) error {
	return s.db.AddNewCategory(category)
}
func (s *SettingService) GetAllCategories(page, limit int) ([]*models.Category, error) {
	offset := (limit * page) - limit
	return s.db.GetAllCategories(limit, offset)
}
func (s *SettingService) EditCategory(category *models.Category) error {
	return s.db.EditCategory(category)
}
func (s *SettingService) DeleteCategory(categoryId int) error {
	return s.db.DeleteCategory(categoryId)
}
