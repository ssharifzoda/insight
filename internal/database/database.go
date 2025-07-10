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

type Users interface {
	AddNewUser(params *models.User) error
	UpdateUserParams(params *models.User) error
	GetAllUsers(limit, offset int) ([]*models.User, error)
	GetUserById(userId int) (*models.User, error)
	GetUserByPhone(phone string) (*models.User, error)
	DeleteUser(userId int) error
}

type Shops interface {
	AddNewShop(params *models.Shop) error
	UpdateShopParams(params *models.Shop) error
	GetAllShops(limit, offset int) ([]*models.Shop, error)
	GetShop(shopId int) (*models.Shop, error)
	DeleteShop(shopId int) error
}

type Database struct {
	Authorization
	Setting
	Users
	Shops
}

func NewDatabase(conn *gorm.DB) *Database {
	return &Database{
		Setting: NewSettingsDb(conn),
		Users:   NewUserDb(conn),
		Shops:   NewShopDb(conn),
	}
}
