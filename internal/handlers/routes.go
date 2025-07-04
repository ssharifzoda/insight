package api

import (
	"github.com/gorilla/mux"
	_ "github.com/swaggo/http-swagger"
	httpSwagger "github.com/swaggo/http-swagger"
	"insight/internal/service"
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
	consoleGr := router.PathPrefix("/home")
	//Workers
	employeeGr := router.PathPrefix("/workers").Subrouter()
	employeeGr.HandleFunc("/list", h.getAllWorkers).Methods(http.MethodGet, http.MethodOptions)
	employeeGr.HandleFunc("/by-id", h.getWorkerById).Methods(http.MethodGet, http.MethodOptions)
	employeeGr.HandleFunc("/edit", h.editWorker).Methods(http.MethodPut, http.MethodOptions)
	employeeGr.HandleFunc("/rm", h.deleteWorker).Methods(http.MethodDelete, http.MethodOptions)
	//Shops
	shopGr := router.PathPrefix("/shops")
	//Supplier
	supplierGr := router.PathPrefix("/suppliers")
	//Application
	applicationGr := router.PathPrefix("/applications")
	//Product
	productsGr := router.PathPrefix("/products")
	//Notification
	notificationsGr := router.PathPrefix("/notifications")
	//Setting
	settingsGr := router.PathPrefix("/settings").Subrouter()
	settingsGr.HandleFunc("/brands", h.addBrand).Methods(http.MethodPost, http.MethodOptions)
	settingsGr.HandleFunc("/brands", h.getAllBrands).Methods(http.MethodGet, http.MethodOptions)
	settingsGr.HandleFunc("/brands", h.editBrand).Methods(http.MethodPut, http.MethodOptions)
	settingsGr.HandleFunc("/brands", h.deleteBrand).Methods(http.MethodDelete, http.MethodOptions)
	settingsGr.HandleFunc("/categories", h.addNewCategory).Methods(http.MethodPost, http.MethodOptions)
	settingsGr.HandleFunc("/categories", h.getAllCategories).Methods(http.MethodGet, http.MethodOptions)
	settingsGr.HandleFunc("/categories", h.editCategory).Methods(http.MethodPut, http.MethodOptions)
	settingsGr.HandleFunc("/categories", h.deleteCategory).Methods(http.MethodDelete, http.MethodOptions)
	settingsGr.HandleFunc("/cities", h.addNewCity).Methods(http.MethodPost, http.MethodOptions)
	settingsGr.HandleFunc("/cities", h.getAllCities).Methods(http.MethodGet, http.MethodOptions)
	settingsGr.HandleFunc("/cities", h.editCity).Methods(http.MethodPut, http.MethodOptions)
	settingsGr.HandleFunc("/cities", h.deleteCity).Methods(http.MethodDelete, http.MethodOptions)
	settingsGr.HandleFunc("/sale-points", h.addNewPoint).Methods(http.MethodPost, http.MethodOptions)
	settingsGr.HandleFunc("/sale-points", h.getAllPoints).Methods(http.MethodGet, http.MethodOptions)
	settingsGr.HandleFunc("/sale-points", h.editSalePoint).Methods(http.MethodPut, http.MethodOptions)
	settingsGr.HandleFunc("/sale-points", h.deletePoint).Methods(http.MethodDelete, http.MethodOptions)
	settingsGr.HandleFunc("/promotions", h.addNewPromotion).Methods(http.MethodPost, http.MethodOptions)
	settingsGr.HandleFunc("/promotions", h.getAllPromotions).Methods(http.MethodGet, http.MethodOptions)
	settingsGr.HandleFunc("/promotion", h.getPromotionById).Methods(http.MethodGet, http.MethodOptions)
	settingsGr.HandleFunc("/promotions", h.editPromotion).Methods(http.MethodPut, http.MethodOptions)
	settingsGr.HandleFunc("/promotions", h.deletePromotion).Methods(http.MethodDelete, http.MethodOptions)
	settingsGr.HandleFunc("/faq", h.addNewFaq).Methods(http.MethodPost, http.MethodOptions)
	settingsGr.HandleFunc("/faq", h.getFaqById).Methods(http.MethodGet, http.MethodOptions)
	settingsGr.HandleFunc("/faq-all", h.getAllFaq).Methods(http.MethodGet, http.MethodOptions)
	settingsGr.HandleFunc("/faq", h.editFaq).Methods(http.MethodPut, http.MethodOptions)
	settingsGr.HandleFunc("/faq", h.deleteFaq).Methods(http.MethodDelete, http.MethodOptions)
	settingsGr.HandleFunc("/roles", h.addNewRole).Methods(http.MethodPost, http.MethodOptions)
	settingsGr.HandleFunc("/roles", h.getAllRoles).Methods(http.MethodGet, http.MethodOptions)
	settingsGr.HandleFunc("/role", h.getRoleById).Methods(http.MethodGet, http.MethodOptions)
	settingsGr.HandleFunc("/roles", h.editRoleById).Methods(http.MethodPut, http.MethodOptions)
	settingsGr.HandleFunc("/roles", h.deleteRole).Methods(http.MethodDelete, http.MethodOptions)
	settingsGr.HandleFunc("/permissions", h.getAllPermissions).Methods(http.MethodGet, http.MethodOptions)

	return router
}
