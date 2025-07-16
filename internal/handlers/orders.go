package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"insight/internal/models"
	"insight/pkg/consts"
	"insight/pkg/utils"
	"net/http"
	"strconv"
)

// @Summary addOrder
// @Security ApiKeyAuth
// @Tags Orders
// @Description Добавление нового заказа
// @ID addOrder
// @Accept json
// @Produce json
// @Param params body models.OrderInput true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /orders/new [post]
func (h *Handler) addOrder(w http.ResponseWriter, r *http.Request) {
	var params *models.OrderInput
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	err = h.service.Orders.AddNewOrder(params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}

	utils.Response(w, consts.Success)
}

// @Summary getAllOrders
// @Security ApiKeyAuth
// @Tags Orders
// @Description Просмотр заказов
// @ID getAllOrders
// @Accept json
// @Produce json
// @Param shop_id query string false "Введите данные"
// @Param supplier_id query string false "Введите данные"
// @Param status query string false "Введите данные"
// @Param date_from query string false "Введите данные"
// @Param date_to query string false "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /orders/list [get]
func (h *Handler) getAllOrders(w http.ResponseWriter, r *http.Request) {
	var filter *models.OrderFilter
	shopStr := r.URL.Query().Get("shop_id")
	supplierStr := r.URL.Query().Get("supplier_id")
	statusStr := r.URL.Query().Get("status")
	shopId, err := strconv.Atoi(shopStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.PageErrResponse, 400, 0)
		return
	}
	supplierId, err := strconv.Atoi(supplierStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.PageErrResponse, 400, 0)
		return
	}
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.CountErrResponse, 400, 0)
		return
	}
	filter.DateTo = r.URL.Query().Get("date_to")
	filter.DateFrom = r.URL.Query().Get("date_from")
	filter.SupplierId = supplierId
	filter.ShopId = shopId
	*filter.Status = status
	orders, err := h.service.Orders.GetAllOrders(filter)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}

	utils.Response(w, orders)
}

// @Summary getOrderById
// @Security ApiKeyAuth
// @Tags Orders
// @Description Просмотр заказа
// @ID getOrderById
// @Accept json
// @Produce json
// @Param order_id query string false "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /orders/by-id [get]
func (h *Handler) getOrderById(w http.ResponseWriter, r *http.Request) {
	orderStr := mux.Vars(r)["order_id"]
	orderId, err := strconv.Atoi(orderStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	order, err := h.service.Orders.GetOrderById(orderId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}

	utils.Response(w, order)
}

// @Summary editOrder
// @Security ApiKeyAuth
// @Tags Orders
// @Description Корректировка заказа
// @ID editOrder
// @Accept json
// @Produce json
// @Param params body models.OrderInput true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /orders/edit [put]
func (h *Handler) editOrder(w http.ResponseWriter, r *http.Request) {
	var params *models.OrderInput
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	err = h.service.Orders.EditOrder(params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}

	utils.Response(w, consts.Success)
}
