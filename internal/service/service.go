package service

import (
	"firebase.google.com/go/v4/messaging"
	"insight/internal/database"
	"insight/internal/models"
)

type Authorization interface {
	GetUserPermission(userId int) ([]int, error)
	UpdateAuthParams(userId int, sessionId string) error
	GetAuthParamsByUserId(userId int) (*models.UserAuth, error)
	ChangeUserPassword(request *models.ChangePassword) error
}

type Settings interface {
	AddBrand(params *models.Brand) (*models.Brand, error)
	GetAllBrands(page, limit int, search string) ([]*models.Brand, int, error)
	EditBrand(brand *models.Brand) error
	DeleteBrand(brandId int) error
	AddNewCategory(category *models.Category) (*models.Category, error)
	GetAllCategories(limit, offset int, search string) (result []*models.Category, total int, err error)
	EditCategory(category *models.Category) error
	DeleteCategory(categoryId int) error
	DeleteCity(cityId int) error
	EditCity(city *models.City) error
	GetAllCities(page, limit int) (result []*models.City, total int, err error)
	AddNewCity(city *models.City) (*models.City, error)
	AddNewPromotion(promotion *models.Promotion) (*models.Promotion, error)
	GetAllPromotions(page, limit int) (result []*models.Promotion, total int, err error)
	GetPromotionById(promotionId int) (result *models.PromotionInfo, err error)
	EditPromotion(promotion *models.Promotion) error
	DeletePromotion(promotionId int) error
	AddNewRole(role *models.RoleInput) (*models.Role, error)
	GetAllRoles(limit, offset int) (result []*models.Role, total int, err error)
	GetRoleById(roleId int) (result *models.RoleInfo, err error)
	EditRole(role *models.RoleInput) error
	DeleteRole(roleId int) error
	GetAllPermissions() ([]*models.Permission, error)
}

type Users interface {
	AddNewUser(params *models.User) (*models.User, error)
	UpdateUserParams(params *models.User) error
	GetAllUsers(page, limit int) ([]*models.User, int, error)
	GetUserById(userId int) (*models.UserInfo, error)
	GetUserByPhone(phone string) (*models.User, error)
	DeleteUser(userId int) error
}

type Shops interface {
	AddNewShop(params *models.Shop) (*models.Shop, error)
	UpdateShopParams(params *models.Shop) error
	GetAllShops(page, limit int, search string) ([]*models.Shop, int, error)
	GetShop(shopId int) (*models.Shop, error)
	DeleteShop(shopId int) error
}

type Suppliers interface {
	AddNewSupplier(params *models.Supplier) (*models.Supplier, error)
	UpdateSupplierParams(params *models.Supplier) error
	GetAllSuppliers(page, limit int, search string) ([]*models.Supplier, int, error)
	GetSupplier(supplierId int) (*models.Supplier, error)
	DeleteSupplier(supplierId int) error
}

type Products interface {
	AddNewProduct(product *models.Product) (*models.Product, error)
	GetAllProducts(limit, offset int, filter *models.ProductFilter) (result []*models.Product, total int, err error)
	GetProductById(productId int) (*models.Product, error)
	EditProduct(product *models.Product) error
	DeleteProduct(productId int) error
}

type Orders interface {
	AddNewOrder(order *models.OrderInput) error
	GetAllOrders(filter *models.OrderFilter) ([]*models.Order, int, error)
	GetOrderById(orderId int) (order *models.OrderInfo, err error)
	EditOrder(order *models.OrderInput) error
}

type Notifications interface {
	CreateNewNotification(message *models.NotificationInput) error
	GetAllNotifications(page, limit int) ([]*models.Notification, int, error)
	GetNotificationById(notificationId int) (*models.NotificationInfo, error)
	DeleteNotification(notificationId int) error
}

type Service struct {
	Authorization
	Settings
	Users
	Shops
	Suppliers
	Products
	Orders
	Notifications
}

func NewService(db *database.Database, client *messaging.Client) *Service {
	return &Service{
		Authorization: NewAuthService(db),
		Settings:      NewSettingService(db),
		Users:         NewUserService(db),
		Shops:         NewShopService(db),
		Suppliers:     NewSupplierService(db),
		Products:      NewProductService(db),
		Orders:        NewOrderService(db),
		Notifications: NewNotificationService(db, client),
	}
}
