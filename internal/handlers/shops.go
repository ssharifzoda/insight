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

// @Summary addNewShop
// @Security ApiKeyAuth
// @Tags Shops
// @Description Добавление нового магазина
// @ID addNewShop
// @Accept json
// @Produce json
// @Param params body models.ShopSW true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /shops/new [post]
func (h *Handler) addNewShop(w http.ResponseWriter, r *http.Request) {
	var params *models.Shop
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	shop, err := h.service.AddNewShop(params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}

	utils.Response(w, shop)
}

// @Summary editShop
// @Security ApiKeyAuth
// @Tags Shops
// @Description Редактирование магазина
// @ID editShop
// @Accept json
// @Produce json
// @Param params body models.ShopSW  true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /shops/edit [put]
func (h *Handler) editShop(w http.ResponseWriter, r *http.Request) {
	var params *models.Shop
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	oldShop, err := h.service.Shops.GetShop(params.Id)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	err = h.service.UpdateShopParams(params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	if oldShop.Status == 3 && params.Status == 1 {
		var user models.User
		user.RoleId = consts.ShopRoleId
		user.FullName = params.Fullname
		user.Phone = params.Phone
		user.Email = params.Email
		user.ShopId = params.Id
		user.Active = 1
		resp, err := h.service.Users.AddNewUser(&user)
		if err != nil {
			h.logger.Error(err)
			utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
			return
		}
		utils.Response(w, map[string]interface{}{
			"shop":     params,
			"login":    resp.Phone,
			"password": resp.Password,
		})
	}
	utils.Response(w, consts.Success)
}

// @Summary getAllShops
// @Security ApiKeyAuth
// @Tags Shops
// @Description Просмотр списока магазинов
// @ID getAllShops
// @Accept json
// @Produce json
// @Param page query string true "Введите данные"
// @Param limit query string true "Введите данные"
// @Param search query string false "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /shops/list [get]
func (h *Handler) getAllShops(w http.ResponseWriter, r *http.Request) {
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
	shops, count, err := h.service.GetAllShops(page, limit, search)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, map[string]interface{}{
		"reports":     shops,
		"total_count": count,
		"count":       len(shops),
		"total_page":  int(math.Ceil(float64(count) / float64(limit))),
	})
}

// @Summary getShop
// @Security ApiKeyAuth
// @Tags Shops
// @Description Просмотр магазина
// @ID getShop
// @Accept json
// @Produce json
// @Param id query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /shops/by-id [get]
func (h *Handler) getShop(w http.ResponseWriter, r *http.Request) {
	shopIdStr := mux.Vars(r)["id"]
	shopId, err := strconv.Atoi(shopIdStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	shop, err := h.service.Shops.GetShop(shopId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, shop)
}

// @Summary deleteShop
// @Security ApiKeyAuth
// @Tags Shops
// @Description Удаление магазина
// @ID deleteShop
// @Accept json
// @Produce json
// @Param id query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /shops/rm [delete]
func (h *Handler) deleteShop(w http.ResponseWriter, r *http.Request) {
	shopIdStr := mux.Vars(r)["id"]
	shopId, err := strconv.Atoi(shopIdStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	err = h.service.DeleteShop(shopId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}
