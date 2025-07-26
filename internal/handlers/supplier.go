package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"insight/internal/models"
	"insight/pkg/consts"
	"insight/pkg/utils"
	"math"
	"net/http"
	"strconv"
)

// @Summary addNewSupplier
// @Security ApiKeyAuth
// @Tags Suppliers
// @Description Добавление нового поставщика
// @ID addNewSupplier
// @Accept json
// @Produce json
// @Param params body models.Supplier true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /suppliers/new [post]
func (h *Handler) addNewSupplier(w http.ResponseWriter, r *http.Request) {
	var params *models.Supplier
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	supplier, err := h.service.AddNewSupplier(params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	var user models.User
	user.Phone = supplier.Phone
	user.SupplierId = supplier.Id
	user.RoleId = consts.SupplierRoleId
	user.FullName = supplier.Fullname
	user.Active = 1
	resp, err := h.service.Users.AddNewUser(&user)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, map[string]interface{}{
		"supplier": supplier,
		"login":    resp.Phone,
		"password": resp.Password,
	})
}

// @Summary editSupplier
// @Security ApiKeyAuth
// @Tags Suppliers
// @Description Редактирование поставщика
// @ID editSupplier
// @Accept json
// @Produce json
// @Param params body models.Supplier true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /suppliers/edit [put]
func (h *Handler) editSupplier(w http.ResponseWriter, r *http.Request) {
	var params *models.Supplier
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	err = h.service.UpdateSupplierParams(params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}

// @Summary getAllSuppliers
// @Security ApiKeyAuth
// @Tags Suppliers
// @Description Просмотр списока поставщиков
// @ID getAllSuppliers
// @Accept json
// @Produce json
// @Param page query string true "Введите данные"
// @Param limit query string true "Введите данные"
// @Param search query string false "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /suppliers/list [get]
func (h *Handler) getAllSuppliers(w http.ResponseWriter, r *http.Request) {
	pageStr := mux.Vars(r)["page"]
	limitStr := mux.Vars(r)["limit"]
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.PageErrResponse, 400, 0)
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.CountErrResponse, 400, 0)
		return
	}
	search := r.URL.Query().Get("search")
	suppliers, count, err := h.service.GetAllSuppliers(page, limit, search)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, map[string]interface{}{
		"reports":     suppliers,
		"total_count": count,
		"count":       len(suppliers),
		"total_page":  int(math.Ceil(float64(count) / float64(limit))),
	})
}

// @Summary getSupplier
// @Security ApiKeyAuth
// @Tags Suppliers
// @Description Просмотр поставщика
// @ID getSupplier
// @Accept json
// @Produce json
// @Param id query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /suppliers/by-id [get]
func (h *Handler) getSupplier(w http.ResponseWriter, r *http.Request) {
	shopIdStr := mux.Vars(r)["id"]
	shopId, err := strconv.Atoi(shopIdStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	supplier, err := h.service.GetShop(shopId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, supplier)
}

// @Summary deleteSupplier
// @Security ApiKeyAuth
// @Tags Suppliers
// @Description Удаление поставщика
// @ID deleteSupplier
// @Accept json
// @Produce json
// @Param id query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /suppliers/rm [delete]
func (h *Handler) deleteSupplier(w http.ResponseWriter, r *http.Request) {
	supplierIdStr := mux.Vars(r)["id"]
	supplierId, err := strconv.Atoi(supplierIdStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	err = h.service.DeleteSupplier(supplierId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}
