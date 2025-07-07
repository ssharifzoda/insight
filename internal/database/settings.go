package database

import (
	"gorm.io/gorm"
	"insight/internal/models"
	"time"
)

type SettingsDb struct {
	conn *gorm.DB
}

func NewSettingsDb(connection *gorm.DB) *SettingsDb {
	return &SettingsDb{conn: connection}
}

func (s *SettingsDb) AddBrand(params *models.Brand) error {
	err := s.conn.Table("brands").Create(&params).Error
	return err
}

func (s *SettingsDb) GetAllBrands(limit, offset int) ([]*models.Brand, error) {
	var brands []*models.Brand
	err := s.conn.Table("brands").Where("status", 1).Limit(limit).Offset(offset).Find(&brands).Error
	return brands, err
}

func (s *SettingsDb) EditBrand(brand *models.Brand) error {
	return s.conn.Table("brands").Updates(&brand).Error
}

func (s *SettingsDb) DeleteBrand(brandId int) error {
	return s.conn.Table("brands").Where("id", brandId).UpdateColumns(map[string]interface{}{
		"status":     0,
		"deleted_at": time.Now(),
	}).Error
}

func (s *SettingsDb) AddNewCategory(category *models.Category) error {
	return s.conn.Table("categories").Create(&category).Error
}
func (s *SettingsDb) GetAllCategories(limit, offset int) (result []*models.Category, err error) {
	err = s.conn.Table("categories").Where("status", 1).Limit(limit).Offset(offset).Find(&result).Error
	return result, err
}
func (s *SettingsDb) EditCategory(category *models.Category) error {
	return s.conn.Table("categories").Updates(&category).Error
}
func (s *SettingsDb) DeleteCategory(categoryId int) error {
	return s.conn.Table("categories").Where("id", categoryId).UpdateColumns(map[string]interface{}{
		"status":     0,
		"deleted_at": time.Now(),
	}).Error
}
