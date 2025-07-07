package service

import (
	"insight/internal/database"
	"insight/internal/models"
)

type Authorization interface {
}

type Settings interface {
	AddBrand(params *models.Brand) error
	GetAllBrands(page, limit int) ([]*models.Brand, error)
	EditBrand(brand *models.Brand) error
	DeleteBrand(brandId int) error
	AddNewCategory(category *models.Category) error
	GetAllCategories(page, limit int) ([]*models.Category, error)
	EditCategory(category *models.Category) error
	DeleteCategory(categoryId int) error
}

type Service struct {
	Authorization
	Settings
}

func NewService(db *database.Database) *Service {
	return &Service{
		Authorization: NewAuthService(db),
		Settings:      NewSettingService(db),
	}
}
