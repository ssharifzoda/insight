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

type Users interface {
	AddNewUser(params *models.User) error
	UpdateUserParams(params *models.User) error
	GetAllUsers(page, limit int) ([]*models.User, error)
	GetUserById(userId int) (*models.User, error)
	GetUserByPhone(phone string) (*models.User, error)
	DeleteUser(userId int) error
}

type Shops interface {
	AddNewShop(params *models.Shop) error
	UpdateShopParams(params *models.Shop) error
	GetAllShops(page, limit int) ([]*models.Shop, error)
	GetShop(shopId int) (*models.Shop, error)
	DeleteShop(shopId int) error
}

type Service struct {
	Authorization
	Settings
	Users
	Shops
}

func NewService(db *database.Database) *Service {
	return &Service{
		Authorization: NewAuthService(db),
		Settings:      NewSettingService(db),
		Users:         NewUserService(db),
		Shops:         NewShopService(db),
	}
}
