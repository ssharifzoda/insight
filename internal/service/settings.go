package service

import (
	"insight/internal/database"
	"insight/internal/models"
	"insight/pkg/consts"
	"insight/pkg/utils"
	"time"
)

type SettingService struct {
	db database.Setting
}

func NewSettingService(db database.Setting) *SettingService {
	return &SettingService{db: db}
}

func (s *SettingService) AddBrand(params *models.Brand) error {
	path := consts.GlobalLogoFilePath + params.Name + time.Now().Format(time.DateOnly)
	err := utils.SaveImageFromBase64(params.Logo, path)
	if err != nil {
		return err
	}
	params.Logo = path
	return s.db.AddBrand(params)
}

func (s *SettingService) GetAllBrands(page, limit int, search string) ([]*models.Brand, error) {
	offset := (limit * page) - limit
	return s.db.GetAllBrands(limit, offset, search)
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
func (s *SettingService) GetAllCategories(page, limit int, search string) (result []*models.Category, err error) {
	offset := (limit * page) - limit
	return s.db.GetAllCategories(limit, offset, search)
}
func (s *SettingService) EditCategory(category *models.Category) error {
	return s.db.EditCategory(category)
}
func (s *SettingService) DeleteCategory(categoryId int) error {
	return s.db.DeleteCategory(categoryId)
}

func (s *SettingService) AddNewCity(city *models.City) error {
	return s.db.AddNewCity(city)
}
func (s *SettingService) GetAllCities(page, limit int) (result []*models.City, err error) {
	offset := (limit * page) - limit
	return s.db.GetAllCities(limit, offset)
}
func (s *SettingService) EditCity(city *models.City) error {
	return s.db.EditCity(city)
}
func (s *SettingService) DeleteCity(cityId int) error {
	return s.db.DeleteCity(cityId)
}
func (s *SettingService) AddNewPromotion(promotion *models.Promotion) error {
	return s.db.AddNewPromotion(promotion)
}
func (s *SettingService) GetAllPromotions(page, limit int) (result []*models.Promotion, err error) {
	offset := (page * limit) - limit
	return s.db.GetAllPromotions(limit, offset)
}
func (s *SettingService) GetPromotionById(promotionId int) (result *models.PromotionInfo, err error) {
	return s.db.GetPromotionById(promotionId)
}

func (s *SettingService) EditPromotion(promotion *models.Promotion) error {
	return s.db.EditPromotion(promotion)
}
func (s *SettingService) DeletePromotion(promotionId int) error {
	return s.db.DeletePromotion(promotionId)
}

func (s *SettingService) AddNewRole(role *models.RoleInput) error {
	return s.db.AddNewRole(role)
}
func (s *SettingService) GetAllRoles(page, limit int) (result []*models.Role, err error) {
	offset := (page * limit) - limit
	return s.db.GetAllRoles(limit, offset)
}
func (s *SettingService) GetRoleById(roleId int) (result *models.RoleInfo, err error) {
	return s.db.GetRoleById(roleId)
}

func (s *SettingService) EditRole(role *models.RoleInput) error {
	return s.db.EditRole(role)
}
func (s *SettingService) DeleteRole(roleId int) error {
	return s.db.DeletePromotion(roleId)
}
func (s *SettingService) GetAllPermissions() ([]*models.Permission, error) {
	return s.db.GetAllPermissions()
}
