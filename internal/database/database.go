package database

import (
	"gorm.io/gorm"
	"insight/internal/models"
)

type Authorization interface {
	GetUserPermission(userId int) ([]int, error)
	UpdateRefreshToken(userId int, accessToken, refreshToken string) error
	GetTokenByUserId(userId int) (*models.UserAuth, error)
	ChangeUserPassword(request *models.ChangePassword) error
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
	DeleteCity(cityId int) error
	EditCity(city *models.City) error
	GetAllCities(limit, offset int) (result []*models.City, err error)
	AddNewCity(city *models.City) error
	AddNewPromotion(promotion *models.Promotion) error
	GetAllPromotions(limit, offset int) (result []*models.Promotion, err error)
	GetPromotionById(promotionId int) (result *models.PromotionInfo, err error)
	EditPromotion(promotion *models.Promotion) error
	DeletePromotion(promotionId int) error
	AddNewRole(role *models.RoleInput) error
	GetAllRoles(limit, offset int) (result []*models.Role, err error)
	GetRoleById(roleId int) (result *models.RoleInfo, err error)
	EditRole(role *models.RoleInput) error
	DeleteRole(roleId int) error
	GetAllPermissions() ([]*models.Permission, error)
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
type Suppliers interface {
	AddNewSupplier(params *models.Supplier) error
	UpdateSupplierParams(params *models.Supplier) error
	GetAllSuppliers(limit, offset int) ([]*models.Supplier, error)
	GetSupplier(supplierId int) (*models.Supplier, error)
	DeleteSupplier(supplierId int) error
}

type Products interface {
	AddNewProduct(product *models.Product) error
	GetAllProducts(limit, offset int) ([]*models.Product, error)
	GetProductById(productId int) (*models.Product, error)
	EditProduct(product *models.Product) error
	DeleteProduct(productId int) error
}

type Database struct {
	Authorization
	Setting
	Users
	Shops
	Suppliers
	Products
}

func NewDatabase(conn *gorm.DB) *Database {
	return &Database{
		Setting:       NewSettingsDb(conn),
		Users:         NewUserDb(conn),
		Shops:         NewShopDb(conn),
		Suppliers:     NewSupplierDb(conn),
		Authorization: NewAuthDb(conn),
		Products:      NewProductDb(conn),
	}
}
