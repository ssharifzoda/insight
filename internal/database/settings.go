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

func (s *SettingsDb) GetAllBrands(limit, offset int, search string) ([]*models.Brand, error) {
	var brands []*models.Brand
	tx := s.conn.Table("brands").Where("status", 1)
	if search != "" {
		tx = tx.Where("name like ?", "%"+search+"%").Find(&brands)
		return brands, tx.Error
	}
	tx = tx.Limit(limit).Offset(offset).Find(&brands)
	return brands, tx.Error
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
func (s *SettingsDb) GetAllCategories(limit, offset int, search string) (result []*models.Category, err error) {
	tx := s.conn.Table("categories").Where("status", 1)
	if search != "" {
		tx = tx.Where("name like ?", "%"+search+"%").Find(&result)
		return result, tx.Error
	}
	tx = tx.Limit(limit).Offset(offset).Find(&result)
	return result, tx.Error
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

func (s *SettingsDb) AddNewCity(city *models.City) error {
	return s.conn.Table("cities").Create(&city).Error
}
func (s *SettingsDb) GetAllCities(limit, offset int) (result []*models.City, err error) {
	err = s.conn.Table("cities").Where("status", 1).Limit(limit).Offset(offset).Find(&result).Error
	return result, err
}
func (s *SettingsDb) EditCity(city *models.City) error {
	return s.conn.Table("cities").Updates(&city).Error
}
func (s *SettingsDb) DeleteCity(cityId int) error {
	return s.conn.Table("cities").Where("id", cityId).UpdateColumns(map[string]interface{}{
		"status":     0,
		"deleted_at": time.Now(),
	}).Error
}

func (s *SettingsDb) AddNewPromotion(promotion *models.Promotion) error {
	return s.conn.Table("promotions").Create(&promotion).Error
}
func (s *SettingsDb) GetAllPromotions(limit, offset int) (result []*models.Promotion, err error) {
	err = s.conn.Table("promotions").Where("status", 1).Limit(limit).Offset(offset).Find(&result).Error
	return result, err
}
func (s *SettingsDb) GetPromotionById(promotionId int) (result *models.PromotionInfo, err error) {
	err = s.conn.Table("promotions").Where("id", promotionId).First(&result).Error
	return result, err
}

func (s *SettingsDb) EditPromotion(promotion *models.Promotion) error {
	return s.conn.Table("promotions").Updates(&promotion).Error
}
func (s *SettingsDb) DeletePromotion(promotionId int) error {
	return s.conn.Table("promotions").Where("id", promotionId).UpdateColumns(map[string]interface{}{
		"status":     0,
		"deleted_at": time.Now(),
	}).Error
}

func (s *SettingsDb) AddNewRole(role *models.RoleInput) error {
	var (
		roleParams     *models.Role
		rolePermission []*models.RolePermission
	)
	roleParams.Name = role.Name
	tx := s.conn.Begin()
	err := tx.Table("roles").Create(&roleParams).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, value := range role.Permissions {
		var param *models.RolePermission
		param.RoleId = roleParams.Id
		param.PermissionId = value
		rolePermission = append(rolePermission, param)
	}
	err = tx.Table("role_permission").Create(&rolePermission).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
func (s *SettingsDb) GetAllRoles(limit, offset int) (result []*models.Role, err error) {
	err = s.conn.Table("roles").Where("status", 1).Limit(limit).Offset(offset).Find(&result).Error
	return result, err
}
func (s *SettingsDb) GetRoleById(roleId int) (result *models.RoleInfo, err error) {
	err = s.conn.Table("roles").Where("status = 1 and id = ?", roleId).First(&result).Error
	if err != nil {
		return nil, err
	}
	err = s.conn.Table("role_permission").Select("permission_id").Find(&result.Permissions).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *SettingsDb) EditRole(role *models.RoleInput) error {
	return nil
}
func (s *SettingsDb) DeleteRole(roleId int) error {
	return s.conn.Table("roles").Where("id", roleId).UpdateColumns(map[string]interface{}{
		"status":     0,
		"deleted_at": time.Now(),
	}).Error
}

func (s *SettingsDb) GetAllPermissions() ([]*models.Permission, error) {
	var permissions []*models.Permission
	err := s.conn.Table("permissions").Where("status", 1).Find(&permissions).Error
	return permissions, err
}
