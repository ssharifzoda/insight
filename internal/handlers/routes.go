package api

import (
	"github.com/gorilla/mux"
	_ "github.com/swaggo/http-swagger"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "insight/docs"
	"insight/internal/service"
	"insight/pkg/logging"
	"net/http"
)

type Handler struct {
	service *service.Service
	logger  logging.Logger
}

func NewHandler(s *service.Service, logger logging.Logger) *Handler {
	return &Handler{service: s, logger: logger}
}

func (h *Handler) InitRoutes() *mux.Router {
	main := mux.NewRouter()
	router := main.PathPrefix("/api/v1").Subrouter()
	//router.Use(CORS, RecoverAllPanic)
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	//Auth
	authGr := router.PathPrefix("/auth").Subrouter()
	//authGr.Use(h.IdentityCheckMiddleware)
	authGr.HandleFunc("/login", h.loginHandler).Methods(http.MethodPost, http.MethodOptions)
	authGr.HandleFunc("/change-password", h.changePassword).Methods(http.MethodPut, http.MethodOptions)
	authGr.HandleFunc("/refresh-token", h.refreshToken).Methods(http.MethodPost, http.MethodOptions)
	authGr.HandleFunc("/log-out", h.logoutHandler).Methods(http.MethodPut, http.MethodOptions)
	//Console
	//consoleGr := router.PathPrefix("/home")
	//Workers
	usersGr := router.PathPrefix("/users").Subrouter()
	usersGr.HandleFunc("/new", h.addNewUser)
	usersGr.HandleFunc("/list", h.getAllUsers).
		Queries("page", "{page}").Queries("limit", "{limit}").Methods(http.MethodGet, http.MethodOptions)
	usersGr.HandleFunc("/by-id", h.getUserById).Queries("id", "{id}").Methods(http.MethodGet, http.MethodOptions)
	usersGr.HandleFunc("/me", h.getMe).Methods(http.MethodGet, http.MethodOptions)
	usersGr.HandleFunc("/edit", h.editUser).Methods(http.MethodPut, http.MethodOptions)
	usersGr.HandleFunc("/rm", h.deleteUser).Methods(http.MethodDelete, http.MethodOptions)
	//Shops
	shopGr := router.PathPrefix("/shops").Subrouter()
	shopGr.HandleFunc("/new", h.addNewShop).Methods(http.MethodPost, http.MethodOptions)
	shopGr.HandleFunc("/list", h.getAllShops).
		Queries("page", "{page}").Queries("limit", "{limit}").Methods(http.MethodGet, http.MethodOptions)
	shopGr.HandleFunc("/by-id", h.getShop).Methods(http.MethodGet, http.MethodOptions)
	shopGr.HandleFunc("/rm", h.deleteShop).Methods(http.MethodDelete, http.MethodOptions)
	shopGr.HandleFunc("/edit", h.editShop).Methods(http.MethodPut, http.MethodOptions)
	//Supplier
	supplierGr := router.PathPrefix("/suppliers").Subrouter()
	supplierGr.HandleFunc("/new", h.addNewSupplier).Methods(http.MethodPost, http.MethodOptions)
	supplierGr.HandleFunc("/list", h.getAllSuppliers).
		Queries("page", "{page}").Queries("limit", "{limit}").Methods(http.MethodGet, http.MethodOptions)
	supplierGr.HandleFunc("/by-id", h.getSupplier).Methods(http.MethodGet, http.MethodOptions)
	supplierGr.HandleFunc("/rm", h.deleteSupplier).Methods(http.MethodDelete, http.MethodOptions)
	supplierGr.HandleFunc("/edit", h.editSupplier).Methods(http.MethodPut, http.MethodOptions)
	//Application
	ordersGr := router.PathPrefix("/orders").Subrouter()
	ordersGr.HandleFunc("/new", h.addOrder).Methods(http.MethodPost, http.MethodOptions)
	ordersGr.HandleFunc("/list", h.getAllOrders).Methods(http.MethodGet, http.MethodOptions)
	ordersGr.HandleFunc("/by-id", h.getOrderById).Queries("order_id", "{order_id}").Methods(http.MethodGet, http.MethodOptions)
	ordersGr.HandleFunc("/edit", h.editOrder).Methods(http.MethodPut, http.MethodOptions)
	//ordersGr.HandleFunc("/list-pdf", h.downloadOrdersPdf).Methods(http.MethodGet, http.MethodOptions)
	//Product
	productsGr := router.PathPrefix("/products").Subrouter()
	productsGr.HandleFunc("/new", h.addNewProduct)
	productsGr.HandleFunc("/list", h.getAllProducts).
		Queries("page", "{page}").Queries("limit", "{limit}").Methods(http.MethodGet, http.MethodOptions)
	productsGr.HandleFunc("/by-id", h.getProduct).Methods(http.MethodGet, http.MethodOptions)
	productsGr.HandleFunc("/edit", h.editProduct).Methods(http.MethodPut, http.MethodOptions)
	productsGr.HandleFunc("/rm", h.deleteProduct).Methods(http.MethodDelete, http.MethodOptions)

	//Notification
	notificationsGr := router.PathPrefix("/notifications").Subrouter()
	notificationsGr.HandleFunc("/new", h.addNewNotification)
	//notificationsGr.HandleFunc("/list", h.getAllNotifications).Methods(http.MethodGet, http.MethodOptions)
	//notificationsGr.HandleFunc("/by-id", h.getNotification).Methods(http.MethodGet, http.MethodOptions)
	//notificationsGr.HandleFunc("/edit", h.editNotification).Methods(http.MethodPut, http.MethodOptions) //todo: здесь же просмотр либо отдельный роут
	//notificationsGr.HandleFunc("/rm", h.deleteNotification).Methods(http.MethodDelete, http.MethodOptions)
	//Setting
	settingsGr := router.PathPrefix("/settings").Subrouter()
	settingsGr.HandleFunc("/brands", h.addBrand).Methods(http.MethodPost, http.MethodOptions)
	settingsGr.HandleFunc("/brands", h.getAllBrands).
		Queries("page", "{page}").Queries("limit", "{limit}").Methods(http.MethodGet, http.MethodOptions)
	settingsGr.HandleFunc("/brands", h.editBrand).Methods(http.MethodPut, http.MethodOptions)
	settingsGr.HandleFunc("/brands", h.deleteBrand).Methods(http.MethodDelete, http.MethodOptions)
	settingsGr.HandleFunc("/categories", h.addNewCategory).Methods(http.MethodPost, http.MethodOptions)
	settingsGr.HandleFunc("/categories", h.getAllCategories).Methods(http.MethodGet, http.MethodOptions)
	settingsGr.HandleFunc("/categories", h.editCategory).Methods(http.MethodPut, http.MethodOptions)
	settingsGr.HandleFunc("/categories", h.deleteCategory).Methods(http.MethodDelete, http.MethodOptions)
	settingsGr.HandleFunc("/cities", h.addNewCity).Methods(http.MethodPost, http.MethodOptions)
	settingsGr.HandleFunc("/cities", h.getAllCities).
		Queries("page", "{page}").Queries("limit", "{limit}").Methods(http.MethodGet, http.MethodOptions)
	settingsGr.HandleFunc("/cities", h.editCity).Methods(http.MethodPut, http.MethodOptions)
	settingsGr.HandleFunc("/cities", h.deleteCity).Methods(http.MethodDelete, http.MethodOptions)
	settingsGr.HandleFunc("/promotions", h.addNewPromotion).Methods(http.MethodPost, http.MethodOptions)
	settingsGr.HandleFunc("/promotions", h.getAllPromotions).
		Queries("page", "{page}").Queries("limit", "{limit}").Methods(http.MethodGet, http.MethodOptions)
	settingsGr.HandleFunc("/promotion", h.getPromotionById).Methods(http.MethodGet, http.MethodOptions)
	settingsGr.HandleFunc("/promotions", h.editPromotion).Methods(http.MethodPut, http.MethodOptions)
	settingsGr.HandleFunc("/promotions", h.deletePromotion).Methods(http.MethodDelete, http.MethodOptions)
	//settingsGr.HandleFunc("/faq", h.addNewFaq).Methods(http.MethodPost, http.MethodOptions)
	//settingsGr.HandleFunc("/faq", h.getFaqById).Methods(http.MethodGet, http.MethodOptions)
	//settingsGr.HandleFunc("/faq-all", h.getAllFaq).Methods(http.MethodGet, http.MethodOptions)
	//settingsGr.HandleFunc("/faq", h.editFaq).Methods(http.MethodPut, http.MethodOptions)
	//settingsGr.HandleFunc("/faq", h.deleteFaq).Methods(http.MethodDelete, http.MethodOptions)
	settingsGr.HandleFunc("/roles", h.addNewRole).Methods(http.MethodPost, http.MethodOptions)
	settingsGr.HandleFunc("/roles", h.getAllRoles).Methods(http.MethodGet, http.MethodOptions)
	settingsGr.HandleFunc("/role", h.getRoleById).Methods(http.MethodGet, http.MethodOptions)
	settingsGr.HandleFunc("/roles", h.editRoleById).Methods(http.MethodPut, http.MethodOptions)
	settingsGr.HandleFunc("/roles", h.deleteRole).Methods(http.MethodDelete, http.MethodOptions)
	settingsGr.HandleFunc("/permissions", h.getAllPermissions).Methods(http.MethodGet, http.MethodOptions)

	return router
}
