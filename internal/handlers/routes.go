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
	consoleGr := router.PathPrefix("/console")
	//Workers
	employeeGr := router.PathPrefix("/employees")
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
	settingsGr := router.PathPrefix("/settings")

	return router
}
