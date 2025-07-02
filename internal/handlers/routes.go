package api

import (
	"github.com/gorilla/mux"
	_ "github.com/swaggo/http-swagger"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "green/docs"
	"green/internal/service"
	"green/pkg/logging"
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
	router.Use(CORS, RecoverAllPanic)
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	//Auth
	authGr := router.PathPrefix("/auth").Subrouter()
	authGr.Use(h.IdentityCheckMiddleware)
	authGr.HandleFunc("/login", h.loginHandler).Methods(http.MethodPost, http.MethodOptions)
	authGr.HandleFunc("/change-password", h.changePassword).Methods(http.MethodPut, http.MethodOptions)
	authGr.HandleFunc("/refresh-token", h.refreshToken).Methods(http.MethodPost, http.MethodOptions)
	authGr.HandleFunc("/log-out", h.logoutHandler).Methods(http.MethodPut, http.MethodOptions)
	return router
}
