package database

import (
	"gorm.io/gorm"
	"insight/internal/models"
)

type Authorization interface {
}

type Setting interface {
	AddBrand(params *models.Brand) error
	GetAllBrands(limit, offset int) ([]*models.Brand, error)
	EditBrand(brand *models.Brand) error
	DeleteBrand(brandId int) error
	AddNewCategory(category *models.Category) error
	GetAllCategories(limit, offset int) ([]*models.Category, error)
	EditCategory(category *models.Category) error
	DeleteCategory(categoryId int) error
}

type Database struct {
	Authorization
	Setting
}

func NewDatabase(conn *gorm.DB) *Database {
	return &Database{
		Setting: NewSettingsDb(conn),
	}
}
